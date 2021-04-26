package cmd

import (
	"errors"
	"fmt"
	giDevice "github.com/electricbubble/gidevice"
	"github.com/electricbubble/gidevice-cli/internal"
	"path/filepath"

	"github.com/spf13/cobra"
)

// crashreportCmd represents the crashreport command
var crashreportCmd = &cobra.Command{
	Use:   "crashreport",
	Short: "Move crash reports from device to a local DIRECTORY.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internal.ErrorExit(errors.New("required parameter missing 'localDirectory'"))
		}
		localDirectory := args[0]
		udid, _ := cmd.Flags().GetString("udid")
		isKeep, _ := cmd.Flags().GetBool("keep")
		isExtract, _ := cmd.Flags().GetBool("extract")

		if !filepath.IsAbs(localDirectory) {
			var err error
			if localDirectory, err = filepath.Abs(localDirectory); err != nil {
				internal.ErrorExit(err)
			}
		}

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		prefix := "Move"
		if isKeep {
			prefix = "Copy"
		}
		err = d.MoveCrashReport(localDirectory,
			giDevice.WithKeepCrashReport(isKeep),
			giDevice.WithExtractRawCrashReport(isExtract),
			giDevice.WithWhenMoveIsDone(func(filename string) {
				fmt.Printf("%s: %s\n", prefix, filename)
			}),
		)
		internal.ErrorExit(err)
	},
}

func init() {
	rootCmd.AddCommand(crashreportCmd)

	crashreportCmd.Flags().BoolP("keep", "k", false, "copy but do not remove crash reports from device")
	crashreportCmd.Flags().BoolP("extract", "e", false, "extract raw crash report into separate '.crash' file")

	crashreportCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
