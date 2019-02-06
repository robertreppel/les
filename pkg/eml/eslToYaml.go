package eml

import (
	yaml "gopkg.in/yaml.v2"
)

// ToYaml converts Event Modeling Language Language to YAML.
func ToYaml(eml Solution) (string, error) {
	d, err := yaml.Marshal(&eml)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
