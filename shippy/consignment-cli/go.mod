module github.com/caseyfu/go/shippy/consignment-cli

go 1.14

require (
	github.com/caseyfu/go/shippy/consignment-service v0.0.0-20210907163652-5faf8a310d9a
	google.golang.org/grpc v1.40.0
)

replace github.com/caseyfu/go/shippy/consignment-service => ../consignment-service
