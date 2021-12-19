package main

import (
	"flag"
	"fmt"
	"log4j-scanner/log4j_folder_scan"
	"log4j-scanner/log4j_url_scan"
	"os"
)

func main() {

	urlCheckCmd := flag.NewFlagSet("url", flag.ExitOnError)
	hfName := urlCheckCmd.String("hf", "", "hf")
	urName := urlCheckCmd.String("ur", "", "ur")
	plName := urlCheckCmd.String("pl", "", "pl")
	hnName := urlCheckCmd.String("hn", "", "hn")

	internalCheckCmd := flag.NewFlagSet("internal", flag.ExitOnError)
	include_zip := urlCheckCmd.Bool("include-zip", false, "include-zip")
	mode := urlCheckCmd.String("mode", "", "mode")
	server := urlCheckCmd.String("server", "", "server")
	online := urlCheckCmd.Bool("online", false, "online")

	// internalCheckCmd := flag.NewFlagSet("internal", flag.ExitOnError)

	// externalCheckCmd := flag.NewFlagSet("external", flag.ExitOnError)
	if len(os.Args) < 2 {
        fmt.Println("expected 'url' or 'internal' or 'external' subcommands")
        os.Exit(1)
    }

	switch os.Args[1] {
		case "url":
			urlCheckCmd.Parse(os.Args[2:])
			url_scan.Url_Execute(*hfName,*plName,*urName,*hnName)
			//fmt.Println("  tail:", urlCheckCmd.Args())
		case "internal":
			internalCheckCmd.Parse(os.Args[2:])
			// fmt.Println("subcommand 'internal'")
			// fmt.Println("  include_zip:", *include_zip)
			// fmt.Println("  mode:", *mode)
			// fmt.Println("  server:", *server)
			// fmt.Println("  online:", *online)
			//fmt.Println("  tail:", internalCheckCmd.Args())

			internal_scan.Internal_Execute(*include_zip, mode, server, online)
			
		default:
			fmt.Println("expected 'url' or 'internal' subcommands")
			os.Exit(1)
		}
}
