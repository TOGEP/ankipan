package main

import (
  "log"
  "net/http"
)

func main(){
  http.HandleFunc("/", func(w http.ResponceWriter, r *http.Request){
    w.Write([]byte(`
      <html>
        <head>
          <title>チャット</title>
        </head>
        <body>
          チャットしましょう!
        </body>
      </html>
    `))
  })
}
