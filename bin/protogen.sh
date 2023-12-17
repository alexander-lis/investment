#!/bin/bash

# Generate all APIs contracts.
echo "protoc user->uthentication."
protoc ./services/user/api/authentication/authentication.proto \
   --go_out . --go_opt paths=source_relative \
   --go-grpc_out . --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./gateway/api \