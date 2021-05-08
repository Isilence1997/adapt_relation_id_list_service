#!/bin/bash
#支持交叉编译
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

env=test              #参考下图如果获取该值
app=video_app_short_video
server=adapt_relation_id_list_service
test=interface_test
instances=cls-rka7p1vd-aad7ef52e4c7d33c4bb7b1859bc4ee95-0
user=isaacmxu             #自己的rtx名字
mkdir bin
mkdir script
mkdir conf
go build -o ${server}
go test ./test -c -o ${test}

cp ${test} ./bin
cp ${server} ./bin
tar zcvf ${server}.tgz bin/ script/ conf/

#本地打包的文件名为{YourServer}.tgz
dtools dpatch -tgz ${server}.tgz -env ${env} -app ${app} -server ${server} -instances ${instances} -user ${user}
rm -rf bin
rm -rf script
rm -rf conf
rm  ${server}.tgz
rm  ${test}
