package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	h "gofr.dev/pkg/gofr/health"
)

const (
	natsBackend            = "NATS"
	jetStreamStatusOK      = "OK"
	jetStreamStatusError   = "Error"
	jetStreamConnected     = "CONNECTED"
	jetStreamConnecting    = "CONNECTING"
	jetStreamDisconnecting = "DISCONNECTED"
	natsHealthCheckTimeout = 5 * time.Second
)

// Health returns the health status of the NATS client.
func (n *NATSClient) Health() h.Health {
	health := h.Health{
		Status:  h.StatusUp,
		Details: make(map[string]interface{}),
	}

	connectionStatus := n.Conn.Status()

	switch connectionStatus {
	case nats.CONNECTING:
		health.Status = h.StatusUp
		health.Details["connection_status"] = jetStreamConnecting

		n.Logger.Debug("NATS health check: Connecting")
	case nats.CONNECTED:
		health.Details["connection_status"] = jetStreamConnected

		n.Logger.Debug("NATS health check: Connected")
	case nats.CLOSED, nats.DISCONNECTED, nats.RECONNECTING, nats.DRAINING_PUBS, nats.DRAINING_SUBS:
		health.Status = h.StatusDown
		health.Details["connection_status"] = jetStreamDisconnecting

		n.Logger.Error("NATS health check: Disconnected")
	default:
		health.Status = h.StatusDown
		health.Details["connection_status"] = connectionStatus.String()

		n.Logger.Error("NATS health check: Unknown status", connectionStatus)
	}

	health.Details["host"] = n.Config.Server
	health.Details["backend"] = natsBackend
	health.Details["jetstream_enabled"] = n.JetStream != nil

	ctx, cancel := context.WithTimeout(context.Background(), natsHealthCheckTimeout)
	defer cancel()

	if n.JetStream != nil && connectionStatus == nats.CONNECTED {
		status := getJetStreamStatus(ctx, n.JetStream)

		health.Details["jetstream_status"] = status

		if status != jetStreamStatusOK {
			n.Logger.Error("NATS health check: JetStream error:", status)
		} else {
			n.Logger.Debug("NATS health check: JetStream enabled")
		}
	} else if n.JetStream == nil {
		n.Logger.Debug("NATS health check: JetStream not enabled")
	}

	return health
}

func getJetStreamStatus(ctx context.Context, js jetstream.JetStream) string {
	_, err := js.AccountInfo(ctx)
	if err != nil {
		return jetStreamStatusError + ": " + err.Error()
	}

	return jetStreamStatusOK
}
