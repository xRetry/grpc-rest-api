package main

import (
	"context"
	"encoding/json"
	pb "grpc-rest-api/proto"
	"io"
	"io/ioutil"
	"net/http"
	"time"
    "protojson"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sendRequest(w http.ResponseWriter, r *http.Request) {
    apiKey := r.Header.Get("apikey")
    if (apiKey == "") {
        io.WriteString(w, "invalid\n")
        return
    }

    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        io.WriteString(w, "invalid body\n")
        return
    }

    conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials())) 
    if err != nil {
        io.WriteString(w, "no grpc server\n")
        return
    }
    defer conn.Close()
    client := pb.NewEchoClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    path := r.URL.Path
    switch path {
    case "/send_one":
        var req pb.OneRequest
        err := json.Unmarshal(reqBody, &req)
        if err != nil {
            io.WriteString(w, "invalid body\n")
            return
        }

        resp, err := client.SendOne(ctx, &req)
        json.Marshal
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
