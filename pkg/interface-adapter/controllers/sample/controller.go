package sample_controller

import (
	sample_usecase "github.com/nenonaoki/clean-architecture-go/pkg/application-business-rule/usecases/sample"
	sample_gateway "github.com/nenonaoki/clean-architecture-go/pkg/interface-adapter/gateways/sample"
)

type IController interface {
	Save() int
}

type SampleInput struct {
	value int
}

func (input *SampleInput) GetValue() int {
	return input.value
}

type SampleController struct {
	gateway *sample_gateway.SampleGateway
}

func (controller *SampleController) Save(value int) sample_usecase.IOutput {
	var usecase sample_usecase.SampleUsecase = *sample_usecase.NewSampleUsecase(controller.gateway)
	var input sample_usecase.IInput = &SampleInput{value}
	var output sample_usecase.IOutput = usecase.Save(input)
	return output
}

func NewSampleController(gateway *sample_gateway.SampleGateway) *SampleController {
	return &SampleController{
		gateway: gateway,
	}
}
