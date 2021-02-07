.SILENT:

docker_tag = orginux/echopod:dev
container_name = echopod-dev

build:
	docker build --tag $(docker_tag) .

run: build stop
	docker run --rm -d -p 80:8080 --name $(container_name) $(docker_tag) \
		&& curl localhost:80

stop:
	docker stop $(container_name) || true

push:
	docker push $(docker_tag)
