package builder

import (
	"os"
	"fmt"
	"encoding/gob"
	"bytes"
	"encoding/json"
	//"strings"
)

type fileCache struct{
	fileMap map[string]string
}

func makeHash( tool,rootPath,filePath string, options interface{})(string){
	j,_ := json.Marshal(options)
	file := rootPath+filePath
	info,_ := os.Stat(file)
	j2 := string(j[:])
	hash:= fmt.Sprintf("%s_%d_%s_%s",string(info.ModTime().String()),info.Size(),tool,j2)
	//fmt.Println(hash)
	return hash
}
func ( fc *fileCache)AddCache( toolId,RootPath,inputFilePath,outputFilePath string, options map[string]interface{} ){

	//fmt.Println("add cache start")
	hash := makeHash ( toolId,RootPath,inputFilePath,options);
	_,err := os.Stat(outputFilePath)
	if err != nil {
		//fmt.Println("add cach fail")
		return
	}
	outputHash:= makeHash(toolId,"",outputFilePath,"")
	if ( fc.fileMap == nil ){
		fc.fileMap = make(map[string]string)
	}
	fc.fileMap[hash] = outputHash
	//fmt.Println("add cache success")
}

func (fc *fileCache)CheckCached( toolId,RootPath,inputFilePath,outputFilePath string, options map[string]interface{} )(bool){
	//fmt.Println("startcache")
	hash := makeHash ( toolId,RootPath,inputFilePath,options);
	outputHash := fc.fileMap[hash]
	if  outputHash == ""{
	//	fmt.Println("blank outputshash")
		return false
	}
	_,err := os.Stat(outputFilePath)
	if err != nil {
	//	fmt.Println("outputfile no exist")
		return false
	}
	actualOutputCache := makeHash(toolId,"",outputFilePath,"")
	if actualOutputCache == outputHash {
	//	fmt.Println("passed output")
		return true
	}
	//fmt.Println("fail7")
	return false

}


func getBytes(key interface{}) ([]byte ) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}