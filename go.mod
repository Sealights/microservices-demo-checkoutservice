module github.com/GoogleCloudPlatform/microservices-demo/src/checkoutservice

go 1.17

require (
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/google/uuid v1.3.0
	github.com/sirupsen/logrus v1.8.1
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.8.0
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
)


require (
	cloud.google.com/go/compute v0.1.0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	go.opentelemetry.io/proto/otlp v0.18.0 // indirect
	golang.org/x/net v0.0.0-20220121210141-e204ce36a2ba // indirect
)

require (
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.29.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.8.0 // indirect
	go.opentelemetry.io/otel/sdk v1.8.0
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5 // indirect
)
