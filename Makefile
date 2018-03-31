default: build

BRANCH := ${shell git rev-parse --abbrev-ref HEAD}
HASH := ${shell git rev-parse HEAD}


build: main.go
	go build -ldflags "-X simple_http_server.version.GitBranch=${BRANCH} -X simple_http_server.version.GitHash=${HASH}"
