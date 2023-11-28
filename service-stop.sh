#!/bin/bash

PROJECT=`echo ${4/} | tr '[:upper:]' '[:lower:]'`

FOLDER=`echo $6_${PROJECT}`

mkdir -p ~/.ssh
chmod 700 ~/.ssh
eval $(ssh-agent -s)
echo -e "Host *\n\tStrictHostKeyChecking no\n\n" > ~/.ssh/config
ssh-add <(echo "$2")

#ssh -oStrictHostKeyChecking=no $5@$1 docker-compose -f $FOLDER/docker-compose.$3.yml down
ssh -oStrictHostKeyChecking=no $5@$1 docker stack rm $6
