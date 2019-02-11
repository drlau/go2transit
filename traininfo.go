package go2transit

func (t TrainStatus) IsDelayed() bool {
	return t.DelayMinute > 0
}
