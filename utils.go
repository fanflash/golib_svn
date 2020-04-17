package svn

import (
	"net/url"
)

//一个目录或文件的svn地址是否跟url相同
func IsSvnPath(target string, svnUrl string, option SvnGlobalOptions) (curUrl string, isSame bool) {
	result, err := Info(option, target)
	if err != nil {
		return "", false
	}
	if len(result.Entrys) < 1 {
		return "", false
	}
	curUrl, _ = url.QueryUnescape(result.Entrys[0].Url)
	if curUrl != svnUrl {
		svnUrl, _ = url.QueryUnescape(svnUrl)
	}
	return curUrl, curUrl == svnUrl
}

//得到一个路径的版本号
func GetRevisionByPath(p string, option SvnGlobalOptions) (int, error) {
	info, err := InfoByOne(p, option)
	if err != nil {
		return -1, err
	}
	return int(info.Entrys[0].Commit.Revision), nil
}
