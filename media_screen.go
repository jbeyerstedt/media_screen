package main

import (
  "fmt"
  "io/ioutil"
  "path/filepath"
  "os/exec"
  "log"
)

func strInSlice(str string, list []string) bool {
  for _, elem := range list {
    if elem == str {
      return true
    }
  }
  return false
}

func main() {
  image_ext := []string{".jpg", ".jpeg",
                        ".gif",
                        ".tiff", ".tif",
                        ".bmp",
                        ".png"}
  video_ext := []string{".avi", ".mov", ".mkv", ".mp4", ".m4v"}

  image_cmd_base := "sudo fbi -a --noverbose -T 1 -t 3 "
  media_folder := "/home/pi/video"


  /* get files and sort by video or image */
  files, err := ioutil.ReadDir(media_folder)
  if err != nil {
    log.Fatal(err)
  }

  images := make([]string, 0, len(files))
  videos := make([]string, 0, len(files))

  for _, f := range files {
    img_true := strInSlice(filepath.Ext(f.Name()), image_ext[:])
    vid_true := strInSlice(filepath.Ext(f.Name()), video_ext[:])
    if (img_true) {
      images = append(images, f.Name())
    } else if (vid_true) {
      videos = append(videos, f.Name())
    }
  }


  /* kill/ disable old instances of slideshow or video_looper */
  log.Println("killing old fbi slideshow")
  cmd_killfbi := exec.Command("/bin/sh", "-c", "sudo killall fbi")
  if err := cmd_killfbi.Run(); err != nil {
    log.Println("killing slideshow failed, perhaps not running")
  }

  log.Println("killing video looper")
  cmd_killvid_str := "sudo supervisorctl stop video_looper"
  cmd_killvid_str += "&& cd /etc/supervisor/conf.d "
  cmd_killvid_str += "&& sudo mv video_looper.conf video_looper.conf.disabled"
  cmd_killvid := exec.Command("/bin/sh", "-c", cmd_killvid_str)
  if err := cmd_killvid.Run(); err != nil {
    log.Println("Error disabling video_looper, perhaps not running")
  }


  /* start new slideshow or video_looper */
  if (len(images) > 0) {
    log.Printf("%d images found, starting slideshow\n", len(images))

    var images_str string
    for _, filename := range images {
      images_str += media_folder + "/"
      images_str += filename
      images_str += " "
    }

    image_cmd := image_cmd_base + images_str
    cmd := exec.Command("/bin/sh", "-c", image_cmd)
    if err := cmd.Run(); err != nil {
      log.Println("Error starting slideshow")
      log.Println(err)
    }

  } else {
    log.Printf("%d videos found, starting videolooper\n", len(videos))

    video_cmd := "cd /etc/supervisor/conf.d "
    video_cmd += "&& sudo mv video_looper.conf.disabled video_looper.conf "
    video_cmd += "&& sudo service supervisor restart"
    cmd := exec.Command("/bin/sh", "-c", video_cmd)
    if err := cmd.Run(); err != nil {
      log.Println("Error starting video_looper")
      log.Println(err)
    }

  }

}

func printSlice(s []string) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

