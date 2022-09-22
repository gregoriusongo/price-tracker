BINARY=price-tracker
test: 
	go test -v -cover -covermode=atomic ./...

app:
	go build -o ${BINARY} cmd/price_tracker/*.go

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
	docker run -d -p 9222:9222 -m 800m --rm --name headless-shell --init chromedp/headless-shell

docker-run:
	docker run -d --rm -v ./config.json:/app/config.json --name price-tracker-shell --init price-tracker:latest

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint