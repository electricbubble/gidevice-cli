package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
)

// infoCmd represents the device info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show device info",
	Run: func(cmd *cobra.Command, args []string) {
		udid, _ := cmd.Flags().GetString("udid")
		jsonFormat, _ := cmd.Flags().GetBool("json")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)
		res, err := d.GetValue("", "")
		internal.ErrorExit(err)
		data, err := json.Marshal(res)
		internal.ErrorExit(err)
		devInfo := new(Device)
		err = json.Unmarshal(data, devInfo)
		internal.ErrorExit(err)

		if jsonFormat {
			str, err := json.Marshal(devInfo)
			internal.ErrorExit(err)
			fmt.Println(string(str))
		} else {
			var keys = reflect.TypeOf(*devInfo)
			var vals = reflect.ValueOf(*devInfo)
			num := keys.NumField()
			for i := 0; i < num; i++ {
				val := vals.Field(i).String()
				if len(val) > 0 {
					fmt.Printf("%s:%s\n", keys.Field(i).Name, vals.Field(i))
				}
			}
		}

	},
}

type Device struct {
	DeviceName           string `json:"DeviceName,omitempty"`
	ProductVersion       string `json:"ProductVersion,omitempty"`
	ProductType          string `json:"ProductType,omitempty"`
	ModelNumber          string `json:"ModelNumber,omitempty"`
	SerialNumber         string `json:"SerialNumber,omitempty"`
	PhoneNumber          string `json:"PhoneNumber,omitempty"`
	CPUArchitecture      string `json:"CPUArchitecture,omitempty"`
	ProductName          string `json:"ProductName,omitempty"`
	ProtocolVersion      string `json:"ProtocolVersion,omitempty"`
	RegionInfo           string `json:"RegionInfo,omitempty"`
	TimeIntervalSince197 string `json:"TimeIntervalSince197,omitempty"`
	TimeZone             string `json:"TimeZone,omitempty"`
	UniqueDeviceID       string `json:"UniqueDeviceID,omitempty"`
	WiFiAddress          string `json:"WiFiAddress,omitempty"`
	BluetoothAddress     string `json:"BluetoothAddress,omitempty"`
	BasebandVersion      string `json:"BasebandVersion,omitempty"`
	DeviceColor          string `json:"DeviceColor,omitempty"`
	DeviceClass          string `json:"DeviceClass,omitempty"`
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().StringP("udid", "u", "", "Device uuid")
	infoCmd.Flags().BoolP("json", "j", false, "Pass true for json format output")
}
