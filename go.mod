module github.com/envoyproxy/protoc-gen-validate

go 1.13

replace github.com/lyft/protoc-gen-star => github.com/TheThingsIndustries/protoc-gen-star v0.4.14-gogo.3

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.3
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/lyft/protoc-gen-star v0.4.14
	golang.org/x/lint v0.0.0-20200130185559-910be7a94367
	golang.org/x/tools v0.0.0-20200221224223-e1da425f72fd
)
