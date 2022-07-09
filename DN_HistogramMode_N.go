package catch22

func DN_HistogramMode_N(f []float32, b int) float32 {
	if len(f) == 0 {
		return 0
	}

	var bin float32
	{
		bin = float32(b)
	}

	var mac float32
	var max float32
	{
		max = 1
	}

	var cou map[int]float32
	var edg map[int]float32
	{
		cou, edg = his(f, bin)
	}

	var out float32
	for i := 0; i < int(bin); i++ {
		if cou[i] > mac {
			mac = cou[i]
			max = 1
			out = (edg[i] + edg[i+1]) * 0.5
		} else if cou[i] == mac {
			max += 1
			out += (edg[i] + edg[i+1]) * 0.5
		}
	}

	return out / max
}

func his(f []float32, b float32) (map[int]float32, map[int]float32) {
	min, max := bou(f)

	ste := (max - min) / float32(b)

	cou := map[int]float32{}
	for i := 0; i < len(f); i++ {
		ind := (f[i] - min) / ste

		if ind < 0 {
			ind = 0
		}

		if ind >= b {
			ind = b - 1
		}

		{
			cou[int(ind)] += 1
		}
	}

	edg := map[int]float32{}
	for i := 0; i < int(b+1); i++ {
		edg[i] = (float32(i) * ste) + min
	}

	return cou, edg
}

func bou(f []float32) (float32, float32) {
	min := f[0]
	max := -min

	for _, l := range f {
		if l < min {
			min = l
		}
		if l > max {
			max = l
		}
	}

	return min, max
}
