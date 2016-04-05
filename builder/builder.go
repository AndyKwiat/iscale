package builder

import (
	"encoding/json"
	"io/ioutil"
	"log"
)
type Builder struct{
	jobsFolder string
	allJobs map[string]Job //https://blog.golang.org/go-maps-in-action
}

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

//-- jobs
// addJob
// listJobs
// getJob
// updateJob
// deleteJob

//-- tools
// addTool
// listTools
// getTool
// updateTool
// deleteTool

// -- filters
// add, list, get, update, delete



func Init( jobsFolder string)(Builder,error){
	var builder Builder
	builder.jobsFolder = jobsFolder
	builder.allJobs = make(map[string]Job)
	return builder,nil
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