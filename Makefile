build:
	go build .

image:
	docker build -t bmcculley/puffy:3.0.2 .

push:
	docker push bmcculley/puffy:3.0.2
