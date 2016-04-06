package builder
import (
	"testing"
	"fmt"
)
type MockOptions struct{
	IntOption int
	StringOption string

}
func MockTool( input,output string, optionsParam interface{}){
	options := optionsParam.(MockOptions)
	fmt.Println("input: " ,input , " output: " , output , "   -->intOpt: " , options.IntOption ,
	" stringOpt:" , options.StringOption )
}
func TestJobs( t *testing.T ){

	var j Job;
	j.ID="resize"
	j.InputFilter="testdata/images"
	j.OutputPath="testdata/output"
	j.OutputFileNamePattern = "{inputFileBase}-{#}.jpg"
	j.OptionsMap = make (map[string]interface{})
	j.OptionsMap["*"]= MockOptions{1,"Test"}
	j.OptionsMap["testdata/images/big.jpg"]=MockOptions{3,"ggoody"}
	j.Tool = MockTool
	j.Run()

}
