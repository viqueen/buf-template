.PHONY: api-codegen
api-codegen:
	@echo "Generating code from schema..."
	./_schema/bin/codegen.sh
