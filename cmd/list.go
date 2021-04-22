package cmd

import (
	"bytes"
	"fmt"
	giDevice "github.com/electricbubble/gidevice"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
	"strings"
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
			devInfo := new(giDevice.DeviceInfo)
			if _devInfo, _err := d.DeviceInfo(); _err != nil {
				if strings.HasSuffix(_err.Error(), "InvalidService") {
					outErr.WriteString(
						fmt.Sprintf("%s: may need to mount Developer Disk Image first `gidevice mount -h`\n", d.Properties().SerialNumber),
					)
				} else {
					outErr.WriteString(fmt.Sprintf("%s: %s\n", d.Properties().SerialNumber, _err))
				}
			} else {
				devInfo = _devInfo
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
