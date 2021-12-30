all: build

build:
	docker-compose -f docker-compose.yml up build

serve:
	docker-compose -f docker-compose.yml up serve

shell:
	docker-compose -f docker-compose.yml run build shell
