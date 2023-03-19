package notif

import (
	"fmt"

	toast "github.com/jacobmarshall/go-toast"
)

// SendSubtitleDownloadSuccess sends a notification when download went well
func SendSubtitleDownloadSuccess(successAPI, videoPath string) {
	bn := path.Base(videoPath)
	_ = Info("I found a subtitle for \""+bn+"\" :)", fmt.Sprintf("Thank you %s <3", successAPI))
}

// SendSubtitleCouldNotBeDownloaded sends a notification when download went bad
func SendSubtitleCouldNotBeDownloaded(noSucessAPIs, videoPath string) {
	bn := path.Base(videoPath)
	_ = Error("!! I didn't found any subtitle for \""+bn+"\" :'(", fmt.Sprintf("No match for your video in : %s. Try later !", noSucessAPIs))
}

func sendMessage(title, message string) error {
	iconPath := downloadIcon()
	notification := toast.Notification{
		AppID:   "Subify",
		Title:   fmt.Sprintf("Subify: %s", title),
		Message: message,
		Audio:   toast.SMS,
	}
	if iconPath != "" {
		notification.Icon = iconPath
	}
	return notification.Push()
}

// Error send a notification error
func Error(title, message string) error {
	return sendMessage(title, message)
}

// Info send a notification information
func Info(title, message string) error {
	return sendMessage(title, message)
}
