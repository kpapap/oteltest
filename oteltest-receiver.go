package oteltest

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

// Implement the methods from the interface component specific to the type of the component in the connector.go file.
// In this tutorial we will implement the Traces connector and therefore must implement the interfaces:
// baseConsumer, Traces and component.Component.

// Define the receiver struct with the desired parameters for your receiver.
type oteltestReceiver struct {
	// host component.Host
	// cancel context.CancelFunc
	config *Config
	logger *zap.Logger
	nextConsumer consumer.Logs
	component.StartFunc
	component.ShutdownFunc
}



// Implement the methods from the interface baseConsumer.

// Implement the methods from the interface component specific to the type of the component.
// Capabilities implements the consumer interface.
func (c *oteltestReceiver) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
// // Start implements the consumer interface.
// func (c *oteltestReceiver) Start(ctx context.Context, host component.Host) error {
// 	c.host = host
// 	ctx = context.Background()
// 	return nil
// }
// // Shutdown implements the consumer interface.
// func (c *oteltestReceiver) Shutdown(ctx context.Context) error {
// 	if c.cancel != nil {
// 		c.cancel()
// 	}
// 	return nil
// }

// ConsumeLogs implements the consumer interface.
func (c *oteltestReceiver) ConsumeLogs(ctx context.Context, ld consumer.Logs) error {
	return nil
}





