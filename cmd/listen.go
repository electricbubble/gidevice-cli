package cmd

import (
	"fmt"
	giDevice "github.com/electricbubble/gidevice"
	"github.com/electricbubble/gidevice-cli/internal"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Monitor the connection status of the device",
	Run: func(cmd *cobra.Command, args []string) {
		usbmux, err := giDevice.NewUsbmux()
		internal.ErrorExit(err)

		devNotifier := make(chan giDevice.Device)
		cancel, err := usbmux.Listen(devNotifier)
		internal.ErrorExit(err)

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, os.Kill)

		for {
			select {
			case <-done:
				cancel()
				return
			case dev := <-devNotifier:
				if dev.Properties().ConnectionType != "" {
					fmt.Printf("%-8s %d %-4s %-40s %d %d\n", "Attached", dev.Properties().DeviceID,
						dev.Properties().ConnectionType, dev.Properties().SerialNumber, dev.Properties().ProductID, dev.Properties().LocationID,
					)
				} else {
					fmt.Printf("%-8s %d\n", "Removed", dev.Properties().DeviceID)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}
