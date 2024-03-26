# shellcheck disable=SC2046
serviceName="attachment-service"
port=9001
env="production"
swaggerDir="third_party/openapiv2"
make build
docker stop -t 0 $(docker ps -a| grep ${serviceName} | awk '{print $1}')
docker stop -t 0 $(docker ps -a | grep "Exited" | awk '{print $1 }')
docker rm -f $(docker ps -a| grep ${serviceName} | awk '{print $1}')
docker rm -f $(docker ps -a | grep "Exited" | awk '{print $1 }')
docker images|grep ${serviceName}|awk '{print $3}'| xargs docker rmi -f
docker rmi -f $(docker images | grep "none" | awk '{print $3}')
docker build --build-arg name=${serviceName} \
             --build-arg env=${env} \
             --build-arg port=${port} \
             --build-arg swaggerDir=${swaggerDir} \
             -t ${serviceName} .
docker run --restart=always -d \
          -p ${port}:${port} \
          --name ${serviceName} ${serviceName}
