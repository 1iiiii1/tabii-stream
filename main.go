package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	webview "github.com/webview/webview_go"
)

type Channel struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

var channels = []Channel{
	{Name: "Selcuk Sports", Slug: "selcuk_sports", URL: "https://df16ea90s1u1080.05949e7ffb6b0bf5.live/live/selcukobs1/playlist.m3u8"},
	{Name: "Tabii Spor 1", Slug: "tabii_spor_1", URL: "https://beert7sqimrk0bfdupfgn6qew.medya.trt.com.tr/master_1080p.m3u8"},
	{Name: "Tabii Spor 2", Slug: "tabii_spor_2", URL: "https://klublsslubcgyiz7zqt5bz8il.medya.trt.com.tr/master_1080p.m3u8"},
	{Name: "Tabii Spor 3", Slug: "tabii_spor_3", URL: "https://ujnf69op16x2fiiywxcnx41q8.medya.trt.com.tr/master_1080p.m3u8"},
	{Name: "Tabii Spor 4", Slug: "tabii_spor_4", URL: "https://bfxy3jgeydpbphtk8qfqwm3hr.medya.trt.com.tr/master_1080p.m3u8"},
	{Name: "Tabii Spor 5", Slug: "tabii_spor_5", URL: "https://z3mmimwz148csv0vaxtphqspf.medya.trt.com.tr/master_1080p.m3u8"},
	{Name: "Tabii Spor 6", Slug: "tabii_spor_6", URL: "https://vbtob9hyq58eiophct5qctxr2.medya.trt.com.tr/master_1080p.m3u8"},
	{Name: "Tabii Spor 7", Slug: "tabii_spor_7", URL: "https://rxne77juptdeyke3tytvgqwyh.medya.trt.com.tr/master_1080p.m3u8"},
	{Name: "TRT 1", Slug: "trt1", URL: "https://tv-trt1.medya.trt.com.tr/master_1080.m3u8"},
}

//go:embed assets/index.html
var indexHTML string

//go:embed assets/hls.min.js
var hlsLib string

func main() {
	// HTTP sunucu
	mux := http.NewServeMux()
	mux.HandleFunc("/hls.min.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		fmt.Fprint(w, hlsLib)
	})
	mux.HandleFunc("/channels", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(channels)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, indexHTML)
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("listener: %v", err)
	}
	go func() {
		if err := http.Serve(ln, mux); err != nil {
			log.Printf("server kapandı: %v", err)
		}
	}()

	url := "http://" + ln.Addr().String() + "/"
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Tabii Stream")
	w.SetSize(1280, 680, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
