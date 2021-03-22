package list

type List struct {
	Ptr     *List
	Content int64
}

func (l *List) Reverse() *List {
	if l == nil {
		return nil
	}
	last := l.Ptr.Reverse()
	l.Ptr.Ptr = l
	l.Ptr = nil
	return last
}

var successPtr *List

func (l *List) ReverseN(n int) *List {
	if n == 1 {
		successPtr = l.Ptr
		return l
	}
	last := l.ReverseN(n - 1)
	l.Ptr.Ptr = l
	l.Ptr = successPtr
	return last
}
