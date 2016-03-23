package main

import(
	_ "image/jpeg"
	"image/png"
	"image"
	"log"
	"os"
	"image/draw"
	"image/color"
)

func main() {

	pic := readPic("castle.png");
	dstPic := resizeImage( pic, 100,100)
	dstPic.Set(10,10,color.Black )
	writePNG("test.png",dstPic );
}
func resizeImage( src image.Image, newWidth int, newHeight int ) *image.RGBA {
	newSrc := image2RGBA(src );
	dest:= image.NewRGBA( image.Rect(0,0,newWidth,newHeight))
	resizeImageRGBA(newSrc,dest)
	return dest
}

func image2RGBA( src image.Image ) *image.RGBA{
	b:= src.Bounds()
	newSrc:=image.NewRGBA(image.Rect(0,0,b.Dx(),b.Dy()))
	draw.Draw(newSrc,newSrc.Bounds(),src,b.Min,draw.Src)
	return newSrc;
}

func resizeImageRGBA( src *image.RGBA, dst *image.RGBA ){
	//boundsSrc := src.Bounds()
	boundsDst := dst.Bounds()
	//xScalar := float32(boundsSrc.Dx()) / float32(boundsDst.Dx())
	//yScalar := float32(boundsSrc.Dy()) / float32(boundsDst.Dy())

	for  y:= boundsDst.Min.Y; y < boundsDst.Max.Y; y++ {
		for x:= boundsDst.Min.X; x< boundsDst.Max.X; x++ {
			dst.Set(x,y,color.RGBA{255,255,0,255})
		}
	}
}
func CalcEdgeFraction( point, start, end float32 ) float32 {
	var frac float32 = 1.0
	if  point -1 < start {	// first iteration
		frac -= point - float32(int(point))
	}
	if  point+1 >=end {	// last iteration
		frac -= 1 - (end - float32(int(point)) )
	}
	return frac
}
func SampleRGBAArea ( src *image.RGBA, startX,startY,endX, endY float32) color.RGBA {
	sx := startX * float32(src.Bounds().Max.X)
	sy := startY * float32(src.Bounds().Max.Y)
	ex := endX * float32(src.Bounds().Max.X)
	ey := endY * float32(src.Bounds().Max.Y)
	//log.Println("sx ",sx," ex",ex)
	var sumr,sumg,sumb,suma float32
	var numSamples float32

	for y:= sy; y < ey; y ++ {
		// check for edges
		sampleFracY := CalcEdgeFraction( y, sy, ey )
		for x:= sx; x < ex; x++ {
			sampleFracX := CalcEdgeFraction( x, sx, ex )

			clr := src.At(int(x),int(y))
			r,g,b,a := clr.RGBA()
			frac := sampleFracY * sampleFracX
			//log.Println("x ",x, " y",y, " r",r," g",g," b",b," a",a," frac",frac," fx",sampleFracX, " fy",sampleFracY )
			sumr += float32(r)*frac
			sumg += float32(g)*frac
			sumb += float32(b)*frac
			suma += float32(a)*frac
			numSamples += frac
		}
	}
	//log.Println("sumr ",sumr )
	//log.Println("numsamples: ",numSamples)
	newColor := color.RGBA{uint8(uint32(sumr/numSamples)>>8),uint8(uint32(sumg/numSamples)>>8),uint8(uint32(sumb/numSamples)>>8),uint8(uint32(suma/numSamples)>>8)}
	//log.Println("color",newColor)
	return newColor
}
func writePNG( filename string , pic image.Image){
	writer,err := os.Create(filename)
	defer writer.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(writer,pic)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("PNG written: ",filename)
}
func readPic( filename string ) image.Image {
	reader,err := os.Open(filename)
	defer reader.Close()
	if err != nil{
		log.Fatal(err)
	}
	pic,format,err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(format," pic read: ",pic.Bounds() )
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