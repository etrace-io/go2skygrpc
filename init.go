package go2sky

import (
	"github.com/etrace-io/go2skygrpc/common"
	language_agent "github.com/etrace-io/go2skygrpc/language-agent"
	"github.com/etrace-io/go2skygrpc/management"
)

var (
	_ = common.CPU{}
	_ = language_agent.ID{}
	_ = management.InstancePingPkg{}
)
