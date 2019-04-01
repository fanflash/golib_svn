package svn

import (
	"fmt"
	"strings"
)

type ErrInfo struct {
	Code    string
	Type    string
	Message string
}

type SvnError struct {
	Cmd       string
	CmdErr    error
	RawErrStr string
	Infos     []ErrInfo
}

func (this *SvnError) HasErrByCode(c string) bool {
	for i := range this.Infos {
		if this.Infos[i].Code == c {
			return true
		}
	}
	return false
}

func (this *SvnError) Error() string {
	return fmt.Sprintf("svn error >> cmd: %s\nstderr: \n%s\ncmd err:%v", this.Cmd, this.RawErrStr, this.CmdErr)
}

//新建一个SVN错误，仅供内部使用
func newError(cmd string, cmdErr error, stderr string) *SvnError {
	svnError := &SvnError{}
	svnError.Cmd = cmd
	svnError.CmdErr = cmdErr
	svnError.RawErrStr = stderr

	if stderr != "" {
		lines := strings.Split(stderr, "\r\n")
		svnError.Infos = make([]ErrInfo, 0, len(lines))
		for i := range lines {
			s := strings.TrimSpace(lines[i])
			if s == "" || len(s) < 14 {
				continue
			}

			errorInfo := ErrInfo{}
			s = s[5:]

			switch s[0] {
			case 'E':
				errorInfo.Type = "Error"
				errorInfo.Code = s[:7]
				errorInfo.Message = s[9:]
			case 'w':
				if len(s) < 17 {
					break
				}
				errorInfo.Type = "Warning"
				errorInfo.Code = s[9:16]
				errorInfo.Message = s[18:]
			}

			if errorInfo.Type == "" {
				errorInfo.Type = "Undefine"
				errorInfo.Code = "0"
				errorInfo.Message = s
			}
			svnError.Infos = append(svnError.Infos, errorInfo)
		}
	}
	return svnError
}
