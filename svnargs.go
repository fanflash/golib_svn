package svn

type SvnArgs struct {
	Cap  int
	Args []string
}

func (this *SvnArgs) Add(args ...string) {
	if this.Args == nil {
		this.Args = make([]string, 0, this.Cap)
	}
	this.Args = append(this.Args, args...)
}

//如果condition为true,则加入args
func (this *SvnArgs) AddIf(condition bool, args ...string) {
	if condition {
		this.Add(args...)
	}
}

//如果value不等空，则加入key and value这两个参数
func (this *SvnArgs) AddIf2(key string, value string) {
	if value != "" {
		this.Add(key, value)
	}
}

func NewArgs(cap int) *SvnArgs {
	return &SvnArgs{Cap: cap}
}
