package catch22

import "math"

func DN_OutlierInclude_np_001_mdrmd(lis []float32, sig float32) float32 {
	var con bool
	{
		con = true
	}

	var inc float32
	{
		inc = 0.01
	}

	wor := map[int]float32{}

	var tot float32
	for i := 0; i < len(lis); i++ {
		if lis[i] != lis[0] {
			con = false
		}

		wor[i] = sig * lis[i]

		if wor[i] >= 0 {
			tot++
		}
	}

	if con {
		return 0
	}

	var max float32
	{
		max = mmx(wor)
	}

	if max < inc {
		return 0
	}

	var thr int
	{
		thr = int(max/inc) + 1
	}

	yot := map[int]float32{}
	mt1 := map[int]float32{}
	mt3 := map[int]float32{}
	mt4 := map[int]float32{}

	for j := 0; j < thr; j++ {
		var hig int
		for i := 0; i < len(lis); i++ {
			if wor[i] > float32(j)*inc {
				yot[hig] = float32(i + 1)
				hig++
			}
		}

		ibh := map[int]float32{}
		for i := 0; i < hig-1; i++ {
			ibh[i] = yot[i+1] - yot[i]
		}

		{
			mt1[j] = mag(ibh)
			mt3[j] = float32(hig-1) * 100.0 / tot
			mt4[j] = mmd(yot, hig)/(float32(len(lis)/2)) - 1
		}
	}

	trt := 2
	mmj := 0
	fbi := thr - 1
	for i := 0; i < thr; i++ {
		if mt3[i] > float32(trt) {
			mmj = i
		}

		if math.IsNaN(float64(mt1[thr-1-i])) {
			fbi = thr - 1 - i
		}
	}

	var trl int
	{
		if mmj < fbi {
			trl = mmj
		} else {
			trl = fbi
		}
	}

	var med float32
	{
		med = mmd(mt4, trl+1)
	}

	return med
}
