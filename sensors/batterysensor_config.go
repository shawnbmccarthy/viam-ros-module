package sensors

import "fmt"

type BatterySensorConfig struct {
	NodeName   string `json:"node_name"`
	PrimaryUri string `json:"primary_uri"`
	Topic      string `json:"topic"`
}

func (cfg *BatterySensorConfig) Validate(path string) ([]string, error) {
	// NodeName will get default value if string is empty
	if cfg.PrimaryUri == "" {
		return nil, fmt.Errorf(`expted "RosMaster" attribute for sensor %q`, path)
	}

	if cfg.Topic == "" {
		return nil, fmt.Errorf(`expted "RosTopic" attribute for sensor %q`, path)
	}

	return nil, nil
}