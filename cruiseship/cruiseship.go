package cruiseship

import base "shipyard/abstract"

type Cruiseship struct {
	base.Shipyard
}

func New(name string) base.ShipyardInterface {
	return &Cruiseship{
		Shipyard: base.Shipyard{
			Name: name,
		},
	}
}

func (c *Cruiseship) GetName() string {
	return c.Name
}
