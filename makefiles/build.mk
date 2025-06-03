.PHONY: codegen
codegen:
	@echo "Generating code from schema..."
	@docker run --rm \
		--volume $(PWD):/workspace \
		--volume $(PWD)/api:/workspace/api \
		--workdir /workspace \
		ghcr.io/viqueen/docker-images-protobuf-gen:main generate --verbose

.PHONY: sqlc
sqlc:
	@echo "sqlc codegen..."
	@docker run --rm \
		--volume $(PWD):/src \
		--workdir /src \
		sqlc/sqlc:1.27.0 \
		generate