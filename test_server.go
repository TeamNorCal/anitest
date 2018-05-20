package main

import (
	"encoding/json"
	"html/template"
	"image/color"
	"io"
	"math/rand"
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
	// Universes []Universe
	Resonators []Universe
	Tower      []Universe
}

// var universeSizes = []uint{30, 60, 15}
var universeSizes []uint //= []uint{30, 30, 30, 30, 30, 30, 30, 30}

// var sr = animation.NewSequenceRunner(universeSizes)
var p = animation.NewPortal()
var status *animation.PortalStatus

func init() {
	universeSizes = make([]uint, 24)
	for idx := 0; idx < 24; idx++ {
		universeSizes[idx] = 30
	}
}

func initPortalStatus(resoLevels []int) {
	resoStatus := make([]animation.ResonatorStatus, 8)
	for idx := range resoStatus {
		resoStatus[idx] = animation.ResonatorStatus{
			Health: 100.0,
			Level:  resoLevels[idx],
		}
	}

	status = &animation.PortalStatus{
		Faction:    animation.ENL,
		Level:      8,
		Resonators: resoStatus,
	}
}

func randomizeAResonator() {
	resoNum := rand.Intn(8)
	resoLevel := rand.Intn(9)
	status.Resonators[resoNum].Level = resoLevel
}

func randomizeLevel() {
	status.Level = rand.Float32() * 8.0
}

// NTimes is a custom template function that creates a slice of nothing to range across
func NTimes(count int) []struct{} {
	return make([]struct{}, count)
}

func writeFrame(w io.Writer) {
	// sr.ProcessFrame(time.Now())
	frameData := p.GetFrame(time.Now())
	datas := make([]UniverseData, 0, 24)
	for id := range universeSizes {
		datas = append(datas, UniverseData{id, frameData[id].Data})
	}
	resp := Response{datas}
	//			[]color.RGBA{color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}}
	ser, _ := json.Marshal(resp)
	w.Write(ser)
}

func getIndexData() IndexData {
	resos := make([]Universe, 0)
	tower := make([]Universe, 0)
	for id, size := range universeSizes {
		if id < 8 {
			resos = append(resos, Universe{id, int(size)})
		} else {
			// Build these in reverse order so first level is at the bottom in the simulator
			tower = append([]Universe{Universe{id, int(size)}}, tower...)
		}
	}
	return IndexData{Resonators: resos, Tower: tower}
}

func renderIndex(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index").Funcs(template.FuncMap{"ntimes": NTimes}).ParseGlob("assets/templates/index/*.gohtml"))
	t.ExecuteTemplate(w, "index", getIndexData())
}

func main() {
	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		// sr.InitSequence(CreateTest3(), time.Now())
		initPortalStatus([]int{1, 2, 3, 4, 5, 6, 7, 8})
		p.UpdateStatus(status)
		ticker := time.NewTicker(5 * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					randomizeAResonator()
					randomizeLevel()
					p.UpdateStatus(status)
				}
			}
		}()
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
