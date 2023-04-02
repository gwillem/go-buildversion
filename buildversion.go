package buildversion

import (
	"runtime/debug"
	"time"
)

type (
	BuildVersion struct {
		Time     time.Time
		Commit   string
		Branch   string
		Modified bool
		VCS      string
	}
)

func Get() BuildVersion {
	bv := BuildVersion{}
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.branch":
				bv.Branch = setting.Value
			case "vcs.time":
				if t, e := time.Parse(time.RFC3339, setting.Value); e == nil {
					bv.Time = t
				} else {
					panic(e)
				}
			case "vcs.revision":
				bv.Commit = setting.Value
			case "vcs.modified":
				bv.Modified = setting.Value == "true"
			case "vcs":
				bv.VCS = setting.Value
			}
		}
	}
	return bv
}
