package rosmovementsensor

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/edaniels/golog"

	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/resource"
)

var Model = resource.NewModel("viamlabs", "yahboom", "imu")

type MyMovementSensor struct {
	resource.Named

	mu        sync.Mutex
	rosmaster string
	rostopic  string
}

type MyMovementSensorConfig struct {
	RosMaster string `json:"rosmaster"`
	RosTopic  string `json:"rostopic"`
}

func init() {
	resource.RegisterComponent(movementsensor.API, Model, resource.Registration[movementsensor.MovementSensor, *MyMovementSensorConfig]{
		Constructor: newMovementSensor,
	})
}

func newMovementSensor(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger golog.Logger) (movementsensor.MovementSensor, error) {
	m := &MyMovementSensor{
		named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}

	if err := m.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *MyMovementSensor) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.rosmaster = conf.Attributes.String("rosmaster")
	m.rostopic = conf.Attributes.String("rostopic")

	if len(strings.TrimSpace(m.rosmaster)) == 0 {
		return errors.New("rosmaster must be set to hostname:port")
	}

	if len(strings.TrimSpace(m.rostopic)) == 0 {
		return errors.New("rostopic must be set to valid imu topic")
	}

	return
}
