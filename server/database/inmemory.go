package database

import (
	"URLShortenerGRPC/server/errorsGRPC"
	pb2 "URLShortenerGRPC/server/proto"
	"URLShortenerGRPC/server/utils"
	"context"
	"fmt"
	checkUrl "net/url"
)

type URLServerInMemory struct {
	pb2.UnimplementedURLShortenerServer
	inMemory map[string]string
}

func (s *URLServerInMemory) SetInMemory(inMemory map[string]string) {
	s.inMemory = inMemory
}

func checkDuplicate(m map[string]string, val string) (bool, string) {
	for key, value := range m {
		if value == val {
			return true, key
		}
	}
	return false, ""
}

func (s *URLServerInMemory) Create(_ context.Context, req *pb2.URLRequest) (resp *pb2.URLResponse, err error) {
	url := req.GetUrl()
	_, err = checkUrl.ParseRequestURI(url)
	if err != nil {
		return errorsGRPC.ErrorCreate("The string is not a url")
	}
	duplicate, duplicateKey := checkDuplicate(s.inMemory, url)
	if duplicate {
		message := fmt.Sprintf("This url is already contained. Short url for this: %s", duplicateKey)
		return errorsGRPC.ErrorCreate(message)
	}
	shortURL := utils.GenerateShortURL()
	s.inMemory[shortURL] = url
	return &pb2.URLResponse{UrlShort: shortURL}, err
}

func (s *URLServerInMemory) Get(_ context.Context, resp *pb2.URLResponse) (req *pb2.URLRequest, err error) {
	if val, ok := s.inMemory[resp.UrlShort]; ok {
		fmt.Println(val)
		return &pb2.URLRequest{Url: val}, nil
	}
	return errorsGRPC.ErrorGet("This short url doesn't exist")
}
