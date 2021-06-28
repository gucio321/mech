package youtube

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
)

const chunk = 10_000_000

var ClientAndroid = Client{"ANDROID_MUSIC", "1.10"}

type Android struct {
   StreamingData struct {
      AdaptiveFormats []Format
   }
   VideoDetails `json:"videoDetails"`
}

func NewAndroid(id string) (*Android, error) {
   res, err := ClientAndroid.newPlayer(id).post()
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   a := new(Android)
   if err := json.NewDecoder(res.Body).Decode(a); err != nil {
      return nil, err
   }
   return a, nil
}

func (a Android) NewFormat(itag int) (*Format, error) {
   for _, format := range a.StreamingData.AdaptiveFormats {
      if format.Itag == itag {
         return &format, nil
      }
   }
   return nil, fmt.Errorf("itag %v", itag)
}

type Format struct {
   Bitrate int64
   ContentLength int64 `json:"contentLength,string"`
   Height int
   Itag int
   MimeType string
   URL string
}

func (f Format) Write(w io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   var pos int64
   fmt.Println(invert, "GET", reset, req.URL)
   for pos < f.ContentLength {
      bytes := fmt.Sprintf("bytes=%v-%v", pos, pos+chunk-1)
      req.Header.Set("Range", bytes)
      fmt.Println(bytes)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      if res.StatusCode != http.StatusPartialContent {
         return fmt.Errorf("status %v", res.Status)
      }
      if _, err := io.Copy(w, res.Body); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}
