package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const runUrl = "http://golang.org/compile?output=json"

func init() {
	http.HandleFunc("/compile", compileHandler)
}

func compileHandler(res http.ResponseWriter, req *http.Request) {
	if err := compile(res, req); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Printf("compile error: %q", err)
		fmt.Fprintf(res, "compile error: %q", err)
	}
}

func compile(w io.Writer, req *http.Request) error {
	client := http.Client{}
	defer req.Body.Close()
	r, err := client.Post(runUrl, req.Header.Get("Content-type"), req.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if _, err := io.Copy(w, r.Body); err != nil {
		return err
	}
	return nil
}
