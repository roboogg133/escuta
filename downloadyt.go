package youtube

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kkdai/youtube/v2"
)

type headerRoundTripper struct {
	rt      http.RoundTripper
	headers map[string]string
}

func (h *headerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range h.headers {
		req.Header.Set(k, v)
	}
	return h.rt.RoundTrip(req)
}

type joson struct {
	Visitor string `json:"visitor"`
}

func CheckVisitor(config string) string {

	_, err := os.Stat(config)
	if err != nil {
		file, err := os.Create(config)
		if err != nil {
			log.Fatal("CREATING FILE: ", err)
		}
		visitorId := GetVisitor()

		v := joson{
			Visitor: visitorId,
		}

		jsonData, err := json.MarshalIndent(v, "", "	")
		if err != nil {
			log.Fatal("ERROR TURNING IN JSON: ", err)
		}

		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatal("ERROR WRITING JSON DATA: ", err)
		}
		file.Close()
		return visitorId

	}

	v := joson{}
	file, err := os.Open(config)
	if err != nil {
		log.Fatal("ERROR OPENING config.json : ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&v)
	if err != nil {
		log.Fatal("CANT READ config.json : ", err)
	}

	file.Close()
	return v.Visitor
}

func Download(link string, path string) error {

	visitorc := CheckVisitor("config.json")

	customTransport := &headerRoundTripper{
		rt: http.DefaultTransport,
		headers: map[string]string{
			"x-goog-visitor-id": visitorc,
			"User-Agent":        "Mozilla/5.0",
		},
	}

	var client = youtube.Client{
		HTTPClient: &http.Client{
			Transport: customTransport,
		},
	}

	video, err := client.GetVideo(link)
	if err != nil {
		log.Fatal("ERROR WHILE GETTING VIDEO: ", err)

	}

	stream, _, err := client.GetStream(video, &video.Formats.WithAudioChannels()[0])
	if err != nil {
		log.Fatal("ERROR WHILE GETTING STREAM: ", err)
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatal("ERROR WHILE CREATING FILE: ", err)
	}

	_, err = io.Copy(file, stream)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	return err
}

func GetTitle(url string) string {

	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		log.Fatal("ERROR WHILE GETTING VIDEO: ", err)

	}
	return video.Title
}

func GetAuthor(url string) string {

	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		log.Fatal("ERROR WHILE GETTING VIDEO: ", err)

	}

	return video.Author

}

func GetDuration(url string) int {

	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		log.Fatal("ERROR WHILE GETTING VIDEO: ", err)

	}

	return int(video.Duration)
}

func GetThumbURL(url string) string {

	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		log.Fatal("ERROR WHILE GETTING VIDEO: ", err)

	}

	return video.Thumbnails[0].URL
}
