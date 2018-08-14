// +build windows

package config

import (
	"os/user"
	"path/filepath"

	"github.com/apahl/utls"
)

// GetConfigDir returns the configuration dir for the current user of the platform
func GetConfigDir() string {
	currUser, err := user.Current()
	utls.QuitOnError(err)
	_, user := filepath.Split(currUser.Username)
	result := "C:/Users/" + user + "/AppData/Roaming/img_viewer"
	return result
}
