package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/itsryuku/gorumbot/botgen"
)

func banner() {
	fmt.Println(`
			┳┓     ┓    
			┣┫┓┏┏┳┓┣┓┏┓╋
			┛┗┗┻┛┗┗┗┛┗┛┗ V3
					Ryuku wishes you a good viewbotting ^_^
				`)
}

func main() {
	urlFlag := flag.String("u", "", "Video URL")
	botsFlag := flag.Int("b", 0, "Number of bots")
	verboseFlag := flag.Bool("v", false, "Verbose mode")

	flag.Parse()

	if *urlFlag == "" || *botsFlag == 0 {
		fmt.Println("usage: go run main.go -u <videoURL> -b <num> [-v]")
		fmt.Println("e.g: go run main.go -u https://rumble.com/v4qtw5r-live-gray-zone-warfare-closed-beta-lvl-3.html -b 50")
		return
	}
	banner()
	videoID, err := botgen.ExtractVideoID(*urlFlag)
	if err != nil {
		fmt.Println(err)
		return
	}

	viewerIDs, extractedVideoID, channelName := botgen.GetViewerIds(videoID, *botsFlag)
		
	botgen.Viewbot(viewerIDs, extractedVideoID, *verboseFlag)
	fmt.Println("(+) Viewbotting Channel:", channelName)
	fmt.Println("(+) Click CTRL + C when you are done to exit.")
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			botgen.Viewbot(viewerIDs, extractedVideoID, *verboseFlag)
		}
	}
	
}
