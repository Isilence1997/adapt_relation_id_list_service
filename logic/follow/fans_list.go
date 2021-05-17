package follow

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"timeline_id_list/common/errorcode"
	"timeline_id_list/config"

	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"
	"git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_data"
	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
	followRead "git.code.oa.com/trpcprotocol/video_app_short_video/trpc_follow_read"
	"git.code.oa.com/video_app_short_video/trpc_go_commonlib/errs"
)

// GetFollowerIndexHelper 拉取FansListIdxCount的个数
func GetFollowerIndexHelper(ctx context.Context,
	proxy followRead.FollowReadClientProxy, inputParam *pb.GetRelationIDListReq) (int32, error) {
	var idxCount int32
	vuid, err := strconv.ParseInt(inputParam.EntityId, 10, 64)
	if err != nil {
		errStr := fmt.Sprintf("GetFollowerIndexHelper %s", err.Error())
		return idxCount, errs.New(errorcode.ParseVuidError, errStr)
	}
	req := &followRead.QueryFansIdxCountRequest{
		Vuid: vuid,
	}
	log.Debugf("QueryFansIdxCountRequest req=%+v", req)
	rsp, err := proxy.QueryFansListIdxCount(ctx, req)
	if err != nil {
		return idxCount, errs.New(errorcode.CallQueryFansListIdxCountError, err.Error())
	}
	log.Debugf("QueryFansIdxCountResponse rsp=%+v", rsp)
	if rsp.ErrCode != 0 {
		err := fmt.Errorf("GetFollowerIndexHelper req[%v] code[%v]", req, rsp.ErrCode)
		log.Errorf("error[%v]", err)
		return idxCount, errs.New(int(rsp.ErrCode), rsp.ErrMsg)
	}
	idxCount = rsp.IdxCount
	log.Debugf("idxCount =%d", idxCount)
	return idxCount, nil
}

// PackUpFansRsp 装配回包
func PackUpFansRsp(rsp *followRead.QueryFansListResponse,
	outputParam *pb.GetRelationIDListRsp, index, idxCount int32) {
	if index >= idxCount {
		outputParam.HasNextPage = false
	} else {
		outputParam.HasNextPage = true
	}
	idx := strconv.Itoa(int(index + 1))
	outputParam.PageInfo = &pb.RelationIDListPageInfo{
		PageContext: map[string]string{
			"index":    idx,
			"idxCount": strconv.FormatInt(int64(idxCount), 10),
		},
	}
	for _, fansVuid := range rsp.FansVuids {
		item := &pb.Item{
			Id:     strconv.FormatInt(fansVuid, 10),
			IdType: video_timeline_timeline_data.IdType_VUID,
		}
		outputParam.Items = append(outputParam.Items, item)
	}
}

// GetFollowersHelper is helper function for GetFollowerAndUpdateRedis
func GetFollowersHelper(ctx context.Context, proxy followRead.FollowReadClientProxy, idxCount int32,
	index int64, inputParam *pb.GetRelationIDListReq) (*followRead.QueryFansListResponse, error) {
	vuid, err := strconv.ParseInt(inputParam.EntityId, 10, 64)
	if err != nil {
		return nil, errs.New(errorcode.ParseVuidError, err.Error())
	}
	req := &followRead.QueryFansListRequest{
		Vuid: vuid,
		Idx:  int32(index),
	}
	rsp, err := proxy.QueryFansList(ctx, req)
	if err != nil {
		return nil, errs.New(errorcode.CallQueryFansListError, err.Error())
	}
	return rsp, nil
}

// GetIDListFollowFansHelper 用于拉取关注的粉丝链
func GetIDListFollowFansHelper(ctx context.Context, inputParam *pb.GetRelationIDListReq,
	outputParam *pb.GetRelationIDListRsp) error {
	userFansConfig := config.GetConfig().UserFansService
	proxy := followRead.NewFollowReadClientProxy(
		client.WithProtocol("trpc"),
		client.WithNetwork("tcp4"),
		client.WithTarget(userFansConfig.ReadServiceName),
		client.WithNamespace(userFansConfig.ReadServiceNamespace),
		client.WithTimeout(time.Duration(userFansConfig.Timeout)*time.Millisecond),
		client.WithDisableServiceRouter(),
	)
	// 获取 idxCount
	var idxCount int32
	var err1 error
	// pageInfo 有可能为nil
	if inputParam.PageInfo == nil {
		inputParam.PageInfo = &pb.RelationIDListPageInfo{}
	}
	val, ok := inputParam.GetPageInfo().PageContext["idxCount"]
	if !ok {
		idxCount, err1 = GetFollowerIndexHelper(ctx, proxy, inputParam)
		if err1 != nil {
			return err1
		}
	} else {
		cnt, err1 := strconv.ParseInt(val, 10, 64)
		if err1 != nil {
			return errs.New(errorcode.ParseIdxCntError, err1.Error())
		}
		idxCount = int32(cnt)
	}
	// 判断是否可以拉取这一页
	index := inputParam.PageInfo.PageContext["index"]
	if index == "" {
		index = "1"
	}
	i, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		return errs.New(errorcode.ParseIndexError, err.Error())
	}
	rsp, err2 := GetFollowersHelper(ctx, proxy, idxCount, i, inputParam)
	if err2 != nil {
		return err2
	}

	PackUpFansRsp(rsp, outputParam, int32(i), idxCount)
	return nil
}
