package resume

//Iterator 迭代器
type Iterator struct {
	List  []*Impl
	Index int
}

//HasNext 判断迭代器是否还有下一个
//Return
//    bool 是否还有下一个元素，true 为是，false 为否
func (i *Iterator) HasNext() bool {
	return i.Index < len(i.List)
}

//Next 返回迭代器当前元素，并移动下标到下一个
//Return
//    *Impl Impl 指针
func (i *Iterator) Next() *Impl {
	if i.HasNext() {
		ans := i.List[i.Index]
		i.Index = i.Index + 1
		return ans
	}
	return nil
}
