// +build windows

package standardpaths

import (
	"errors"
	"os"
)

// WritableLocation Returns the directory where files of type should be written to, or an empty string if the location cannot be determined.
func WritableLocation(tp StandardLocation) (string, error) {
	switch tp {
	case AppDataLocation:
		dir := os.Getenv("AppData")
		if dir == "" {
			return "", errors.New("%AppData% is not defined")
		}
		return dir, nil
	case AppLocalDataLocation:
		return os.UserCacheDir()
	}
	return "", nil
}
