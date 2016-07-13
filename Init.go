package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func Init(glpmJsonFile string) {
	if _, err := os.Stat(glpmJsonFile); os.IsNotExist(err) {
		jsonData, err := json.MarshalIndent(GlpmJson{},"","    ")
		if err != nil {
			log.Panicf("could not parse empty glpm json: %v", err)
		} else {
			err = ioutil.WriteFile(glpmJsonFile, jsonData, 0644)
			if err != nil {
				log.Panicf("could not write glpm.json file: %v", err)
			} else {
				log.Println("created glpm.json file succesfully")
			}
		}
	} else {
		log.Fatalf("json file %v already exists",glpmJsonFile)
	}
}
