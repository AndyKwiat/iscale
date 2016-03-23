package main
import (
	"iscale"
	"log"
	"time"
)

func main() {
	stime:= time.Now()
	pic := iscale.ReadPic("test.jpg");
	log.Println("read duration ",time.Since(stime).Seconds())
	stime= time.Now()
	//dstPic :=
	iscale.ResizeImage( pic, 200,200)
	log.Println("resize duration ",time.Since(stime).Seconds())

	//writePNG("test.png",dstPic );
}
