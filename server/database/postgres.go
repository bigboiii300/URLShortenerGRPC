package database

import (
	"URLShortenerGRPC/server/errorsGRPC"
	pb2 "URLShortenerGRPC/server/proto"
	"URLShortenerGRPC/server/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"log"
	"net/url"
)

type Config struct {
	username,
	password,
	host,
	port,
	databaseName string
}

type URLServerPostgres struct {
	pb2.UnimplementedURLShortenerServer
	connect *pgxpool.Pool
}

func (s *URLServerPostgres) SetConnect(connect *pgxpool.Pool) {
	s.connect = connect
}

func (s *URLServerPostgres) CloseConnection() {
	s.connect.Close()
}

func (s *URLServerPostgres) Get(_ context.Context, resp *pb2.URLResponse) (req *pb2.URLRequest, err error) {
	q := `SELECT long_url from urls WHERE short_url LIKE $1`
	var res string
	row := s.connect.QueryRow(context.Background(), q, resp.GetUrlShort())
	err = row.Scan(&res)
	if err != nil {
		return errorsGRPC.ErrorGet("This short url doesn't exist")
	}
	return &pb2.URLRequest{Url: res}, nil
}

func (s *URLServerPostgres) Create(_ context.Context, req *pb2.URLRequest) (resp *pb2.URLResponse, err error) {
	_, err = url.ParseRequestURI(req.Url)
	if err != nil {
		return errorsGRPC.ErrorCreate("The string is not a url")
	}
	shortURL := utils.GenerateShortURL()
	q := fmt.Sprintf(`INSERT INTO urls (short_url,long_url) VALUES ($1, $2)`)
	row := s.connect.QueryRow(context.Background(), q, shortURL, req.GetUrl())
	errSql := row.Scan(&row)
	if errSql != nil && errSql != pgx.ErrNoRows {
		return errorsGRPC.ErrorCreate("This url is already contained")
	}
	return &pb2.URLResponse{UrlShort: shortURL}, nil
}

func (s *URLServerPostgres) CreateTable() {
	create := `CREATE TABLE IF NOT EXISTS urls
(
    short_url varchar(255) PRIMARY KEY,
    long_url varchar(255) not null unique
);`
	_, err := s.connect.Exec(context.Background(), create)
	if err != nil {
		return
	}
}

func NewClient() *pgxpool.Pool {
	config := Config{
		username:     viper.GetString("database.username"),
		password:     viper.GetString("database.password"),
		host:         viper.GetString("database.host"),
		port:         viper.GetString("database.port"),
		databaseName: viper.GetString("database.databaseName")}
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.username, config.password, config.host, config.port, config.databaseName)
	connection, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n\n", err)
	}
	return connection
}
