package main

import "fmt"
import "github.com/alexflint/go-arg"
import gidit "./lib"
import "os"

func main() {
	var args struct {
		Input     []string `arg:"positional"`
		Operation string   `arg:"-o,separate","help:"resize,cat or append"`
	}

	args.Operation = "resize"
	arg.MustParse(&args)
	switch oper := args.Operation; oper {
	case "cat":
		gidit.Gicat(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
	case "append":
		gidit.Giappend(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
	default:
		// freebsd, openbsd,
		// plan9, windows...
		//fmt.Printf("%s.\n", os)
		if len(os.Args) == 2 || args.Operation == "resize" {

			gidit.Gisize(os.Args[len(os.Args)-1], os.Args[1], "")
		} else if len(os.Args) == 3 || args.Operation == "resize" {
			gidit.Gisize(os.Args[len(os.Args)-1], os.Args[1], os.Args[2])
		} else {
			fmt.Println("Must Specify a png file and a size to scale to, or specify different operation")
		}
	}
}
