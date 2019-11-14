module github.com/envoyproxy/protoc-gen-validate

go 1.13

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.4.12-gogo

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/lyft/protoc-gen-star v0.4.12
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de
	golang.org/x/tools v0.0.0-20191114200427-caa0b0f7d508 // indirect
)
