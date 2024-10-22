package main

// r=requests.post(
//     'http://localhost:9080/exec',
//     json={"command": "sdf"}
// )

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type CommandRequest struct {
	Command string `json:"command"`
}

type CommandResponse struct {
	Output string `json:"output"`
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

func executeCommand(command string) (string, int, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()

	statusCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			statusCode = exitErr.ExitCode()
		} else {
			statusCode = 1
		}
	} else {
		statusCode = 0
	}
	return string(output), statusCode, err
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Parse the JSON request body
	var req CommandRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Execute the command
	output, statusCode, err := executeCommand(req.Command)

	// Prepare the response
	var response CommandResponse
	if err != nil {
		response = CommandResponse{
			Output: output,
			Status: statusCode,
			Error:  err.Error(),
		}
	} else {
		response = CommandResponse{
			Output: output,
			Status: statusCode,
		}
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/exec", commandHandler)

	// health check
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// body message "ok"
		w.Write([]byte("ok"))
	})

	port := ":31020"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

