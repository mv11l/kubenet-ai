SRC=$(shell find . -name "*.go")

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

.PHONY: go_fmt go_lint go_test go_install_deps go_clean go_build go_run

default: all

all: go_fmt go_test

go_fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

go_lint:
	$(info ******************** running lint tools ********************)
	cd golang && golangci-lint run -v ./...

go_test:
	$(info ******************** running tests ********************)
	cd golang && go test -v ./...

go_install_deps:
	$(info ******************** downloading dependencies ********************)
	cd golang && go mod tidy

go_clean:
	cd golang && rm -rf bin/

go_build:
	cd golang && go build -o bin/kubenet-ai main.go

go_run:
	cd golang && ./bin/kubenet-ai

