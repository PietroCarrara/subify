package cmd

import (
	"strings"

	"github.com/matcornic/subify/common/utils"
	"github.com/matcornic/subify/subtitles"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	logger "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

var openVideo bool

var notify bool

// dlCmd represents the dl command
var dlCmd = &cobra.Command{
	Use:     "dl <video-path>",
	Aliases: []string{"download"},
	Short:   "Download the subtitles for your video - 'subify dl --help'",
	Long: `Download the subtitles for your video (movie or TV Shows)
Give the path of your video as first parameter and let's go !`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			utils.Exit("Video file needed. See usage : 'subify help' or 'subify dl --help'")
		}
		videoPath := args[0]
		utils.VerbosePrintln(logger.INFO, "Given video file is "+videoPath)

		apis := strings.Split(viper.GetString("download.apis"), ",")
		languages := strings.Split(viper.GetString("download.languages"), ",")
		err := subtitles.Download(videoPath, apis, languages, notify)
		if err != nil {
			utils.ExitPrintError(err, "Sadly, we could not download any subtitle for you. Try another time or contribute to the apis. See 'subify upload -h'")
		}
		if openVideo {
			err := open.Run(videoPath)
			if err != nil {
				utils.ExitPrintError(err, "Sadly, we could not open video: %s", videoPath)
			}
		}
	},
}

func init() {
	dlCmd.Flags().StringP("languages", "l", "en", "Languages of the subtitle separate by a comma (First to match is downloaded). Available languages at 'subify list languages'")
	dlCmd.Flags().StringP("apis", "a", "SubDB,OpenSubtitles,Addic7ed", "Overwrite default searching APIs behavior, hence the subtitles are downloaded. Available APIs at 'subify list apis'")
	dlCmd.Flags().BoolVarP(&openVideo, "open", "o", false,
		"Once the subtitle is downloaded, open the video with your default video player"+
			` (OSX: "open", Windows: "start", Linux/Other: "xdg-open")`)
	dlCmd.Flags().BoolVarP(&notify, "notify", "n", true, "Display desktop notification")
	_ = viper.BindPFlag("download.languages", dlCmd.Flags().Lookup("languages"))
	_ = viper.BindPFlag("download.apis", dlCmd.Flags().Lookup("apis"))

	RootCmd.AddCommand(dlCmd)
}
