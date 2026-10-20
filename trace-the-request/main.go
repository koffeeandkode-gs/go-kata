package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	requestID, ok := r.Context().Value(requestIDKey).(string)
	if ok {
		log.Printf("handler request ID: %s", requestID)
	}
	fmt.Fprintln(w, "hello world")
}

type contextKey string

const requestIDKey contextKey = "request-id"

// 16-byte random value and returns a 32-character hexadecimal ID
func generateRequestID() (string, error) {

	bytes := make([]byte, 16)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil

}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request-ID logic
		requestID, err := generateRequestID()
		if err != nil {
			http.Error(w, "could not generate ID", http.StatusInternalServerError)
			return
		}
		log.Printf("requestID %s", requestID)
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		r = r.WithContext(ctx)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", requestIDMiddleware(mux)))

}
