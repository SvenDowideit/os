package init

import (
	"github.com/rancher/os/cmd/control"
	"github.com/rancher/os/init/runc"
	"github.com/rancher/os/util"

	log "github.com/sirupsen/logrus"
)

// Non-containerized way to get a console
func recovery(containerized bool) error {
	if containerized {
		log.Infof("Start recovery using Runc")
		return runc.RunSet("recovery_services", util.RootFsIsNotReal())
	}

	log.Infof("Start recovery")
	return control.AutoLogin("root", "", "recovery", "default")
}
