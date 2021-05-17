package logic

import (
	"context"
	"testing"
	"timeline_id_list/common/errorcode"
	"timeline_id_list/config"

	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/errs"
	relationship_read "git.code.oa.com/trpcprotocol/component_plat/account_service_user_relationship_read"
	"git.code.oa.com/trpcprotocol/component_plat/common_comm"
	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
	followRead "git.code.oa.com/trpcprotocol/video_app_short_video/trpc_follow_read"
	cFollowInnerJce "git.code.oa.com/video_app_short_video/short_video_trpc_proto/ugc_follow_inner"
	cFollowInnerJceMock "git.code.oa.com/video_app_short_video/short_video_trpc_proto/ugc_follow_inner/mock_proxy"
	"git.code.oa.com/vlib/go/video_common_api/componenthead"
	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
)

func TestGetIDList(t *testing.T) {
	p := gomonkey.ApplyFunc(config.GetConfig, func() config.ServiceConfig {
		return config.ServiceConfig{
			UserRelationshipService: config.RelationshipService{
				ReadServiceName:      "polaris://trpc.account_service.trpc_user_relationship_read.UserRelationshipRead",
				ReadServiceNamespace: "Development",
			},
			BizScene: struct {
				SubsRelScene    string `json:"subs_rel_scene" yaml:"subs_rel_scene"`
				FollowRelScene  string `json:"follow_rel_scene" yaml:"follow_rel_scene"`
				SubsFansScene   string `json:"subs_fans_scene" yaml:"subs_fans_scene"`
				FollowFansScene string `json:"follow_fans_scene" yaml:"follow_fans_scene"`
			}{
				SubsRelScene:    "subs_rel",
				FollowRelScene:  "follow_rel",
				SubsFansScene:   "subs_fans",
				FollowFansScene: "follow_fans",
			},
		}
	})
	defer p.Reset()

	// mock componenthead.SetComponentReqHead
	p1 := gomonkey.ApplyFunc(componenthead.SetComponentReqHead,
		func(ctx context.Context, componentHead *common_comm.ComponentReqHead) error {
			return nil
		})
	defer p1.Reset()

	// mock RPC Call
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()
	relationShipReadMock := relationship_read.NewMockUserRelationShipReadClientProxy(mockCtr)
	relationShipReadMock.EXPECT().GetFollowList(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, req *relationship_read.GetFollowListReq,
			opts ...client.Option) (rsp *relationship_read.GetFollowListRsp, err error) {
			if req.Id == "Vuid1" && req.Start == 0 {
				return &relationship_read.GetFollowListRsp{
					RetCode:     0,
					HasNextPage: false,
					UserInfos: []*relationship_read.UserInfo{
						{
							UserId:     "Vcuid1",
							DetailInfo: map[string]string{},
						},

						{
							UserId:     "Vcuid2",
							DetailInfo: map[string]string{},
						},

						{
							UserId:     "Vcuid3",
							DetailInfo: map[string]string{},
						},
					},
				}, nil
			} else if req.Id == "Vuid2" {
				return &relationship_read.GetFollowListRsp{
					RetCode:     -1,
					HasNextPage: false,
					UserInfos:   []*relationship_read.UserInfo{},
				}, nil
			}
			return &relationship_read.GetFollowListRsp{
				RetCode:     0,
				HasNextPage: false,
				UserInfos:   []*relationship_read.UserInfo{},
			}, errs.New(-1, "err_msg")
		}).AnyTimes()

	relationShipReadMock.EXPECT().GetFansList(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, req *relationship_read.GetFansListReq,
			opts ...client.Option) (rsp *relationship_read.GetFansListRsp, err error) {
			if req.Id == "Vcuid1" && req.Start == 0 {
				return &relationship_read.GetFansListRsp{
					RetCode:     0,
					HasNextPage: false,
					UserInfos: []*relationship_read.UserInfo{
						{
							UserId:     "Vuid1",
							DetailInfo: map[string]string{},
						},

						{
							UserId:     "Vuid2",
							DetailInfo: map[string]string{},
						},

						{
							UserId:     "Vuid3",
							DetailInfo: map[string]string{},
						},
					},
				}, nil
			} else if req.Id == "Vcuid2" {
				return &relationship_read.GetFansListRsp{
					RetCode:     -1,
					HasNextPage: false,
					UserInfos:   []*relationship_read.UserInfo{},
				}, nil
			}
			return &relationship_read.GetFansListRsp{
				RetCode:     0,
				HasNextPage: false,
				UserInfos:   []*relationship_read.UserInfo{},
			}, errs.New(-1, "err_msg")
		}).AnyTimes()

	p2 := gomonkey.ApplyFunc(relationship_read.NewUserRelationShipReadClientProxy,
		func(opts ...client.Option) relationship_read.UserRelationShipReadClientProxy {
			return relationShipReadMock
		})
	defer p2.Reset()

	// mock RPC Call
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	UgcFollowInnerMock := cFollowInnerJceMock.NewMockUgcFollowInnerServiceProxy(mockCtrl)
	UgcFollowInnerMock.EXPECT().QueryFollowVpps(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, req *cFollowInnerJce.QueryFollowVppsReq,
			opts ...client.Option) (*cFollowInnerJce.QueryFollowVppsRsp, error) {
			if req.User.UserId == 1 {
				return &cFollowInnerJce.QueryFollowVppsRsp{
					HasNextPage: false,
					PageContext: "",
					VecVppIds: []cFollowInnerJce.User{
						{
							UserId: 123,
						},
						{
							UserId: 234,
						},
					},
				}, nil
			} else if req.User.UserId == 2 {
				return &cFollowInnerJce.QueryFollowVppsRsp{
					Result:      -1,
					HasNextPage: false,
					PageContext: "",
					VecVppIds: []cFollowInnerJce.User{
						{
							UserId: 123,
						},
						{
							UserId: 234,
						},
					},
				}, nil
			}
			return &cFollowInnerJce.QueryFollowVppsRsp{}, errs.New(1, "err_msg")
		}).AnyTimes()

	p3 := gomonkey.ApplyFunc(cFollowInnerJce.NewUgcFollowInnerServiceProxy,
		func(name string) cFollowInnerJce.UgcFollowInnerServiceProxy {
			return UgcFollowInnerMock
		})
	defer p3.Reset()

	// mock RPC Call
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	followReadMock := followRead.NewMockFollowReadClientProxy(mockctrl)
	followReadMock.EXPECT().QueryFansListIdxCount(context.Background(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, req *followRead.QueryFansIdxCountRequest) (
			*followRead.QueryFansIdxCountResponse, error) {
			if req.Vuid == 1234567 {
				return &followRead.QueryFansIdxCountResponse{
					ErrCode:  0,
					ErrMsg:   "Success",
					IdxCount: 5,
				}, nil
			} else if req.Vuid == 7654321 {
				return &followRead.QueryFansIdxCountResponse{
					ErrCode:  1000,
					ErrMsg:   "ErrMsg",
					IdxCount: 5,
				}, nil
			} else if req.Vuid == 76543210 {
				return &followRead.QueryFansIdxCountResponse{
					ErrCode:  1000,
					ErrMsg:   "ErrMsg",
					IdxCount: 5,
				}, errs.New(errorcode.CallQueryFansListIdxCountError, "Error")
			} else if req.Vuid == 12345678 || req.Vuid == 123456789 {
				return &followRead.QueryFansIdxCountResponse{
					ErrCode:  0,
					ErrMsg:   "Success",
					IdxCount: 1,
				}, nil
			}
			return &followRead.QueryFansIdxCountResponse{}, nil
		}).AnyTimes()

	followReadMock.EXPECT().QueryFansList(context.Background(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, req *followRead.QueryFansListRequest) (
			*followRead.QueryFansListResponse, error) {
			if req.Vuid == 1234567 && req.Idx >= 1 && req.Idx <= 5 {
				return &followRead.QueryFansListResponse{
					ErrCode:   0,
					ErrMsg:    "",
					FansVuids: []int64{1},
				}, nil
			} else if req.Vuid == 12345678 {
				return &followRead.QueryFansListResponse{
					ErrCode:   0,
					ErrMsg:    "",
					FansVuids: []int64{1},
				}, errs.New(errorcode.CallQueryFansListError, "Error")
			} else if req.Vuid == 123456789 {
				return &followRead.QueryFansListResponse{
					ErrCode:   1000,
					ErrMsg:    "",
					FansVuids: []int64{1},
				}, nil
			}
			return &followRead.QueryFansListResponse{}, nil
		}).AnyTimes()

	p4 := gomonkey.ApplyFunc(followRead.NewFollowReadClientProxy,
		func(opt ...client.Option) followRead.FollowReadClientProxy {
			return followReadMock
		})
	defer p4.Reset()

	type args struct {
		ctx         context.Context
		inputParam  *pb.GetRelationIDListReq
		outputParam *pb.GetRelationIDListRsp
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "NormalCase",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "Vuid1",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
					},
					Scene: "subs_rel",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: false,
		},

		{
			name: "AbnormalCase",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "Vuid2",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
					},
					Scene: "subs_rel",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: true,
		},

		{
			name: "AbnormalCase2",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "Vuid3",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
					},
					Scene: "subs_rel",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: true,
		},

		{
			name: "NormalCase1",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "Vcuid1",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
					},
					Scene: "subs_fans",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: false,
		},

		{
			name: "AbnormalCase3",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "Vcuid2",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
					},
					Scene: "subs_rel",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: true,
		},

		{
			name: "AbnormalCase4",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "Vcuid3",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
					},
					Scene: "subs_rel",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: true,
		},

		{
			name: "NormalCaseFollow_rel",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "1",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
						PageContext: map[string]string{
							"page_context": "",
						},
					},
					Scene: "follow_rel",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: false,
		},

		{
			name: "AbnormalCaseFollow_rel",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "2",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
						PageContext: map[string]string{
							"page_context": "",
						},
					},
					Scene: "follow_rel",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: true,
		},

		{
			name: "AbnormalCaseFollow_rel2",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "3",
					PageInfo: &pb.RelationIDListPageInfo{
						Offset:   0,
						PageSize: 100,
						PageContext: map[string]string{
							"page_context": "",
						},
					},
					Scene: "follow_rel",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: true,
		},

		{
			name: "NormalCaseFollow_fans",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "1234567",
					PageInfo: &pb.RelationIDListPageInfo{
						PageContext: map[string]string{},
					},
					Scene: "follow_fans",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: false,
		},

		{
			name: "NormalCaseFollow_fans1",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "1234567",
					PageInfo: &pb.RelationIDListPageInfo{
						PageContext: map[string]string{
							"index":    "2",
							"idxCount": "2",
						},
					},
					Scene: "follow_fans",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: false,
		},

		{
			name: "AbnormalCaseFollow_fans",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "7654321",
					PageInfo: &pb.RelationIDListPageInfo{
						PageContext: map[string]string{},
					},
					Scene: "follow_fans",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: true,
		},

		{
			name: "AbnormalCaseFollow_fans1",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "76543210",
					PageInfo: &pb.RelationIDListPageInfo{
						PageContext: map[string]string{},
					},
					Scene: "follow_fans",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: true,
		},

		{
			name: "NormalCaseFollow_fans",
			args: args{
				ctx: context.Background(),
				inputParam: &pb.GetRelationIDListReq{
					EntityId: "1234567",
					Scene:    "follow_fans",
				},
				outputParam: &pb.GetRelationIDListRsp{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetIDList(tt.args.ctx, tt.args.inputParam, tt.args.outputParam); (err != nil) != tt.wantErr {
				t.Errorf("GetIDList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
