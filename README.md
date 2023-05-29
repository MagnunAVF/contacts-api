# Contacts API
A simple contacts book REST API. Using Go, Serverless Framework and AWS.

## Requirements
* Go lang
* Node
* Serverless Framework
* Docker
* AWS account

## Setup
Run in terminal:
```
npm install
```

## How to run in localhost
Run in terminal:
```
make start
```

## Deploy
Run in terminal:
```
make deploy AWS_PROFILE=<your_aws_profile>
```

To remove the deployed stack:
```
make destroy AWS_PROFILE=<your_aws_profile>
```
