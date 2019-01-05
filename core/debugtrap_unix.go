// +build !windows

package core

import (
	"os"
	"os/signal"

	stackdump "github.com/go-apps/pkg/signal"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

func (d *CoreApplication) setupDumpStackTrap(root string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, unix.SIGUSR1)
	go func() {
		for range c {
			path, err := stackdump.DumpStacks(root)
			if err != nil {
				logrus.WithError(err).Error("failed to write goroutines dump")
			} else {
				logrus.Infof("goroutine stacks written to %s", path)
			}
		}
	}()
}