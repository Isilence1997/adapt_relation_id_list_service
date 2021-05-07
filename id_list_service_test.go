package main

import (
	"context"
	"testing"

	trpc "git.code.oa.com/trpc-go/trpc-go"
	_ "git.code.oa.com/trpc-go/trpc-go/http"
	_ "git.code.oa.com/trpc-go/trpc-selector-cl5"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
)

var iDListServiceService = &iDListServiceServiceImpl{}

//go:generate go mod tidy
//go:generate mockgen -destination=stub/git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list/i_d_list_service_mock.go -package=video_timeline_timeline_id_list -self_package=git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list IDListServiceClientProxy

func Test_IDListService_GetRelationIDList(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	iDListServiceClientProxy := pb.NewMockIDListServiceClientProxy(ctrl)

	// 预期行为
	m := iDListServiceClientProxy.EXPECT().GetRelationIDList(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.GetRelationIDListReq)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.GetRelationIDListRsp{}
		err := iDListServiceService.GetRelationIDList(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.GetRelationIDListReq{}

	rsp, err := iDListServiceClientProxy.GetRelationIDList(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("IDListService_GetRelationIDList req: %v", req)
	t.Logf("IDListService_GetRelationIDList rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}

func Test_IDListService_GetWorksIDList(t *testing.T) {

	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	iDListServiceClientProxy := pb.NewMockIDListServiceClientProxy(ctrl)

	// 预期行为
	m := iDListServiceClientProxy.EXPECT().GetWorksIDList(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	m.DoAndReturn(func(ctx context.Context, req interface{}, opts ...interface{}) (interface{}, error) {

		r, ok := req.(*pb.GetWorksIDListReq)
		if !ok {
			panic("invalid request")
		}

		rsp := &pb.GetWorksIDListRsp{}
		err := iDListServiceService.GetWorksIDList(trpc.BackgroundContext(), r, rsp)
		return rsp, err
	})

	// 开始写单元测试逻辑
	req := &pb.GetWorksIDListReq{}

	rsp, err := iDListServiceClientProxy.GetWorksIDList(trpc.BackgroundContext(), req)

	// 输出入参和返回 (检查t.Logf输出，运行 `go test -v`)
	t.Logf("IDListService_GetWorksIDList req: %v", req)
	t.Logf("IDListService_GetWorksIDList rsp: %v, err: %v", rsp, err)

	assert.Nil(t, err)
}
