package builder
import (
	"testing"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

var inputs []string
var outputs []string
var optionsList []map[string]interface{}

func MockTool( input,output string, options map[string]interface{}){

	fmt.Println("input: " ,input , " output: " , output , "   -->intOpt: " , options["IntOption"] ,
	" stringOpt:" , options["StringOption"] )
	inputs = append(inputs,input)
	outputs = append(outputs,output)
	optionsList = append(optionsList,options)
}


func TestJobs( t *testing.T ){

	var j Job;
	rootPath := "../"
	j.ID="resize"
	j.InputFilter="testdata/images"
	j.OutputPath="testdata/output"
	j.OutputFileNamePattern = "{inputFileBase}-{#}.jpg"
	j.OptionsMap = make (map[string]map[string]interface{})
	j.OptionsMap["."]= map[string]interface{}{"IntOption":1,"StringOption":"Test"}
	j.OptionsMap["testdata/images/big.jpg"]=map[string]interface{}{"IntOption":3}

	j.Run(rootPath, MockTool)

	json,_ := json.MarshalIndent(j,"","    ")
	fmt.Println("json")
	fmt.Print(string(json))



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
