all: build

build:
	docker-compose -f docker-compose.yml up build

serve:
	docker-compose -f docker-compose.yml up serve

bash:
	docker-compose -f docker-compose.yml exec serve bash

shell:
	docker-compose -f docker-compose.yml run build shell

sync: up
up: build
	rsync -e ssh --progress --delete -lprtvvzog src/public/ me:/home/www/oscarmlage.com/www/

buildall: build
	cp -r src/public/ me:/home/www/oscarmlage.com/www

