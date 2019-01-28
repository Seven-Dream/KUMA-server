#!/bin/sh

sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

export PATH=$PATH:/usr/local/bin

sudo chmod +x /usr/local/bin/docker-compose
