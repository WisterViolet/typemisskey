package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/WisterViolet/typemisskey/typo"
	"github.com/WisterViolet/typemisskey/util"
	"github.com/carlescere/scheduler"
)

type note string

type postData struct {
	I         string `json:"i"`
	Text      string `json:"text"`
	ViaMobile bool   `json:"viaMobile"`
	LocalOnly bool   `json:"localOnly"`
}

func newPostData(i, t string, v, l bool) *postData {
	p := new(postData)
	p.I = i
	p.Text = t
	p.ViaMobile = v
	p.LocalOnly = l
	return p
}

func main() {
	var (
		notes   []note
		payload = newPostData("", "", false, false)
	)
	token := os.Getenv("MISSKEY_PERSONAL_API_TOKEN")
	if token == "" {
		log.Fatal("MISSKEY_PERSONAL_API_TOKEN is not difined")
	}
	appSecret := os.Getenv("MISSKEY_APPLICATION_API_TOKEN")
	if appSecret == "" {
		log.Fatal("MISSKEY_APPLICATION_API_TOKEN is not defined")
	}
	byteI := sha256.Sum256([]byte(token + appSecret))
	payload.I = hex.EncodeToString(byteI[:])
	err := util.LoadJSON("notes.json", &notes)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(notes), func(i, j int) {
		notes[i], notes[j] = notes[j], notes[i]
	})
	index := 0
	scheduler.Every(1).Hours().NotImmediately().Run(func() {
		baseNote := notes[index]
		payload.Text, err = typo.Generate(string(baseNote))
		if err != nil {
			log.Fatal(err)
		}
		postData, err := util.WrapJSONString(payload)
		if err != nil {
			log.Fatal(err)
		}
		baseURL := os.Getenv("MISSKEY_URL")
		if baseURL == "" {
			log.Fatal("MISSKEY_APPLICATION_API_TOKEN is not defined")
		}
		destURL, err := util.JoinURLPath(baseURL, "/notes/create")
		if err != nil {
			log.Fatal(err)
		}
		response, err := util.PostJSON(destURL, bytes.NewReader(postData))
		fmt.Printf("%s\n", response)
		if index++; index == len(notes) {
			index = 0
		}
	})
	runtime.Goexit()
}
