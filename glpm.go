package glpm

import (
	"os"
	"log"
	"flag"
)

const glpmJsonFile string = "glpm.json"

func init() {
	currentDir, err := os.Getwd()
	if (err != nil) {
		log.Panicf("couldn't get current directory's path: %v",err)
	}
	if Options.IsShowVersion {
		ShowVersion()
	}
	if Options.IsInit {
		Init(glpmJsonFile)
	}
	if Options.IsInstall {
		InstallAllPackage(glpmJsonFile)
	}
	if Options.IsAutoAddDependencies {
		AutoAddDependencies(glpmJsonFile,currentDir)
	}
	if len(*Options.InstallPackage)>0 {
		InstallPackage(glpmJsonFile,*Options.InstallPackage)
	}
	if len(*Options.RemovePackage)>0 {
		RemovePackage(glpmJsonFile,*Options.RemovePackage)
	}
	if (flag.NFlag() == 0) {
		ShowVersion()
		flag.PrintDefaults()
	}
}
