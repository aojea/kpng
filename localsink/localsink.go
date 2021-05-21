package localsink

import (
	"os"

	"github.com/spf13/pflag"

	"sigs.k8s.io/kpng/pkg/server/watchstate"
)

type Sink interface {
	// WaitRequest waits for the next diff request, returning the requested node name. If an error is returned, it will cancel the job.
	WaitRequest() (nodeName string, err error)

	// Reset the state of the Sink (ie: when the client is disconnected and reconnects)
	Reset()

	watchstate.OpSink
}

type Config struct {
	NodeName string
}

func (c *Config) BindFlags(flags *pflag.FlagSet) {
	flags.StringVar(&c.NodeName, "node-name", func() string {
		s, _ := os.Hostname()
		return s
	}(), "Node name override")
}
