.DEFAULT_GOAL := help

PACKAGE_VERSION=v0.1.0-beta.1

help: # generate annotations of each target
	@grep -hE '^[a-zA-Z_-]+:.*?#+ .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?#+ "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

version: # Print package version
	echo $(PACKAGE_VERSION)

release: # Release current code
	git tag $(PACKAGE_VERSION)
	git push origin $(PACKAGE_VERSION)

build: # Build sportscrape-development docker image
	docker compose -p sportscrape -f docker-compose.yml --project-directory . build --force-rm sportscrape-local

build-no-cache: # Build sportscrape-development docker image no cache
	docker compose -p sportscrape -f docker-compose.yml --project-directory . build --force-rm --no-cache sportscrape-local

enter: # spin up sportscrape-local-container docker container and shell into it
	docker compose -p sportscrape -f docker-compose.yml --project-directory . up -d
	docker exec -it sportscrape-local-container sh

down: # stop and kill sportscrape-local-container and sportscrape-local network
	docker compose -p sportscrape --project-directory . -f docker-compose.yml down

pre-commit-all: # Run pre-commit on all files
	pre-commit run --all-files

coverage-html: # Converts coverage.out to coverage.html
	go tool cover -html=coverage.out -o coverage.html

unit-tests: # Run unit tests
	go test -v -short -tags=unit -coverprofile=coverage.out ./...
	$(MAKE) coverage-html

all-tests: # Run all tests regardless of tags
	go test -v -tags="unit integration" -coverprofile=coverage.out ./...
	$(MAKE) coverage-html
