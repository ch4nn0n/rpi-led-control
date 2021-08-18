package cmd

import (
	ledControl "github.com/ch4nn0n/rpi-led-control/pkg/rpi-led-control"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable LEDs",

	Run: func(cmd *cobra.Command, args []string) {
		err := ledControl.Enable()
		if err != nil {
			println("Error when trying to turn on LEDs")
			panic(err)
		}
		println("LEDs on!")
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
