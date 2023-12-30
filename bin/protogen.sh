#!/bin/bash

# Generate all APIs contracts.
echo "protoc user->uthentication."
protoc ./app/protobuf/user/authentication/authentication.proto \
   --go_out . --go_opt paths=source_relative \
   --go-grpc_out . --go-grpc_opt paths=source_relative \
   --grpc-gateway_out . --grpc-gateway_opt paths=source_relative \