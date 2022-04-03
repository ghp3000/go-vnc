package version

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/suutaku/go-vnc/internal/buffer"
)

// Protocol version strings. We don't support V3.
const (
	V7 = "RFB 003.007\n"
	V8 = "RFB 003.008\n"
)

// NegotiateProtocolVersion will negotiate the protocol version with the given connection.
func NegotiateProtocolVersion(buf *buffer.ReadWriter) (string, error) {
	logrus.Infof("Sending version: %q", V8)
	buf.Dispatch([]byte(V8))

	sl, err := buf.Reader().ReadSlice('\n')
	if err != nil {
		return "", fmt.Errorf("reading client protocol version: %v", err)
	}
	ver := string(sl)
	logrus.Infof("Client wants: %q", ver)
	switch ver {
	case V7, V8: // cool.
	default:
		return "", fmt.Errorf("unsupported client-requested version %q", ver)
	}
	return ver, nil
}