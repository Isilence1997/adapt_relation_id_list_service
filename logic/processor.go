package logic

import (
	//内部包
	"context"
	"fmt"
	"timeline_id_list/common/errorcode"
	"timeline_id_list/config"
	"timeline_id_list/logic/subs"
	"timeline_id_list/logic/follow"

	//第三方包

	"git.code.oa.com/trpc-go/trpc-go/errs"
	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
)

// GetIDList 拉取关系和
func GetIDList(ctx context.Context, inputParam *pb.GetRelationIDListReq,
	outputParam *pb.GetRelationIDListRsp) error {
	config := config.GetConfig()
	fmt.Printf("config=[%+v]", config)
	switch inputParam.Scene {
	case config.BizScene.SubsRelScene:
		return subs.GetIDListSubsRelHelper(ctx, inputParam, outputParam)
	case config.BizScene.SubsFansScene:
		return subs.GetIDListSubsFansHelper(ctx, inputParam, outputParam)
	case config.BizScene.FollowRelScene:
		return follow.GetIDListFollowRelHelper(ctx, inputParam, outputParam)
	default:
		return errs.New(errorcode.UnknownParamError, "UnknownParamError")
	}
}
