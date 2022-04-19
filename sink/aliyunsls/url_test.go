package aliyunsls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUrl(t *testing.T) {
	assert := assert.New(t)
	u, err := ParseURL("aliyun+sls://user:passwd@endpoint/projectA/logstoreA")
	assert.NoError(err)
	assert.Equal("user", u.AccessKeyID)
	assert.Equal("passwd", u.AccessKeySecret)
	assert.Equal("projectA", u.Project)
	assert.Equal("logstoreA", u.LogStore)
}
