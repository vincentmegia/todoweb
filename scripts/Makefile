run:
	@air

compile:
	@templ generate
	@go build ../cmd/todo -ldflags="-X github.com/vincentmegia/go-todo-web/build.Environment=development" -X "github.com/vincentmegia/go-todo-web/build.Version=0.0.1"
