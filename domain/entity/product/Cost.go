package productEntity

type Cost int64

func CostFrom(costInt int64) (Cost, error) {
	if costInt < 0 {
		return -1, ErrCostIsNegative
	}

	return Cost(costInt), nil
}

func (t Cost) Int() int64 {
	return int64(t)
}
