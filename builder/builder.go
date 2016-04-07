package builder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"strconv"
	//"iscale/imgresize"
)
type Builder struct{
	jobsFolder string
	allJobs map[string]Job //https://blog.golang.org/go-maps-in-action
}

type ToolFunc func(inputFile,outputFile string, options interface{})


type Job struct{
	InputFilter string
	OutputPath string
	OutputFileNamePattern string
	ID string
	OptionsMap map[string]interface{}
	Tool ToolFunc
	RootPath string
}

func (job *Job ) Run(){
	files:= parseInputFilter( job.RootPath+job.InputFilter)
	for i,f := range files{
		filename:= parsePattern(i,f,job.OutputFileNamePattern )
		var filePath string= job.RootPath+job.OutputPath + "/" + filename
		job.Tool(f,filePath,job.OptionsMap["*"])
	}
}
func parsePattern(number int, inputFile string, pattern string )(string){
	index := strings.LastIndex(inputFile,path.Ext(inputFile))
	inputFileBase := inputFile
	if (index >=0 ){
		inputFileBase = inputFile[:index]
	}

	inputFileBase = path.Base(inputFileBase)

	file:= strings.Replace(pattern,"{inputFileBase}",inputFileBase,-1)
	file= strings.Replace(file,"{#}",strconv.Itoa(number),-1)
	return file
}
func parseInputFilter( filter string ) (fileList []string ){
	fileInfos,err := ioutil.ReadDir(filter)
	var files []string;
	if ( err == nil ){
		for _,f := range fileInfos{
			if ( !f.IsDir() ){
				files = append(files,filter+"/"+f.Name())
			}
		}
	} else {
		log.Println(err)
	}



	return files
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




func Init( jobsFolder string)(*Builder,error){
	return nil,nil
}

func (b *Builder) getFileNameForJobId( id string)string{
	return b.jobsFolder+"/jobs/"+id+".job.json";
}

func (b *Builder)LoadJob( jobId string) (Job,error){
	j ,err := ioutil.ReadFile(b.getFileNameForJobId( jobId))
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

func (b *Builder)SaveJob( job Job ) error{
	log.Println("saving job ",job.ID)
	j,err := json.MarshalIndent(job,"","    ")
	if err != nil{
		log.Print(err)
		return err
	}
	err = ioutil.WriteFile(b.getFileNameForJobId(job.ID),j,0644)
	if err != nil{
		log.Print(err)
		return err
	}
	return nil
}