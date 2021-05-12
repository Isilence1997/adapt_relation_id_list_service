package subs

import (
	//内部包
	"context"
	"fmt"
	"timeline_id_list/common/errorcode"
	"timeline_id_list/config"

	//第三方包

	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/errs"
	"git.code.oa.com/trpc-go/trpc-go/log"
	relationship_read "git.code.oa.com/trpcprotocol/component_plat/account_service_user_relationship_read"
	"git.code.oa.com/trpcprotocol/component_plat/common_comm"
	"git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_data"
	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
	"git.code.oa.com/vlib/go/video_common_api/componenthead"
)

// PackUpRelRsp 装配关系链回包
func PackUpRelRsp(rsp *relationship_read.GetFollowListRsp,
	inputParam *pb.GetRelationIDListReq, outputParam *pb.GetRelationIDListRsp) {
	// 从拉取的 UserInfos array里面提取vcuids
	for _, usrInfo := range rsp.UserInfos {
		item := &pb.Item{
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
}

// GetIDListSubsRelHelper 拉取订阅关系链
func GetIDListSubsRelHelper(ctx context.Context, inputParam *pb.GetRelationIDListReq,
	outputParam *pb.GetRelationIDListRsp) error {
	userRelationshipConfig := config.GetConfig().UserRelationshipService
	// set componentHead
	componentHead := &common_comm.ComponentReqHead{
		AppInfo: &common_comm.AppInfo{
			Appid:  userRelationshipConfig.AppID,
			Appkey: userRelationshipConfig.AppKey,
		},
	}
	_ = componenthead.SetComponentReqHead(ctx, componentHead)

	proxy := relationship_read.NewUserRelationShipReadClientProxy(
		client.WithProtocol("trpc"),
		client.WithNetwork("tcp4"),
		client.WithTarget(userRelationshipConfig.ReadServiceName),
		client.WithNamespace(userRelationshipConfig.ReadServiceNamespace),
		client.WithDisableServiceRouter())
	req := &relationship_read.GetFollowListReq{
		Id:                inputParam.EntityId,
		Start:             inputParam.PageInfo.Offset,   // 拉取订阅关系列表的起始位置
		Limit:             inputParam.PageInfo.PageSize, // 一次拉取订阅链的最大限度
		NeedUserExtraInfo: false,                        // 是否需要额外信息
	}
	log.Debugf("GetIDListSubsRelHelper req=%+v", req)
	rsp, err := proxy.GetFollowList(ctx, req)
	log.Debugf("GetIDListSubsRelHelper rsp=%+v", rsp)
	// 判断返回error是否为nil
	if err != nil {
		err = errs.New(errorcode.SubsRelRPCFuncCallError, err.Error())
		return err
	}
	// 判断RetCode是否为0
	if rsp.RetCode != 0 {
		errMsg := fmt.Sprintf("req[%v] rsp[%v] code[%v]", req, rsp, rsp.RetCode)
		err = errs.New(errorcode.SubsRelReturnCodeError, errMsg)
		return err
	}
  PackUpRelRsp(rsp, inputParam, outputParam)
	return nil
}
