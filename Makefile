github-test:
	go run ./cmd/init
	go test ./... --tags=github

test:
	go fmt ./...
	go test ./... --cover --count=1

build:
	yarn --cwd ui
	go-assets-builder ui/public ui/templates -o internal/infra/delivery/ui.go -p delivery -v ui
	go build -o daanretard ./cmd/daanretard/main.go