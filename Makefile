schema-codegen:
	@echo "Generating code from schema..."
	@docker run --rm \
		--volume $(PWD):/workspace \
		--volume $(PWD)/api:/workspace/api \
		--workdir /workspace \
		ghcr.io/viqueen/docker-images-protobuf-gen:main generate --verbose
