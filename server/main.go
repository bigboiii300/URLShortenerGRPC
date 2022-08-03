package main

import (
	db "URLShortenerGRPC/server/database"
	"URLShortenerGRPC/server/proto"
	"URLShortenerGRPC/server/utils"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	mode := flag.String("mode", "postgres", "input mode")
	flag.Parse()
	postgresDatabase := &db.URLServerPostgres{}
	inMemoryDatabase := &db.URLServerInMemory{}
	errViper := utils.InitViper()
	if err != nil {
		log.Fatal(errViper)
	}
	switch *mode {
	case "postgres":
		pb.RegisterURLShortenerServer(s, postgresDatabase)
		connection := db.NewClient()
		postgresDatabase.SetConnect(connection)
		postgresDatabase.CreateTable()
	case "local":
		pb.RegisterURLShortenerServer(s, inMemoryDatabase)
		inMemoryDatabase.SetInMemory(make(map[string]string))
	default:
		pb.RegisterURLShortenerServer(s, inMemoryDatabase)
		inMemoryDatabase.SetInMemory(make(map[string]string))
	}
	defer postgresDatabase.CloseConnection()
	errServe := s.Serve(lis)
	if errServe != nil {
		log.Fatal(errServe)
	}
}
