package remove

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"../../glpmjson"
	"sort"
)

func RemovePackage(glpmJsonFile string,goPackage string) {
	jsonBytes, err := ioutil.ReadFile(glpmJsonFile)
	if (err != nil) {
		log.Fatalf("could not read json file %v: %v", glpmJsonFile, err)
	}
	var glpmJsonObj = glpmjson.GlpmJson{}
	err = json.Unmarshal(jsonBytes, &glpmJsonObj)
	if (err != nil) {
		log.Fatalf("could not parse json file %v: %v", glpmJsonFile, err)
	}
	log.Printf("removing package %v...", goPackage)
	sort.Strings(glpmJsonObj.Packages)
	idx := sort.SearchStrings(glpmJsonObj.Packages, goPackage)
	if idx < len(glpmJsonObj.Packages) {
		copy(glpmJsonObj.Packages[idx:], glpmJsonObj.Packages[idx + 1:])
		glpmJsonObj.Packages[len(glpmJsonObj.Packages) - 1] = ""
		glpmJsonObj.Packages = glpmJsonObj.Packages[:len(glpmJsonObj.Packages) - 1]
		jsonBytes, err = json.MarshalIndent(glpmJsonObj,"","    ")
		if err != nil {
			log.Fatalf("internal error, could not parse json object: %v", err)
		}
		err = ioutil.WriteFile(glpmJsonFile, jsonBytes, 0644)
		if (err != nil) {
			log.Fatalf("could not write modified json file at %v", glpmJsonFile)
		} else {
			log.Printf("package %v removed from %v", goPackage, glpmJsonFile)
		}
	} else {
		log.Fatalf("could not find package %v in %v", goPackage, glpmJsonFile)
	}
}
