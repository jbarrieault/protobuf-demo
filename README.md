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
`protoc --go_out=pkg --ruby_out=./ruby user.proto`

That generates go code in `./pkg/user.pb.go`, and ruby code in `./ruby/user_pb.rb`

### Passing protobuf data from Go to Ruby

You can experiment schema evolution using the included `/cmd/user_client.go` and `ruby/user_server.rb`,
which sends `user` messages over a socket.

Schema evolution can be performed by using `user.proto` for one program, and `user-v2.proto` for the other.

To test forward compatibility (writing with new schema, reading with old):
`protoc --go_out=pkg user-v2.proto`
`protoc --ruby_out=./ruby user.proto`

Conversely, backward compatibility  (writing with old schema, reading with new):
`protoc --go_out=pkg user.proto`
`protoc --ruby_out=./ruby user-v2.proto`
