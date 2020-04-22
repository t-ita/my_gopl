package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	// http://localhost:8000/?cycles=5&res=0.001&size=100&nframes=64&delay=8 のように指定しないといけないのでデフォルト値を入れたい
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.FormValue("cycles")) // 発信器 x が完了する周回の回数
		if err != nil {
			log.Print(err)
			return
		}
		res, err := strconv.ParseFloat(r.FormValue("res"), 64) // 回転の分解能
		if err != nil {
			log.Print(err)
			return
		}
		size, err := strconv.Atoi(r.FormValue("size")) // 画像キャンパスは[-size..+size] の範囲を扱う
		if err != nil {
			log.Print(err)
			return
		}
		nframes, err := strconv.Atoi(r.FormValue("nframes")) // アニメーションのフレーム数
		if err != nil {
			log.Print(err)
			return
		}
		delay, err := strconv.Atoi(r.FormValue("delay")) // 10ms 単位でのフレーム間の遅延
		if err != nil {
			log.Print(err)
			return
		}
		lissajous(w, cycles, res, size, nframes, delay)
	}) //個々のリクエストに対してhandlerが呼ばれる
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // パレットの最初の色
	blackIndex = 1 // パレットの次の色
)

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0 // 発信器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)*0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}
