/*
 *Author  fanflash
 *Date    2023/6/10
 *Desc
 */
package svn

func Delete(localPath string, option SvnGlobalOptions) (*SvnResult, *SvnError) {
	return exeSvn("del", option, localPath)
}

func Deletes(files []string, options SvnGlobalOptions) (*SvnResult, *SvnError) {
	return exeSvn("del", options, files...)
}
