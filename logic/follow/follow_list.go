package follow

import (
	// 内部包
	"context"
	"fmt"
	"strconv"
	"time"
	"timeline_id_list/common/errorcode"
	"timeline_id_list/config"

	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"
	"git.code.oa.com/video_app_short_video/trpc_go_commonlib/errs"

	// 协议文件
	"git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_data"
	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
	cFollowInnerJce "git.code.oa.com/video_app_short_video/short_video_trpc_proto/ugc_follow_inner"
	videopacketHelper "git.code.oa.com/video_app_short_video/trpc_go_commonlib/videopacket-helper"
)

// getJCEOptions 获取jce调用客户端配置
func getJCEOptions(ctx context.Context, followConfig config.QueryFollowService) []client.Option {
	// 调用jce服务，需要构造vp包头
	videoPacket := videopacketHelper.GetVideoPacketCopy(ctx)
	videoPacket.CommHeader.ServerRoute.ToServerName = followConfig.Service
	videoPacket.CommHeader.BasicInfo.Command = int16(followConfig.Cmd)
	return []client.Option{
		client.WithServiceName(followConfig.Service),
		client.WithReqHead(&videoPacket),
		client.WithProtocol("videopacket"),
		client.WithNetwork("tcp4"),
		client.WithTarget(followConfig.ReadServiceName),
		client.WithNamespace(followConfig.ReadServiceNamespace),
		client.WithTimeout(time.Duration(followConfig.Timeout) * time.Millisecond),
		client.WithDisableServiceRouter(),
	}
}

// PackUpRelRsp 装配回包关注粉丝链
func PackUpRelRsp(rsp *cFollowInnerJce.QueryFollowVppsRsp,
	outputParam *pb.GetRelationIDListRsp) {
	for _, IDs := range rsp.VecVppIds {
		item := &pb.Item{
			Id:     strconv.FormatInt(IDs.UserId, 10),
			IdType: video_timeline_timeline_data.IdType_VUID,
		}
		outputParam.Items = append(outputParam.Items, item)
	}
	outputParam.HasNextPage = rsp.HasNextPage
	outputParam.PageInfo = &pb.RelationIDListPageInfo{
		PageContext: map[string]string{
			"page_context": rsp.PageContext,
		},
	}
}

// GetIDListFollowRelHelper 拉关注业务关系链的逻辑
func GetIDListFollowRelHelper(ctx context.Context, inputParam *pb.GetRelationIDListReq,
	outputParam *pb.GetRelationIDListRsp) error {
	followConfig := config.GetConfig().UserQueryFollowService
	log.Debugf("followConfig=[%+v]", followConfig)
	opts := getJCEOptions(ctx, followConfig)
	// set request param
	vuid, err := strconv.Atoi(inputParam.EntityId)
	if err != nil {
		return errs.New(errorcode.ParseVuidError, err.Error())
	}
	req := &cFollowInnerJce.QueryFollowVppsReq{
		User: cFollowInnerJce.User{
			UserId:   int64(vuid),
			UserType: cFollowInnerJce.USER_TYPE_USER_TYPE_VUID,
		},
		PageContext: inputParam.GetPageInfo().PageContext["page_context"],
	}
	proxy := cFollowInnerJce.NewUgcFollowInnerServiceProxy("UgcFollowInnerService")
	rsp, err := proxy.QueryFollowVpps(ctx, req, opts...)
	if err != nil {
		errMsg := fmt.Sprintf("QueryFollowVpps failed, vuid:%d, err:%v", vuid, err)
		return errs.New(errorcode.QueryFollowVppsError, errMsg)
	}
	if rsp.Result != 0 {
		errMsg := fmt.Sprintf("QueryFollowVpps failed, vuid:%d, rsp.Result=%d, rsp.ErrMsg=%s",
			vuid, rsp.Result, rsp.Errmsg)
		return errs.New(errorcode.QueryFollowVppsError, errMsg)
	}
	PackUpRelRsp(rsp, outputParam)
	return nil
}
