package svn

import(
	"os"
)

func Revert(localTarget string, isRecursive bool, delNewFile bool, option SvnGlobalOptions)(*SvnResult, *SvnError) {
	args := NewArgs(2)
	args.Add(localTarget)
	args.AddIf(isRecursive,"-R")
	result, err := exeSvn("revert", option, args.Args...)
	if err != nil{
		return result, err
	}

	if delNewFile {
		//清空多出来的文件
		status, _ := Status(localTarget,option)
		for i := 0; i < len(status); i++ {
			s := status[i];
			if s.Status == "?"{
				os.Remove(s.Path)
			}
		}
	}
	return exeSvn("revert", option, args.Args...)
}