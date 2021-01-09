package main

import (
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
)

type Changeable interface {
	Set(x, y int, c color.Color)
}

func main() {
	rd, err := os.Open("img.png")
	if err != nil {
		panic(err)
	}
	defer rd.Close()
	img, _, err := image.Decode(rd)
	if err != nil {
		panic(err)
	}
	imgs := RepeatedSmoothImage(img, 200)
	ff, err := os.Create("smothed_image.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(ff, imgs)
	if err != nil {
		panic(err)
	}
}

func Scale32To8(c0 uint32) uint8 {
	cc := c0 >> 8
	c1 := uint8(cc)
	return c1
}

func SmoothImage(img0 image.Image) image.Image {
	img := img0
	bounds := img.Bounds()
	var n int64
	if cimg, ok := img.(Changeable); ok {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
				n = 1
				r, g, b, a := img.At(i, j).RGBA()
				ri := int64(r)
				gi := int64(g)
				bi := int64(b)
				ai := int64(a)
				if i != bounds.Min.X {
					n += 1
					rn, gn, bn, an := img.At(i-1, j).RGBA()
					ri += int64(rn)
					gi += int64(gn)
					bi += int64(bn)
					ai += int64(an)
				}
				if i != bounds.Max.X {
					n += 1
					rn, gn, bn, an := img.At(i+1, j).RGBA()
					ri += int64(rn)
					gi += int64(gn)
					bi += int64(bn)
					ai += int64(an)
				}
				if i != bounds.Min.Y {
					n += 1
					rn, gn, bn, an := img.At(i, j-1).RGBA()
					ri += int64(rn)
					gi += int64(gn)
					bi += int64(bn)
					ai += int64(an)
				}
				if i != bounds.Max.Y {
					n += 1
					rn, gn, bn, an := img.At(i, j+1).RGBA()
					ri += int64(rn)
					gi += int64(gn)
					bi += int64(bn)
					ai += int64(an)
				}
				cimg.Set(i, j, color.RGBA{Scale32To8(uint32(ri / n)),
					Scale32To8(uint32(gi / n)), Scale32To8(uint32(bi / n)),
					Scale32To8(uint32(ai / n))})
			}
		}
	}
	return img
}

func RepeatedSmoothImage(img image.Image, rn int) image.Image {
	for i := 0; i < rn; i++ {
		img = SmoothImage(img)
	}
	return img
}
