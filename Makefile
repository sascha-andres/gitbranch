.PHONY: build install

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s'

install: build
	go install

docker: build
	docker build -t briefbote/gitbranch:latest .
	docker push briefbote/gitbranch:latest