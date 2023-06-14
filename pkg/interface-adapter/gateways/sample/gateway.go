package sample_gateway

import (
	"fmt"

	sample_database "github.com/nenonaoki/clean-architecture-go/pkg/frameworks-drivers/database"
)

type SampleGateway struct {
	database sample_database.SampleDatabase
}

func (gateway *SampleGateway) Save(x int) {
	fmt.Printf("Data process comes here\n")
}

func NewSampleGateway(database sample_database.SampleDatabase) *SampleGateway {
	return &SampleGateway{
		database: database,
	}
}
