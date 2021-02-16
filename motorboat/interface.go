package motorboat

import base "shipyard/abstract"

type MotorboatInterface interface {
	base.ShipyardInterface
	GetSpeed() int
}
