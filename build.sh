#!/bin/bash

# NOTE: You must have the AWS CLI installed and configured to run this script!

GOOS=linux go build main.go
zip bookie.zip main
aws lambda update-function-code --function-name perilarea-export --zip-file fileb://./bookie.zip
rm -f main

# We are passing a request only for the invocation. No input currently required by the function
aws lambda invoke \
	--invocation-type RequestResponse \
	--function-name perilarea-export \
	--log-type Tail \
	--payload '{"key1":"value1"}' \
	/dev/null