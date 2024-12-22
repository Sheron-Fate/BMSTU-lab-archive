package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/server/calculator"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error"`
}

// Handler
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func calculate(w http.ResponseWriter, r *http.Request) {
	var req Request
	var res Response
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		if errors.Is(err, errors.New("expression is not valid")) {
			res = Response{0, "expression is not valid"}
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			res = Response{0, "internal server error"}
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		res = Response{result, "nil"}
	}
	
	w.Header().Set("Content-Type", "application/json")
	jRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jRes)

}

// Middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("User %s Hit Endpoint", r.FormValue("user"))
		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", loggingMiddleware(getRoot))
	mux.HandleFunc("/hello", loggingMiddleware(getHello))
	mux.HandleFunc("POST /calculate", loggingMiddleware(calculate))

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
