// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		x, err := qToF(query, "x")
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		y, err := qToF(query, "y")
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		zoom, err := qToF(query, "zoom")
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		encode(rw, x, y, zoom)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func qToF(query url.Values, key string) (float64, error) {
	xstr := query.Get(key)
	x, err := strconv.ParseFloat(xstr, 64)
	return x, err
}

func encode(w http.ResponseWriter, x, y, zoom float64) {
	m := math.Exp2(1 - zoom)
	xmin, xmax := x-m, x+m
	ymin, ymax := y-m, y+m
	const (
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
