/*
Copyright © 2022 Rohan Prabhu rohan.prabhu05@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thumbify",
	Short: "A CLI app to generate thumbnails of videos",
	Long: "This CLI app uses ffmpeg to generate the thumbnails",
	Example: `thumbify gen -f "./video.mp4" -c false -d 120 -o "output.png"`,	
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


