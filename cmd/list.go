package cmd

import (
	"bytes"
	"fmt"
	giDevice "github.com/electricbubble/gidevice"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all devices",
	Run: func(cmd *cobra.Command, args []string) {
		usbmux, err := giDevice.NewUsbmux()
		internal.ErrorExit(err)

		devices, err := usbmux.Devices()
		internal.ErrorExit(err)

		outErr := new(bytes.Buffer)
		for _, d := range devices {
			devInfo, err := d.DeviceInfo()
			if err != nil {
				outErr.WriteString(fmt.Sprintf("%s: %s\n", d.Properties().SerialNumber, err))
				devInfo = new(giDevice.DeviceInfo)
			}
			fmt.Println(d.Properties().SerialNumber, devInfo.DisplayName)
		}
		if outErr.Len() != 0 {
			fmt.Println(outErr.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
