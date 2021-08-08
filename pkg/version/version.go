package version

import (
	"fmt"
)

// Version describes version
// Tag show the git tag for this version
// BuildTime show the compile time
var (
	Version   = "17.03.28"
	Tag       = "2017-03-28 Release"
	BuildTime = "2017-03-28 19:50:00"
	GitHash   = "unknown"
)

// ShowVersion is the default handler which match the --version flag
func ShowVersion() {
	fmt.Printf("%s", GetVersion())
}

func GetVersion() string {
	version := fmt.Sprintf("Version  :%s\nTag      :%s\nBuildTime:  %s\nGitHash:  %s\n", Version, Tag, BuildTime, GitHash)
	return version
}
