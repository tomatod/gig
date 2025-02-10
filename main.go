package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	cfg := NewConfig()
	ctx := context.TODO()
	cli := GetCLI(ctx, cfg.SetConfigFromCLI)
	if err := cli.Run(ctx, os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if cli.Bool("help") {
		os.Exit(0)
	}
	if err := Execute(cfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
