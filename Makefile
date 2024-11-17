ifneq (,$(wildcard ./.env))
    include .env
    export
endif

APP_NAME=multi-file-mailer

build:
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

test:
	go test ./...

docker-build:
	docker build -t $(APP_NAME) .

docker-run: docker-build
	docker run -p $(PORT)$(PORT) --env-file .env $(APP_NAME)

stop:
	docker stop $(APP_NAME)