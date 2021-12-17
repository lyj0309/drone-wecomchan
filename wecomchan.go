package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	key := os.Getenv("PLUGIN_KEY")
	path := os.Getenv("PLUGIN_PATH")
	msg := os.Getenv("PLUGIN_MSG")

	if key == "" || msg == "" || path == "" {
		log.Fatalln("key or title[text] or path is required")
	}

	res, err := http.PostForm(path, url.Values{
		"sendkey":  []string{key},
		"msg":      []string{msg},
		"msg_type": []string{"text"},
	})
	if err != nil {
		log.Fatalln("post error:", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalln("status code:", res.StatusCode)
	}
}
