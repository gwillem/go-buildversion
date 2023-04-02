package buildversion

import (
	"fmt"
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

const (
	timestampFormat = "20060102"
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

	if bv.Time.IsZero() {
		bv.Time = time.Now()
		bv.Commit = "DEV"
	}
	return bv
}

func String() string {
	bv := Get()
	return fmt.Sprintf("%s-%s", bv.Time.Format(timestampFormat), bv.Commit)
}
