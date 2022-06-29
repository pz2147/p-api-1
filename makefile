SERVICE_FILE_NAME?=pApi1
SERVICE_NAME?=pApi1
SERVICE_PREFIX?=pApi1
#每次发版本需要维护版本
SERVICE_VERSION?=v1.0.0


service-name:
	@echo ${SERVICE_NAME} ${SERVICE_VERSION}

# go快捷命令
tidy:
	go mod tidy

# goctl 快捷命令
goctl-api:
	 goctl api plugin -p goctl-go-compact -api ${SERVICE_FILE_NAME}.api -dir . -style goZero

swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename docs/${SERVICE_PREFIX}.json" -api api/${SERVICE_FILE_NAME}.api -dir .

swagger-file:
	@echo docs/${SERVICE_PREFIX}.json

test-rpc:
	go test ./rpc/test -v

test:
	make test-rpc

version:
	@echo ${SERVICE_VERSION}


image-api-prefix:
	@echo ${SERVICE_NAME}-api:${SERVICE_VERSION}

image-rpc-prefix:
	@echo ${SERVICE_NAME}-rpc:${SERVICE_VERSION}
