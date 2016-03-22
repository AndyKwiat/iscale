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
	b:= pic.Bounds()
	m:=image.NewRGBA(image.Rect(0,0,b.Dx(),b.Dy()));
	draw.Draw(m,m.Bounds(),pic,b.Min,draw.Src)
	m.Set(10,10,color.Black )
	writePNG("test.png",m );
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