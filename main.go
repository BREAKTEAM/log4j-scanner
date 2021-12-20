package main

import (
	"flag"
	"fmt"
	"log4j-scanner/log4j_folder_scan"
	"log4j-scanner/log4j_folder_scan/server"
	"log4j-scanner/log4j_url_scan"
	"os"
)

func main() {

	urlCheckCmd := flag.NewFlagSet("url", flag.ExitOnError)
	hfName := urlCheckCmd.String("hf", "", "path to list header file")
	urName := urlCheckCmd.String("ur", "", "path to list url file")
	plName := urlCheckCmd.String("pl", "", "path to list payload file")
	hnName := urlCheckCmd.String("hn", "", "your hostname to connect to (dnslog.cn/burp collabarator...)")

	internalCheckCmd := flag.NewFlagSet("internal", flag.ExitOnError)
	include_zip := internalCheckCmd.Bool("include-zip", false, "include zip files in the scan")
	mode := internalCheckCmd.String("mode", "", "the output mode, either 'report' (every java archive pretty printed) or 'list' (list of potentially vulnerable files) (default 'report')")
	server := internalCheckCmd.String("server", "", "server to listen result return")
	online := internalCheckCmd.Bool("online", false, "go with server opton")

	externalCheckCmd := flag.NewFlagSet("external", flag.ExitOnError)
	port := externalCheckCmd.String("port", "8080", "port to listen result from client")

	// externalCheckCmd := flag.NewFlagSet("external", flag.ExitOnError)
	if len(os.Args) < 2 {
        fmt.Println("expected 'url' or 'internal' or 'external' subcommands")
        os.Exit(1)
    }

	switch os.Args[1] {
		case "url":
			urlCheckCmd.Parse(os.Args[2:])
			if *hfName == "" || *plName == "" || *urName == "" || *hnName == "" {
				fmt.Printf("Usage : go run . url -hf headers.txt -pl payloads.txt -ur urls.txt -hn xxx.burpcollaborator.net")
			}else{
				url_scan.Url_Execute(*hfName,*plName,*urName,*hnName)
			}

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
		case "external":
			externalCheckCmd.Parse(os.Args[2:])
			// fmt.Println("subcommand 'internal'")
			// fmt.Println("  include_zip:", *include_zip)
			// fmt.Println("  mode:", *mode)
			// fmt.Println("  server:", *server)
			// fmt.Println("  online:", *online)
			//fmt.Println("  tail:", internalCheckCmd.Args())

			external_scan.External_Execute(*port)
			
		default:
			fmt.Println("expected 'url' or 'internal' or 'external' subcommands")
			os.Exit(1)
		}
}
