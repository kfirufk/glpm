package glpm


import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"io/ioutil"
	"regexp"
	"encoding/json"
	"sort"
)


var findSingles = regexp.MustCompile("[^\\w]import\\s*\"(.*?)\"")
var findContainer = regexp.MustCompile("[^\\w]import\\s*\\(([^)]+)")
var findSpaces = regexp.MustCompile("(\\\")?\\s+(\\\")?")
var validUrl = regexp.MustCompile("^(?:[-A-Za-z0-9]+\\.)+[A-Za-z]{2,6}$")
var dependencies = map[string]uint{}


func addValidImport(importName string) {
	log.Printf("import name %v", importName)
	if (!strings.HasPrefix(importName, "../")) {
		if idx := strings.Index(importName, "/"); idx > -1 {
			url := importName[:idx]
			log.Printf("url: %v", url)
			if validUrl.MatchString(url) {
				dependencies[importName]++
			}
		}
	}
}

func visit(path string, f os.FileInfo, err error) error {
	base := filepath.Base(path)
	if f.IsDir() {
		if strings.HasPrefix(base,".") {
			return filepath.SkipDir
		}
	} else if strings.HasSuffix(base,".go") && !strings.HasPrefix(base,".") {
		srcCode, err := ioutil.ReadFile(path)
		if (err != nil) {
			log.Panicf("could not read go file %v: %v",path,err)
		} else {
			codeStr := string(srcCode)
			results1 := findSingles.FindAllStringSubmatch(codeStr,-1)
			results2 := findContainer.FindAllStringSubmatch(codeStr,-1)
			log.Printf("Visited: %s", path)
			for _, singleSlice := range results1 {
				for _, singleMatch := range singleSlice[1:] {
					addValidImport(singleMatch)
				}
			}
			for _, containerSlice := range results2 {
				for _, containerMatch := range containerSlice[1:] {
					containerParsed := strings.TrimSpace(findSpaces.ReplaceAllString(containerMatch," "))
					for _, singleImport := range strings.Split(containerParsed," ") {
						addValidImport(singleImport)
					}
				}
			}
		}
	}
	return nil
}

func AutoAddDependencies(glpmJsonFile string,path string) {
	log.Printf("about to find dependencies from project's root directory at: %v",path)
	filepath.Walk(path,visit)
	jsonBytes, err := ioutil.ReadFile(glpmJsonFile)
	if (err != nil) {
		log.Fatalf("could not read json file %v: %v",glpmJsonFile,err)
	}
	glpmJsonObj := GlpmJson{}
	err = json.Unmarshal(jsonBytes,&glpmJsonObj)
	if (err != nil) {
		log.Fatalf("could not parse json file: %v: %v",glpmJsonFile,err)
	}
	sort.Strings(glpmJsonObj.Packages)
	for goPackage,goPackageUsageCount := range dependencies {
		log.Printf("found package %v - required %v times",goPackage,goPackageUsageCount)
		if idx := sort.SearchStrings(glpmJsonObj.Packages,goPackage); idx < len(glpmJsonObj.Packages) && glpmJsonObj.Packages[idx]==goPackage {
			log.Printf("ignore %v, already included in %v json file at position %v",goPackage,glpmJsonFile,idx)
		} else {
			glpmJsonObj.Packages = append(glpmJsonObj.Packages, goPackage)
			sort.Strings(glpmJsonObj.Packages)
		}
	}
	if len(glpmJsonObj.Packages) == 0 {
		log.Printf("did not find any 3rd party dependencies in your code!")
	} else {
		jsonBytes, err = json.MarshalIndent(glpmJsonObj,"","    ")
		if (err != nil) {
			log.Fatalf("internal error, could not parse json Object")
		}
		err = ioutil.WriteFile(glpmJsonFile, jsonBytes, 0644)
		if (err != nil) {
			log.Fatalf("could not write modified json file %v: %v", glpmJsonFile, err)
		}
		log.Fatalf("written modified json file at %v", glpmJsonFile)
	}
}
