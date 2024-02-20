ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=set
else
    SET=export
endif

build:
	go fmt
	go build

test:
	go test -v

readme:
	$(SET) "GOEXPERIMENT=rangefunc" && example-into-readme

.PHONY: build test readme
