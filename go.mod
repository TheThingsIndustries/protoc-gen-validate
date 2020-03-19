module github.com/envoyproxy/protoc-gen-validate

go 1.14

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.4.14-gogo.4

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.5
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/lyft/protoc-gen-star v0.4.14
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/tools v0.0.0-20200318150045-ba25ddc85566
)
