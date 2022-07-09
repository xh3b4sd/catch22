package catch22

func SB_BinaryStats_diff_longstretch0(lis []float32) float32 {
	b := map[int]float32{}
	for i := 1; i < len(lis); i++ {
		p := i - 1

		if lis[i]-lis[p] < 0 {
			b[p] = 0
		} else {
			b[p] = 1
		}
	}

	var l float32
	var s float32
	for i := 0; i < len(lis); i++ {
		if (b[i] == 1) || (i == len(lis)-2) {
			c := float32(i) - l

			if c > s {
				s = c
			}

			l = float32(i)
		}
	}

	return s
}
