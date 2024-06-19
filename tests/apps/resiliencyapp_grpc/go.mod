module github.com/dapr/dapr/tests/apps/resiliencyapp_grpc

go 1.22.3

require (
	github.com/dapr/dapr v1.7.4
	google.golang.org/grpc v1.64.0
	google.golang.org/grpc/examples v0.0.0-20230224211313-3775f633ce20
	google.golang.org/protobuf v1.33.0
)

require (
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
)

replace github.com/dapr/dapr => ../../../