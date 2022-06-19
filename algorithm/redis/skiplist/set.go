package skiplist

func zslValueGteMin(value float64, spec *zrangespec) bool {
	if spec.minex != 0 {
		return value > spec.min
	} else {
		return value >= spec.min
	}
}

func zslValueLteMax(value float64, spec *zrangespec) bool {
	if spec.maxex != 0 {
		return value < spec.max
	} else {
		return value <= spec.max
	}
}
