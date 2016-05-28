.PHONY: vendor
vendor:
	godep save ./...

server:
	cd ./docker;docker-compose up -d

log:
	cd ./docker;docker-compose logs beego