package builder

import (
	"os"
	"fmt"
)

type fileCache struct{
	fileMap map[string]string
}

func makeHash( tool,rootPath,filePath string)(string){
	file := rootPath+filePath
	info,_ := os.Stat(file)
	hash:= fmt.Sprintf("%s_%d_%s",string(info.ModTime().String()),info.Size(),tool)
	return hash
}
func ( fc *fileCache)AddCache( toolId,RootPath,inputFilePath,outputFilePath string, options map[string]interface{} ){
	hash := makeHash ( toolId,RootPath,inputFilePath);
	_,err := os.Stat(outputFilePath)
	if err != nil {
		return
	}
	outputHash:= makeHash(toolId,"",outputFilePath)
	if ( fc.fileMap == nil ){
		fc.fileMap = make(map[string]string)
	}
	fc.fileMap[hash] = outputHash
}

func (fc *fileCache)CheckCached( toolId,RootPath,inputFilePath,outputFilePath string, options map[string]interface{} )(bool){
	hash := makeHash ( toolId,RootPath,inputFilePath);
	outputHash := fc.fileMap[hash]
	if  outputHash == ""{
		return false
	}
	_,err := os.Stat(outputFilePath)
	if err != nil {
		return false
	}
	actualOutputCache := makeHash(toolId,"",outputFilePath)
	if actualOutputCache == outputHash {
		return true
	}
	return false

}