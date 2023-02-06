package main

import (
        "context"
        "fmt"
        "log"

        "hello-plugin/proto"
        "google.golang.org/grpc"
)

func main() {
        conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
                log.Fatal(err)
        }
        client := proto.NewHelloPluginClient(conn)

        res, err := client.GetPerson(context.TODO(), &proto.RequestPerson{
                Name: "Jackson",
        })
        fmt.Println(res.Name, err)
}
