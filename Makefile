format:
	gofmt -s -w .

clean:
	go clean
	rm -rf ./bin

build: clean
	sh build.sh

zip: build
	sh zip.sh

start: build
	sls offline start --useDocker start --host 0.0.0.0

deploy: build
	sls deploy --verbose --aws-profile $(AWS_PROFILE) --stage $(STAGE)

destroy:
	sls remove --verbose --aws-profile $(AWS_PROFILE) --stage $(STAGE)