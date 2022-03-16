package internal

import (
	"fmt"
	"os"
	"syscall"

	"mydocker/pkg/container"
	"mydocker/pkg/image"
	"mydocker/pkg/network"
	"mydocker/pkg/reexec"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Run runs a command inside a new container.
func Run(cmd *cobra.Command, args []string) error {
	fmt.Println("start new container")
	ctr := container.NewContainer()
	defer ctr.Remove()
	fmt.Println("end new container")

	// setup bridge
	fmt.Println("start setup bridge")
	if err := network.SetupBridge("vessel0"); err != nil {
		return err
	}
	fmt.Println("end setup bridge")

	// setup network
	fmt.Println("start setup network")
	deleteNetwork, err := ctr.SetupNetwork("vessel0")
	if err != nil {
		return err
	}
	defer deleteNetwork()
	fmt.Println("end setup network")

	// setup filesystem from image
	fmt.Println("start get image")
	img, err := getImage(args[0])
	fmt.Println("image: ", img)
	if err != nil {
		return err
	}
	fmt.Println("end get image")

	fmt.Println("start mount from image")
	unmount, err := ctr.MountFromImage(img)
	if err != nil {
		return errors.Wrap(err, "can't Mount image filesystem")
	}
	defer unmount()
	fmt.Println("end mount from image")

	// Format fork options
	options := append([]string{}, rawFlags(cmd.Flags())...)
	options = append(options, fmt.Sprintf("--root=%s", ctr.RootFS))
	options = append(options, fmt.Sprintf("--container=%s", ctr.Digest))
	newArgs := []string{"fork"}
	newArgs = append(newArgs, options...)
	newArgs = append(newArgs, args[1:]...)
	fmt.Println("newArgs: ", newArgs)
	newCmd := reexec.Command(newArgs...)
	newCmd.Stdin, newCmd.Stdout, newCmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	newCmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID,
	}
	if err := newCmd.Run(); err != nil {
		return errors.Wrap(err, "failed run fork process")
	}

	return nil
}

// rawFlags convert a pflag.FlagSet to a slice of string.
func rawFlags(flags *pflag.FlagSet) []string {
	var flagList []string
	flags.VisitAll(func(flag *pflag.Flag) {
		if flag.Value.String() == "" {
			return
		}
		flagList = append(flagList, fmt.Sprintf("--%s=%v", flag.Name, flag.Value))
	})
	return flagList
}

// getImage pulls an image and download its layers if it image
// does not exist locally.
func getImage(src string) (*image.Image, error) {
	img, err := image.NewImage(src)
	if err != nil {
		return img, errors.Wrapf(err, "Can't pull %q", src)
	}
	exists, err := img.Exists()
	if err != nil {
		return img, err
	}
	if !exists {
		fmt.Printf("Unable to find image %s:%s locally\n", img.Repository, img.Tag)
		fmt.Printf("downloading the image from %s\n", img.Registry)
		err = img.Download()
	}

	return img, err
}
