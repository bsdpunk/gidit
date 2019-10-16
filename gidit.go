package main

import "fmt"

import gidit "github.com/bsdpunk/gidit/lib"

//import "./lib"
import "os"

func main() {
	switch oper := os.Args[1]; oper {
	case "cat":
		gidit.Gicat(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
		os.Exit(0)
	case "append":
		gidit.Giappend(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
		os.Exit(0)
	default:
		// freebsd, openbsd,
		// plan9, windows...
		//fmt.Printf("%s.\n", os)
		if len(os.Args) > 1 {
			if len(os.Args) > 2 {
				gidit.Gisize(os.Args[len(os.Args)-1], os.Args[1], os.Args[2], "out.png")
			} else {

				gidit.Gisize(os.Args[len(os.Args)-1], os.Args[1], "", "out.png")
			}
		} else {
			fmt.Println("Must Specify a png file and a size to scale to, or specify different operation")
			os.Exit(1)
		}
	}
}
