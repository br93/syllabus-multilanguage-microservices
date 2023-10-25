#!/bin/sh

compose_files=("docker-network.yml" "docker-db.yml" "docker-kafka.yml" "docker-services.yml")

for file in "${compose_files[@]}"; do
  docker-compose -f "$file" up -d && docker-compose -f "$file" rm -f
  docker rmi `docker images | grep "<none>" | awk {'print $3'}`
done
