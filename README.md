# protobuf-demo

Trying out protobuf, that's it.

### Setup

Install the protoc and runtime libraries:
`brew install protobuf`

Install the Go protobuf codegen plugin:
`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

Ensure your go package bin is in your $PATH, to `protoc` can call it:

`export PATH=$PATH:$(go env GOPATH)/bin`

### Development

Compile user.proto file:
`protoc --go_out=pkg user.proto`

That generates go implementation code in `./pkg/user.pb.go`, which can be imported within the project like so:
`import "github.com/jbarrieault/protobuf-demo/pkg/user"`
