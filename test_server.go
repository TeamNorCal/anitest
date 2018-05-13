package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"net/http"
	"time"

	"github.com/TeamNorCal/animation"
)

// A server to expose animation frame data over a JSON interface

// Response encapsulates a frame data response
type Response struct {
	Universe int
	Data     []color.RGBA
}

var sr = animation.NewSequenceRunner([]uint{30})

func writeFrame(w io.Writer) {
	sr.ProcessFrame(time.Now())
	strand := sr.UniverseData(0)
	data := Response{0, strand}
	//			[]color.RGBA{color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}}
	ser, _ := json.Marshal(data)
	w.Write(ser)
}

func main() {
	http.HandleFunc("/ello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Wello, Horld!")
	})
	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		sr.InitSequence(Test1, time.Now())
		//		writeFrame(w)
	})
	http.HandleFunc("/getFrame", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.Header().Add("Cache-Control", "no-store")
		writeFrame(w)
	})
	http.Handle("/", http.FileServer(http.Dir("assets/static")))
	http.ListenAndServe(":8080", nil)
}
