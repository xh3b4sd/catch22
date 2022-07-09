package catch22

func SB_BinaryStats_mean_longstretch1(lis []float32) float32 {
	var a float32
	{
		a = avg(lis)
	}

	b := map[int]float32{}
	for i := 0; i < len(lis); i++ {
		if lis[i]-a <= 0 {
			b[i] = 0
		} else {
			b[i] = 1
		}
	}

	var l float32
	var s float32
	for i := 0; i < len(lis); i++ {
		if (b[i] == 0) || (i == len(lis)-1) {
			c := float32(i) - l

			if c > s {
				s = c
			}

			l = float32(i)
		}
	}

	return s
}
