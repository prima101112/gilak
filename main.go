package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {
	type Options struct {
		IsServer bool `short:"s" long:"is-server" description:"define as server or generator"`
	}
	var argsval Options
	args, err := flags.ParseArgs(&argsval, os.Args)
	if err != nil {
		log.Printf("failed parse args %s", err)
	}

	log.Printf("args from parse %s", args)

	log.Println(argsval.IsServer)

	if argsval.IsServer {
		http.HandleFunc("/", Root)
		http.HandleFunc("/ping", Pong)
		http.HandleFunc("/checkfile", CheckFile)

		log.Println("server running")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		//generate here
		envex := os.Getenv("EXAMPLE_ENV")
		log.Println("============================")
		log.Println(envex)
		f, err := os.Create("/tmp/generatefromenv.html")
		if err != nil {
			log.Printf("failed create file %s", err)
		}
		defer f.Close()

		_, err = f.WriteString("writes\n")
		_, err = f.WriteString("<h1>" + envex + "</h1>\n")
		f.Sync()
	}

}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to gilak</h1>")
}

func CheckFile(w http.ResponseWriter, r *http.Request) {
	dat, err := os.ReadFile("/tmp/generatefromenv.html")
	if err != nil {
		fmt.Fprintf(w, "no file")
	}
	fmt.Fprintf(w, string(dat))
}
