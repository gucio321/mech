package main

import (
   "fmt"
   "github.com/89z/mech/pandora"
   "net/http"
   "os"
)

func playback(cache, addr string, info bool) error {
   user, err := pandora.OpenUserLogin(cache)
   if err != nil {
      return err
   }
   id, err := pandora.AppLink(addr)
   if err != nil {
      return err
   }
   det, err := pandora.NewDetails(id)
   if err != nil {
      return err
   }
   anno := det.Result.Annotations[id]
   play, err := user.PlaybackInfo(id)
   if err != nil {
      return err
   }
   if play.Result != nil {
      if info {
         fmt.Printf("%+v\n", play.Result.AudioUrlMap)
      } else {
         addr := play.Result.AudioUrlMap.HighQuality.AudioUrl
         fmt.Println("GET", addr)
         res, err := http.Get(addr)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         ext, err := play.Ext()
         if err != nil {
            return err
         }
         file, err := os.Create(anno.ArtistName + "-" + anno.Name + ext)
         if err != nil {
            return err
         }
         defer file.Close()
         if _, err := file.ReadFrom(res.Body); err != nil {
            return err
         }
      }
   } else {
      fmt.Printf("%+v\n", play)
   }
   return nil
}

func login(cache, username, password string) error {
   part, err := pandora.NewPartnerLogin()
   if err != nil {
      return err
   }
   user, err := part.UserLogin(username, password)
   if err != nil {
      return err
   }
   fmt.Println("Create", cache)
   return user.Create(cache)
}