help:
	@echo 'Makefile help                                                             '
	@echo '                                                                          '
	@echo 'Usage:                                                                    '
	@echo '   make docker-compose-build                         构建所有image       '
	@echo '   make docker-compose-up                           创建并启动所有容器       '
	@echo '   make docker-compose-stop                         停止容器               '
	@echo '   make docker-compose-start                        启动                   '
	@echo '   make docker-compose-down                         删除                   '
	@echo '   make docker-compose-ps                           查看状态               '
	@echo '                                                                          '
	@echo '   make start-redis-single                           单独启动redis         '
	@echo '                                                                          '

docker-compose-build: build
	docker-compose build

docker-compose-up:
	docker-compose up -d

docker-compose-stop:
	docker-compose stop

docker-compose-start:
	docker-compose start

docker-compose-down:
	docker-compose down

docker-compose-ps:
	docker-compose ps

docker-compose-logs:
	docker-compose logs

proto:
	for d in api srv; do \
		for f in $$d/**/proto/*/*.proto; do \
			protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done


build:
	for d in api srv web; do \
		echo $$d; \
		for f in $$d/*; do \
			echo $$f; \
			cd $$f; make build; cd ../../; \
		done \
	done


docker:
	for d in api srv web; do \
		echo $$d; \
		for f in $$d/*; do \
			echo $$f; \
			cd $$f; make docker; cd ../../; \
		done \
	done


run:
	docker-compose build
	docker-compose up
