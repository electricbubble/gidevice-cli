package cmd

import (
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"
	"time"

	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		udid, _ := cmd.Flags().GetString("udid")
		isAll, _ := cmd.Flags().GetBool("all")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		apps, err := d.AppList()
		internal.ErrorExit(err)

		maxName := 0

		mapper := make(map[string]interface{})
		for _, app := range apps {
			mapper[app.ExecutableName] = app.CFBundleIdentifier
			if len(app.ExecutableName) > maxName {
				maxName = len(app.ExecutableName)
			}
		}

		runningProcesses, err := d.AppRunningProcesses()
		internal.ErrorExit(err)

		for _, p := range runningProcesses {
			if !isAll && !p.IsApplication {
				continue
			}
			bundleID, ok := mapper[p.Name]
			if !ok {
				bundleID = ""
			}
			d := time.Unix(int64(time.Since(p.StartDate).Seconds()), 0).Format("15:01:05")
			fmt.Printf("%4d %-"+fmt.Sprintf("%d", maxName)+"s %8s %s\n", p.Pid, p.Name, d, bundleID)
		}
	},
}

func init() {
	rootCmd.AddCommand(psCmd)

	psCmd.Flags().BoolP("all", "a", false, "All application types")

	psCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
