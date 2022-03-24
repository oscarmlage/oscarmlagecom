all: build

build:
	docker-compose -f docker-compose.yml up build

serve:
	docker-compose -f docker-compose.yml up serve

bash:
	docker-compose -f docker-compose.yml exec serve bash

shell:
	docker-compose -f docker-compose.yml run build shell
