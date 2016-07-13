package glpm

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"os/exec"
)

func InstallAllPackage(glpmJsonFile string) {
	jsonBytes, err := ioutil.ReadFile(glpmJsonFile)
	if (err != nil) {
		log.Fatalf("could not read json file %v: %v",glpmJsonFile,err)
	}
	var glpmJsonObj = GlpmJson{}
	err = json.Unmarshal(jsonBytes,&glpmJsonObj)
	if (err != nil) {
		log.Fatalf("could not parse json file %v: %v",glpmJsonFile,err)
	}
	for _,goPackage := range glpmJsonObj.Packages {
		log.Printf("installing package %v...", goPackage)
		err = exec.Command("go", "get" ,goPackage).Run()
		if (err != nil) {
			log.Fatalf("could not install package %v: %v", goPackage, err)
		} else {
			log.Printf("finished installing package %v", goPackage)
		}
	}
	log.Printf("finished installing packages from %v",glpmJsonFile)
}
