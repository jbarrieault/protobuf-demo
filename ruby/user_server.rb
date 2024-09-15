require 'socket'
require 'base64'
require 'google/protobuf'

require_relative 'user_pb'

SOCKET_PATH = "/tmp/protobuf-demo-socket.sock"

File.delete(SOCKET_PATH) if File.exist?(SOCKET_PATH)

socket = Socket.new(:UNIX, :STREAM)
addr = Socket.pack_sockaddr_un(SOCKET_PATH)
socket.bind(addr)
socket.listen(5)
puts "listening for (base64 encoded) protobuf user data on socket #{SOCKET_PATH}"

loop do
    client, _ = socket.accept
    puts "client connected"

    data = client.recv(1024)
    raw = Base64.decode64(data)
    user = User::User.decode(raw)
    puts "decoded user: #{user}"

    client.send("thank you", 0)
    client.close
end
