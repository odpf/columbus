package handlers_test

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/odpf/columbus/api/handlers"
	"github.com/odpf/columbus/lib/mock"
	"github.com/odpf/columbus/record"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeHandler(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		type testCase struct {
			Description  string
			ExpectStatus int
			Setup        func(tc *testCase, er *mock.TypeRepository)
			PostCheck    func(t *testing.T, tc *testCase, resp *http.Response) error
		}

		var types = []record.TypeName{
			"table",
			"topic",
			"job",
		}

		var testCases = []testCase{
			{
				Description:  "should return 500 status code if failing to fetch types",
				ExpectStatus: http.StatusInternalServerError,
				Setup: func(tc *testCase, er *mock.TypeRepository) {
					er.On("GetAll", context.Background()).Return([]record.TypeName{}, errors.New("failed to fetch type"))
				},
			},
			{
				Description:  "should return 500 status code if failing to fetch counts",
				ExpectStatus: http.StatusInternalServerError,
				Setup: func(tc *testCase, er *mock.TypeRepository) {
					er.On("GetAll", context.Background()).Return(types, nil)
					er.On("GetRecordsCount", context.Background()).Return(map[string]int{}, errors.New("failed to fetch records count"))
				},
			},
			{
				Description:  "should return all types with its record count",
				ExpectStatus: http.StatusOK,
				Setup: func(tc *testCase, er *mock.TypeRepository) {
					er.On("GetAll", context.Background()).Return(types, nil)
					er.On("GetRecordsCount", context.Background()).Return(map[string]int{
						"table":         10,
						"topic":         30,
						"job":           15,
						"to_be_ignored": 100,
					}, nil)
				},
				PostCheck: func(t *testing.T, tc *testCase, resp *http.Response) error {
					actual, err := ioutil.ReadAll(resp.Body)
					require.NoError(t, err)

					expected, err := json.Marshal([]map[string]interface{}{
						{"name": "table", "count": 10},
						{"name": "topic", "count": 30},
						{"name": "job", "count": 15},
					})
					require.NoError(t, err)

					assert.JSONEq(t, string(expected), string(actual))

					return nil
				},
			},
		}
		for _, tc := range testCases {
			t.Run(tc.Description, func(t *testing.T) {
				er := new(mock.TypeRepository)
				defer er.AssertExpectations(t)
				tc.Setup(&tc, er)

				handler := handlers.NewTypeHandler(new(mock.Logger), er)
				rr := httptest.NewRequest("GET", "/", nil)
				rw := httptest.NewRecorder()

				handler.Get(rw, rr)
				if rw.Code != tc.ExpectStatus {
					t.Errorf("expected handler to return %d status, was %d instead", tc.ExpectStatus, rw.Code)
					return
				}

				if tc.PostCheck != nil {
					if err := tc.PostCheck(t, &tc, rw.Result()); err != nil {
						t.Error(err)
					}
				}
			})
		}
	})
}
