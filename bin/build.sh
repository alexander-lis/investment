#!/bin/sh

docker build -f ./src/app/gateway/Dockerfile -t lisitsynalex/investment/gateway ./src/app
#docker push lisitsynalex/investment/gateway

docker build -f ./src/app/services/user/Dockerfile -t lisitsynalex/investment/services/user ./src/app
#docker push lisitsynalex/investment/user_uservice

docker compose -f ./bin/compose.yaml up -d