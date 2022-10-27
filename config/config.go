package config

import (
	"bytes"
	"fmt"
	"log"
)

var (
	buf    bytes.Buffer
	Logger = log.New(&buf, "logger: ", log.Lshortfile)
)

const (
	Host     = "localhost"
	User     = "postgres"
	Password = "mysecretpassword"
	Name     = "postgres"
	Port     = "5432"
)

var ConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	Host,
	Port,
	User,
	Name,
	Password,
)
