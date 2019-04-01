package svn

func Add(localPath string, option SvnGlobalOptions) (*SvnResult, *SvnError) {
	return exeSvn("add", option, localPath)
}

func Adds(files []string, options SvnGlobalOptions) (*SvnResult, *SvnError) {
	return exeSvn("add", options, files...)
}
