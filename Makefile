default: build

BRANCH := ${shell git rev-parse --abbrev-ref HEAD}
HASH := ${shell git rev-parse HEAD}
IMPORT_PATH = github.com/admarcel/go_http_server/
LDFLAGS = \
	-X '${IMPORT_PATH}version.GitBranch=${BRANCH}' \
	-X '${IMPORT_PATH}version.GitHash=${HASH}'


build: main.go
	go build -ldflags "$(LDFLAGS)"
