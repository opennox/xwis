package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/opennox/xwis"
)

var (
	Root = &cobra.Command{
		Use:   "xwis",
		Short: "Nox XWIS tools",
	}
	fRootHost = Root.PersistentFlags().String("host", xwis.DefaultAddress, "lobby server address")
	fRootName = Root.PersistentFlags().String("login", "", "user login to use")
	fRootPass = Root.PersistentFlags().String("pass", "", "user password to use")
)

func newClient(ctx context.Context) (*xwis.Client, error) {
	return xwis.NewClientWithAddress(ctx, *fRootHost, *fRootName, *fRootPass)
}

func main() {
	if err := Root.Execute(); err != nil {
		os.Exit(1)
	}
}
