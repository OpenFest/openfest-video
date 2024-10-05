package main

import (
	"fmt"
	"os"

	"git.openfest.org/Video/openfest-video/break_player/mainservice"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "usage: %s <yaml config filename>\n", os.Args[0])
		os.Exit(1)
	}

	svc, err := mainservice.New(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not init main service: %s\n", err)
		os.Exit(1)
	}

	err = svc.Serve()
	if err != nil {
		fmt.Fprintf(os.Stdout, "server died: %s\n", err)
		os.Exit(1)
	}
}
