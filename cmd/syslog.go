package cmd

import (
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// syslogCmd represents the syslog command
var syslogCmd = &cobra.Command{
	Use:   "syslog",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		udid, _ := cmd.Flags().GetString("udid")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		lines, err := d.Syslog()
		internal.ErrorExit(err)

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, os.Kill)

		go func() {
			for line := range lines {
				fmt.Println(line)
			}
			done <- os.Interrupt
		}()

		<-done
		d.SyslogStop()
		fmt.Println()
		log.Println("DONE")
	},
}

func init() {
	rootCmd.AddCommand(syslogCmd)

	syslogCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
