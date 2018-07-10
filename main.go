package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	cd, _ := os.Getwd()
	port := ":8080"
	url := strings.Join([]string{"http://localhost", port, "/"}, "")

	fmt.Println("Ctrl + C キーで終了します。")
	open.Run(url)

	http.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir(cd+"/web/"))))
	log.Fatal(http.ListenAndServe(port, nil))
}
