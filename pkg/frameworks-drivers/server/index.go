package sample_server

import (
	"fmt"

	sample_database "github.com/nenonaoki/clean-architecture-go/pkg/frameworks-drivers/database"
	sample_controller "github.com/nenonaoki/clean-architecture-go/pkg/interface-adapter/controllers/sample"
	sample_gateway "github.com/nenonaoki/clean-architecture-go/pkg/interface-adapter/gateways/sample"
	sample_presenter "github.com/nenonaoki/clean-architecture-go/pkg/interface-adapter/presenters/sample"
)

type SampleServer struct {
	database sample_database.SampleDatabase
}

func (server *SampleServer) Sample(x int) *sample_presenter.ViewModel {
	gateway := sample_gateway.NewSampleGateway(server.database)
	controller := sample_controller.NewSampleController(gateway)
	output := controller.Save(x)
	presenter := sample_presenter.NewSamplePresenter(output)
	viewModel := presenter.ToViewModel()
	return viewModel
}

func (server *SampleServer) Start() {
	var value = 1
	fmt.Printf("Start server application %#v\n", server.Sample(value))
}

func NewSampleServer(database sample_database.SampleDatabase) *SampleServer {
	fmt.Printf("Initialize server application\n")
	return &SampleServer{
		database: database,
	}
}
