#!/bin/bash

docker-compose down
docker-compose up -d postgres
echo "Waiting for databases to be created..."
sleep 5
docker-compose up -d bet_api