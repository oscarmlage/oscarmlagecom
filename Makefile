all: build

build:
	docker compose -f docker-compose.yml up build

serve:
	docker compose -f docker-compose.yml up serve

bash:
	docker compose -f docker-compose.yml exec serve bash

shell:
	docker compose -f docker-compose.yml run build shell

deploy: up
sync: up
up: build
	rsync -e 'ssh -p 235' --progress --delete -lprtvvzog src/public/ root@151.80.35.190:/root/docker/docker-static-nginx-oscarmlage/_data/

buildall: build
	cp -r src/public/ root@151.80.35.190:/root/docker/docker-static-nginx-oscarmlage/_data/

