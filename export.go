package svn

func Export(svnPath string, localPath string, option SvnGlobalOptions)(*SvnResult, *SvnError) {
	return exeSvn("export", option, svnPath, localPath)
}
