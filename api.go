package main

//Notice v2.x

type v2Notice struct {
	Version           string            `xml:"version,attr"`
	ApiKey            string            `xml:"api-key"`
	Notifier          Notifier          `xml:"notifier"`
	Error             Errors            `xml:"error"`
	Request           Request           `xml:"request"`
	ServerEnvironment ServerEnvironment `xml:"server-environment"`
	App               string
}

type Notifier struct {
	Name    string `xml:"name"`
	Version string `xml:"version"`
	Url     string `xml:"url"`
}

type Errors struct {
	Class     string    `xml:"class"`
	Message   string    `xml:"message"`
	Backtrace Backtrace `xml:"backtrace"`
}

type Backtrace struct {
	Lines []Line `xml:"line"`
}

type Line struct {
	Method string `xml:"method,attr"`
	File   string `xml:"file,attr"`
	Number string `xml:"number,attr"`
}

type Request struct {
	Url       string  `xml:"url"`
	Component string  `xml:"component"`
	Action    string  `xml:"action"`
	CgiData   CgiData `xml:"cgi-data"`
}

type CgiData struct {
	Vars []Var `xml:"var"`
}

type Var struct {
	Key     string `xml:"key,attr"`
	Content string `xml:",innerxml"`
}

type ServerEnvironment struct {
	ProjectRoot string `xml:"project-root"`
	Name        string `xml:"environment-name"`
	AppVersion  string `xml:"app-version"`
}

//End Notice v2.x
