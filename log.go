package svn

import (
	"encoding/xml"
	"strconv"
)

type LogEntry struct {
	Version string `xml:"revision,attr"`
	Author  string `xml:"author"`
	Date    string `xml:"date"`
	Msg     string `xml:"msg"`
	Paths   []PathInfo `xml:"paths>path"`
}
type PathInfo struct {
	Action   string `xml:"action,attr"`
	PropsMod bool   `xml:"prop-mods,attr"`
	TextMods bool   `xml:"text-mods,attr"`
	Kind     string `xml:"kind,attr"`
	File     string `xml:",chardata"` // 使用 ",chardata" 来捕获元素内的文本内容作为字段值
}
type LogResult struct {
	Cmd       string
	LogEntrys []LogEntry `xml:"logentry"`
}

//得到一个文件或路径的LOG记录
//@param pathOrUrl 本地路径或地址
//@param revision  默认值为空，ARG (some commands also take ARG1:ARG2 range)
//                             A revision argument can be one of:
//                                NUMBER       revision number
//                                '{' DATE '}' revision at start of the date
//                                'HEAD'       latest in repository
//                                'BASE'       base rev of item's working copy
//                                'COMMITTED'  last commit at or before BASE
//                                'PREV'       revision just before COMMITTED
//@param limit "最大的查询行数"
//@param option: svn的选项，默认为空
func Log(pathOrUrl string, revision string, limit int, option SvnGlobalOptions, others ...string) (*LogResult, *SvnError) {
	args := NewArgs(5)
	args.AddIf2("-r", revision)
	args.Add("--xml")
	args.Add(others[0])
	args.AddIf(limit > 0, "-l", strconv.Itoa(limit))
	args.Add(pathOrUrl);
	result, err := exeSvn("log", option, args.Args...)
	if err != nil {
		return nil, err
	}

	logResult := &LogResult{}
	logResult.Cmd = result.Cmd
	unmaErr := xml.Unmarshal([]byte(result.Result), logResult)
	if unmaErr != nil {
		return nil, newError(result.Cmd, unmaErr, "")
	}
	return logResult, nil
}
