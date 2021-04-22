package cmd

import (
	"errors"
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.ErrorExit(errors.New("required parameter missing 'ipaPath'"))
		}
		ipaPath := args[0]
		udid, _ := cmd.Flags().GetString("udid")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		err = d.AppInstall(ipaPath)
		internal.ErrorExit(err)

		fmt.Println("successful install")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
