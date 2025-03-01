#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/IM-CHAT/user-rpc-dev'
tag='latest'

container_name="IM-CHAT-user-rpc-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network IM-CHAT -v /im-chat/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 10000:10000  --name=${container_name} -d ${reso_addr}:${tag}