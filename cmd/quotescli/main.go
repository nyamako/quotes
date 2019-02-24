package main

import (
	"os"
	"flag"
	"fmt"
	"github.com/nyamako/quotes"
)

var dbcreds = flag.String("c", "", "database credentials")

func die(err error){
	fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
	os.Exit(1)
}

func usage(){
	//init
	//insert author text
	//get [id]

	usage := `usage: %s -c=user:password@/database cmd
	random
	add text author
`
	fmt.Fprintf(os.Stderr, usage, os.Args[0])
	os.Exit(1)
}

func main() {
	flag.Parse()

	if *dbcreds == "" {
		usage()
	}

	if flag.NArg() < 1 {
		usage()
	}

	qdb, err := quotes.New(*dbcreds)
	if err != nil {
		die(err)
	}
	cmd := flag.Arg(0)
	switch cmd {
	case "random":
		q, err := qdb.Random()
		if err != nil {
			die(err)
		}

		fmt.Println(q.Text)
		fmt.Println("        --", q.Author)

	case "add":
		if flag.NArg() != 3 {
			usage()
		}

		text := flag.Arg(1)
		author := flag.Arg(2)

		q := &quotes.Quote{
			Author: author,
			Text: text,
		}

		if err := qdb.Add(q); err != nil {
			die(err)
		}
	default:
		usage()
	}
}

