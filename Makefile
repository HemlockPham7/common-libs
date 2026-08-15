
COVERAGE_EXCLUDE=mocks|main.go|test|config.go|client.go|level.go|mock.go|request.go|parser.go|postgres_data|infrastructure|migration.go
COVERAGE_THRESHOLD=70
COVERAGE_FOLDER=./test-output

docker-test:
	mkdir -p $(COVERAGE_FOLDER)
	docker buildx build --build-arg COVERAGE_EXCLUDE="${COVERAGE_EXCLUDE}" --progress=plain --target test -t test:test --output $(COVERAGE_FOLDER) .
	@total=$$(go tool cover -func=$(COVERAGE_FOLDER)/coverage.out | grep total: | awk '{print $$3}' | sed 's/%//'); \
	if [ $$(echo "$$total < $(COVERAGE_THRESHOLD)" | bc -l) -eq 1 ]; then \
	   echo "❌ Coverage ($$total%) is below threshold ($(COVERAGE_THRESHOLD)%)"; \
	   exit 1; \
	else \
	   echo "✅ Coverage ($$total%) meets threshold ($(COVERAGE_THRESHOLD)%)"; \
	fi