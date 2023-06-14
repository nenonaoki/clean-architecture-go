package sample_presenter

import (
	sample_usecase "github.com/nenonaoki/clean-architecture-go/pkg/application-business-rule/usecases/sample"
)

type ViewModel struct {
	value int
}

type IPresenter interface {
	ToViewModel() ViewModel
}

type SamplePresenter struct {
	value int
}

func (presenter *SamplePresenter) ToViewModel() *ViewModel {
	var viewModel *ViewModel = &ViewModel{presenter.value}
	return viewModel
}

func NewSamplePresenter(output sample_usecase.IOutput) *SamplePresenter {
	return &SamplePresenter{
		value: output.GetValue(),
	}
}
