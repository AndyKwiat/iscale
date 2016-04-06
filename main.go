package main
import (
	//"time"
	"iscale/builder"
	"log"
	//"fmt"
)

func errorFunc( err error ){
	log.Fatal(err)
}
func main() {
	//jobs,_ :=
	builder.Init( "testdata/builder")
	//job,_ := jobs.getJob("resizeTestImages")
	//job.run( func (){ log.Println("Done job")}, errorFunc)


	/*stime:= time.Now()
	pic := iscale.ReadPic("src1.png");
	log.Println("read duration ",time.Since(stime).Seconds())
	stime= time.Now()
	dstPic :=
	iscale.ResizeImage( pic, -1,256,16)
	log.Println("resize duration ",time.Since(stime).Seconds())

	iscale.WritePNG("test.png",dstPic );*/
	/*c := make(chan int)
	go test( 4,"p1" ,c)
	go test( 1,"p2" ,c)
	var x int
	for i:=0; i < 2; i++{
		x+=<-c
	}
	fmt.Println("main done",x)*/
}
