#!/bin/bash

echo "enabling https"
sudo yum install certbot python3-certbot-nginx -y
sudo certbot --nginx -n --domains gochat.us-east-1.elasticbeanstalk.com --agree-tos --email miruna.palaghean@gmail.com