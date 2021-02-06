.SILENT:

build:
	docker build --tag echopod:dev .

run: build stop
	docker run --rm -d -p 80:8080 --name echopod-dev echopod:dev

stop:
	docker stop echopod-dev || true
