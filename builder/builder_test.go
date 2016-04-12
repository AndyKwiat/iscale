package builder
import (
	"testing"
	//"encoding/json"
	//"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

var inputs []string
var outputs []string
var optionsList []map[string]interface{}

func MockTool( input,output string, options map[string]interface{}){

	//fmt.Println("input: " ,input , " output: " , output , "   -->intOpt: " , options["IntOption"] ," stringOpt:" , options["StringOption"] )
	inputs = append(inputs,input)
	outputs = append(outputs,output)
	optionsList = append(optionsList,options)
}


// if file exists single.jpg delete single.jpg
// test single, file exists (maybe use another mocktool)
// run again test caching (job returns # of caches) // cache hash: tool_mergedoptions_filesize_filehash
func clearResult(){
	inputs = nil
	outputs = nil
	optionsList = nil
}

func TestMultipleFiles( t *testing.T ){
	clearResult()
	var j Job
	rootPath := "../testdata/"
	j.ID="resize"
	j.InputFilter="images/src1.png;images/src2.png"
	j.OutputPath="output"
	j.OutputFileNamePattern = "{inputFileBase}-{#}.jpg"
	j.OptionsMap = make (map[string]map[string]interface{})
	j.OptionsMap["."]= map[string]interface{}{"IntOption":1,"StringOption":"Test"}
	j.OptionsMap["images"]= map[string]interface{}{"IntOption":2}
	j.OptionsMap["images/src2.png"]=map[string]interface{}{"StringOption":"Test2"}
	j.Run(rootPath,MockTool);
	Convey("test two", t, func() {
		So(len(outputs),ShouldEqual,2)
		So(outputs[0],ShouldEqual,rootPath+j.OutputPath+"/src1-0.jpg")
		So(outputs[1],ShouldEqual,rootPath+j.OutputPath+"/src2-1.jpg")
		So(optionsList[0]["IntOption"],ShouldEqual,2)
		So(optionsList[0]["StringOption"],ShouldEqual,"Test")
		So(optionsList[1]["IntOption"],ShouldEqual,2)
		So(optionsList[1]["StringOption"],ShouldEqual,"Test2")

	})
}
func TestSingle( t *testing.T){
	clearResult()
	var j Job;
	rootPath := "../testdata/"
	j.ID="resize"
	j.InputFilter="images/src1.png"
	j.OutputPath="output"
	j.OutputFileNamePattern = "{inputFileBase}-{#}.jpg"
	j.OptionsMap = make (map[string]map[string]interface{})
	j.OptionsMap["."]= map[string]interface{}{"IntOption":1,"StringOption":"Test"}
	j.Run(rootPath,MockTool);
	Convey("test single", t, func() {
		So(len(outputs),ShouldEqual,1)
		So(outputs[0],ShouldEqual,rootPath+j.OutputPath+"/src1-0.jpg")
	})
}
func TestJobs( t *testing.T ){
	clearResult()
	var j Job;
	rootPath := "../testdata/"
	j.ID="resize"
	j.InputFilter="images"
	j.OutputPath="output"
	j.OutputFileNamePattern = "{inputFileBase}-{#}.jpg"
	j.OptionsMap = make (map[string]map[string]interface{})
	j.OptionsMap["."]= map[string]interface{}{"IntOption":1,"StringOption":"Test"}
	j.OptionsMap["images/big.jpg"]=map[string]interface{}{"IntOption":3}

	j.Run(rootPath, MockTool)

/*	json,_ := json.MarshalIndent(j,"","    ")
	fmt.Println("json")
	fmt.Print(string(json))
*/
	Convey("Reading input Filter and options",t,func(){
		So(inputs[0],ShouldStartWith,rootPath+j.InputFilter)
		So(outputs[0],ShouldStartWith,rootPath+j.OutputPath)
		So(optionsList[1]["IntOption"],ShouldEqual,1)
		So(optionsList[4]["IntOption"],ShouldEqual,1)
		So(optionsList[5]["StringOption"],ShouldEqual,"Test")
		So(optionsList[0]["IntOption"],ShouldEqual,3)
		So(optionsList[0]["StringOption"],ShouldEqual,"Test") //cascading

		// file output
		So(outputs[0],ShouldEndWith,"big-0.jpg")


	})

}
