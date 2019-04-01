package svn

//一个目录或文件的svn地址是否跟url相同
func IsSvnPath(target string, url string,option SvnGlobalOptions)(curUrl string, isSame bool){
	result,err := Info(option,target)
	if err != nil{
		return "",false;
	}
	if len(result.Entrys)<1{
		return  "",false;
	}

	curUrl = result.Entrys[0].Url;
	return curUrl, curUrl == url;
}
