package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	flagVersion = flag.Bool("v", false, "Show version infomation")
)

func main() {
	flag.Parse()

	if *flagVersion {
		flag.PrintDefaults()
		fmt.Print("version: 0.1")
		return
	}

	log.Print("WI is serving at port :3333")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}
