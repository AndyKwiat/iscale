package builder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"strconv"
	//"iscale/imgresize"
	"os"
)
type Builder struct{
	jobsFolder string
	allJobs map[string]Job //https://blog.golang.org/go-maps-in-action
}

type ToolFunc func(inputFile,outputFile string, options map[string]interface{})

//http://play.golang.org/p/FL5eEgSXTT
type Job struct{
	ID string
	ToolID string
	InputFilter string
	OutputPath string
	OutputFileNamePattern string
	OptionsMap map[string]map[string]interface{}
}

func (job *Job ) Run(RootPath string, tool ToolFunc, cacher *fileCache ){
	files:= parseInputFilter( RootPath,job.InputFilter)
	for i,f := range files{
		filename:= parsePattern(i,f,job.OutputFileNamePattern )
		var filePath string= RootPath+job.OutputPath + "/" + filename

		// figure out cascading options, we go through to the root folder and add any options along the way
		optPath := f
		var optionStack []map[string]interface{}
		for  {
			options := job.OptionsMap[optPath]
			if ( options != nil ) {
				optionStack = append([]map[string]interface{}{options},optionStack...)	// prepend
			}
			if ( optPath == ".") {
				break;
			}
//			log.Println(optPath)
			optPath = path.Dir(optPath)
		}
		mergedOptions := make(map[string]interface{})
		for _,v := range optionStack{
			for key,value := range v{
				mergedOptions[key] = value
			}
		}
		if ( !cacher.CheckCached( job.ToolID,RootPath, f, filePath,mergedOptions ) ) {
			tool(RootPath + f, filePath, mergedOptions)
			cacher.AddCache(job.ToolID,RootPath, f, filePath,mergedOptions )
		}
	}
}

func checkCache( toolId,RootPath,inputFilePath,outputFilePath string, options map[string]interface{} )(bool){

	return false;
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
func parseInputFilter( root,filter string ) (fileList []string ){
	var files []string;
	filters := strings.Split(filter,";")
	for _,filterToken := range filters {
		fullPath := root + filterToken
		stat, err := os.Stat(fullPath)
		if ( err != nil ) {
			log.Println(err)
		} else {
			if ( stat.Mode().IsDir()) {
				fileInfos, err := ioutil.ReadDir(root + filter)

				if ( err == nil ) {
					for _, f := range fileInfos {
						if ( !f.IsDir() ) {
							files = append(files, filterToken + "/" + f.Name())
						}
					}
				} else {
					log.Println(err)
				}
			} else {
				files = append(files, filterToken)
			}
		}

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