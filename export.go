package svn

func Export(svnPath string, localPath string, option SvnGlobalOptions,others ...string)(*SvnResult, *SvnError) {
	return exeSvn("export", option, svnPath, localPath,others)
}
