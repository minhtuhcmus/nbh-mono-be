docker-compose stop
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker image rm $(docker image ls -q)
docker volume rm $(docker volume ls -q)