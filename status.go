package svn

import "strings"

type FileStatus struct {
	Status string
	Path   string
	Raw    string
}

func Status(localPath string,option SvnGlobalOptions) ([]FileStatus, *SvnError){
	result,err :=  exeSvn("status", option, localPath)
	if err != nil{
		return nil,err
	}

	if result.Result == ""{
		return make([]FileStatus, 0), nil
	}

	lines := strings.Split(result.Result, "\r\n")
	status := make([]FileStatus, 0, len(lines))
	for _, v := range lines {
		s := strings.TrimSpace(v)
		if len(s) < 9 {
			continue
		}

		statusItem := FileStatus{}
		statusItem.Status = s[:1]
		statusItem.Path = s[8:]
		statusItem.Raw = s
		status = append(status, statusItem)
	}
	return status, err
}