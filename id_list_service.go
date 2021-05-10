package main

import (
	"context"
	"fmt"
	"timeline_id_list/common"
	"timeline_id_list/logic"

	"git.code.oa.com/trpc-go/trpc-go/errs"
	"git.code.oa.com/trpc-go/trpc-go/log"
	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
	"git.code.oa.com/video_app_short_video/trpc_go_commonlib/atta-report"
)

// GetRelationIDList 获取关系链列表
func (s *iDListServiceServiceImpl) GetRelationIDList(ctx context.Context,
	req *pb.GetRelationIDListReq, rsp *pb.GetRelationIDListRsp) error {
	// 检查参数req, rsp是否的nil
	log.Debugf("req[%+v]", req)
	if req == nil || rsp == nil {
		errStr := fmt.Sprintf("GetRelationIDList req=[%+v] or rsp=[%+v] is nil", req, rsp)
		err := errs.New(common.ParamsInvalidError, errStr)
		log.Error(err)
		atta.AttaLogTrpcCtxWithVuid(ctx, fmt.Sprintf("GetRelationIDList req=[%+v] or rsp=[%+v] is nil", req, rsp),
			"GetRelationIDList failed", req.EntityId)
		return err
	}
	// 检查入参数是否合法
	if req.EntityId == "" {
		errStr := fmt.Sprintf("GetRelationIDList with empty user_id, req=[%+v], rsp=[%+v]", req, rsp)
		err := errs.New(common.EmptyInputIDError, errStr)
		log.Error(err)
		atta.AttaLogTrpcCtxWithVuid(ctx, fmt.Sprintf("GetRelationIDList with empty user_id, req=[%+v], rsp=[%+v]",
			req, rsp), "GetRelationIDList failed", req.EntityId)
		return err
	}
	// 执行拉取关系列表主逻辑
	err := logic.GetIDList(ctx, req, rsp)
	if err != nil {
		log.Errorf("GetRelationIDList req=[%+v], rsp=[%+v], err=[%v]", req, rsp, err)
		atta.AttaLogTrpcCtxWithVuid(ctx, fmt.Sprintf("GetRelationIDList req=[%+v], rsp=[%+v], err=[%v]",
			req, rsp, err), "GetRelationIDList failed", req.EntityId)
		return err
	}
	log.Debugf("rsp[%+v]", rsp)
	atta.AttaLogTrpcCtxWithVuid(ctx, fmt.Sprintf("GetRelationIDList req=[%+v], rsp=[%+v]",
			req, rsp), "GetRelationIDList succeed", req.EntityId)
	return nil
}

// GetWorksIDList 获取作品列表
func (s *iDListServiceServiceImpl) GetWorksIDList(ctx context.Context,
	req *pb.GetWorksIDListReq, rsp *pb.GetWorksIDListRsp) error {
	// implement business logic here ...
	// ...

	return nil
}
