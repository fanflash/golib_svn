/*
 *Author  fangao
 *Date    2023/6/10
 *Desc
 */
package svn

func Commit(localPath string, message string, option SvnGlobalOptions) (*SvnResult, *SvnError) {
	return exeSvn("commit", option, "-m", message, localPath)
}
