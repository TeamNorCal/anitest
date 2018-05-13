package main

import (
	"encoding/json"
	"html/template"
	"image/color"
	"io"
	"net/http"
	"time"

	"github.com/TeamNorCal/animation"
)

// A server to expose animation frame data over a JSON interface

// UniverseData contains data for a universe for a frame, sent to the browser
type UniverseData struct {
	ID   int
	Data []color.RGBA
}

// Response encapsulates a frame data response
type Response struct {
	Universes []UniverseData
}

// Universe defines a universe from the perspective of the simulator
type Universe struct {
	ID   int
	Size int
}

// IndexData contains data to populate the index template
type IndexData struct {
	Universes []Universe
}

var sr = animation.NewSequenceRunner([]uint{30, 60})

// NTimes is a custom template function that creates a slice of nothing to range across
func NTimes(count int) []struct{} {
	return make([]struct{}, count)
}

func writeFrame(w io.Writer) {
	sr.ProcessFrame(time.Now())
	data := Response{[]UniverseData{UniverseData{0, sr.UniverseData(0)}, UniverseData{1, sr.UniverseData(1)}}}
	//			[]color.RGBA{color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}}
	ser, _ := json.Marshal(data)
	w.Write(ser)
}

func getIndexData() IndexData {
	return IndexData{[]Universe{Universe{0, 30}, Universe{1, 60}}}
}

func renderIndex(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index").Funcs(template.FuncMap{"ntimes": NTimes}).ParseGlob("assets/templates/index/*.gohtml"))
	t.ExecuteTemplate(w, "index", getIndexData())
}

func main() {
	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		sr.InitSequence(Test1, time.Now())
		//		writeFrame(w)
	})
	http.HandleFunc("/getFrame", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.Header().Add("Cache-Control", "no-store")
		writeFrame(w)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets/static"))))
	http.HandleFunc("/index.html", renderIndex)
	http.Handle("/", http.RedirectHandler("/index.html", http.StatusFound))
	http.ListenAndServe(":8080", nil)
}