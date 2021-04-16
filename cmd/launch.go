package cmd

import (
	"errors"
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch application",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.ErrorExit(errors.New("required parameter missing 'bundleID'"))
		}
		bundleID := args[0]
		udid, _ := cmd.Flags().GetString("udid")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		pid, err := d.AppLaunch(bundleID)
		internal.ErrorExit(err)

		fmt.Printf("pid: %d\tbundleID: %s\n", pid, bundleID)
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)

	launchCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
