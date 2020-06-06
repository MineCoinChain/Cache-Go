package cache

//缓存实现接口，添加set,get,delete方法
type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	//获取缓存信息
	GetStat() Stat
}
