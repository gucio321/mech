package main

import (
   "bytes"
   "fmt"
   "image"
   "image/jpeg"
   "image/png"
   "io"
   "net/http"
   "strconv"
   "time"
)

var ids = []int64{
   0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 20, 21, 22, 23, 24,
   25, 26, 27, 28, 29, 31, 32, 33, 36, 37, 38, 41, 42, 43, 44, 50, 65, 66, 67,
   68, 69,
}

type imageType struct {
   image.Image
   typ string
}

func newImage(buf []byte) (*imageType, error) {
   img, err := jpeg.Decode(bytes.NewReader(buf))
   if err != nil {
      img, err := png.Decode(bytes.NewReader(buf))
      if err != nil {
         return nil, err
      }
      return &imageType{img, "png"}, nil
   }
   return &imageType{img, "jpeg"}, nil
}

func main() {
   for _, id := range ids {
      addr := []byte("http://f4.bcbits.com/img/a3809045440_")
      addr = strconv.AppendInt(addr, id, 10)
      res, err := http.Get(string(addr))
      if err != nil {
         panic(err)
      }
      defer res.Body.Close()
      buf, err := io.ReadAll(res.Body)
      if err != nil {
         panic(err)
      }
      img, err := newImage(buf)
      if err != nil {
         panic(err)
      }
      bound := img.Bounds()
      fmt.Println(id, bound, img.typ)
      time.Sleep(time.Millisecond)
   }
}