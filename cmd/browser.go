package cmd

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/sempr/cf/client"
	"github.com/sempr/cf/config"
	"github.com/skratchdot/open-golang/open"
)

var is_wsl bool = isWSL()

func isWSL() bool {

	if runtime.GOOS != "linux" {
		return false
	}

	cmd := exec.Command("uname", "-a")
	if output, err := cmd.Output(); err == nil {
		if strings.Contains(strings.ToLower(string(output)), "microsoft") {
			return true
		}
	}
	return false
}

func openURL(url string) error {
	color.Green("Open %v", url)
	if is_wsl {
		return exec.Command("wslview", url).Start()
	}
	return open.Run(url)
}

// Open command
func Open() (err error) {
	URL, err := Args.Info.OpenURL(config.Instance.Host)
	if err != nil {
		return
	}
	return openURL(URL)
}

// Stand command
func Stand() (err error) {
	URL, err := Args.Info.StandingsURL(config.Instance.Host)
	if err != nil {
		return
	}
	return openURL(URL)
}

// Sid command
func Sid() (err error) {
	info := Args.Info
	if info.SubmissionID == "" && client.Instance.LastSubmission != nil {
		info = *client.Instance.LastSubmission
	}
	URL, err := info.SubmissionURL(config.Instance.Host)
	if err != nil {
		return
	}
	return openURL(URL)
}
