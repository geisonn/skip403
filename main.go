package main

import (
	"flag"
	"fmt"

	"github.com/geisonn/skip403/functions"
)

func main() {
	help := flag.Bool("help", false, "Display this help message")
	flag.BoolVar(help, "h", false, "Display this help message (shorthand)")

	url := flag.String("url", "", "Single target URL (e.g., https://example.com)")
	flag.StringVar(url, "u", "", "Single target URL (shorthand)")

	list := flag.String("list", "", "File containing a list of target URLs")
	flag.StringVar(list, "l", "", "List of target URLs (shorthand)")

	hideFails := flag.Bool("hide-fails", false, "Only show successful bypass attempts")
	flag.BoolVar(hideFails, "hf", false, "Only show successful bypass attempts (shorthand)")
	flag.Parse()

	fmt.Println("\033[32m" + `     _    _       _  _    ___ _____ 
 ___| | _(_)_ __ | || |  / _ \___ / 
/ __| |/ / | '_ \| || |_| | | ||_ \ 
\__ \   <| | |_) |__   _| |_| |__) |
|___/_|\_\_| .__/   |_|  \___/____/ 
           |_|                      

Created by: Geison Nunes | Version: 1.0
` + "\033[0m")

	if *help {
		functions.Help()
	}

	if *url != "" {
		functions.BypassUrl(*url, *hideFails)
	}

	if *list != "" {
		functions.BypassList(*list, *hideFails)
	}
}
