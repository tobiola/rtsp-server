package main

import "fmt"
import "html"
import "log"
import "net/http"
import "strings"
import "os/exec"

type camera struct  {
  name        string
  url         string
  liveStarted bool
  liveffmpeg *Cmd
}

cameras := []camera {
  camera {
    name: "cam1",
    url: "rtsp://192.168.0.124:554/mpeg4?username=admin&password=E10ADC3949BA59ABBE56E057F20F883E",
    liveStarted: false,
  },
}



func liveStream(w http.ResponseWriter, r *http.Request) {

  reqUrl := r.URL.Query()

  // TODO: parse Url request
  cameraName := reqUrl(

  w.Header().Set("Connection": "keep-alive")
  w.Header().Set("Content-Type": "video/mp4")
  w.Header().Set("Accept-Ranges": "bytes")

  for _, cam := range cameras {
    // if cam.name === cameraName {

    if !cam.liveStarted {
      stdin := cam.liveffmpeg.StdinPipe()
      stdout := cam.liveffmpeg.StdoutPipe()

    }
  }
}

func shutStream(event string) {

  for _, cam := range cameras {

    if cam.liveStarted {
      cam.liveffmpeg.Process.Kill()
      cam.liveStarted = false
    }

  }
}

func main() {

  http.HandleFunc("/cam1", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hi")
  })

  log.Fatal(http.ListenAndServe(":3000", nil))

}
