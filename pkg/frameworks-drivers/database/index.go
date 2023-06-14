package sample_database

import "fmt"

type SampleDatabase struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func NewSampleDatabase() *SampleDatabase {
	fmt.Printf("Initialize database client\n")
	return &SampleDatabase{
		Host:     "127.0.0.1",
		Port:     5432,
		Username: "root",
		Password: "",
		Database: "sample",
	}
}
