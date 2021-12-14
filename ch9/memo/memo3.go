package memo

// 一些URL被获取了两次   两个以上的goroutine同一时刻调用Get来请求同样的URL

func (memo *Memo) Get3(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
