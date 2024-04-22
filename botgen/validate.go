package botgen

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func ExtractVideoID(RumbleURL string) (string, error) {
 
    livestream, err := url.Parse(RumbleURL)
    host := livestream.Hostname()
    if host != "rumble.com" && host != "www.rumble.com" {
      return "", errors.New("(-) please enter a valid rumble livestream link")
    }

    if err != nil {
      return "", err
    }
    fmt.Println("(+) Getting the video id...")
    client := &http.Client{}

    req, err := http.NewRequest("GET", RumbleURL, nil)
    if err != nil {
        return "", fmt.Errorf("error: %v", err)
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0")

    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("error: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error: %v", err)
    }

    html := string(body)

    re := regexp.MustCompile(`"embedUrl":"https://rumble.com/embed/([^"]+)/"`)
    match := re.FindStringSubmatch(html)

    if len(match) < 2 {
        return "", fmt.Errorf("(-) Couldn't retrieve the video id, please review the live stream link you provided")
    }

    videoID := match[1]
    fmt.Println("(+) Retrieved video id:", videoID)
    return videoID, nil
}

