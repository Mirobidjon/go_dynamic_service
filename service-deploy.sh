#!/bin/sh

PROJECT=`echo ${6/} | tr '[:upper:]' '[:lower:]'`
REGISTRY=`$3`
FOLDER=`echo ${PROJECT}`
IMAGENAME=`echo $3/${PROJECT}:$4`

mkdir -p ~/.ssh
chmod 700 ~/.ssh
eval $(ssh-agent -s)
echo -e "Host *\n\tStrictHostKeyChecking no\n\n" > ~/.ssh/config
ssh-add <(echo "$2")

ssh -oStrictHostKeyChecking=no $7@$1 mkdir $FOLDER -p
scp -oStrictHostKeyChecking=no ./docker-compose.$5.yml $7@$1:$FOLDER/
ssh -oStrictHostKeyChecking=no $7@$1 env IMAGE=$IMAGENAME env APP_NAME=$8 docker stack deploy -c $FOLDER/docker-compose.$5.yml $8
