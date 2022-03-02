# envoy-grpc-json-example
envoy-grpc-json-example

example for envoy grpc-json transcode

# Usage

- `make run` to run services
- `sh verify.sh` to check

# Tech Stack

- [envoy](https://github.com/envoyproxy/envoy)
- [grpc-go](https://github.com/grpc/grpc-go)
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [buf](https://github.com/bufbuild/buf)
- [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)

# Issues

### Route config not right

[Envoy doc](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter#route-configs-for-transcoded-requests)
say path will be `<package>.<service>/<method>` , but it not act like that.

config already set `match_incoming_request_route=false`

For example:
`Status` only match `/status`, do not match `/helloworld.Greeter/Status`
```
service Greeter {
  rpc Status(Empty) returns (StatusResponse) {
    option (google.api.http) = {
      get: "/status"
    };
  }
}
```

`SayHello` match `/helloworld.Greeter/Status` only because I set the path in `google.api.http`
```
service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {
    option (google.api.http) = {
      get: "/helloworld.Greeter/SayHello"
    };
  }
}
```


### config not working

[config doc](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/grpc_json_transcoder/v3/transcoder.proto#envoy-v3-api-msg-extensions-filters-http-grpc-json-transcoder-v3-grpcjsontranscoder)

- match_incoming_request_route
- auto_mapping

