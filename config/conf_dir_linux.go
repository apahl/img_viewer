// +build linux

package config

import (
	"os/user"

	"github.com/apahl/utls"
)

// GetConfigDir returns the configuration dir for the current user of the platform
func GetConfigDir() string {
	currUser, err := user.Current()
	utls.QuitOnError(err)
	result := "/home/" + currUser.Username + "/.config/img_viewer"
	return result
}
