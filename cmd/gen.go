/*
Copyright © 2022 Rohan Prabhu rohan.prabhu05@gmail.com

*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var file string
var output string
var duration string
var crop bool

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a thumbnai for the specified video",
	Long: `Use the -f or the --file flag to specify a file to generate thumbnail`,
	Example: `thumbify gen -f "./video.mp4" -c false -d 120 -o "output.png"`,	
	Run: func(ccmd *cobra.Command, args []string) {
		//00:00:01.000

		if duration == "" {
			durationCMD := `bc -l <<< "$(ffprobe -loglevel error -of csv=p=0 -show_entries format=duration video.mp4) * 0.5"`

			dCmd := exec.Command("bash", "-c", durationCMD)

			cmdOutput, err := dCmd.CombinedOutput()
			if err != nil {
				log.Println(fmt.Sprint(err) + ": " + string(output))
			}
			duration = string(cmdOutput)
		}
		var cmd string
		if crop {
			cmd = `ffmpeg -ss %s -i "%s" -an -q 0 -vf scale="'if(gt(iw,ih),-1,1080):if(gt(iw,ih),1080,-1)', crop=1080:1080:exact=1" -vframes 1 "%s" -y`
		} else {
			cmd = `ffmpeg -ss %s -i "%s" -an -q 0  -vframes 1 "%s" -y`
		}

		cmdSubstituted := fmt.Sprintf(cmd, strings.TrimSpace(string(duration)), file, output)

		ffCmd := exec.Command("bash", "-c", cmdSubstituted)
		cmdOutput, err := ffCmd.CombinedOutput()
		if err != nil {
			log.Println(fmt.Sprint(err) + ": " + string(cmdOutput))

		}
	},
}

func init() {

	genCmd.Flags().StringVarP(&file, "file", "f", "", "Input video file for generating thumbnail")
	genCmd.Flags().StringVarP(&output, "output", "o", "output.png", "Output file name")
	genCmd.Flags().BoolVarP(&crop, "crop", "c", false, "Boolean to decide if the thumbnail should be cropped or not")
	genCmd.Flags().StringVarP(&duration, "time", "t", "", "Time Frame(in seconds) at which thumnail is needed. If not specified, default will be considered as 50% of the total time")

	genCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(genCmd)

}
