package main

import (
	pb2 "URLShortenerGRPC/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("input url")
	var url string
	_, err := fmt.Scanf("%s\n", &url)
	if err != nil {
		return
	}
	connection, err := grpc.Dial(":9000", grpc.WithInsecure())
	client := pb2.NewURLShortenerClient(connection)
	urlResponse, err := client.Create(context.Background(), &pb2.URLRequest{Url: url})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(urlResponse.GetUrlShort())
	fmt.Println(urlResponse)
	fmt.Println("input url")
	var shortUrl string
	_, err1 := fmt.Scanf("%s\n", &shortUrl)
	if err1 != nil {
		fmt.Println(err1)
	}
	urlResp, err12 := client.Get(context.Background(), &pb2.URLResponse{UrlShort: shortUrl})
	if err12 != nil {
		fmt.Println(err12)
	}
	fmt.Println(urlResp.GetUrl())
	fmt.Println(urlResp)
	fmt.Println("bb")
}
