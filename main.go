package main

import (
	"context"
	"fmt"
	"os"

	"github.com/superfly/flyctl/flaps"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.TODO()
	appName := os.Args[1]
	flapsClient, err := flaps.NewFromAppName(ctx, appName)
	if err != nil {
		return err
	}
	machines, err := flapsClient.ListActive(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%d machines in %s app\n", len(machines), appName)
	return nil
}
