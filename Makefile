.PHONY: help
help:
	@echo '                                                                          '
	@echo 'Makefile help                                                             '
	@echo '                                                                          '
	@echo 'Usage:                                                                    '
	@echo '   make proto                         编译所有probuf文件                   '
	@echo '   make build                         构建所有二进制文件                    '
	@echo '                                                                          '
	@echo '   make docker-compose-build                        构建所有image       '
	@echo '   make docker-compose-up                           创建并启动所有容器       '
	@echo '   make docker-compose-stop                         停止容器               '
	@echo '   make docker-compose-start                        启动                   '
	@echo '   make docker-compose-down                         删除                   '
	@echo '   make docker-compose-ps                           查看状态               '
	@echo '   make docker-compose-logs                         查看日志              '
	@echo '                                                                          '
	@echo '   make docker                        构建所有image，用docker-compose-build替代     '
	@echo '                                                                          '
	@echo '   make run                           编译并启动容器，build, up             '
	@echo '                                                                          '
	@echo '                                                                          '
	@echo '                                                                          '

.PHONY: proto
proto:
	for d in api srv; do \
		for f in $$d/**/proto/*/*.proto; do \
			protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done


.PHONY: build
build: proto
	for d in api srv web; do \
		echo $$d; \
		for f in $$d/*; do \
			echo $$f; \
			cd $$f; make build; cd ../../; \
		done; \
	done; \
	cd micro-cors/; make build; cd ../; \


.PHONY: docker-compose-build
docker-compose-build: build
	docker-compose build


.PHONY: docker-compose-up
docker-compose-up:
	docker-compose up -d



.PHONY: docker-compose-start
docker-compose-start:
	docker-compose start


.PHONY: docker-compose-stop
docker-compose-stop:
	docker-compose stop


.PHONY: docker-compose-down
docker-compose-down:
	docker-compose down


.PHONY: docker-compose-ps
docker-compose-ps:
	docker-compose ps


.PHONY: docker-compose-logs
docker-compose-logs:
	docker-compose logs


.PHONY: docker
docker:
	for d in api srv web; do \
		echo $$d; \
		for f in $$d/*; do \
			echo $$f; \
			cd $$f; make docker; cd ../../; \
		done \
	done


.PHONY: run
run: docker-compose-build docker-compose-up

