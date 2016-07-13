package glpm

import (
	"flag"
)

type options struct {
	IsInit                bool
	IsAutoAddDependencies bool
	InstallPackage        *string
	IsInstall             bool
	IsShowVersion	      bool
	RemovePackage	      *string
}

var Options options

func init() {
	flag.BoolVar(&Options.IsShowVersion,"version",false,"show version of glpm")
	flag.BoolVar(&Options.IsInit,"init",false,"initialize a new glpm.json file")
	flag.BoolVar(&Options.IsAutoAddDependencies,"auto",false,"auto add project dependencies to glpm.json file")
	Options.InstallPackage=flag.String("ipkg","","install specific go package (and save it in glpm.json file)")
	flag.BoolVar(&Options.IsInstall,"install",false,"install go packages that are listed in glpm.json file")
	Options.RemovePackage=flag.String("remove","","remove a package from the glpm.json file. (does NOT uninstall it)")
	flag.Parse()
}