package imgresize

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

const DoneCode = 1000

func ResizeImage(src image.Image, newWidth int, newHeight int, threadCount int) *image.RGBA {
	newSrc := image2RGBA(src)
	if( newWidth < 0 ){// code for aspect ratio
		newWidth = src.Bounds().Dx() * newHeight /src.Bounds().Dy()
	} else if ( newHeight < 0 ){
		newHeight = src.Bounds().Dy() * newWidth / src.Bounds().Dx()
	}

	dest := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	ResizeImageRGBA(newSrc, dest, threadCount)
	return dest
}

func image2RGBA(src image.Image) *image.RGBA {
	b := src.Bounds()
	newSrc := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(newSrc, newSrc.Bounds(), src, b.Min, draw.Src)
	return newSrc
}

func resizeSubImage(src, dst *image.RGBA, sy, ey int, update chan int) {
	boundsDst := dst.Bounds()
	pixelSizeX := 1 / float32(boundsDst.Dx())
	pixelSizeY := 1 / float32(boundsDst.Dy())

	div := (ey - sy) / 10
	for y := sy; y < ey; y++ {
		if div > 0 && y%div == 0 {
			update <- (y - sy) * 100 / (ey - sy)
		}

		for x := boundsDst.Min.X; x < boundsDst.Max.X; x++ {
			sx := float32(x) / float32(boundsDst.Max.X)
			sy := float32(y) / float32(boundsDst.Max.Y)
			ex := sx + pixelSizeX
			ey := sy + pixelSizeY
			dst.Set(x, y, SampleRGBAArea(src, sx, sy, ex, ey))
		}
	}
	update <- DoneCode
}
func ResizeImageRGBA(src *image.RGBA, dst *image.RGBA, threadCount int) {
	boundsDst := dst.Bounds()

	channel := make(chan int)
	ht := boundsDst.Dy()
	stride := ht / threadCount
	for i := 0; i < threadCount; i++ {
		sy := i * stride
		ey := i*stride + stride
		go resizeSubImage(src, dst, sy, ey, channel)
	}
	go resizeSubImage(src, dst, stride*threadCount, ht, channel) // cleanup the left over lines

	doneCount := 0
	var maxPct int
	for {
		pct := <-channel
		if pct == DoneCode {
			doneCount++
			if doneCount >= threadCount+1 {
				break
			}
		} else {
			if pct > maxPct {
				maxPct = pct
				log.Println("pct ", pct)
			}
		}

	}

}

func CalcEdgeFraction(point, start, end float32) float32 {
	var frac float32 = 1.0
	if point-1 < start { // first iteration
		frac -= point - float32(int(point))
	}
	if point+1 >= end { // last iteration
		frac -= 1 - (end - float32(int(point)))
	}

	return frac
}
func SampleRGBAArea(src *image.RGBA, startX, startY, endX, endY float32) color.RGBA {
	sx := startX * float32(src.Bounds().Max.X)
	sy := startY * float32(src.Bounds().Max.Y)
	ex := endX * float32(src.Bounds().Max.X)
	ey := endY * float32(src.Bounds().Max.Y)
	//log.Println("sx ",sx," ex",ex)
	var sumr, sumg, sumb, suma float32
	var numSamples float32

	for y := sy; y < ey; y++ {
		// check for edges
		sampleFracY := CalcEdgeFraction(y, sy, ey)
		for x := sx; x < ex; x++ {
			sampleFracX := CalcEdgeFraction(x, sx, ex)

			clr := src.At(int(x), int(y))
			r, g, b, a := clr.RGBA()
			frac := sampleFracY * sampleFracX
			//log.Println("x ",x, " y",y, " r",r," g",g," b",b," a",a," frac",frac," fx",sampleFracX, " fy",sampleFracY )
			sumr += float32(r) * frac
			sumg += float32(g) * frac
			sumb += float32(b) * frac
			suma += float32(a) * frac
			numSamples += frac
		}
	}
	//log.Println("sumr ",sumr )
	//log.Println("numsamples: ",numSamples)
	newColor := color.RGBA{uint8(uint32(sumr/numSamples) >> 8), uint8(uint32(sumg/numSamples) >> 8), uint8(uint32(sumb/numSamples) >> 8), uint8(uint32(suma/numSamples) >> 8)}
	//log.Println("color",newColor)
	return newColor
}
func WritePNG(filename string, pic image.Image) {
	writer, err := os.Create(filename)
	defer writer.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(writer, pic)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("PNG written: ", filename)
}
func ReadPic(filename string) image.Image {
	reader, err := os.Open(filename)
	defer reader.Close()
	if err != nil {
		log.Fatal(err)
	}
	pic, format, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(format, " pic read: ", pic.Bounds())
	return pic
}

/*func readJPEG( filename string ) image.Image {
	reader,err := os.Open(filename)
	defer reader.Close()
	if err != nil{
		log.Fatal(err)
	}
	pic,err := jpeg.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("JPEG read: ",pic.Bounds() )
	return pic
}*/
