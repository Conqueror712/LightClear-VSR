package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8000", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// handle file upload
		filePath := "/path/to/save/uploaded/file"
		cmd := exec.Command("python", "path/to/processing/script.py", filePath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Fprintf(w, "File processed successfully : %s", output)
	} else {
		fmt.Fprintf(w, "Invalid request method")
	}
}
