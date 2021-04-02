package main

import (
        "io"
        "log"
        "net/http"
		"io/ioutil"
		"os"
)
func main() {
        // Set routing rules
        http.HandleFunc("/", home)

        //Use the default DefaultServeMux.
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal(err)
        }
}

func home(w http.ResponseWriter, r *http.Request) {
	programName := os.Args[1]
	io.WriteString(w, readFile(programName))
}

func readFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
    if err != nil {
        return "500 ERROR."
    }
    
	return string(data)
}
