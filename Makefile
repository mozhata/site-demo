.PHONY: vendor
vendor:
	godep save ./...

.PHONY: server
server:
	cd ./docker;docker-compose up -d

.PHONY: log
log:
	cd ./docker;docker-compose logs -f goserver