package interface_test

import (
	//内部包
	"context"
	"flag"
	"fmt"
	"testing"

	//第三方包
	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/client"

	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
)

var (
	//配置
	option []client.Option
	//命令字
	svrConfigPath = flag.String("conf", "/usr/local/trpc/bin/trpc_go.yaml", "config file path")
)

func TestMain(m *testing.M) {
	err := trpc.LoadGlobalConfig(*svrConfigPath)
	if err != nil {
		panic(err)
	}

	option = []client.Option{
		client.WithProtocol("trpc"),
		client.WithNetwork("tcp4"),
		client.WithDisableServiceRouter(),
	}
	m.Run()
}

func TestInterfaceCaseSubsRel(t *testing.T) {
	t.Logf("%+v", trpc.GlobalConfig())
	testing.Init()
	if !flag.Parsed() {
		flag.Parse()
	}
	if len(trpc.GlobalConfig().Server.Service) == 0 {
		panic("global config error")
	}

	target := fmt.Sprintf("ip://%s:%d", trpc.GlobalConfig().Server.Service[0].IP,
		trpc.GlobalConfig().Server.Service[0].Port)
	option = append(option, client.WithTarget(target))
	t.Logf("targetStr: %s", target)

	proxy := pb.NewIDListServiceClientProxy(option...)
	req := &pb.GetRelationIDListReq{
		EntityId: "161425204",
		Scene:    "subs_rel",
		PageInfo: &pb.RelationIDListPageInfo{
			Offset:   0,
			PageSize: 5,
		},
	}
	rsp, err := proxy.GetRelationIDList(context.Background(), req)
	if err != nil {
		t.Errorf("error:%v", err)
		return
	}

	t.Logf("rsp[%v]", rsp)
	if rsp != nil {
		for rsp.HasNextPage {
			req.PageInfo = rsp.PageInfo
			rsp, err = proxy.GetRelationIDList(context.Background(), req)
			if err != nil {
				t.Errorf("error:%v", err)
			}

			t.Logf("rsp[%+v], nums[%d],hasNextPage[%v]", rsp, len(rsp.Items), rsp.HasNextPage)
		}
	}
}

func TestInterfaceCaseSubsFans(t *testing.T) {
	t.Logf("%+v", trpc.GlobalConfig())
	testing.Init()
	if !flag.Parsed() {
		flag.Parse()
	}
	if len(trpc.GlobalConfig().Server.Service) == 0 {
		panic("global config error")
	}

	target := fmt.Sprintf("ip://%s:%d", trpc.GlobalConfig().Server.Service[0].IP,
		trpc.GlobalConfig().Server.Service[0].Port)
	option = append(option, client.WithTarget(target))
	t.Logf("targetStr: %s", target)

	proxy := pb.NewIDListServiceClientProxy(option...)
	req := &pb.GetRelationIDListReq{
		EntityId: "9000026631",
		Scene:    "subs_fans",
		PageInfo: &pb.RelationIDListPageInfo{
			Offset:   0,
			PageSize: 5,
		},
	}
	rsp, err := proxy.GetRelationIDList(context.Background(), req)
	if err != nil {
		t.Errorf("error:%v", err)
		return
	}

	t.Logf("rsp[%v]", rsp)
	if rsp != nil {
		for rsp.HasNextPage {
			req.PageInfo = rsp.PageInfo
			rsp, err = proxy.GetRelationIDList(context.Background(), req)
			if err != nil {
				t.Errorf("error:%v", err)
			}

			t.Logf("rsp[%+v], nums[%d],hasNextPage[%v]", rsp, len(rsp.Items), rsp.HasNextPage)
		}
	}
}

func TestInterfaceCaseFollowRel(t *testing.T) {
	t.Logf("%+v", trpc.GlobalConfig())
	testing.Init()
	if !flag.Parsed() {
		flag.Parse()
	}
	if len(trpc.GlobalConfig().Server.Service) == 0 {
		panic("global config error")
	}

	target := fmt.Sprintf("ip://%s:%d", trpc.GlobalConfig().Server.Service[0].IP,
		trpc.GlobalConfig().Server.Service[0].Port)
	option = append(option, client.WithTarget(target))
	t.Logf("targetStr: %s", target)

	proxy := pb.NewIDListServiceClientProxy(option...)
	req := &pb.GetRelationIDListReq{
		EntityId: "2184715911",
		Scene:    "follow_rel",
		PageInfo: &pb.RelationIDListPageInfo{
			Offset:   0,
			PageSize: 5,
		},
	}
	rsp, err := proxy.GetRelationIDList(context.Background(), req)
	if err != nil {
		t.Errorf("error:%v", err)
		return
	}

	t.Logf("rsp[%v]", rsp)
	if rsp != nil {
		for rsp.HasNextPage {
			req.PageInfo = rsp.PageInfo
			rsp, err = proxy.GetRelationIDList(context.Background(), req)
			if err != nil {
				t.Errorf("error:%v", err)
			}

			t.Logf("rsp[%+v], nums[%d],hasNextPage[%v]", rsp, len(rsp.Items), rsp.HasNextPage)
		}
	}
}

func TestInterfaceCaseFollowFans(t *testing.T) {
	t.Logf("%+v", trpc.GlobalConfig())
	testing.Init()
	if !flag.Parsed() {
		flag.Parse()
	}
	if len(trpc.GlobalConfig().Server.Service) == 0 {
		panic("global config error")
	}

	target := fmt.Sprintf("ip://%s:%d", trpc.GlobalConfig().Server.Service[0].IP,
		trpc.GlobalConfig().Server.Service[0].Port)
	option = append(option, client.WithTarget(target))
	t.Logf("targetStr: %s", target)

	proxy := pb.NewIDListServiceClientProxy(option...)
	req := &pb.GetRelationIDListReq{
		EntityId: "2358495800",
		Scene:    "follow_fans",
		PageInfo: &pb.RelationIDListPageInfo{
			PageContext: map[string]string{},
		},
	}
	rsp, err := proxy.GetRelationIDList(context.Background(), req)
	if err != nil {
		t.Errorf("error:%v", err)
		return
	}

	t.Logf("rsp[%v]", rsp)
	if rsp != nil {
		for rsp.HasNextPage {
			req.PageInfo = rsp.PageInfo
			rsp, err = proxy.GetRelationIDList(context.Background(), req)
			if err != nil {
				t.Errorf("error:%v", err)
			}

			t.Logf("rsp[%+v], nums[%d],hasNextPage[%v]", rsp, len(rsp.Items), rsp.HasNextPage)
		}
	}
}
