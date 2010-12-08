package crawler

import (
  "fmt"
  "http"
  "os"
)

const (
  BufSize = 1024 * 8
)

func Crawl(url string) (result string, finalUrl string, err os.Error) {
  resp, finalUrl, err := http.Get(url)
  if err == nil {
    fmt.Println("Status: " + resp.Status)
    fmt.Printf("Close: %b\n", resp.Close)
    body := resp.Body
    defer body.Close()

    buf := make([]byte, BufSize)
    var n int
    for {
      n, err = body.Read(buf);
      if n != 0 {
        result += string(buf[0:n])
      }
      if err != nil {
        break
      }
    }
  }

  if err == os.EOF {
    err = nil
  }
  return result, finalUrl, err
}

//func createUrl(stock String) (url *URL) {
//  url := new(URL)
//  url.Raw = "http://stockhtm.finance.qq.com/sstock/ggcx/" + stock + ".shtml?pgv_ref=smartbox&_ver=1.16"
//  url.Schema = "http"
//  url.Host = "stockhtm.finance.qq.com"
//  url.RawPath = "sstock/ggcx/" + stock + ".shtml?pgv_ref=smartbox&_ver=1.16"
//  url.Path = "sstock/ggcx/" + stock
//  url.RawQuery = "pgv_ref=smartbox&_ver=1.16"
//  return &url
//}
