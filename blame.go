package svn

import (
	"encoding/xml"
)

type SvnCommitInfo struct {
	Revision int    `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
}

type LineVersionInfo struct {
	Line   int           `xml:"line-number,attr"`
	Commit SvnCommitInfo `xml:"commit"`
}

type BlameResult struct {
	Cmd       string
	LineInfos []LineVersionInfo `xml:"target>entry"`
}

//追溯
//@param pathOrUrl: 文件路径或URL, 必需有值
//@param revision: 版本参数可以是如下之一: 默认值为空
//                                NUMBER       版本号
//                                '{' DATE '}' 在指定时间以后的版本
//                                'HEAD'       版本库中的最新版本
//                                'BASE'       工作副本的基线版本
//                                'COMMITTED'  最后提交或基线之前
//                                'PREV'       COMMITTED的前一版本
//@param verbose: 得到附加信息， 就是多了是时间数据
//@param option: svn的选项，默认为空
func Blame(pathOrUrl string, revision string, option SvnGlobalOptions) (*BlameResult, *SvnError) {
	args := NewArgs(4)
	args.Add(pathOrUrl)
	args.AddIf2("-r", revision)
	args.Add("--xml")

	result, err := exeSvn("blame", option, args.Args...)
	if err != nil {
		return nil, err
	}

	blameResult := &BlameResult{}
	blameResult.Cmd = result.Cmd
	unmaErr := xml.Unmarshal([]byte(result.Result), blameResult)
	if unmaErr != nil {
		return nil, newError(result.Cmd, unmaErr, "")
	}

	return blameResult, nil
}
