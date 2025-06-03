DOCKER_FLAGS := --rm --volume ${PWD}:/go/src --workdir /go/src/backend
ifeq ($(CI),)
	DOCKER_FLAGS += -it
endif

.PHONY: lint
lint:
	docker run $(DOCKER_FLAGS) \
		golangci/golangci-lint:v2.1.2 \
		golangci-lint run

.PHONY: lint-fix
lint-fix:
	docker run $(DOCKER_FLAGS) \
    		golangci/golangci-lint:v2.1.2 \
    		golangci-lint run --fix