package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<html>
		<head><title>Welcome to GDG!</title></head>
		<body style="font-family: sans-serif; text-align: center; margin-top: 100px;">
			<h1>Hello, GDG Soongsil 👋</h1>
			<p>You're running a Go app on Kubernetes ☸️ via GitOps ⚙️ with Argo CD 🚀</p>
			<p style="color: gray;">%s</p>
		</body>
		</html>
	`, r.UserAgent())
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Printf("🌐 Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}