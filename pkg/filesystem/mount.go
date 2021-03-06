package filesystem

import (
	"fmt"
	"syscall"

	"github.com/pkg/errors"
)

type MountOption struct {
	Source string
	Target string
	Type   string
	Flag   uintptr
	Option string
}

type Unmounter func() error

// Mount mounts list of mountOptions and returns a function to unmount them.
func Mount(mountOpts ...MountOption) (Unmounter, error) {
	unmounter := func() error {
		for _, p := range mountOpts {
			if err := syscall.Unmount(p.Target, 0); err != nil {
				return errors.Wrapf(err, "unable to umount %q", p.Target)
			}
		}
		return nil
	}

	for _, p := range mountOpts {
		fmt.Println("p: ", p)
		if err := syscall.Mount(p.Source, p.Target, p.Type, p.Flag, p.Option); err != nil {
			return unmounter, errors.Wrapf(err, "unable to mount %s to %s", p.Source, p.Target)
		}
	}

	return unmounter, nil
}
