.PHONY: buid app client swagger test controller dev local-env \
	local-env-down

PROJECT_DIR=github.com/luanngominh/goa-example
DESIGN_DIR=${PROJECT_DIR}/goa/design

app:
	goagen app -d ${DESIGN_DIR}

client:
	goagen client -d ${DESIGN_DIR}

swagger:
	goagen swagger -d ${DESIGN_DIR}

controller:
	goagen controller -d ${DESIGN_DIR} -o ext/controller

goa: app client swagger controller

build:
	@mkdir -p bin
	go build -o bin/meocon	ext/cmd/goa-example/main.go

dev: build
	@ENV=DEV ./bin/meocon

test:
	go test ./...

local-env:
	docker-compose up -d

local-env-down:
	docker-compose down
