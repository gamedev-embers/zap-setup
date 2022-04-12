package zapsetup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestRootLogger(t *testing.T) {
	assert := assert.New(t)

	log := RootLogger()
	log.Info("ok")

	root1 := RootLogger()
	root2 := RootLogger()
	assert.Equal(root1.level.Level(), zap.InfoLevel)

	new := NewLogger()
	assert.Equal(new.level.Level(), zap.InfoLevel)
	new.SetLevel(zap.DebugLevel)
	assert.Equal(root1.level.Level(), zap.InfoLevel)
	assert.Equal(new.level.Level(), zap.DebugLevel)

	root1.SetLevel(zap.ErrorLevel)
	assert.Equal(root1.level.Level(), zap.ErrorLevel)
	assert.Equal(root2.level.Level(), zap.ErrorLevel)
}
