WORK_DIR=$(shell pwd)
OUTPUT_DIR=$(WORK_DIR)/bin
GOPATH=$(WORK_DIR)
VENDOR_DIR=$(WORK_DIR)/src/vendor
LOG_DIR=$(WORK_DIR)/log
XLSX_DIR=$(WORK_DIR)/xlsx
########################################################################################################################
.PHONY: all clean clean-log glide-up unzip-vendor zip-vendor publish rpc model gateway login-msg login game-sd game-msg game-cache game robot
########################################################################################################################
all: server
########################################################################################################################
clean:
	rm -rf $(WORK_DIR)/pkg
	rm -f $(OUTPUT_DIR)/gateway
	rm -f $(OUTPUT_DIR)/game
	rm -f $(OUTPUT_DIR)/login
	rm -f $(OUTPUT_DIR)/robot

clean-log:
	rm -rf $(LOG_DIR)/gateway*/*.log
	rm -rf $(LOG_DIR)/game*/*.log
	rm -rf $(LOG_DIR)/login*/*.log
	rm -rf $(LOG_DIR)/robot*/*.log

glide-up:
	go_tool.sh glide-up

unzip-vendor:
	go_tool.sh unzip-vendor

zip-vendor:
	rm -rf $(VENDOR_DIR)/gitee.com/lwj8507/nggs/vendor.zip
	rm -rf $(VENDOR_DIR)/gitee.com/lwj8507/light-protoactor-go/vendor.zip
	rm -rf $(VENDOR_DIR)/github.com/coreos/etcd/cmd
	go_tool.sh zip-vendor
########################################################################################################################
publish:
	rm -rf $(WORK_DIR)/publish/bin
	mkdir -p $(WORK_DIR)/publish/bin
	cp $(OUTPUT_DIR)/gateway $(WORK_DIR)/publish/bin/gateway
	cp $(OUTPUT_DIR)/game $(WORK_DIR)/publish/bin/game
	cp $(OUTPUT_DIR)/login $(WORK_DIR)/publish/bin/login
	rm -rf $(WORK_DIR)/publish/xlsx
	mkdir -p $(WORK_DIR)/publish/xlsx
	cp $(XLSX_DIR)/*.xlsx $(WORK_DIR)/publish/xlsx
	cd $(WORK_DIR)/publish;tar czf hh-server.$(shell date "+%Y%m%d%k%M").tar.gz ./bin ./xlsx;rm -rf ./bin ./xlsx
	rm -rf $(WORK_DIR)/publish/bin
	rm -rf $(WORK_DIR)/publish/xlsx

rpc:
	@echo $(shell date "+%F %R:%S")
	cd $(WORK_DIR)/src/rpc; go generate; go test

model:
	@echo $(shell date "+%F %R:%S")
	cd $(WORK_DIR)/src/model; ./gen.sh

gateway:
	@echo $(shell date "+%F %R:%S")
	go build -o $(OUTPUT_DIR)/gateway gateway

login-msg:
	@echo $(shell date "+%F %R:%S")
	cd $(WORK_DIR)/src/login/msg; go generate; go test

login:
	@echo $(shell date "+%F %R:%S")
	go build -o $(OUTPUT_DIR)/login login

msg:
	@echo $(shell date "+%F %R:%S")
	cd $(WORK_DIR)/src/msg; ./gen.sh

sd:
	@echo $(shell date "+%F %R:%S")
	cd $(WORK_DIR)/src/sd; go generate

game-cache:
	@echo $(shell date "+%F %R:%S")
	cd $(WORK_DIR)/src/game/cache; go generate; go test

server:
	@echo $(shell date "+%F %R:%S")
	cd $(WORK_DIR)/src;go build -o $(OUTPUT_DIR)/server

robot:
	@echo $(shell date "+%F %R:%S")
	go build -o $(OUTPUT_DIR)/robot robot
########################################################################################################################
