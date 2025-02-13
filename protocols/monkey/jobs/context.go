package jobs

import (
	"context"
	"time"

	protocolCommon "github.com/taubyte/tau/protocols/common"
)

// Used in tests for setting the unexported contexts
func (c *Context) ForceContext(ctx context.Context) {
	c.ctx, c.ctxC = context.WithCancel(ctx)
	if c.Node != nil && c.OdoClientNode == nil {
		c.OdoClientNode = c.Node
	}
}

func (c *Context) ForceGitDir(dir string) {
	c.gitDir = dir
}

func (c *Context) startTimeout(ctx context.Context, ctxC context.CancelFunc) {
	defaultWaitTime := protocolCommon.DefaultLockTime
	if protocolCommon.TimeoutTest {
		defaultWaitTime = 5
	}

	<-time.After(time.Duration(defaultWaitTime) * time.Second)
	err := c.Patrick.Timeout(c.Job.Id)
	if err != nil {
		logger.Errorf("Sending timeout for job %s failed with %s", c.Job.Id, err.Error())
		return
	}

	c.Monkey.Delete(c.Job.Id)
}
