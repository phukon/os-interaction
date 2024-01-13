package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func machineHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	if cmd == "" {
		http.Error(w, "Missing 'cmd' parameter", http.StatusBadRequest)
		return
	}

	args := strings.Fields(cmd)

	// This unpacks the slice elements
	output, err := exec.Command(args[0], args[1:]...).
		CombinedOutput()
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Error executing command: %s", err),
			http.StatusInternalServerError,
		)
		return
	}

	// Send the command output as the response
	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}

func main() {
	http.HandleFunc("/machine", machineHandler)

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
