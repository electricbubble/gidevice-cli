package cmd

import (
	"fmt"

	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch [-u udid] bundleID",
	Short: "Launch application",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("required parameter missing 'bundleID'")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
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
