package botgen

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

func GenerateUserAgents(number int) []string {

  rand.Seed(time.Now().UnixNano())

  var userAgents []string
  me := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
  for i := 0; i < number; i++ {
    var sb strings.Builder
    sb.WriteString(me)
    for j := 0; j < 8; j++ {
      sb.WriteByte(byte(rand.Intn(36) + 48))
    }

    userAgents = append(userAgents, sb.String())
  }

  return userAgents
}



func GetViewerIds(videoID string, number int) (map[string]string, string, string) {
	fmt.Println("(+) Getting viewer ids...")

	userAgents := GenerateUserAgents(number)
	viewerIds := make(map[string]string)
	var channelName string
	var extractedVideoID string

	url := "https://rumble.com/embedJS/u3/?request=video&v=" + videoID

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}

	var mu sync.Mutex
	progress := make(chan int)
	total := len(userAgents)

	go func() {
		for n := range progress {
			fmt.Printf("\r(+) Retrieved %d/%d viewer id", n, total)
		}
		fmt.Println("\n(+) Viewer IDs retrieval completed...")
	}()

	var wg sync.WaitGroup
	for _, userAgent := range userAgents {
		wg.Add(1)
		go func(ua string) {
			defer wg.Done()

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
			req.Header.Set("User-Agent", ua)

			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
			defer resp.Body.Close()

			var data map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				return
			}

			mu.Lock()
			defer mu.Unlock()

			if vid, ok := data["vid"].(float64); ok {
				extractedVideoID = fmt.Sprintf("%.0f", vid)
			}

			if author, ok := data["author"].(map[string]interface{}); ok {
				channelName = author["name"].(string)
			}

			viewerID, ok := data["viewer_id"].(string)
			if !ok {
				return
			}
			viewerIds[viewerID] = ua
			progress <- len(viewerIds)
		}(userAgent)
	}

	wg.Wait()
	close(progress)

	return viewerIds, extractedVideoID, channelName
}