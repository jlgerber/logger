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

	log.Debugf("%s %s", "Shoot!", "this is rad")
	log.Errorf("%s %d", "That was a problem:", 1)
	log.Warningf("%s %d", "Bad Value", bv[0])
	log.Infof("%s", "So this doesn't work")
}
