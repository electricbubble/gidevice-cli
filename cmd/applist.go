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

		if !internal.IsDeveloper(d) {
			internal.ErrorExit(fmt.Errorf("%s: may need to mount Developer Disk Image first", d.Properties().SerialNumber))
		}

		apps, err := d.AppList()
		internal.ErrorExit(err)

		for _, app := range apps {
			if appType == "all" {
				fmt.Println(app.Type, app.DisplayName, app.CFBundleIdentifier, app.Version)
				continue
			}

			if appType == strings.ToLower(app.Type) {
				fmt.Println(app.Type, app.DisplayName, app.CFBundleIdentifier, app.Version)
				continue
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(applistCmd)

	applistCmd.Flags().StringP("udid", "u", "", "Device uuid")
	applistCmd.Flags().StringP("type", "t", "all", "Application Type")
}
