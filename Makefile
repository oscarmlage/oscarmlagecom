.DEFAULT_GOAL := help

MASTODON_ENV_FILE ?= utils/mastodon-sync/.env

.PHONY: all help build serve bash shell deploy sync up buildall clean-ds medialog-add medialog-add-deploy mastodon-sync

all: build ## Alias for build

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_.-]+:.*##/ {printf "  %-12s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the website
	docker compose -f docker-compose.yml up build

serve: ## Start local web server
	docker compose -f docker-compose.yml up serve

bash: ## Open shell in serve container
	docker compose -f docker-compose.yml exec serve sh

shell: ## Run shell in build container
	docker compose -f docker-compose.yml run --entrypoint sh build

deploy: up ## Alias for up deployment
sync: up ## Alias for up synchronization
up: build ## Build and sync to remote server
	rsync -e 'ssh -p 235' --progress --delete -lprtvvzog src/public/ root@151.80.35.190:/root/docker/docker-static-nginx-oscarmlage/_data/
	$(MAKE) mastodon-sync

mastodon-sync: ## Publish new posts/microposts to Mastodon when configured [ARGS="--from YYYY-MM-DD --dry-run"]
	@if [ -f "$(MASTODON_ENV_FILE)" ]; then \
		docker compose --env-file "$(MASTODON_ENV_FILE)" -f docker-compose.yml run --rm mastodon-sync go run . $(ARGS); \
	else \
		docker compose -f docker-compose.yml run --rm mastodon-sync go run . $(ARGS); \
	fi

buildall: build ## Build and copy to remote server
	cp -r src/public/ root@151.80.35.190:/root/docker/docker-static-nginx-oscarmlage/_data/

clean-ds: ## Remove .DS_Store files
	find . -name '.DS_Store' -type f -delete

medialog-add: ## Add/enrich/export a medialog item, then build [category title rate date]
	cd ../medialog && $(MAKE) medialog-add-hugo \
		category="$(category)" \
		title="$(title)" \
		rate="$(rate)" \
		date="$(date)" \
		released="$(released)" \
		progress="$(progress)" \
		episode="$(episode)" \
		thoughts="$(thoughts)" \
		link="$(link)" \
		no_fill="$(no_fill)" \
		no_export="$(no_export)" \
		overwrite_images="$(overwrite_images)" \
		force_new="$(force_new)"
	$(MAKE) build

medialog-add-deploy: medialog-add ## Add a medialog item, build and deploy
	$(MAKE) up
