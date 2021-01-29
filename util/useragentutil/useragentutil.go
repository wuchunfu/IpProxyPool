package useragentutil

import "github.com/mssola/user_agent"

type UserAgent struct {
	OsName         string
	OsVersion      string
	Platform       string
	EngineName     string
	EngineVersion  string
	BrowserName    string
	BrowserVersion string
}

func GetUserAgent(ua string) *UserAgent {
	userAgent := user_agent.New(ua)

	osName := userAgent.OSInfo().Name
	osVersion := userAgent.OSInfo().Version
	platform := userAgent.Platform()
	engineName, engineVersion := userAgent.Engine()
	browserName, browserVersion := userAgent.Browser()
	return &UserAgent{
		OsName:         osName,
		OsVersion:      osVersion,
		Platform:       platform,
		EngineName:     engineName,
		EngineVersion:  engineVersion,
		BrowserName:    browserName,
		BrowserVersion: browserVersion,
	}
}
