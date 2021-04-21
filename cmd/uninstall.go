package cmd

import (
	"errors"
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"

	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
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

		err = d.AppUninstall(bundleID)
		internal.ErrorExit(err)

		fmt.Println("successful uninstall")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
