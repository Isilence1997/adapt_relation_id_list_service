package config

import (
	"git.code.oa.com/trpc-go/trpc-go/config"
	"git.code.oa.com/trpc-go/trpc-go/log"
)

var (
	// serviceConfig 配置信息对象
	serviceConfig ServiceConfig
)

// 账号中台服务路由配置
type RelationshipService struct {
	ReadServiceName       string `json:"read_service_name" yaml:"read_service_name"`
	ReadServiceNamespace  string `json:"read_service_namespace" yaml:"read_service_namespace"`
}

// ServiceConfig 配置信息
type ServiceConfig struct {
	UserRelationshipService  RelationshipService `json:"user_relationship_service" yaml:"user_relationship_service"`
	// BizScene 业务场景配置
	BizScene struct {
		SubsRelScene    string `json:"subs_rel_scene" yaml:"subs_rel_scene"`
		FollowRelScene  string `json:"follow_rel_scene" yaml:"follow_rel_scene"`
		SubsFansScene   string `json:"subs_fans_scene" yaml:"subs_fans_scene"`
		FollowFansScene string `json:"follow_fans_scene" yaml:"follow_fans_scene"`
	} `json:"biz_scene" yaml:"biz_scene"`
}

// InitServiceConfig 初始化服务配置
func InitServiceConfig() {
	// 加载配置文件
	confName := "adapt_relation_id_list_service.yaml"
	serviceConfig = ServiceConfig{}
	err := config.GetYAML(confName, &serviceConfig)
	if err != nil {
		log.Errorf("get yaml conf error,err:%v", err)
		panic(err)
	} else {
		log.Infof("yaml conf, conf:%+v", serviceConfig)
	}
}

// GetConfig 获取配置
func GetConfig() ServiceConfig {
	return serviceConfig
}
