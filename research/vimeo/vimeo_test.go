package vimeo

import (
   "fmt"
   "testing"
)

// vimeo.com/_rv/title?path=/581039021/9603038895
// vimeo.com/581039021/9603038895
const path = "/videos/581039021:9603038895"

func TestVimeo(t *testing.T) {
   logLevel = 1
   web, err := newJsonWeb()
   if err != nil {
      t.Fatal(err)
   }
   video, err := web.video(path)
   if err != nil {
      t.Fatal(err)
   }
   for _, down := range video.Download {
      fmt.Printf("%+v\n", down)
   }
}
