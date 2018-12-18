REPO?=onuryilmaz/$(notdir $(CURDIR))
VERSION=$(shell git rev-parse --verify HEAD)


.PHONY: all

static-code-check:
	docker build --rm -f docker/Dockerfile.static-code-check .

unit-test:
	docker build --rm -f docker/Dockerfile.unit-test .

smoke-test:
	docker build -f docker/Dockerfile.smoke-test -t $(REPO)/smoke-test:$(VERSION) .
	docker run -d -p 5432:5432 --name postgres postgres
	docker run --rm --link postgres:postgres gesellix/wait-for postgres:5432
	docker run -e DATABASE="postgresql://postgres:postgres@postgres:5432/postgres?sslmode=disable" --link postgres  $(REPO)/smoke-test:$(VERSION)

	docker stop postgres
	docker rm postgres

integration-test:
	docker build -f docker/Dockerfile.integration-test -t $(REPO)/integration-test:$(VERSION) .

	docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=default --name mysql mysql
	docker run --rm --link mysql:mysql gesellix/wait-for mysql:3306 -t 30
	docker run -e "DATABASE=mysql://root:password@mysql:3306/default" --link mysql $(REPO)/integration-test:$(VERSION)
	docker stop mysql
	docker rm mysql

	docker run -d -p 5432:5432 --name postgres postgres
	docker run --rm --link postgres:postgres gesellix/wait-for postgres:5432
	docker run -e DATABASE="postgresql://postgres:postgres@postgres:5432/postgres?sslmode=disable" --link postgres  $(REPO)/integration-test:$(VERSION)
	docker stop postgres
	docker rm postgres

	docker run -d -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=Password!" -p 1433:1433 --name mssql mcr.microsoft.com/mssql/server:2017-latest
	docker run --rm --link mssql:mssql gesellix/wait-for mssql:1433
	docker run -e DATABASE="mssql://sa:Password!@mssql:1433" --link mssql  $(REPO)/integration-test:$(VERSION)
	docker stop mssql
	docker rm mssql

build:
	docker build --target builder --build-arg VERSION=$(VERSION) -t $(REPO)/build:$(VERSION) .

prod:
	docker build --target production --build-arg VERSION=$(VERSION) -t $(REPO)/production:$(VERSION) .
