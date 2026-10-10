default: test

# Load environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Run all acceptance tests against whatever IOSXR_* environment is
# currently exported (no version defaulting, no .env sub-target lookup).
# Usage: make testall [TESTARGS="-run IosxrLogging"]
.PHONY: testall
testall:
	TF_ACC=1 go test -v $(TESTARGS) -timeout 120m ./internal/provider

# Run acceptance tests against the default versions (24.4, 25.4, 26.2)
# Usage: make test [NAME=TestName] [DEBUG=1]
.PHONY: test
test: test-244 test-254 test-262
	@echo ""
	@echo "All multi-version tests completed!"

# Test against 24.4 XRV9K
# Usage: make test-244 [NAME=TestName] [DEBUG=1]
.PHONY: test-244
test-244:
	@echo "========================================="
	@echo "Testing against 24.4 XRV9K..."
	@echo "========================================="
	@if [ -z "$(IOSXR_244_HOST)" ]; then \
		echo "SKIPPED: IOSXR_244_HOST is not configured"; \
		echo "To enable this test, configure IOSXR_244_HOST in your .env file"; \
	else \
		$(if $(NAME),echo "Running tests matching: $(NAME)";) \
		TF_ACC=1 \
		IOSXR_HOST=$(IOSXR_244_HOST) \
		IOSXR_USERNAME=$(or $(IOSXR_244_USERNAME),$(IOSXR_USERNAME)) \
		IOSXR_PASSWORD=$(or $(IOSXR_244_PASSWORD),$(IOSXR_PASSWORD)) \
		IOSXR_VERSION=24.4 \
		XRV9K=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $(if $(NAME),-run $(NAME)) $(TESTARGS) -count 1 -timeout 120m ./internal/provider/ $(if $(DEBUG),2>&1 | tee test-output-244.log); \
	fi

# Test against 25.4 XRV9K
# Usage: make test-254 [NAME=TestName] [DEBUG=1]
.PHONY: test-254
test-254:
	@echo "========================================="
	@echo "Testing against 25.4 XRV9K..."
	@echo "========================================="
	@if [ -z "$(IOSXR_254_HOST)" ]; then \
		echo "SKIPPED: IOSXR_254_HOST is not configured"; \
		echo "To enable this test, configure IOSXR_254_HOST in your .env file"; \
	else \
		$(if $(NAME),echo "Running tests matching: $(NAME)";) \
		TF_ACC=1 \
		IOSXR_HOST=$(IOSXR_254_HOST) \
		IOSXR_USERNAME=$(or $(IOSXR_254_USERNAME),$(IOSXR_USERNAME)) \
		IOSXR_PASSWORD=$(or $(IOSXR_254_PASSWORD),$(IOSXR_PASSWORD)) \
		IOSXR_VERSION=25.4 \
		XRV9K=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $(if $(NAME),-run $(NAME)) $(TESTARGS) -count 1 -timeout 120m ./internal/provider/ $(if $(DEBUG),2>&1 | tee test-output-254.log); \
	fi

# Test against 26.2 XRV9K
# Usage: make test-262 [NAME=TestName] [DEBUG=1]
.PHONY: test-262
test-262:
	@echo "========================================="
	@echo "Testing against 26.2 XRV9K..."
	@echo "========================================="
	@if [ -z "$(IOSXR_262_HOST)" ]; then \
		echo "SKIPPED: IOSXR_262_HOST is not configured"; \
		echo "To enable this test, configure IOSXR_262_HOST in your .env file"; \
	else \
		$(if $(NAME),echo "Running tests matching: $(NAME)";) \
		TF_ACC=1 \
		IOSXR_HOST=$(IOSXR_262_HOST) \
		IOSXR_USERNAME=$(or $(IOSXR_262_USERNAME),$(IOSXR_USERNAME)) \
		IOSXR_PASSWORD=$(or $(IOSXR_262_PASSWORD),$(IOSXR_PASSWORD)) \
		IOSXR_VERSION=26.2 \
		XRV9K=1 \
		$(if $(DEBUG),TF_LOG=Trace) \
		go test -v $(if $(NAME),-run $(NAME)) $(TESTARGS) -count 1 -timeout 120m ./internal/provider/ $(if $(DEBUG),2>&1 | tee test-output-262.log); \
	fi

# Update files from a single definition
# Usage: make gen NAME="Logging"
# NAME: The name of the definition, e.g. "Logging"
.PHONY: gen
gen:
	go run gen/load_models.go
	go run ./gen/generator.go "$(NAME)"
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	GOFLAGS=-buildvcs=false go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name iosxr --rendered-provider-name terraform-provider-iosxr
	go run gen/doc_category.go
	go run gen/doc_version_changes.go

# Update all files
.PHONY: genall
genall:
	go run gen/load_models.go
	go run ./gen/generator.go
	go run golang.org/x/tools/cmd/goimports -w internal/provider/
	terraform fmt -recursive ./examples/
	GOFLAGS=-buildvcs=false go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name iosxr --rendered-provider-name terraform-provider-iosxr
	go run gen/doc_category.go
	go run gen/doc_version_changes.go

# Run unit tests (no device required)
.PHONY: test-unit
test-unit:
	go test -v -cover -timeout 10m ./internal/provider/helpers/...
	go test -v -cover -timeout 10m -run . gen/generator_test.go gen/generator.go

# Run acceptance tests (legacy target for backward compatibility)
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
