.PHONY: api-codegen
api-codegen:
	@echo "Generating code from schema..."
	./_schema/bin/codegen.sh

.PHONY: sqlc-codegen
sqlc-codegen:
	@echo "sqlc codegen..."
	@docker run --rm \
		--volume $(PWD):/src \
		--workdir /src \
		sqlc/sqlc:1.27.0 \
		generate