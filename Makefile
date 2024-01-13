docker-dev = docker-compose -f docker-compose-dev.yaml

#Start docker
run:
	$(docker-dev) up $(args)

#Stop docker
stop:
	$(docker-dev) down

#Clean rebuild
rebuild:
	$(docker-dev) down --volumes && \
	$(docker-dev) rm --all && \
	$(docker-dev) pull --ignore-buildable && \
	$(docker-dev) build --no-cache && \
	$(docker-dev) up --force-recreate

#Exec into main app
exec:
	docker exec -it gopi-dev sh
