package cache

type Stat struct {
	//当前map存储数量
	Count     int64
	//键值长度
	KeySize   int64
	//value值长度
	ValueSize int64
}
//添加时调用
func (s *Stat) add(k string, v []byte) {
	s.Count += 1
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}
//删除时调用
func (s *Stat) del(k string, v []byte) {
	s.Count -= 1
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
