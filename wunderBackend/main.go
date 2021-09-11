package main

import (
	// "fmt"
	// "fmt"
	"image/color"
	"log"

	"net/http"

	"gocv.io/x/gocv"
)

// dur := 2590 //hardcoded
var a [2590]bool
var c int

var count int

func main() {
	c = 3

	go cvvv()
	http.HandleFunc("/track/", track)
	// http.HandleFunc("/prog/", prog)

	log.Fatal(http.ListenAndServe("127.0.0.1:3300", nil))

}

func cvvv() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("error opening web cam: %v", err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	window := gocv.NewWindow("webcamwindow")
	defer window.Close()

	harrcascade := "tt.xml"
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(harrcascade)
	defer classifier.Close()

	color := color.RGBA{0, 255, 0, 0}
	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the device")
			continue
		}

		rects := classifier.DetectMultiScale(img)
		for _, r := range rects {
			// fmt.Println("detected", r)
			gocv.Rectangle(&img, r, color, 3)
			c = 20
		}

		window.IMShow(img)
		window.WaitKey(10)

		if c != 0 {
			c--
		}

	}

}
