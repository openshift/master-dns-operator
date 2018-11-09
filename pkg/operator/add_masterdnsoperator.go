package operator

import (
	"github.com/openshift/master-dns-operator/pkg/operator/masterdns"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, masterdns.Add)
}
