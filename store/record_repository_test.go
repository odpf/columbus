package store_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/odpf/columbus/models"
	"github.com/odpf/columbus/store"
	"github.com/stretchr/testify/assert"
)

func TestRecordRepository(t *testing.T) {
	ctx := context.Background()

	t.Run("CreateOrReplaceMany", func(t *testing.T) {
		var testCases = []struct {
			Title      string
			ShouldFail bool
			Setup      func(cli *elasticsearch.Client, records []models.Record, recordType models.Type) error
			PostCheck  func(cli *elasticsearch.Client, records []models.Record, recordType models.Type) error
			Type       models.Type
			Records    []models.Record
		}{
			{
				Title:      "should return an error if the type under which the records are inserted does not exist",
				ShouldFail: true,
				Type: models.Type{
					Name: "i-dont-exist",
				},
				Records: []models.Record{
					{
						Urn: "sample-urn",
					},
				},
			},
			{
				Title: "should succesfully write all the documents to the index for a valid type",
				Type: models.Type{
					Name:           "dagger",
					Classification: models.TypeClassificationResource,
				},
				Records: []models.Record{
					{
						Urn: "dagger1",
						Data: map[string]interface{}{
							"foo": "bar",
						},
					},
					{
						Urn: "dagger2",
						Data: map[string]interface{}{
							"foo": "bar",
						},
					},
					{
						Urn: "dagger3",
						Data: map[string]interface{}{
							"foo": "bar",
						},
					},
				},
				Setup: func(cli *elasticsearch.Client, records []models.Record, recordType models.Type) error {
					return store.NewTypeRepository(cli).CreateOrReplace(ctx, recordType)
				},
				PostCheck: func(cli *elasticsearch.Client, records []models.Record, recordType models.Type) error {
					searchReq := esapi.SearchRequest{
						Index: []string{recordType.Name},
						Body:  strings.NewReader(`{"query":{"match_all":{}}}`),
					}
					res, err := searchReq.Do(context.Background(), cli)
					if err != nil {
						return fmt.Errorf("error querying elasticsearch: %w", err)
					}
					defer res.Body.Close()
					if res.IsError() {
						return fmt.Errorf("elasticsearch query returned error: %s", res.Status())
					}

					var response = struct {
						Hits struct {
							Hits []interface{} `json:"hits"`
						} `json:"hits"`
					}{}
					err = json.NewDecoder(res.Body).Decode(&response)
					if err != nil {
						return fmt.Errorf("error parsing elasticsearch response: %w", err)
					}
					if len(records) != len(response.Hits.Hits) {
						return fmt.Errorf("expected elasticsearch index to contain %d records, but had %d records instead", len(records), len(response.Hits.Hits))
					}

					return nil
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.Title, func(t *testing.T) {
				cli := esTestServer.NewClient()
				if testCase.Setup != nil {
					err := testCase.Setup(cli, testCase.Records, testCase.Type)
					if err != nil {
						t.Errorf("error setting up testcase: %v", err)
					}
				}
				factory := store.NewRecordRepositoryFactory(cli)
				repo, err := factory.For(testCase.Type)
				if err != nil {
					t.Fatalf("error creating record repository: %s", err)
				}

				err = repo.CreateOrReplaceMany(ctx, testCase.Records)
				if testCase.ShouldFail {
					assert.Error(t, err)
				} else if err != nil {
					t.Errorf("repository returned unexpected error: %v", err)
					return
				}
				if testCase.PostCheck != nil {
					if err := testCase.PostCheck(cli, testCase.Records, testCase.Type); err != nil {
						t.Error(err)
						return
					}
				}
			})
		}
	})

	// the following block of code setups
	// an type repository, initialised with the daggerType
	// as well as records from the file ./testdata/dagger-populate.json
	// this is used by test cases of `GetAll` and `GetByID`
	cli := esTestServer.NewClient()
	typeRepo := store.NewTypeRepository(cli)
	err := typeRepo.CreateOrReplace(ctx, daggerType)
	if err != nil {
		t.Fatalf("failed to create dagger type: %v", err)
		return
	}

	rrf := store.NewRecordRepositoryFactory(cli)
	recordRepo, err := rrf.For(daggerType)
	if err != nil {
		t.Fatalf("failed to construct record repository: %v", err)
		return
	}

	records := insertRecord(ctx, t, recordRepo)

	t.Run("GetAllIterator", func(t *testing.T) {
		type testCase struct {
			Description string
			Filter      models.RecordFilter
			ResultsFile string
		}

		t.Run("should return record iterator to iterate records", func(t *testing.T) {
			expectedResults := []models.Record{}
			raw, err := ioutil.ReadFile("./testdata/dagger.json")
			if err != nil {
				t.Fatalf("error reading results file: %v", err)
				return
			}
			err = json.Unmarshal(raw, &expectedResults)
			if err != nil {
				t.Fatalf("error parsing results file: %v", err)
				return
			}

			var actualResults []models.Record
			iterator, err := recordRepo.GetAllIterator(ctx)
			if err != nil {
				t.Fatalf("error executing GetAllIterator: %v", err)
				return
			}
			for iterator.Scan() {
				actualResults = append(actualResults, iterator.Next()...)
			}
			iterator.Close()

			if reflect.DeepEqual(expectedResults, actualResults) == false {
				t.Error(incorrectResultsError(expectedResults, actualResults))
				return
			}
		})
	})
	t.Run("GetAll", func(t *testing.T) {
		type testCase struct {
			Description string
			Filter      models.RecordFilter
			ResultsFile string
		}

		var testCases = []testCase{
			{
				Description: "should handle nil filter",
				Filter:      nil,
				ResultsFile: "./testdata/dagger.json",
			},
			{
				Description: "should handle filter by service",
				Filter: map[string][]string{
					"service": {"kafka"},
				},
				ResultsFile: "./testdata/dagger-service.json",
			},
			{
				Description: "should support a single value filter",
				Filter: map[string][]string{
					"data.country": {"id"},
				},
				ResultsFile: "./testdata/dagger-id.json",
			},
			{
				Description: "should support multi value filter",
				Filter: map[string][]string{
					"data.country": {"id", "vn"},
				},
				ResultsFile: "./testdata/dagger-vn-id.json",
			},
			{
				Description: "should support multiple terms",
				Filter: map[string][]string{
					"data.country": {"th"},
					"data.title":   {"test_grant2"},
				},
				ResultsFile: "./testdata/dagger-th-deployed.json",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.Description, func(t *testing.T) {
				expectedResults := []models.Record{}
				raw, err := ioutil.ReadFile(tc.ResultsFile)
				if err != nil {
					t.Fatalf("error reading results file: %v", err)
					return
				}
				err = json.Unmarshal(raw, &expectedResults)
				if err != nil {
					t.Fatalf("error parsing results file: %v", err)
					return
				}

				actualResults, err := recordRepo.GetAll(ctx, tc.Filter)
				if err != nil {
					t.Fatalf("error executing GetAll: %v", err)
					return
				}

				assert.Equal(t, len(expectedResults), len(actualResults))
				if reflect.DeepEqual(expectedResults, actualResults) == false {
					t.Error(incorrectResultsError(expectedResults, actualResults))
					return
				}
			})
		}
	})
	t.Run("GetByID", func(t *testing.T) {
		t.Run("data-based tests", func(t *testing.T) {
			for _, record := range records {
				recordFromRepo, err := recordRepo.GetByID(ctx, record.Urn)
				if err != nil {
					t.Errorf("unexpected error: GetByID(%q): %v", record.Urn, err)
					return
				}
				if reflect.DeepEqual(record, recordFromRepo) == false {
					t.Error(incorrectResultsError(record, recordFromRepo))
				}
			}
		})
		t.Run("should return an error if a non-existent record is requested", func(t *testing.T) {
			var id = "this-doesnt-exists"
			_, err := recordRepo.GetByID(ctx, id)
			_, ok := err.(models.ErrNoSuchRecord)
			assert.True(t, ok)
		})
	})
	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete record from index", func(t *testing.T) {
			id := "delete-id-01"
			err := recordRepo.CreateOrReplaceMany(ctx, []models.Record{
				{
					Urn:  id,
					Name: "To be deleted",
					Data: map[string]interface{}{
						"title": "To be deleted",
						"urn":   id,
					},
				},
			})
			if err != nil {
				t.Fatal(err)
			}

			err = recordRepo.Delete(ctx, id)
			assert.Nil(t, err)

			record, err := recordRepo.GetByID(ctx, id)
			assert.NotNil(t, err)
			assert.Equal(t, models.Record{}, record)
		})

		t.Run("should return custom error when record could not be found", func(t *testing.T) {
			err := recordRepo.Delete(ctx, "not-found-id")
			assert.NotNil(t, err)
			assert.IsType(t, models.ErrNoSuchRecord{}, err)
		})
	})
}

func insertRecord(ctx context.Context, t *testing.T, repo models.RecordRepository) (records []models.Record) {
	src, err := ioutil.ReadFile("./testdata/dagger.json")
	err = json.Unmarshal(src, &records)
	if err != nil {
		t.Fatalf("error reading testdata: %v", err)
		return
	}
	err = repo.CreateOrReplaceMany(ctx, records)
	if err != nil {
		t.Fatalf("error writing testdata to elasticsearch: %v", err)
		return
	}

	return
}
