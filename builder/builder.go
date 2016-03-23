package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	SaveJob( myJob )
	job,_ := LoadJob( "Job1")
	log.Print(job)

}
func getFileNameForJobId( id string)string{
	return "jobs/"+id+".job.json";
}
func LoadJob( jobId string) (Job,error){
	j ,err := ioutil.ReadFile(getFileNameForJobId( jobId))
	if err == nil {
		var myJob2 Job
		err = json.Unmarshal(j, &myJob2)
		if ( err != nil ){
			log.Print(err)
			return myJob2,err
		}
		return myJob2,nil
	}else {
		log.Print(err)
		var empty Job
		return empty,err
	}

}

func SaveJob( job Job ) error{
	log.Println("saving job ",job.ID)
	j,err := json.MarshalIndent(job,"","    ")
	if err != nil{
		log.Print(err)
		return err
	}
	err = ioutil.WriteFile(getFileNameForJobId(job.ID),j,0644)
	if err != nil{
		log.Print(err)
		return err
	}
	return nil
}