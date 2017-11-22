package dal

import (
	"vmango/domain"
)

type StubProvider struct {
	TName     string
	TMachines Machinerep
	TImages   Imagerep

	StatusResponse struct {
		Err    error
		Status *domain.StatusInfo
	}
}

func (p *StubProvider) Name() string {
	return p.TName
}
func (p *StubProvider) Machines() Machinerep {
	return p.TMachines
}
func (p *StubProvider) Images() Imagerep {
	return p.TImages
}

func (p *StubProvider) Status(status *domain.StatusInfo) error {
	if p.StatusResponse.Err != nil {
		return p.StatusResponse.Err
	}
	*status = *p.StatusResponse.Status
	return nil
}
