
all: deps build

glide-install:
	go get github.com/Masterminds/glide

deps: glide-install
	glide install

build:
	go build .
