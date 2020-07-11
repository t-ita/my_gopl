package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // キャンバスの大きさ（画素数）
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲（-xyrange .. +xyrange)
	xyscale       = width / 2 / xyrange // ｘ 単位 および y 単位あたりの画素数
	zscale        = height * 0.4        // z単位当たりの画素数
	angle         = math.Pi / 6         // x, y 軸の角度（=30度）

	red  = 0xff0000
	blue = 0x0000ff
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度）, cos(30度）

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	zmax := 0.0
	zmin := 0.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, az, ok := point(i+1, j)
			if !ok {
				continue
			}
			_, _, bz, ok := point(i, j)
			if !ok {
				continue
			}
			_, _, cz, ok := point(i, j+1)
			if !ok {
				continue
			}
			_, _, dz, ok := point(i+1, j+1)
			if !ok {
				continue
			}
			zavg := (az + bz + cz + dz) / 4
			if i == 0 && j == 0 {
				zmax = zavg
				zmin = zavg
			}
			zmax = math.Max(zmax, zavg)
			zmin = math.Min(zmin, zavg)
		}
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var zsum float64
			x, y, z, ok := point(i+1, j)
			if !ok {
				continue
			}
			ax, ay := corner(x, y, z)
			zsum += z

			x, y, z, ok = point(i, j)
			if !ok {
				continue
			}
			bx, by := corner(x, y, z)
			zsum += z

			x, y, z, ok = point(i, j+1)
			if !ok {
				continue
			}
			cx, cy := corner(x, y, z)
			zsum += z

			x, y, z, ok = point(i+1, j+1)
			if !ok {
				continue
			}
			dx, dy := corner(x, y, z)
			zsum += z

			color := color(zsum/4, zmax, zmin)
			fmt.Printf("<polygon style='fill: #%06x;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func color(z, zmax, zmin float64) uint32 {
	zRange := zmax - zmin
	zPos := z - zmin
	return uint32((red-blue)*zPos/zRange) + blue
}

func point(i, j int) (x, y, z float64, ok bool) {
	// マス目のかどの点 （x, y) をみつける
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)

	// 面の高さ z を計算する
	z = f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}

	ok = true
	return
}

func corner(x, y, z float64) (sx float64, sy float64) {
	// (x, y, z) を 2-D SVG キャンバス (sx, sy) へ等角的に投影
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0) からの距離
	return math.Sin(r) / r
}
