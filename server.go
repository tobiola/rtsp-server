package main

// import "fmt"
// import "html"
import "log"
import "net/http"
// import "strings"
import "os/exec"

type camera struct  {
  name        string
  rtsp        string
  liveStarted bool
  liveffmpeg  *exec.Cmd
}

/*
░░░░░░░░░▄░░░░░░░░░░░░░░▄░░░░
░░░░░░░░▌▒█░░░░░░░░░░░▄▀▒▌░░░
░░░░░░░░▌▒▒█░░░░░░░░▄▀▒▒▒▐░░░
░░░░░░░▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐░░░
░░░░░▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐░░░
░░░▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌░░░
░░▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌░░
░░▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐░░
░▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌░
░▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌░
▀▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒▐░
▐▒▒▐▀▐▀▒░▄▄▒▄▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
▐▒▒▒▀▀▄▄▒▒▒▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▐░
░▌▒▒▒▒▒▒▀▀▀▒▒▒▒▒▒░▒░▒░▒░▒▒▒▌░
░▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▄▒▒▐░░
░░▀▄▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▄▒▒▒▒▌░░
░░░░▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀░░░
░░░░░░▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀░░░░░
░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▀▀░░░░░░░░
*/

func main() {

  // Create constants of available cameras
  cameras := []camera {
    camera {
      name: "cam1",
      rtsp: "rtsp://192.168.0.124:554/mpeg4?username=admin&password=E10ADC3949BA59ABBE56E057F20F883E",
      liveStarted: false,
    },
  }

  // Generic http post request
  http.HandleFunc("/cam1", func(w http.ResponseWriter, r *http.Request) {

    // reqUrl := r.URL.Query()

    // TODO: parse Url request
    // cameraName := reqUrl(

    w.Header().Set("Connection", "keep-alive")
    w.Header().Set("Content-Type", "video/mp4")
    w.Header().Set("Accept-Ranges", "bytes")

    for _, cam := range cameras {
      if !cam.liveStarted {

        // Start ffmpeg command
        cam.liveffmpeg = exec.Command("ffmpeg", "-rtsp_transport", "tcp", "-i",
                                   cam.rtsp, "-vcodec", "copy", "-f",
                                   "mp4", "-movflags", "frag_keyframe+empty_moov",
                                   "-reset_timestamps", "1", "-vsync", "1",
                                   "-flags", "global_header", "-bsf:v",
                                   "dump_extra", "-y", "-")

        cam.liveffmpeg.Stdout = w
        // stdin, _ := cam.liveffmpeg.StdinPipe()
        // stdout, _ := cam.liveffmpeg.StdoutPipe()
        cam.liveffmpeg.Run()

      }
      break
    }

    if false {
      for _, cam := range cameras {
        if cam.liveStarted {
          cam.liveffmpeg.Process.Kill()
          cam.liveStarted = false
        }
      }
    }
  })

  log.Println("Listening on server 3000")
  log.Fatal(http.ListenAndServe(":3000", nil))

}
