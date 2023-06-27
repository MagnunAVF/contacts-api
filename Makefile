format:
	gofmt -s -w .

setup-local-env:
	cp .env.sample .env
	sed -i '/^HOST_IP/d' .env && echo "\nHOST_IP=$(shell hostname -I | awk '{print $$1}')" >> .env

clean:
	go clean
	rm -rf ./bin

build: clean
	sh build.sh

zip: build
	sh zip.sh

start: build
	sls offline start --useDocker --host 0.0.0.0

deploy: build
	sls deploy --verbose --aws-profile $(AWS_PROFILE) --stage $(STAGE)

destroy:
	sls remove --verbose --aws-profile $(AWS_PROFILE) --stage $(STAGE)