package dal

import (
	"vmango/domain"
)

type Machinerep interface {
	List(*domain.VirtualMachineList) error
	Get(*domain.VirtualMachine) (bool, error)
	Create(*domain.VirtualMachine, *domain.Image, *domain.Plan) error
	Start(*domain.VirtualMachine) error
	Stop(*domain.VirtualMachine) error
	Remove(*domain.VirtualMachine) error
	Reboot(*domain.VirtualMachine) error
}

type Imagerep interface {
	List(*domain.ImageList) error
	Get(*domain.Image) (bool, error)
}

type Planrep interface {
	List(*[]*domain.Plan) error
	Get(*domain.Plan) (bool, error)
}

type SSHKeyrep interface {
	List(*[]*domain.SSHKey) error
	Get(*domain.SSHKey) (bool, error)
}

type Authrep interface {
	Get(*domain.User) (bool, error)
}
