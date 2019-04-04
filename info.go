package svn

import "encoding/xml"
import "errors"

type InfoResult struct {
	Cmd    string
	Entrys []InfoEntry `xml:"entry"`
}

type InfoEntry struct {
	Kind        string         `xml:"kind,attr"`
	Revision    uint           `xml:"revision,attr"`
	Path        string         `xml:"path,attr"`
	Url         string         `xml:"url"`
	RelativeUrl string         `xml:"relative-url"`
	Repository  RepositoryInfo `xml:"repository"`
	WcInfo      WcInfo         `xml:"wc-info"`
	Commit      SvnCommitInfo  `xml:"commit"`
}

type RepositoryInfo struct {
	Root string `xml:"root"`
	UUID string `xml:"uuid"`
}

type WcInfo struct {
	RootAbspath string `xml:"wcroot-abspath"`
	Schedule    string `xml:"schedule"`
	TextUpdated string `xml:"text-updated"`
	CheckSum    string `xml:"checksum"`
}

func Info(option SvnGlobalOptions, targets ...string) (*InfoResult, *SvnError) {
	args := NewArgs(len(targets) + 1)
	args.Add(targets...)
	args.Add("--xml")

	result, err := exeSvn("info", option, args.Args...)
	if err != nil {
		return nil, err
	}

	infoResult := &InfoResult{}
	infoResult.Cmd = result.Cmd
	if result.Result == "" {
		return infoResult, nil
	}

	unmaErr := xml.Unmarshal([]byte(result.Result), infoResult)
	if unmaErr != nil {
		return nil, newError(result.Cmd, unmaErr, "")
	}

	return infoResult, nil
}

func InfoByOne(path string, option SvnGlobalOptions) (*InfoResult, *SvnError) {
	result, err := Info(option, path)
	if err != nil {
		return nil, err
	}

	if len(result.Entrys) < 1 {
		return nil, newError(result.Cmd,errors.New("len(result.Entrys) < 1 "),"")
	}

	return result, err
}
