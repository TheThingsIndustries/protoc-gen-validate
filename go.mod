module github.com/envoyproxy/protoc-gen-validate

go 1.14

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.4.15-gogo

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/lyft/protoc-gen-star v0.4.15
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/tools v0.0.0-20200519015757-0d0afa43d58a
)
