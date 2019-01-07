// +build !linux,!darwin,!freebsd,!windows

package crashdmp

func SetupDumpStackTrap(_ string) {
	return
}
