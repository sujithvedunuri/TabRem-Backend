#!/bin/sh -x

RED='\033[0;31m'
LIGHTRED='\033[1;31m'
CYAN='\033[0;36m'
GREEN='\033[0;32m'

echo -e "${RED}stopping container"
docker stop mysql-db
echo -e "${GREEN}restart"
docker compose up -d