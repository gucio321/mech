package instagram

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "os"
   "strings"
)


func jsonChannel(name string) error {
   req, err := http.NewRequest("GET", origin + "/" + name + "/channel/", nil)
   if err != nil {
      return err
   }
   q := req.URL.Query()
   q.Set("__a", "1")
   req.URL.RawQuery = q.Encode()
   req.Header.Set("User-Agent", "Mozilla")
   dum, err := httputil.DumpRequest(req, false)
   if err != nil {
      return err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}

const body1 = `
{
   "query_hash": "2c4c2e343a8f64c625ba02b2aa12c7f8",
   "variables": {"shortcode":"CT-cnxGhvvO"}
}
`

// severe rate limit
func jsonGraphQL1(id string) error {
   req, err := http.NewRequest(
      "POST", origin + "/graphql/query/", strings.NewReader(body1),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0)"},
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}

const body2 = `
{
   "query_hash": "cf28bf5eb45d62d4dc8e77cdb99d750d",
   "variables": {"shortcode":"CT-cnxGhvvO"}
}
`

// severe rate limit
func jsonGraphQL2(id string) error {
   req, err := http.NewRequest(
      "POST", origin + "/graphql/query/", strings.NewReader(body2),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0)"},
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}

// severe rate limit
func jsonP(id string) error {
   req, err := http.NewRequest("GET", origin + "/p/" + id + "/", nil)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("__a", "1")
   req.URL.RawQuery = val.Encode()
   req.Header.Set("User-Agent", "Mozilla")
   dum, err := httputil.DumpRequest(req, false)
   if err != nil {
      return err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}

// severe rate limit
func jsonTV(id string) error {
   req, err := http.NewRequest("GET", origin + "/tv/" + id + "/", nil)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   val.Set("__a", "1")
   req.URL.RawQuery = val.Encode()
   req.Header.Set("User-Agent", "Mozilla")
   dum, err := httputil.DumpRequest(req, false)
   if err != nil {
      return err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}

func jsonUsers(id string) error {
   req, err := http.NewRequest(
      "GET", "https://i.instagram.com/api/v1/users/" + id + "/info/", nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set("User-Agent", "Instagram 1.1.1")
   dum, err := httputil.DumpRequest(req, false)
   if err != nil {
      return err
   }
   os.Stdout.Write(dum)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      dum, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      return fmt.Errorf("%s", dum)
   }
   return nil
}