package catch22

import "sort"

func lag(lis []float32) float32 {
	var s float32

	for _, l := range lis {
		s += l
	}

	return s / float32(len(lis))
}

func mag(dic map[int]float32) float32 {
	var s float32

	for _, v := range dic {
		s += v
	}

	return s / float32(len(dic))
}

func mmd(dic map[int]float32, siz int) float32 {
	var m float32

	var lis []float32
	for _, v := range dic {
		lis = append(lis, v)
	}

	sort.SliceStable(lis, func(i, j int) bool {
		return lis[i] < lis[j]
	})

	if siz%2 == 1 {
		m = lis[siz/2]
	} else {
		u := siz / 2
		l := u - 1
		m = (lis[u] + lis[l]) / 2
	}

	return m
}

func mmx(dic map[int]float32) float32 {
	m := dic[0]

	for _, v := range dic {
		if v > m {
			m = v
		}
	}

	return m
}
