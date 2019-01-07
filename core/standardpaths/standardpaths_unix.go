// +build !windows

package standardpaths

import (
	"os"
)

// WritableLocation Returns the directory where files of type should be written to, or an empty string if the location cannot be determined.
func WritableLocation(tp StandardLocation) (string, error) {
	switch tp {
	case AppDataLocation:
	case AppLocalDataLocation:
		return os.UserCacheDir()
	}
	return "", nil
}
