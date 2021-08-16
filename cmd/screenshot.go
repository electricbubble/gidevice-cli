package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
)

// screenshotCmd represents the screenshot command
var screenshotCmd = &cobra.Command{
	Use:   "screenshot",
	Short: "Device screenshot",
	Run: func(cmd *cobra.Command, args []string) {
		udid, _ := cmd.Flags().GetString("udid")
		outputFile, _ := cmd.Flags().GetString("outputFile")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		raw, err := d.Screenshot()
		if err != nil {
			log.Fatalln(err)
		}

		img, format, err := image.Decode(raw)
		if err != nil {
			log.Fatalln(err)
		}
		if len(outputFile) == 0 {
			now := time.Now()
			outputFile = fmt.Sprintf("%d%d%d%d%d%d", now.Year(), int(now.Month()),
				now.Day(), now.Hour(), now.Minute(), now.Second()) + "." + format
		}
		file, err := os.Create(outputFile)
		if err != nil {
			log.Fatalln(err)
		}
		defer func() { _ = file.Close() }()
		switch format {
		case "png":
			err = png.Encode(file, img)
		case "jpeg":
			err = jpeg.Encode(file, img, nil)
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(file.Name())
	},
}

func init() {
	rootCmd.AddCommand(screenshotCmd)
	screenshotCmd.Flags().StringP("udid", "u", "", "Device uuid")
	screenshotCmd.Flags().StringP("outputFile", "o", "", "Output file")
}
