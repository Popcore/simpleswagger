SWAGGER = swagger
SWAG_GENERATE = $(SWAGGER) generate
SWAG_SERVE = $(SWAGGER) serve

doc_spec:
	$(SWAG_GENERATE) spec -o ./swagger.json

doc_client:
	$(SWAG_GENERATE) Client

doc_ui:
	$(SWAG_SERVE) swagger.json -F redoc

.PHONY: doc_spec doc_client doc_ui
