module github.com/viqueen/buf-template/backend

go 1.23.4

require (
	connectrpc.com/connect v1.18.1
	connectrpc.com/cors v0.1.0
	connectrpc.com/otelconnect v0.7.2
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/lib/pq v1.10.9
	github.com/rs/cors v1.11.1
	github.com/viqueen/buf-template/api/go-sdk v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.41.0
)

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	go.opentelemetry.io/otel/metric v1.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.37.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/viqueen/buf-template/api/go-sdk => ../api/go-sdk
