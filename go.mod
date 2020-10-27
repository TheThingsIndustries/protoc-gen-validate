module github.com/envoyproxy/protoc-gen-validate

go 1.15

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.5.1-gogo.2

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.5
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/iancoleman/strcase v0.1.2
	github.com/spf13/afero v1.3.3 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/net v0.0.0-20201027133719-8eef5233e2a1
	golang.org/x/tools v0.0.0-20201026223136-e84cfc6dd5ca
	google.golang.org/grpc v1.27.0 // indirect
)
