module github.com/envoyproxy/protoc-gen-validate/tests

go 1.15

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.5.2-gogo.1

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.5

require (
	github.com/envoyproxy/protoc-gen-validate v0.0.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.5
	golang.org/x/net v0.0.0-20201027133719-8eef5233e2a1
)

replace github.com/envoyproxy/protoc-gen-validate => ../
