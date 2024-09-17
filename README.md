# protobuf-demo

Trying out protobuf, that's it.

## Setup

Install the protoc and runtime libraries:

`brew install protobuf`

Install the Go protobuf codegen plugin:

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

Ensure your go package bin is in your $PATH, to `protoc` can call it:

`export PATH=$PATH:$(go env GOPATH)/bin`

Compile user.proto file:

`protoc --go_out=pkg --ruby_out=./ruby user.proto`

That generates go code in `./pkg/user.pb.go`, and ruby code in `./ruby/user_pb.rb`

Do the same for the v2 schema:

`protoc --go_out=pkg --ruby_out=./ruby user_v2.proto`

## Passing protobuf data from Go to Ruby

You can experiment schema evolution using the included `/cmd/[v1|v2]/write-user.go` and `ruby/read_user_[v1|v2].rb`,
which sends/receives `user` messages over a socket.

Schema evolution can be performed by using mismatching version between the go and ruby programs.

To test backward compatibility (decoding messages encoded using older schema):

`ruby read_user_v2.rb`
`go run cmd/v1/write_user.go`

Conversely, for forward compatibility (decoding message encoded in newer schema):

`ruby read_user_v1.rb`
`go run cmd/v2/write_user.go`

You can observe how certain field changes such as renames are both backwards and forwards compatible! Additionally, added/removed fields are quietly ignored on the read side.

Some changes, however, have unexpected results.

Changing `email` from `string` to a custom `Email` type appears to _almost_ backwards compatible—decoding via the v1 schema can parse without error, but the resulting value contains leading characters `\n\u0016` (🤔). The change isn't forward comptaible whatsoever—decoding via the v2 schema raises a parse exception. The behavior seems exactly opposite from what I expected.
