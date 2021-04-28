package cmd

import (
	"errors"
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"strings"
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
		contains, _ := cmd.Flags().GetStringArray("contains")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		out, cancel, err := d.XCTest(bundleID)
		internal.ErrorExit(err)

		done := make(chan os.Signal, 1)
		// signal.Notify(done, os.Interrupt, os.Kill, syscall.SIGTERM)
		signal.Notify(done, os.Interrupt, os.Kill)

		go func() {
			for s := range out {
				// show all
				if len(contains) == 0 {
					fmt.Print(s)
					continue
				}

				for _, sub := range contains {
					if strings.Contains(s, sub) {
						fmt.Print(s)
					}
				}
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

	xctestCmd.Flags().StringArrayP("contains", "c", []string{}, "Only the logs contained in the text are displayed")

	xctestCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
