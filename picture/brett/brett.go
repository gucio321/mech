package main

import (
   "fmt"
   "github.com/89z/mech/youtube"
   "github.com/brett-lempereur/ish"
   "image/jpeg"
   "net/http"
   "time"
)

const (
   width = 8
   height = 8
)

func brett(addr string, img *youtube.Image) ([]byte, error) {
   r, err := http.Get(addr)
   if err != nil {
      return nil, err
   }
   defer r.Body.Close()
   i, err := jpeg.Decode(r.Body)
   if err != nil {
      return nil, err
   }
   if img != nil {
      i = img.SubImage(i)
   }
   return ish.NewDifferenceHash(width, height).Hash(i)
}

func brett_main(img youtube.Image) error {
   a, err := brett(mb, nil)
   if err != nil {
      return err
   }
   for _, id := range ids {
      b, err := brett(img.Address(id), &img)
      if err != nil {
         return err
      }
      fmt.Println(ish.NewDifferenceHash(width, height).Distance(a, b), id)
      time.Sleep(100 * time.Millisecond)
   }
   return nil
}