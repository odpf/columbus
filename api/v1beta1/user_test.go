package v1beta1_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/odpf/compass/api"
	compassv1beta1 "github.com/odpf/compass/api/proto/odpf/compass/v1beta1"
	"github.com/odpf/compass/asset"
	"github.com/odpf/compass/discussion"
	"github.com/odpf/compass/lib/mocks"
	"github.com/odpf/compass/star"
	"github.com/odpf/compass/user"
	"github.com/odpf/salt/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGetUserStarredAssets(t *testing.T) {
	var (
		userID = uuid.NewString()
		offset = 2
		size   = 10
	)
	type testCase struct {
		Description  string
		ExpectStatus codes.Code
		Setup        func(context.Context, *mocks.StarRepository)
		PostCheck    func(resp *compassv1beta1.GetUserStarredAssetsResponse) error
	}

	var testCases = []testCase{
		{
			Description:  "should return internal server error if failed to fetch starred",
			ExpectStatus: codes.Internal,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAllAssetsByUserID(ctx, star.Filter{Offset: offset, Size: size}, userID).Return(nil, errors.New("failed to fetch starred"))
			},
		},
		{
			Description:  "should return invalid argument if star repository return invalid error",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAllAssetsByUserID(ctx, star.Filter{Offset: offset, Size: size}, userID).Return(nil, star.InvalidError{})
			},
		},
		{
			Description:  "should return not found if starred not found",
			ExpectStatus: codes.NotFound,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAllAssetsByUserID(ctx, star.Filter{Offset: offset, Size: size}, userID).Return(nil, star.NotFoundError{})
			},
		},
		{
			Description:  "should return starred assets of a user if no error",
			ExpectStatus: codes.OK,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAllAssetsByUserID(ctx, star.Filter{Offset: offset, Size: size}, userID).Return([]asset.Asset{
					{ID: "1", URN: "asset-urn-1", Type: "asset-type"},
					{ID: "2", URN: "asset-urn-2", Type: "asset-type"},
					{ID: "3", URN: "asset-urn-3", Type: "asset-type"},
				}, nil)
			},
			PostCheck: func(resp *compassv1beta1.GetUserStarredAssetsResponse) error {
				expected := &compassv1beta1.GetUserStarredAssetsResponse{
					Data: []*compassv1beta1.Asset{
						{
							Id:   "1",
							Urn:  "asset-urn-1",
							Type: "asset-type",
						},
						{
							Id:   "2",
							Urn:  "asset-urn-2",
							Type: "asset-type",
						},
						{
							Id:   "3",
							Urn:  "asset-urn-3",
							Type: "asset-type",
						},
					},
				}

				if diff := cmp.Diff(resp, expected, protocmp.Transform()); diff != "" {
					return fmt.Errorf("expected response to be %+v, was %+v", expected, resp)
				}
				return nil
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			ctx := user.NewContext(context.Background(), userID)

			logger := log.NewNoop()
			mockStarRepo := new(mocks.StarRepository)
			if tc.Setup != nil {
				tc.Setup(ctx, mockStarRepo)
			}
			defer mockStarRepo.AssertExpectations(t)

			handler := api.NewGRPCHandler(logger, &api.Dependencies{
				StarRepository: mockStarRepo,
			})

			got, err := handler.GetUserStarredAssets(ctx, &compassv1beta1.GetUserStarredAssetsRequest{
				UserId: userID,
				Offset: uint32(offset),
				Size:   uint32(size),
			})
			code := status.Code(err)
			if code != tc.ExpectStatus {
				t.Errorf("expected handler to return Code %s, returned Code %sinstead", tc.ExpectStatus.String(), code.String())
				return
			}
			if tc.PostCheck != nil {
				if err := tc.PostCheck(got); err != nil {
					t.Error(err)
					return
				}
			}
		})
	}
}

func TestGetMyStarredAssets(t *testing.T) {
	var (
		userID = uuid.NewString()
		offset = 2
		size   = 10
	)
	type testCase struct {
		Description  string
		ExpectStatus codes.Code
		Setup        func(context.Context, *mocks.StarRepository)
		PostCheck    func(resp *compassv1beta1.GetMyStarredAssetsResponse) error
	}

	var testCases = []testCase{
		{
			Description:  "should return internal server error if failed to fetch starred",
			ExpectStatus: codes.Internal,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAllAssetsByUserID(ctx, star.Filter{Offset: offset, Size: size}, userID).Return(nil, errors.New("failed to fetch starred"))
			},
		},
		{
			Description:  "should return invalid argument if star repository return invalid error",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAllAssetsByUserID(ctx, star.Filter{Offset: offset, Size: size}, userID).Return(nil, star.InvalidError{})
			},
		},
		{
			Description:  "should return not found if starred not found",
			ExpectStatus: codes.NotFound,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAllAssetsByUserID(ctx, star.Filter{Offset: offset, Size: size}, userID).Return(nil, star.NotFoundError{})
			},
		},
		{
			Description:  "should return starred assets of a user if no error",
			ExpectStatus: codes.OK,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAllAssetsByUserID(ctx, star.Filter{Offset: offset, Size: size}, userID).Return([]asset.Asset{
					{ID: "1", URN: "asset-urn-1", Type: "asset-type"},
					{ID: "2", URN: "asset-urn-2", Type: "asset-type"},
					{ID: "3", URN: "asset-urn-3", Type: "asset-type"},
				}, nil)
			},
			PostCheck: func(resp *compassv1beta1.GetMyStarredAssetsResponse) error {
				expected := &compassv1beta1.GetMyStarredAssetsResponse{
					Data: []*compassv1beta1.Asset{
						{
							Id:   "1",
							Urn:  "asset-urn-1",
							Type: "asset-type",
						},
						{
							Id:   "2",
							Urn:  "asset-urn-2",
							Type: "asset-type",
						},
						{
							Id:   "3",
							Urn:  "asset-urn-3",
							Type: "asset-type",
						},
					},
				}

				if diff := cmp.Diff(resp, expected, protocmp.Transform()); diff != "" {
					return fmt.Errorf("expected response to be %+v, was %+v", expected, resp)
				}
				return nil
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			ctx := user.NewContext(context.Background(), userID)

			logger := log.NewNoop()
			mockStarRepo := new(mocks.StarRepository)
			if tc.Setup != nil {
				tc.Setup(ctx, mockStarRepo)
			}
			defer mockStarRepo.AssertExpectations(t)
			handler := api.NewGRPCHandler(logger, &api.Dependencies{
				StarRepository: mockStarRepo,
			})

			got, err := handler.GetMyStarredAssets(ctx, &compassv1beta1.GetMyStarredAssetsRequest{
				Offset: uint32(offset),
				Size:   uint32(size),
			})
			code := status.Code(err)
			if code != tc.ExpectStatus {
				t.Errorf("expected handler to return Code %s, returned Code %sinstead", tc.ExpectStatus.String(), code.String())
				return
			}
			if tc.PostCheck != nil {
				if err := tc.PostCheck(got); err != nil {
					t.Error(err)
					return
				}
			}
		})
	}
}

func TestGetMyStarredAsset(t *testing.T) {
	var (
		userID    = uuid.NewString()
		assetID   = uuid.NewString()
		assetType = "an-asset-type"
		assetURN  = "dummy-asset-urn"
	)
	type testCase struct {
		Description  string
		ExpectStatus codes.Code
		Setup        func(context.Context, *mocks.StarRepository)
		PostCheck    func(resp *compassv1beta1.GetMyStarredAssetResponse) error
	}

	var testCases = []testCase{
		{
			Description:  "should return invalid argument if asset id is empty",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAssetByUserID(ctx, userID, assetID).Return(asset.Asset{}, star.ErrEmptyAssetID)
			},
		},
		{
			Description:  "should return invalid argument if repository return invalid error",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAssetByUserID(ctx, userID, assetID).Return(asset.Asset{}, star.InvalidError{})
			},
		},
		{
			Description:  "should return not found if star not found",
			ExpectStatus: codes.NotFound,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAssetByUserID(ctx, userID, assetID).Return(asset.Asset{}, star.NotFoundError{})
			},
		},
		{
			Description:  "should return internal server error if failed to fetch a starred asset",
			ExpectStatus: codes.Internal,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAssetByUserID(ctx, userID, assetID).Return(asset.Asset{}, errors.New("failed to fetch starred"))
			},
		},
		{
			Description:  "should return a starred assets of a user if no error",
			ExpectStatus: codes.OK,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().GetAssetByUserID(ctx, userID, assetID).Return(asset.Asset{Type: asset.Type(assetType), URN: assetURN}, nil)
			},
			PostCheck: func(resp *compassv1beta1.GetMyStarredAssetResponse) error {
				expected := &compassv1beta1.GetMyStarredAssetResponse{
					Data: &compassv1beta1.Asset{
						Urn:  assetURN,
						Type: assetType,
					},
				}

				if diff := cmp.Diff(resp, expected, protocmp.Transform()); diff != "" {
					return fmt.Errorf("expected response to be %+v, was %+v", expected, resp)
				}
				return nil
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			ctx := user.NewContext(context.Background(), userID)

			logger := log.NewNoop()
			mockStarRepo := new(mocks.StarRepository)
			if tc.Setup != nil {
				tc.Setup(ctx, mockStarRepo)
			}
			defer mockStarRepo.AssertExpectations(t)

			handler := api.NewGRPCHandler(logger, &api.Dependencies{
				StarRepository: mockStarRepo,
			})

			got, err := handler.GetMyStarredAsset(ctx, &compassv1beta1.GetMyStarredAssetRequest{
				AssetId: assetID,
			})
			code := status.Code(err)
			if code != tc.ExpectStatus {
				t.Errorf("expected handler to return Code %s, returned Code %sinstead", tc.ExpectStatus.String(), code.String())
				return
			}
			if tc.PostCheck != nil {
				if err := tc.PostCheck(got); err != nil {
					t.Error(err)
					return
				}
			}
		})
	}
}

func TestStarAsset(t *testing.T) {
	var (
		userID  = uuid.NewString()
		assetID = uuid.NewString()
	)
	type testCase struct {
		Description  string
		ExpectStatus codes.Code
		Setup        func(context.Context, *mocks.StarRepository)
	}

	var testCases = []testCase{
		{
			Description:  "should return invalid argument if asset id in param is invalid",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Create(ctx, userID, assetID).Return("", star.ErrEmptyAssetID)
			},
		},
		{
			Description:  "should return invalid argument if star repository return invalid error",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Create(ctx, userID, assetID).Return("", star.InvalidError{})
			},
		},
		{
			Description:  "should return invalid argument if user not found",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Create(ctx, userID, assetID).Return("", star.UserNotFoundError{UserID: userID})
			},
		},
		{
			Description:  "should return internal server error if failed to star an asset",
			ExpectStatus: codes.Internal,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Create(ctx, userID, assetID).Return("", errors.New("failed to star an asset"))
			},
		},
		{
			Description:  "should return ok if starring success",
			ExpectStatus: codes.OK,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Create(ctx, userID, assetID).Return("1234", nil)
			},
		},
		{
			Description:  "should return ok if asset is already starred",
			ExpectStatus: codes.OK,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Create(ctx, userID, assetID).Return("", star.DuplicateRecordError{})
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			ctx := user.NewContext(context.Background(), userID)

			logger := log.NewNoop()
			mockStarRepo := new(mocks.StarRepository)
			if tc.Setup != nil {
				tc.Setup(ctx, mockStarRepo)
			}
			defer mockStarRepo.AssertExpectations(t)

			handler := api.NewGRPCHandler(logger, &api.Dependencies{
				StarRepository: mockStarRepo,
			})

			_, err := handler.StarAsset(ctx, &compassv1beta1.StarAssetRequest{
				AssetId: assetID,
			})
			code := status.Code(err)
			if code != tc.ExpectStatus {
				t.Errorf("expected handler to return Code %s, returned Code %sinstead", tc.ExpectStatus.String(), code.String())
				return
			}
		})
	}
}

func TestUnstarAsset(t *testing.T) {
	var (
		userID  = uuid.NewString()
		assetID = uuid.NewString()
	)
	type testCase struct {
		Description  string
		ExpectStatus codes.Code
		Setup        func(context.Context, *mocks.StarRepository)
	}

	var testCases = []testCase{
		{
			Description:  "should return invalid argument if asset id in param is empty",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Delete(ctx, userID, assetID).Return(star.ErrEmptyAssetID)
			},
		},
		{
			Description:  "should return invalid argument if star repository return invalid error",
			ExpectStatus: codes.InvalidArgument,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Delete(ctx, userID, assetID).Return(star.InvalidError{})
			},
		},
		{
			Description:  "should return internal server error if failed to unstar an asset",
			ExpectStatus: codes.Internal,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Delete(ctx, userID, assetID).Return(errors.New("failed to star an asset"))
			},
		},
		{
			Description:  "should return ok if unstarring success",
			ExpectStatus: codes.OK,
			Setup: func(ctx context.Context, er *mocks.StarRepository) {
				er.EXPECT().Delete(ctx, userID, assetID).Return(nil)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			ctx := user.NewContext(context.Background(), userID)

			logger := log.NewNoop()
			mockStarRepo := new(mocks.StarRepository)
			if tc.Setup != nil {
				tc.Setup(ctx, mockStarRepo)
			}
			defer mockStarRepo.AssertExpectations(t)

			handler := api.NewGRPCHandler(logger, &api.Dependencies{
				StarRepository: mockStarRepo,
			})

			_, err := handler.UnstarAsset(ctx, &compassv1beta1.UnstarAssetRequest{
				AssetId: assetID,
			})
			code := status.Code(err)
			if code != tc.ExpectStatus {
				t.Errorf("expected handler to return Code %s, returned Code %sinstead", tc.ExpectStatus.String(), code.String())
				return
			}
		})
	}
}

func TestGetMyDiscussions(t *testing.T) {
	var (
		userID = uuid.NewString()
	)
	type testCase struct {
		Description  string
		ExpectStatus codes.Code
		Request      *compassv1beta1.GetMyDiscussionsRequest
		Setup        func(context.Context, *mocks.DiscussionRepository)
		PostCheck    func(resp *compassv1beta1.GetMyDiscussionsResponse) error
	}

	var testCases = []testCase{
		{
			Description:  `should return internal server error if fetching fails`,
			Request:      &compassv1beta1.GetMyDiscussionsRequest{},
			ExpectStatus: codes.Internal,
			Setup: func(ctx context.Context, dr *mocks.DiscussionRepository) {
				dr.EXPECT().GetAll(ctx, discussion.Filter{
					Type:                  "all",
					State:                 discussion.StateOpen.String(),
					Assignees:             []string{userID},
					SortBy:                "created_at",
					SortDirection:         "desc",
					DisjointAssigneeOwner: false,
				}).Return([]discussion.Discussion{}, errors.New("unknown error"))
			},
		},
		{
			Description: `should parse querystring to get filter`,
			Request: &compassv1beta1.GetMyDiscussionsRequest{
				Type:      "issues",
				State:     "closed",
				Labels:    "label1,label2,label4",
				Asset:     "e5d81dcd-3046-4d33-b1ac-efdd221e621d",
				Sort:      "updated_at",
				Direction: "asc",
				Size:      30,
				Offset:    50,
			},
			ExpectStatus: codes.OK,
			Setup: func(ctx context.Context, dr *mocks.DiscussionRepository) {
				dr.EXPECT().GetAll(ctx, discussion.Filter{
					Type:                  "issues",
					State:                 "closed",
					Assignees:             []string{userID},
					Assets:                []string{"e5d81dcd-3046-4d33-b1ac-efdd221e621d"},
					Labels:                []string{"label1", "label2", "label4"},
					SortBy:                "updated_at",
					SortDirection:         "asc",
					Size:                  30,
					Offset:                50,
					DisjointAssigneeOwner: false,
				}).Return([]discussion.Discussion{}, nil)
			},
		}, {
			Description: `should search by assigned or created if filter is all`,
			Request: &compassv1beta1.GetMyDiscussionsRequest{
				Filter: "all",
			},
			ExpectStatus: codes.OK,
			Setup: func(ctx context.Context, dr *mocks.DiscussionRepository) {
				dr.EXPECT().GetAll(ctx, discussion.Filter{
					Type:                  "all",
					State:                 "open",
					Assignees:             []string{userID},
					Owner:                 userID,
					SortBy:                "created_at",
					SortDirection:         "desc",
					DisjointAssigneeOwner: true,
				}).Return([]discussion.Discussion{}, nil)
			},
		},
		{
			Description:  `should set filter to default if empty`,
			ExpectStatus: codes.OK,
			Request:      &compassv1beta1.GetMyDiscussionsRequest{},
			Setup: func(ctx context.Context, dr *mocks.DiscussionRepository) {
				dr.EXPECT().GetAll(ctx, discussion.Filter{
					Type:                  "all",
					State:                 "open",
					Assignees:             []string{userID},
					SortBy:                "created_at",
					SortDirection:         "desc",
					Size:                  0,
					Offset:                0,
					DisjointAssigneeOwner: false,
				}).Return([]discussion.Discussion{}, nil)
			},
		},
		{
			Description:  "should return ok along with list of discussions",
			ExpectStatus: codes.OK,
			Request:      &compassv1beta1.GetMyDiscussionsRequest{},
			Setup: func(ctx context.Context, dr *mocks.DiscussionRepository) {
				dr.EXPECT().GetAll(ctx, discussion.Filter{
					Type:                  "all",
					State:                 discussion.StateOpen.String(),
					Assignees:             []string{userID},
					SortBy:                "created_at",
					SortDirection:         "desc",
					DisjointAssigneeOwner: false,
				}).Return([]discussion.Discussion{
					{ID: "1122"},
					{ID: "2233"},
				}, nil)
			},
			PostCheck: func(resp *compassv1beta1.GetMyDiscussionsResponse) error {
				expected := &compassv1beta1.GetMyDiscussionsResponse{
					Data: []*compassv1beta1.Discussion{
						{Id: "1122"},
						{Id: "2233"},
					},
				}

				if diff := cmp.Diff(resp, expected, protocmp.Transform()); diff != "" {
					return fmt.Errorf("expected response to be %+v, was %+v", expected, resp)
				}
				return nil
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			ctx := user.NewContext(context.Background(), userID)

			logger := log.NewNoop()
			mockDiscussionRepo := new(mocks.DiscussionRepository)
			if tc.Setup != nil {
				tc.Setup(ctx, mockDiscussionRepo)
			}
			defer mockDiscussionRepo.AssertExpectations(t)

			handler := api.NewGRPCHandler(logger, &api.Dependencies{
				DiscussionRepository: mockDiscussionRepo,
			})

			got, err := handler.GetMyDiscussions(ctx, tc.Request)
			code := status.Code(err)
			if code != tc.ExpectStatus {
				t.Errorf("expected handler to return Code %s, returned Code %sinstead", tc.ExpectStatus.String(), code.String())
				return
			}
			if tc.PostCheck != nil {
				if err := tc.PostCheck(got); err != nil {
					t.Error(err)
					return
				}
			}
		})
	}
}
