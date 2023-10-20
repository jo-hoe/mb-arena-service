.DEFAULT_GOAL := all

.PHONY: all
all: create helm-install

.PHONY: restart
restart: clean all

.PHONY: create
create:
	@k3d cluster create --port '8080:80@loadbalancer'

.PHONY: clean
clean:
	@k3d cluster delete

.PHONY: docker-run
docker-run:
	@docker-compose up --build

.PHONY: docker-build
docker-build:
	@docker-compose build

.PHONY: helm-test
helm-test:
	@helm test mb-arena-schedule-api

.PHONY: helm-install
helm-install:
	@helm install  mb-arena-schedule-api ./charts/mb-arena-schedule-api/