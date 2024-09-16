package binarysearch

type TimeMap struct {
	store map[string][]ValStamp
}

type ValStamp struct {
	Val  string
	Time int
}

func Constructor() TimeMap {
	return TimeMap{store: map[string][]ValStamp{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.store[key]; !ok {
		this.store[key] = []ValStamp{}
	}

	this.store[key] = append(this.store[key], ValStamp{
		Val:  value,
		Time: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	res := ""
	values := []ValStamp{}

	if _, ok := this.store[key]; ok {
		values = this.store[key]
	}

	l, r := 0, len(values)-1
	for l <= r {
		m := (r + l) / 2
		if values[m].Time <= timestamp {
			res = values[m].Val
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}
