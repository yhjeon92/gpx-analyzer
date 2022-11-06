# gpx-analyzer

- Build with grpc stub

```shell
$ protoc -I . --go_out=gpx/. --go-grpc_out=gpx/. gpx/gpx.proto
```