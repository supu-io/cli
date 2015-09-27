build:
		go build
deps:
	go get -u github.com/fatih/color
	go get -u github.com/codegangsta/cli
dev-deps:
	go get -u github.com/golang/lint/golint
lint:
	golint
