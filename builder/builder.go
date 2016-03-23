package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type JobStep struct {
	CommandLine string
}
type Job struct{
	Name string
	InputFilter string
	OutputPath string
	Steps []JobStep
	ID string
}

func main() {
	myJob := Job{ "Resize Stuff",
		"test.jpg",
		"output",
		[]JobStep{JobStep{"main"}},"Job1"}
	j,_ := json.MarshalIndent(myJob,"","    ")
	err := ioutil.WriteFile("jobs/"+myJob.ID+".job.json",j,0644)
	if err != nil {
		fmt.Println("error writing")
	}
	j ,err = ioutil.ReadFile("myjob.json")
	if err == nil {
		var myJob2 Job
		json.Unmarshal(j, &myJob2)
		fmt.Println(myJob2)
	}

}
