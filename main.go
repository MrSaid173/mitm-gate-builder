package main

import (
    "log"
    "strings"
    "github.com/lqqyt2423/go-mitmproxy/proxy"
)

type BlockAddon struct {
    proxy.BaseAddon
}

var blocked = []string{
    // YouTube Shorts
    "/shorts",
    "youtube.com/shorts",
    "m.youtube.com/shorts",
    "www.youtube.com/shorts",
    "youtubei.googleapis.com/youtubei/v1/next",
    "shorts.youtube.com",

    // TikTok
    "tiktok.com",
    "vm.tiktok.com",
    "vt.tiktok.com",
    "musically.com",
    "tiktokcdn.com",
    "tiktokv.com",
    "tiktok.com/",

    // Instagram Reels
    "instagram.com/reel",
    "instagram.com/reels",
    "cdninstagram.com/reel",

    // Likee
    "likee.video",
    "likee.com",

    // Snapchat Spotlight
    "snapchat.com/spotlight",
    "snapchat.com/discover",

    // Facebook Reels
    "facebook.com/reel",
    "facebook.com/watch",
}

func (b *BlockAddon) Request(f *proxy.Flow) {
    host := f.Request.URL.Host
    path := f.Request.URL.Path
    target := host + path

    for _, pattern := range blocked {
        if strings.Contains(target, pattern) {
            f.Response = &proxy.Response{
                StatusCode: 403,
                Body:       []byte("Blocked"),
            }
            return
        }
    }
}

func main() {
    opts := &proxy.Options{
        Addr: ":8080",
    }
    p, err := proxy.NewProxy(opts)
    if err != nil {
        log.Fatal(err)
    }
    p.AddAddon(&BlockAddon{})
    log.Fatal(p.Start())
}
