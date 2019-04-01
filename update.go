package svn

func Update(localPath string, option SvnGlobalOptions)(*SvnResult, *SvnError) {
	return exeSvn("update", option, localPath)
}