.PHONY: vendor
vendor:
	godep save ./...

.PHONY: server
server:
	cd ./docker;docker-compose up -d

.PHONY: log
log:
	cd ./docker;docker-compose logs -f sitedemo_server

clean_containers:
	docker rm $$(docker stop $$(docker ps -q -a))

.PHONY: connect_mysql
connect_mysql:
	mysql -h127.0.0.1 -P3308 -uroot -p
