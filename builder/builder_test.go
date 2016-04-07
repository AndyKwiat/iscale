package builder
import (
	"testing"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)
type MockOptions struct{
	IntOption int
	StringOption string

}

var inputs []string
var outputs []string
var optionsList []MockOptions

func MockTool( input,output string, optionsParam interface{}){
	options := optionsParam.(MockOptions)
	fmt.Println("input: " ,input , " output: " , output , "   -->intOpt: " , options.IntOption ,
	" stringOpt:" , options.StringOption )
	inputs = append(inputs,input)
	outputs = append(outputs,output)
	optionsList = append(optionsList,options)
}


func TestJobs( t *testing.T ){

	var j Job;
	j.ID="resize"
	j.RootPath ="../"
	j.InputFilter="testdata/images"
	j.OutputPath="testdata/output"
	j.OutputFileNamePattern = "{inputFileBase}-{#}.jpg"
	j.OptionsMap = make (map[string]interface{})
	j.OptionsMap["."]= MockOptions{1,"Test"}
	j.OptionsMap["testdata/images/big.jpg"]=MockOptions{3,"ggoody"}
	j.Tool = MockTool
	j.Run()

	json,_ := json.MarshalIndent(j,"","    ")
	fmt.Println("json")
	fmt.Print(string(json))



	Convey("Reading input Filter and options",t,func(){
		So(inputs[0],ShouldStartWith,j.RootPath+j.InputFilter)
		So(outputs[0],ShouldStartWith,j.RootPath+j.OutputPath)
		So(optionsList[1].IntOption,ShouldEqual,1)
		So(optionsList[4].IntOption,ShouldEqual,1)
		So(optionsList[5].StringOption,ShouldEqual,"Test")
		So(optionsList[0].IntOption,ShouldEqual,3)
		So(optionsList[0].StringOption,ShouldEqual,"ggoody")


	})

}
