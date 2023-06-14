package sample_usecase

import (
	sample_entity "github.com/nenonaoki/clean-architecture-go/pkg/enterprise-business-rule/entities/sample"
)

type SampleOutput struct {
	value int
}

func (output *SampleOutput) GetValue() int {
	return output.value
}

type SampleUsecase struct {
	dataAccess IDataAccess
}

func (usecase *SampleUsecase) Save(input IInput) IOutput {
	var computedValue int = sample_entity.RunSample(input.GetValue())
	usecase.dataAccess.Save(computedValue)
	var output IOutput = &SampleOutput{computedValue}
	return output
}

func NewSampleUsecase(dataAccess IDataAccess) *SampleUsecase {
	return &SampleUsecase{
		dataAccess: dataAccess,
	}
}
