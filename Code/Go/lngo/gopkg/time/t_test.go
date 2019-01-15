package time_test

import (
	"testing"
	"time"
)

var beijing = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))

type generalTime struct {
	*time.Time
}

func Now() *generalTime {
	now := time.Now().UTC()
	return &generalTime{&now}
}
func (gt *generalTime) String() string {
	return gt.UTC().String()
}
func (gt *generalTime) Beijing() string {
	return gt.UTC().In(beijing).String()
}
func Test_Time(t *testing.T) {
	t.Error(Now())
	t.Error(Now().Beijing())
}
