#!/bin/bash

# NOTE: You must have the AWS CLI installed and configured to run this script!

GOOS=linux go build main.go
zip bookie.zip main
aws lambda update-function-code --function-name perilarea-export --zip-file fileb://./bookie.zip