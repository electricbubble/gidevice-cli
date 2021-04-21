package cmd

import (
	"errors"
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
)

// xctestCmd represents the xctest command
var xctestCmd = &cobra.Command{
	Use:   "xctest",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.ErrorExit(errors.New("required parameter missing 'bundleID'"))
		}
		bundleID := args[0]
		udid, _ := cmd.Flags().GetString("udid")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		if !internal.IsDeveloper(d) {
			internal.ErrorExit(fmt.Errorf("%s: may need to mount Developer Disk Image first", d.Properties().SerialNumber))
		}

		out, cancel, err := d.XCTest(bundleID)
		internal.ErrorExit(err)

		done := make(chan os.Signal, 1)
		// signal.Notify(done, os.Interrupt, os.Kill, syscall.SIGTERM)
		signal.Notify(done, os.Interrupt, os.Kill)

		go func() {
			for s := range out {
				fmt.Print(s)
			}
			done <- os.Interrupt
		}()

		<-done
		cancel()
		fmt.Println()
		log.Println("DONE")
	},
}

func init() {
	rootCmd.AddCommand(xctestCmd)

	xctestCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
