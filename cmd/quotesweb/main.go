package main

import (
	"flag"
	"net/http"
	"log"
	"fmt"

	"github.com/nyamako/quotes"
)

var dbcreds = flag.String("c", "", "database credentials")

func main() {
	flag.Parse()

	qdb, err := quotes.New(*dbcreds)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		q, err := qdb.Random()
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			log.Print(err)
			return
		}

		fmt.Fprintf(w, "%s\n-- %s\n", q.Text, q.Author)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

