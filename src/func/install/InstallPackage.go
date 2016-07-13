package install

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"../../glpmjson"
	"os/exec"
)

func InstallPackage(glpmJsonFile string,goPackage string) {
	jsonBytes, err := ioutil.ReadFile(glpmJsonFile)
	if (err != nil) {
		log.Fatalf("could not read json file %v: %v",glpmJsonFile,err)
	}
	var glpmJsonObj = glpmjson.GlpmJson{}
	err = json.Unmarshal(jsonBytes,&glpmJsonObj)
	if (err != nil) {
		log.Fatalf("could not parse json file %v: %v",glpmJsonFile,err)
	}
	log.Printf("installing package %v...",goPackage)
	err = exec.Command("go","get",goPackage).Run()
	if (err != nil) {
		log.Fatalf("could not install package %v: %v",goPackage,err)
	} else {
		log.Printf("finished installing package %v",goPackage)
	}
	isFoundPackage := false
	for _, packageName := range glpmJsonObj.Packages {
		if !isFoundPackage && packageName == goPackage {
			isFoundPackage=true
			break
		}
	}
	if isFoundPackage {
		log.Printf("package %v already exists in %v", goPackage,glpmJsonFile)
	} else {
		glpmJsonObj.Packages=append(glpmJsonObj.Packages,goPackage)
		jsonBytes, err := json.MarshalIndent(glpmJsonObj,"","    ")
		if (err != nil) {
			log.Fatalf("internal error, could not parse modified json object")
		} else {
			err = ioutil.WriteFile(glpmJsonFile,jsonBytes,0644)
			if (err != nil) {
				log.Fatalf("could not write modified json file at %v",glpmJsonFile)
			} else {
				log.Printf("package %v added to %v",goPackage,glpmJsonFile)
			}
		}
	}

}
