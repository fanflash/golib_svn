package svn

import(
	"time"
)

type OwnerInfo struct {
	LastAuthor  string
	MostAuthor  string
	FirstAuthor string
}

//得到一个路径的所有者信息
//如果返回的信息里没有带信息，说明没有查询到条目
//如果queryLimit足够大，那么FirstAuthor也可以当作第一个提交的人
func GetOwnerByPath(filePath string, queryLimit int, option SvnGlobalOptions) (owner OwnerInfo, err *SvnError) {
	result, err := Log(filePath, "", queryLimit, option)
	if err != nil {
		return OwnerInfo{}, err
	}

	if len(result.LogEntrys) < 1 {
		return OwnerInfo{}, nil
	}

	ownerInfo := OwnerInfo{}
	ownerInfo.LastAuthor = result.LogEntrys[0].Author
	ownerInfo.FirstAuthor = result.LogEntrys[len(result.LogEntrys)-1].Author
	authorCountMap := make(map[string]int, len(result.LogEntrys))
	for _, v := range result.LogEntrys {
		authorCountMap[v.Author]++
	}
	most := 0
	for key, value := range authorCountMap {
		if value > most {
			most = value
			ownerInfo.MostAuthor = key
		}
	}
	return ownerInfo, nil
}

var ownerMap map[string]SvnOwnerCache = make(map[string]SvnOwnerCache)

//所有者的超时时间
var OwnerCacheTime time.Duration = 5

func GetOwnerbyPathUseCache(filePath string, queryLimit int, option SvnGlobalOptions) (owner OwnerInfo, err *SvnError) {
	now := time.Now()
	v, ok := ownerMap[filePath]
	if ok{
		if now.Before(v.Timeout) {
			return v.Owner, nil
		}else{
			delete(ownerMap,filePath)
		}
	}

	owner, err = GetOwnerByPath(filePath, queryLimit, option)
	if err != nil{
		return owner, err
	}

	cache := SvnOwnerCache{}
	cache.Path = filePath
	cache.Owner = owner
	cache.Timeout = now.Add(OwnerCacheTime * time.Second)
	ownerMap[filePath] = cache
	return owner, err
}

type SvnOwnerCache struct {
	Path    string
	Owner   OwnerInfo
	Timeout time.Time
}
