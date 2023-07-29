package main

import (
	pb "grpc-rest-api/protoc"
	"io"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sendRequest(w http.ResponseWriter, r *http.Request) {
    apiKey := r.Header.Get("apikey")
    if (apiKey == "") {
        io.WriteString(w, "invalid\n")
        return
    }

    conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials())) 
    if err != nil {
        io.WriteString(w, "no grpc server\n")
    }
    defer conn.Close()
    //client := pb.NewEcho()

    path := r.URL.Path
    switch path {
    case "/send_one":
        io.WriteString(w, "one\n")
    case "/send_two":
        io.WriteString(w, "two\n")
    default:
        io.WriteString(w, "default\n")
    }
}

func main() {
    http.HandleFunc("/", sendRequest)

    http.ListenAndServe(":3000", nil)
}
