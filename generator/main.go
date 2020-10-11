package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/murphynet/wolfcolor"
)

var (
	size = flag.Int("size", 16, "size (in pixels) of each color square")
	out  = flag.String("out", "palette.png", "filename for the png output")
)

func main() {
	pal := wolfcolor.Wolf3d

	max := int(math.Ceil(math.Sqrt(float64(len(pal))))) * *size
	log.Println(max)

	img := image.NewNRGBA(image.Rect(0, 0, max, max))

	for y := 0; y < max; y++ {
		cy := y / *size
		for x := 0; x < max; x++ {
			cx := x / *size
			colorIndex := cx + (cy * (max / *size))
			img.Set(x, y, pal[colorIndex])
		}
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}

	// Print the image with lossless png compression.
	enc := png.Encoder{
		CompressionLevel: png.BestCompression,
	}
	err = enc.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
	log.Println("done")
}
