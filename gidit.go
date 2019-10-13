package main

import "fmt"
import "github.com/alexflint/go-arg"
import gidit "github.com/bsdpunk/gidit/lib"
import "os"

func main() {
	var args struct {
		Operation  string `arg:"-o,separate","help:"resize,cat or append"`
		OutputFile string `arg:"-f,separate","help:"output file"`
		Width      string `arg:"-w,separate","help:"width"`
		Height     string `arg:"-h,separate","help:"width"`
		InputFile  string `arg:"-i,separate","help:"input file"`
	}
	args.OutputFile = ""
	args.InputFile = ""
	args.Height = ""
	args.Width = ""
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
		if args.InputFile != "" && args.Width != "" {

			gidit.Gisize(args.InputFile, args.Width, args.Height, args.OutputFile)
		} else {
			fmt.Println("Must Specify a png file and a size to scale to, or specify different operation")
			os.Exit(1)
		}
	}
}
