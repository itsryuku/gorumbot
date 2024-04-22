package botgen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)
func SendView(client *http.Client, url string, body []byte, viewerID, userAgent string, verbose bool, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
			fmt.Printf("error creating request for viewer %s: %v\n", viewerID, err)
			return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)

	if err != nil {
			fmt.Printf("error sending viewer %s: %v\n", viewerID, err)
			return
	}

	defer resp.Body.Close()

	if verbose {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
					fmt.Printf("error reading response body for viewer %s: %v\n", viewerID, err)
					return
			}
			fmt.Printf("Viewer %s response: %s\n", viewerID, string(body))
	}
}


func Viewbot(viewerIDs map[string]string, videoID string, verbose bool) {


	url := "https://wn0.rumble.com/service.php?api=7&name=video.watching-now"
	client := &http.Client{}

	var wg sync.WaitGroup
	for viewerID, userAgent := range viewerIDs {
			body := []byte(fmt.Sprintf("video_id=%s&viewer_id=%s", videoID, viewerID))
			wg.Add(1)
			go SendView(client, url, body, viewerID, userAgent, verbose, &wg)
	}

	wg.Wait()
}
