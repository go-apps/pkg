// +build !linux,!darwin,!freebsd,!windows

package core

func (d *CoreApplication) setupDumpStackTrap(_ string) {
	return
}
