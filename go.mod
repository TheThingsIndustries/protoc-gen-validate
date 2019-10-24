module github.com/envoyproxy/protoc-gen-validate

go 1.13

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.4.11-gogo.3

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/lyft/protoc-gen-star v0.4.11
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de
)
