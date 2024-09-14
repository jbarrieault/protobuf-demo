require 'base64'
require 'google/protobuf'


require_relative 'user_pb'

user = User::User.new(first_name: "Annie", last_name: "Barrieault")
puts "user: #{user}"

data = User::User.encode(user)
puts "encoded: #{Base64.encode64(data)}"

u2 = User::User.decode(data)
puts "decoded: #{u2}"
