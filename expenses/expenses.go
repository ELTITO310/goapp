package expenses

func Avarage(expns ...float32) float32 {
	var sum float32

	for _, v := range expns {
		sum += v
	}

	return sum / float32(len(expns))
}

func Sum(expns ...float32) float32 {
	var sum float32

	for _, v := range expns {
		sum += v
	}

	return sum / float32(len(expns))
}

func Max(expns ...float32) float32 {
	var max float32

	for _, v := range expns {
		if v > max {
			max = v
		}
	}

	return max

}

func Min(expns ...float32) float32 {

	if len(expns) == 0 {
		return 0
	}

	var min float32 = expns[0]

	for _, v := range expns {
		if v < min {
			min = v
		}
	}

	return min

}
