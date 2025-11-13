OPENAPI_FILE = openapi/openapi.yaml
OUTPUT_DIR = internal/Web

gen:
	oapi-codegen -config openapi/.openapi $(OPENAPI_FILE) > $(OUTPUT_DIR)/api.gen.go



clean:
	del /Q $(OUTPUT_DIR)\api.gen.go

regen: clean generate