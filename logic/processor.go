package logic

import (
	//内部包
	"context"
	"fmt"
	"timeline_id_list/common"
	"timeline_id_list/config"

	//第三方包
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/errs"
	"git.code.oa.com/trpc-go/trpc-go/log"
	relationship_read "git.code.oa.com/trpcprotocol/component_plat/account_service_user_relationship_read"
	"git.code.oa.com/trpcprotocol/component_plat/common_comm"
	"git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_data"
	"git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
	"git.code.oa.com/vlib/go/video_common_api/componenthead"
)

// GetIDListSubsRelHelper 拉取订阅关系链
func GetIDListSubsRelHelper(ctx context.Context, inputParam *pb.GetRelationIDListReq,
	outputParam *pb.GetRelationIDListRsp) error {
	// 目前订阅最多订阅500个用户
	if inputParam.PageInfo.Offset > common.MaxSubNum {
		outputParam.PageInfo = &pb.RelationIDListPageInfo{
			Offset:   inputParam.PageInfo.Offset,
			PageSize: inputParam.PageInfo.PageSize,
		}
		outputParam.HasNextPage = false
		return nil
	}
	// set componentHead
	componentHead := &common_comm.ComponentReqHead{
		AppInfo: &common_comm.AppInfo{
			Appid:  "tencent_video.mobile_client.user_relation_sub",
			Appkey: "567da9dc-2cb8-436e-9322-1ca3302e331e",
		},
	}
	err := componenthead.SetComponentReqHead(ctx, componentHead)
	if err != nil {
		log.Fatal(err)
	}
	proxy := relationship_read.NewUserRelationShipReadClientProxy(
		client.WithServiceName("trpc.account_service.trpc_user_relationship_read.UserRelationshipRead"))
	req := &relationship_read.GetFollowListReq{
		Id:                inputParam.EntityId,
		Start:             inputParam.PageInfo.Offset,   // 拉取订阅关系列表的起始位置
		Limit:             inputParam.PageInfo.PageSize, // 一次拉取订阅链的最大限度
		NeedUserExtraInfo: false,                        // 是否需要额外信息
	}
	log.Debugf("GetIDList req=%+v", req)
	rsp, err := proxy.GetFollowList(ctx, req)
	log.Debugf("GetIDList rsp=%+v", rsp)
	// 判断返回error是否为nil
	if err != nil {
		err = errs.New(common.SubsRelRPCFuncCallError, err.Error())
		return err
	}
	// 判断RetCode是否为0
	if rsp != nil || rsp.RetCode != 0 {
		errMsg := fmt.Sprintf("req[%v] rsp[%v] code[%v]", req, rsp, rsp.RetCode)
		err = errs.New(common.SubsRelReturnCodeError, errMsg)
		return err
	}
	// 从拉取的 UserInfos array里面提取vcuids
	for _, usrInfo := range rsp.UserInfos {
		item := &video_timeline_timeline_id_list.Item{
			Id:     usrInfo.UserId,
			IdType: video_timeline_timeline_data.IdType_VCUID,
		}
		outputParam.Items = append(outputParam.Items, item)
	}
	// 返回是否有下一页
	if !rsp.HasNextPage || len(rsp.UserInfos) == 0 {
		outputParam.HasNextPage = false
	} else {
		outputParam.HasNextPage = true
	}
	// 返回下一次请求的PageInfo
	outputParam.PageInfo = &pb.RelationIDListPageInfo{
		Offset:   inputParam.PageInfo.Offset + inputParam.PageInfo.PageSize,
		PageSize: inputParam.PageInfo.PageSize,
	}
	return nil
}

// GetIDList 拉取关系和
func GetIDList(ctx context.Context, inputParam *pb.GetRelationIDListReq,
	outputParam *pb.GetRelationIDListRsp) error {
	config := config.GetConfig()
	var err error
	switch inputParam.Scene {
	case config.BizScene.SubsRelScene:
		err = GetIDListSubsRelHelper(ctx, inputParam, outputParam)
		return err
	}
	return nil
}
