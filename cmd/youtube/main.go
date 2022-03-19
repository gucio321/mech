package main

import (
   "flag"
   "github.com/89z/mech/youtube"
   "os"
)

func main() {
   var vid video
   // a
   flag.StringVar(&vid.address, "a", "", "address")
   // b
   flag.StringVar(&vid.id, "b", "", "video ID")
   // c
   flag.BoolVar(&vid.construct, "c", false, "OAuth construct request")
   // e
   flag.BoolVar(&vid.embed, "e", false, "use embedded player")
   // f
   // youtube.com/watch?v=7WTEB7Qbt4U
   flag.IntVar(&vid.bitrate, "f", 513_601, "target video bitrate")
   // g
   flag.StringVar(&vid.audio, "g", "AUDIO_QUALITY_MEDIUM", "target audio")
   // i
   flag.BoolVar(&vid.info, "i", false, "information")
   // r
   var refresh bool
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   // x
   var exchange bool
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.Parse()
   if verbose {
      youtube.LogLevel = 1
   }
   if exchange {
      err := doExchange()
      if err != nil {
         panic(err)
      }
   } else if refresh {
      err := doRefresh()
      if err != nil {
         panic(err)
      }
   } else if vid.id != "" || vid.address != "" {
      err := vid.do()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (v video) player() (*youtube.Player, error) {
   client := youtube.Android
   if v.embed {
      client = youtube.Embed
   }
   if v.id == "" {
      var err error
      v.id, err = youtube.VideoID(v.address)
      if err != nil {
         return nil, err
      }
   }
   if v.construct {
      cache, err := os.UserCacheDir()
      if err != nil {
         return nil, err
      }
      exc, err := youtube.OpenExchange(cache, "/mech/youtube.json")
      if err != nil {
         return nil, err
      }
      return client.PlayerHeader(exc.Header(), v.id)
   }
   return client.Player(v.id)
}
