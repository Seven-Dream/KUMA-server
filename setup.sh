#!/bin/sh

echo "----- start yum update ------"
sudo yum update -y
echo "----- finish yum update ------"
echo "----- start install docker ------"
sudo yum install -y docker
echo "----- finish install docker ------"
echo "----- start install docker compose ------"
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
echo "----- finish install docker compose ------"
sudo chmod +x /usr/local/bin/docker-compose

echo "----- start docker service ------"
sudo service docker start

echo "if you want to use KUMA-server,"
echo "export PATH=\$PATH:/usr/local/bin"
echo "docker-compose up" 
