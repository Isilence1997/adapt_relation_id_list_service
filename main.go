package main

import (
	_ "go.uber.org/automaxprocs"

	_ "git.code.oa.com/trpc-go/trpc-config-tconf"
	_ "git.code.oa.com/trpc-go/trpc-filter/debuglog"
	_ "git.code.oa.com/trpc-go/trpc-filter/recovery"
	_ "git.code.oa.com/trpc-go/trpc-log-atta"
	_ "git.code.oa.com/trpc-go/trpc-metrics-m007"
	_ "git.code.oa.com/trpc-go/trpc-metrics-runtime"
	_ "git.code.oa.com/trpc-go/trpc-naming-polaris"
	_ "git.code.oa.com/trpc-go/trpc-opentracing-tjg"
	_ "git.code.oa.com/trpc-go/trpc-selector-cl5"

	"git.code.oa.com/trpc-go/trpc-go/log"

	trpc "git.code.oa.com/trpc-go/trpc-go"
	pb "git.code.oa.com/trpcprotocol/component_plat/video_timeline_timeline_id_list"
)

type iDListServiceServiceImpl struct{}

func main() {

	s := trpc.NewServer()

	pb.RegisterIDListServiceService(s, &iDListServiceServiceImpl{})

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
