package RInt

import (
)

type RocketInterface interface {
	Gogo() bool
}

func Launch(r RocketInterface) {
    r.Gogo()
}