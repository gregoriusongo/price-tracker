BINARY=price-tracker
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} cmd/*/*.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t ${BINARY} .

run:
	docker-compose up --build -d

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

headless-shell:
	docker run -d -p 9222:9222 --rm --name headless-shell --init chromedp/headless-shell

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint