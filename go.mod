module github.com/envoyproxy/protoc-gen-validate

go 1.15

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.5.2-gogo.1

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.5

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.5
	github.com/iancoleman/strcase v0.1.2
	github.com/lyft/protoc-gen-star v0.5.2
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/net v0.0.0-20201027133719-8eef5233e2a1
	golang.org/x/tools v0.0.0-20201026223136-e84cfc6dd5ca
)
