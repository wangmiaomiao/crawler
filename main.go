package main

import (
  "./crawler"
  "./print"
  "fmt"
)

func createUrl(stock string) (url string) {
  return "http://stockhtm.finance.qq.com/sstock/ggcx/" + stock + ".shtml?pgv_ref=smartbox&_ver=1.16"
}


func main() {
  if res, url, err := crawler.Crawl(createUrl("601318")); err != nil {
    print.PrintOrDie("Error: " + err.String())
  } else {
    print.PrintOrDie(fmt.Sprintf("Final url: %s\nContent: %s", url, res))
  }
}
