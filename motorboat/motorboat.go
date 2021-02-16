package motorboat

import base "shipyard/abstract"

type Motorboat struct {
	base.Shipyard
	Speed int
}

func New(name string, speed int) MotorboatInterface {
	return &Motorboat{
		Shipyard: base.Shipyard{
			Name: name,
		},
		Speed: speed,
	}
}

func (m *Motorboat) GetName() string {
	return m.Name
}

func (m *Motorboat) GetSpeed() int {
	return m.Speed
}
