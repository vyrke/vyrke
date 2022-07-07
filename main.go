package main

import (
	"flag"
)

func main() {

	var h bool
	flag.BoolVar(&h, "h", false, "display help")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 && !h {
		deploy()
	}

	if h {
		help()
	} else {
		for _, arg := range args {

			if arg == "deploy" {
				deploy()
			}
	
		}
	}

}