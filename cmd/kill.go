package cmd

import (
	"errors"
	"fmt"
	giDevice "github.com/electricbubble/gidevice"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.ErrorExit(errors.New("required parameter missing 'bundleID'"))
		}
		bundleID := args[0]
		udid, _ := cmd.Flags().GetString("udid")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		lookup, err := d.InstallationProxyLookup(giDevice.WithBundleIDs(bundleID))
		internal.ErrorExit(err)

		lookupResult := lookup.(map[string]interface{})
		lookupResult = lookupResult[bundleID].(map[string]interface{})
		execName := lookupResult["CFBundleExecutable"]

		runningProcesses, err := d.AppRunningProcesses()
		internal.ErrorExit(err)

		pid := 0
		for _, p := range runningProcesses {
			if p.Name == execName {
				pid = p.Pid
			}
		}

		if pid == 0 {
			internal.ErrorExit(fmt.Errorf("didn't running: %s", bundleID))
		}

		err = d.AppKill(pid)
		internal.ErrorExit(err)

		fmt.Printf("pid: %d\tbundleID: %s\n", pid, bundleID)
	},
}

func init() {
	rootCmd.AddCommand(killCmd)

	killCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
