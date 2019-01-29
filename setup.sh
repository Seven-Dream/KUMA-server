#!/bin/sh

sudo yum update -y
sudo yum install -y docker
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
export PATH=$PATH:/usr/local/bin
echo "export PATH=$PATH:/usr/local/bin" >> ~/.bashrc
sudo chmod +x /usr/local/bin/docker-compose
sudo service docker start
