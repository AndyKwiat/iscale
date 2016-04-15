package builder
import (
	"testing"
	"os"
	"io/ioutil"
	. "github.com/smartystreets/goconvey/convey"
)
func TestCache( t *testing.T ) {
	Convey("file cache", t, func() {
		var cacher fileCache
		inputFile := "images/src1.png"
		outputFile := "../testdata/cache.txt"
		tool := "myTool"
		root := "../testdata/"
		mergedOptions := map[string]interface{}{"MyOption":1,"StringOption":"Test"}
		os.Remove(outputFile)
		cached:= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
		So(cached,ShouldBeFalse)
		ioutil.WriteFile(outputFile,[]byte("test"),0644)
		cacher.AddCache(tool,root,inputFile,outputFile,mergedOptions)
		cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
		So(cached,ShouldBeTrue)
		os.Remove(outputFile)
		cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
		So(cached,ShouldBeFalse)
		ioutil.WriteFile(outputFile,[]byte("test"),0644)
		cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
		So(cached,ShouldBeFalse)
		cacher.AddCache(tool,root,inputFile,outputFile,mergedOptions)
		cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
		So(cached,ShouldBeTrue)

		Convey("new tool",func(){
		tool = "newTool"
		cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
		So(cached,ShouldBeFalse)})

		Convey("new option",func(){
			mergedOptions["MyOption"]=2
			cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
			So(cached,ShouldBeFalse)
			mergedOptions["MyOption"]=1
			cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
			So(cached,ShouldBeTrue)
		})
		Convey("new input",func(){
			inputFile := "images/src2.png"
			cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
			So(cached,ShouldBeFalse)
			inputFile = "images/src1.png"
			cached= cacher.CheckCached(tool,root,inputFile,outputFile,mergedOptions)
			So(cached,ShouldBeTrue)})


	})
}