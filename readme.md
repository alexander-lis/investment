# General
About this app

# Tips
## Development
### Protobuf compilation

In VS Code compile ```.proto``` files using **vscode-proto3** extension with ```.vscode/settings.json```.

Or manually from the workspace root:
```bash
# Protobuf compile example.
protoc ./src/app/services/user/proto/authentication.proto \
   --go_out ./src/app/shared/protobuf \
   --go-grpc_out ./src/app/shared/protobuf \
   --grpc-gateway_out ./src/app/shared/protobuf \
```

TODO: добавить конфиг для https://github.com/bufbuild/buf.

## Local deployment
### Docker build
### Kubernetes
### Helm charts
### Database