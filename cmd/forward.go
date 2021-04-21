package cmd

import (
	"fmt"
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

// forwardCmd represents the forward command
var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		udid, _ := cmd.Flags().GetString("udid")
		// netType, _ := cmd.Flags().GetString("type")
		// netType = strings.ToLower(netType)
		localPort, _ := cmd.Flags().GetInt("lport")
		remotePort, _ := cmd.Flags().GetInt("rport")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		fmt.Printf("local port: %d\tremote port: %d\tdevice: %s\n", localPort, remotePort, d.Properties().SerialNumber)

		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", localPort))
		internal.ErrorExit(err)

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, os.Kill)

		go func(listener net.Listener) {
			for {
				var accept net.Conn
				if accept, err = listener.Accept(); err != nil {
					log.Println("accept:", err)
				}

				fmt.Println("accept", accept.RemoteAddr())

				rInnerConn, err := d.NewConnect(remotePort)
				internal.ErrorExit(err)

				rConn := rInnerConn.RawConn()
				_ = rConn.SetDeadline(time.Time{})

				go func(lConn net.Conn) {
					go func(lConn, rConn net.Conn) {
						if _, err := io.Copy(lConn, rConn); err != nil {
							log.Println("copy local -> remote:", err)
						}
					}(lConn, rConn)
					go func(lConn, rConn net.Conn) {
						if _, err := io.Copy(rConn, lConn); err != nil {
							log.Println("copy local <- remote:", err)
						}
					}(lConn, rConn)
				}(accept)
			}
		}(listener)

		<-done
		fmt.Println()
		log.Println("DONE")
	},
}

func init() {
	rootCmd.AddCommand(forwardCmd)

	forwardCmd.Flags().StringP("udid", "u", "", "Device uuid")
	// forwardCmd.Flags().StringP("type", "t", "http", "Network Type")
	forwardCmd.Flags().IntP("lport", "l", 8100, "Local Port")
	forwardCmd.Flags().IntP("rport", "r", 8100, "Remote Port")
}
