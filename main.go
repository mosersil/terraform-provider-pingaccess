package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	tf5server "github.com/hashicorp/terraform-plugin-go/tfprotov5/server"
	tfmux "github.com/hashicorp/terraform-plugin-mux"
	protocol "github.com/iwarapter/terraform-provider-pingaccess/internal/protocolprovider"
	sdkv2 "github.com/iwarapter/terraform-provider-pingaccess/internal/sdkv2provider"
)

func main() {
	ctx := context.Background()
	sdkv2 := sdkv2.Provider().GRPCProvider
	factory, err := tfmux.NewSchemaServerFactory(ctx, sdkv2, protocol.Server)
	if err != nil {
		panic(err)
	}
	err = tf5server.Serve("registry.terraform.io/iwarapter/pingaccess", func() tfprotov5.ProviderServer {
		return factory.Server()
	})
	if err != nil {
		panic(err)
	}
}
