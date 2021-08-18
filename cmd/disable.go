package cmd

import (
	ledControl "github.com/ch4nn0n/rpi-led-control/pkg/rpi-led-control"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable LEDs",

	Run: func(cmd *cobra.Command, args []string) {
		err := ledControl.Disable()
		if err != nil {
			println("Error when trying to turn off LEDs")
			panic(err)
		}
		println("LEDs off!")
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
