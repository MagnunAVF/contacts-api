# Contacts API
A simple contacts book REST API. Using Go, Serverless Framework and AWS.

## Requirements
* Go lang
* Node
* Serverless Framework
* Java
* AWS account

## Setup
Run in terminal:
```
npm install
sls dynamodb install
```

If you are getting errors in install command check:
* problem: https://github.com/99x/serverless-dynamodb-local/issues/294
* fix: https://github.com/99x/dynamodb-localhost/pull/78

## How to run in localhost
Run in terminal:
```
make start
```

## Deploy
Run in terminal:
```
make deploy AWS_PROFILE=<your_aws_profile> STAGE=<stage>
```

To remove the deployed stack:
```
make destroy AWS_PROFILE=<your_aws_profile> STAGE=<stage>
```
