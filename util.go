package catch22

func avg(lis []float32) float32 {
	var s float32

	for _, l := range lis {
		s += l
	}

	return s / float32(len(lis))
}
