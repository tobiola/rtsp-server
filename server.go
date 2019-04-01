package main

import "fmt"
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

    fmt.Println("Request made")

    // reqUrl := r.URL.Query()

    // TODO: parse Url request
    // cameraName := reqUrl(

    w.Header().Set("Connection", "keep-alive")
    w.Header().Set("Content-Type", "video/mp4")
    w.Header().Set("Accept-Ranges", "bytes")

    for _, cam := range cameras {
      if !cam.liveStarted {



        /*
        // ffmpeg -i in.avi -f image2 -frames:v 1 img.jpeg
                cam.liveffmpeg = exec.Command("ffmpeg -rtsp_transport udp -i", cam.rtsp, "-f image2 -r 1/5 -updatefirst 1 img.jpg")

                */
        // Start ffmpeg command

        cam.liveffmpeg = exec.Command("ffmpeg -i rtsp://192.168.0.124:554/mpeg4?username=admin&password=E10ADC3949BA59ABBE56E057F20F883E -ar 32000 -threads 4 -f mpegts -codec:a mp2 -b:a 128k -bufsize 8192k -muxdelay 0.001 http://127.0.0.1/cam1/480/640")
        // cam.liveffmpeg = exec.Command("ffmpeg -i [url_for_ip_camera_stream] -ar 32000 -threads 4 -f mpegts -codec:a mp2 -b:a 128k -bufsize 8192k -muxdelay 0.001 http://[putbaseurlhere]/[steamkey]/[xres]/[yres]")
        // cam.liveffmpeg = exec.Command("ffmpeg", "-rtsp_transport", "udp",
        //                              "-i", cam.rtsp, "-vcodec", "copy",
        //                               "-tune", "zerolatency", "-f", "mp4",
        //                               "-movflags", "frag_keyframe+empty_moov",
        //                               "-reset_timestamps", "1", "-vsync", "1",
        //                               "-flags", "global_header", "-y", "-")

        // cam.liveffmpeg = exec.Command("ffmpeg -i rtsp://192.168.0.124:554/mpeg4?username=admin&password=E10ADC3949BA59ABBE56E057F20F883E -f mpegts -codec:v mpeg1video -s 320x240 -b:v 100k -bf 0 http://localhost:3000/cam1")

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
