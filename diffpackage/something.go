package something1

import (
	something "github.com/egon12/dep-drawer/diffpackage/diffpackage2"
	"github.com/egon12/dep-drawer/diffpackage/diffpackage3"
)

func Hola() string {
	diffpackage3.Something()
	return something.Hola()
}
