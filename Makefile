.PHONY: api

api: ## Generate pb.go api file
	protoc -I api/spec-gen/v1alpha1 \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:api/spec-gen/v1alpha1 \
		api/spec-gen/v1alpha1/gen.proto

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dep: ## Get the dependencies
	@go get -v -d ./...

