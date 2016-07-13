package main

import (
	"../params"
	"../func/initglpm"
	"os"
	"../func/dependencies"
	"../func/version"
	"../func/install"
	"../func/remove"
	"log"
	"flag"
)

const glpmJsonFile string = "glpm.json"

func main() {
	currentDir, err := os.Getwd()
	if (err != nil) {
		log.Panicf("couldn't get current directory's path: %v",err)
	}
	if params.Options.IsShowVersion {
		version.ShowVersion()
	}
	if params.Options.IsInit {
		initglpm.Init(glpmJsonFile)
	}
	if params.Options.IsInstall {
		install.InstallAllPackage(glpmJsonFile)
	}
	if params.Options.IsAutoAddDependencies {
		dependencies.AutoAddDependencies(glpmJsonFile,currentDir)
	}
	if len(*params.Options.InstallPackage)>0 {
		install.InstallPackage(glpmJsonFile,*params.Options.InstallPackage)
	}
	if len(*params.Options.RemovePackage)>0 {
		remove.RemovePackage(glpmJsonFile,*params.Options.RemovePackage)
	}
	if (flag.NFlag() == 0) {
		version.ShowVersion()
		flag.PrintDefaults()
	}
}
