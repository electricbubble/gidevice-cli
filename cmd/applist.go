package cmd

import (
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
	"strings"
)

// applistCmd represents the applist command
var applistCmd = &cobra.Command{
	Use:   "applist",
	Short: "List of all installed applications",
	Run: func(cmd *cobra.Command, args []string) {
		udid, _ := cmd.Flags().GetString("udid")
		appType, _ := cmd.Flags().GetString("type")
		appType = strings.ToLower(appType)

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		apps, err := d.AppList()
		internal.ErrorExit(err)

		for _, app := range apps {
			if appType == "all" || appType == strings.ToLower(app.Type) {
				fmt.Printf("%s %s %s %s\n", app.Type, app.DisplayName, app.CFBundleIdentifier, app.Version)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(applistCmd)

	applistCmd.Flags().StringP("udid", "u", "", "Device uuid")
	applistCmd.Flags().StringP("type", "t", "user", "Application Type")
}
