clean:
	go clean
	rm -rf ./bin

build: clean
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/health functions/health/main.go

start: build
	sls offline --useDocker start --host 0.0.0.0

zip: build
	zip -j -9 ./bin/health.zip ./bin/health

format:
	gofmt -s -w .

deploy: build
	sls deploy --verbose --aws-profile $(AWS_PROFILE)

destroy:
	sls remove --verbose --aws-profile $(AWS_PROFILE)