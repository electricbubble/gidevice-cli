package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		udid, _ := cmd.Flags().GetString("udid")
		dir, _ := cmd.Flags().GetString("dir")
		show, _ := cmd.Flags().GetBool("list")

		if show {
			d, err := internal.GetDeviceFromCommand(udid)
			internal.ErrorExit(err)

			imageSignatures, err := d.Images()
			internal.ErrorExit(err)

			for i, imgSign := range imageSignatures {
				fmt.Printf("[%d] %s\n", i+1, base64.StdEncoding.EncodeToString(imgSign))
			}
			return
		}

		dmgPath, signaturePath := "", ""
		if len(dir) != 0 {
			entries, err := os.ReadDir(dir)
			internal.ErrorExit(err)

			if !filepath.IsAbs(dir) {
				abs, err := filepath.Abs(dir)
				internal.ErrorExit(err)
				dir = abs
			}

			for _, e := range entries {
				if strings.HasSuffix(strings.ToLower(e.Name()), ".dmg") {
					dmgPath = filepath.Join(dir, e.Name())
				}
				if strings.HasSuffix(strings.ToLower(e.Name()), ".signature") {
					signaturePath = filepath.Join(dir, e.Name())
				}
			}
		} else {
			if len(args) < 2 {
				internal.ErrorExit(errors.New("required parameter missing 'dmg' & 'signature'"))
			}
			dmgPath, signaturePath = args[1], args[2]
		}

		if len(dmgPath) == 0 {
			internal.ErrorExit(errors.New("required parameter missing 'dmg'"))
		}
		if len(signaturePath) == 0 {
			internal.ErrorExit(errors.New("required parameter missing 'signature'"))
		}

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		// dmgPath := "/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/DeviceSupport/14.4/DeveloperDiskImage.dmg"
		// signaturePath := "/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/DeviceSupport/14.4/DeveloperDiskImage.dmg.signature"

		err = d.MountDeveloperDiskImage(dmgPath, signaturePath)
		internal.ErrorExit(err)

		fmt.Println("successful mount")
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)

	mountCmd.Flags().BoolP("list", "l", true, "DeveloperDiskImage list")

	mountCmd.Flags().StringP("udid", "u", "", "Device uuid")
	mountCmd.Flags().StringP("dir", "d", "", "DeveloperDiskImage directory")
}
