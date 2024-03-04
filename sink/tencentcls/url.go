package tencentcls

import (
	"fmt"
	"net/url"
	"strings"
)

type URL struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Project         string
	LogStore        string
}

// ParseUrl parse the url of tencentcls
// url format: tencent+cls://accessKeyId:accessKeySecret@endpoint/projectId/logstore
func ParseURL(_url string) (*URL, error) {
	u, err := url.Parse(_url)
	if err != nil {
		return nil, fmt.Errorf("invalid url of aliyunsls: %w", err)
	}
	if u.Scheme != "tencent+cls" {
		return nil, fmt.Errorf("unknonw scheme(%s) for tencentcls", u.Scheme)
	}
	tmpArr := strings.SplitN(u.Path, "/", 3)
	if len(tmpArr) != 3 {
		return nil, fmt.Errorf("invalid path(%s), it must be /{project}/{logstore}", u.Path)
	}
	tmpArr = tmpArr[1:]
	password, _ := u.User.Password()
	rs := &URL{
		Endpoint:        u.Host,
		AccessKeyID:     u.User.Username(),
		AccessKeySecret: password,
		Project:         tmpArr[0],
		LogStore:        tmpArr[1],
	}
	if err := rs.check(); err != nil {
		return nil, err
	}
	return rs, nil
}

func (u *URL) check() error {
	if u.Project == "" {
		return fmt.Errorf("missing project setup")
	}
	if u.LogStore == "" {
		return fmt.Errorf("missing logstore setup")
	}
	return nil
}
