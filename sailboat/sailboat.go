package sailboat

import base "shipyard/abstract"

type Sailboat struct {
	base.Shipyard
}

func New(name string) base.ShipyardInterface {
	return &Sailboat{
		Shipyard: base.Shipyard{
			Name: name,
		},
	}
}

func (s *Sailboat) GetName() string {
	return s.Name
}
