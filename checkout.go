package svn

func Checkout(svnPath string, localPath string, option SvnGlobalOptions)(*SvnResult, *SvnError) {
	return exeSvn("co", option, svnPath, localPath)
}