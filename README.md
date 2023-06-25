# Contacts API
A simple contacts book REST API. Using Go, Serverless Framework and AWS.

## Requirements
* Go lang
* Node
* Serverless Framework
* Java
* AWS account

## Setup
Create .env file:
```shell
cp .env.sample .env
```

Run in terminal:
```shell
npm install
serverless dynamodb install
```

If you are getting this error:
```
Error:
Error: Error getting DynamoDb local latest tar.gz location undefined: 403
    at ClientRequest.<anonymous> (/home/matrix/Projects/projectx/projectx-resources-api/node_modules/dynamodb-localhost/dynamodb/installer.js:29:15)
    at Object.onceWrapper (node:events:628:26)
    at ClientRequest.emit (node:events:513:28)
    at ClientRequest.emit (node:domain:489:12)
    at HTTPParser.parserOnIncomingClient [as onIncoming] (node:_http_client:701:27)
    at HTTPParser.parserOnHeadersComplete (node:_http_common:119:17)
    at Socket.socketOnData (node:_http_client:542:22)
    at Socket.emit (node:events:513:28)
    at Socket.emit (node:domain:489:12)
    at addChunk (node:internal/streams/readable:324:12)
    at readableAddChunk (node:internal/streams/readable:297:9)
    at Readable.push (node:internal/streams/readable:234:10)
    at TCP.onStreamRead (node:internal/stream_base_commons:190:23)
```

Use this commands:
```shell
sed -i 's|http://s3-us-west-2|https://s3-us-west-2|g' ./node_modules/dynamodb-localhost/dynamodb/config.json && sed -i 's|require(\"http\")|require(\"https\")|g' ./node_modules/dynamodb-localhost/dynamodb/installer.js

serverless dynamodb install
```

References to this problem:
* https://github.com/99x/serverless-dynamodb-local/issues/294
* https://github.com/99x/dynamodb-localhost/pull/78

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
