// The new receiver has to provide a receiver.Factory implementation, and although you will find a receiver.Factory interface
// (you can find its definition in the receiver/receiver.go file within the Collector’s project),
// the right way to provide the implementation is by using the functions available within the go.opentelemetry.io/collector/receiver package.

// To instantiate the object, you will need to use the NewFactory function associated with each of the components.
// We will use the connector.NewFactory function.
// The connector.NewFactory function instantiates and returns a connector.Factory and it requires the following parameters:

//     component.Type: a unique string identifier for your connector across all collector’s components of the same type. This string also acts as the name to refer to the connector by.
//     component.CreateDefaultConfigFunc: a reference to a function that returns the default component.Config instance for your connector.
//     ...FactoryOption: the slice of connector.FactoryOptions will determine what type of signal your connector is capable of processing.

package oteltest

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// Create factory.go file and define the unique string to identify your connector as a global constant
const (
	defaultVal = "1m"
)

var (
	// this is the name used to refer to the connector in the config.yaml
	typeStr = component.MustNewType("oteltest")
)

// Create the default configuration function. This is how you choose to initialize your connector object with default values.
func createDefaultConfig() component.Config {
	return &Config{
		Interval: defaultVal,
	}
}

// Define the type of the receiver
func createLogsReceiver(ctx context.Context, set Settings, cfg component.Config, nextConsumer consumer.Logs,) (Logs, error) {
	return nil, nil
}

// Write a NewFactory function that instantiates your custom factory for your connector(component).
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithTraces(createLogsReceiver, component.StabilityLevelAlpha))
}
