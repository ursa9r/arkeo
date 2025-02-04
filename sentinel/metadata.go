package sentinel

import (
	"github.com/arkeonetwork/arkeo/sentinel/conf"
	"github.com/arkeonetwork/arkeo/x/arkeo/configs"
)

var Version = "0.0.0"

type Metadata struct {
	Configuration conf.Configuration `json:"config"`
	Version       string             `json:"version"`
}

func NewMetadata(config conf.Configuration) Metadata {
	return Metadata{
		Version:       configs.Version,
		Configuration: config,
	}
}
