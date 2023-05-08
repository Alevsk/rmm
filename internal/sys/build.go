package sys

import (
	"runtime/debug"
	"strings"
	"sync"
)

type BuildInfo struct {
	Version  string
	CommitID string
}

func BinaryInfo() BuildInfo {
	readBinaryOnce.Do(func() { binaryInfo = readBinaryInfo() })
	return binaryInfo
}

func readBinaryInfo() BuildInfo {
	const (
		DefaultVersion  = "<unknown>"
		DefaultCommitID = "<unknown>"
	)
	binaryInfo := BuildInfo{
		Version:  DefaultVersion,
		CommitID: DefaultCommitID,
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return binaryInfo
	}

	const (
		GitTimeKey     = "vcs.time"
		GitRevisionKey = "vcs.revision"
	)
	for _, setting := range info.Settings {
		if setting.Key == GitTimeKey {
			binaryInfo.Version = strings.ReplaceAll(setting.Value, ":", "-")
		}
		if setting.Key == GitRevisionKey {
			binaryInfo.CommitID = setting.Value
		}
	}
	return binaryInfo
}

var (
	readBinaryOnce sync.Once
	binaryInfo     BuildInfo
)
