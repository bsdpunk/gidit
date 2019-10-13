package main

import "fmt"
import "github.com/alexflint/go-arg"
import gidit "github.com/bsdpunk/gidit/lib"
import "os"

func main() {
	switch oper := args.Operation; oper {
	case "cat":
		gidit.Gicat(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
	case "append":
		gidit.Giappend(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
	default:
		// freebsd, openbsd,
		// plan9, windows...
		//fmt.Printf("%s.\n", os)
		if len(os.Args) > 1 {
			if len(os.Args) > 3 {
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
