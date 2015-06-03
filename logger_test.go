package logger

import (
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	bv := []int{1, 2, 3}
	log := GetLogger()
	log.Level = DEBUG
	log.Debug("Shoot!", "this is rad")
	log.Error("That was a problem:", 1)
	log.Warning("Bad Value", bv)
	log.Info("So this doesn't work")
}
