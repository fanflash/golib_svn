package svn

func Cleanup(pathOrUrl string, removeUnversioned bool, removeIgnored bool, vacuumPristines bool,option SvnGlobalOptions) (*SvnResult, *SvnError) {
	args := NewArgs(4)
	args.AddIf(removeUnversioned,"--remove-unversioned")
	args.AddIf(removeIgnored,"--remove-ignored")
	args.AddIf(vacuumPristines,"--vacuum-pristines")
	args.Add(pathOrUrl)
	result, err := exeSvn("cleanup", option, args.Args...)
	if err != nil {
		return nil, err
	}
	return result,nil;
}
