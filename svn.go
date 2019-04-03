package svn

import (
	"bytes"
	"os/exec"
	"strings"
	"io/ioutil"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

//默认的设置值
var DefOption SvnGlobalOptions

type SvnResult struct {
	Cmd    string
	Result string
}

type SvnGlobalOptions struct {
	Username string
	Password string
	//为true则不缓存认证信息
	NoAuthCache bool
	//svn程序地址
	Svn string
}

//执行svn命令
//@param svnCmd: svn命令
//@param option: 选项
func exeSvn(svnCmd string, option SvnGlobalOptions, args ...string) (*SvnResult, *SvnError) {
	//6为SvnGlobalOption中字段个数*2
	svnArgs := NewArgs(len(args) + 6)
	svnArgs.Add(svnCmd)
	svnArgs.Add(args...)
	if option.Username != "" {
		svnArgs.Add("--username", option.Username)
		svnArgs.AddIf2("--password", option.Password)
		svnArgs.AddIf(option.NoAuthCache, "--no-auth-cache")
	}
	//不要交互提法（API形式，这个是肯定的）
	svnArgs.Add("--non-interactive")
	if option.Svn == "" {
		option.Svn = "svn"
	}
	cmd := exec.Command(option.Svn, svnArgs.Args...)
	cmdStr := strings.Join(cmd.Args, " ")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, newError(cmdStr, err, Gbk2Utf8(stderr.String()))
	}
	return &SvnResult{Cmd: cmdStr, Result: Gbk2Utf8(stdout.String())}, nil
}

func Gbk2Utf8(src string) string {
	tar, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return src
	}
	return string(tar)
}
