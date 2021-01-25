package basicauth

import (
	"encoding/json"
	"errors"
	"github.com/devopsfaith/krakend/config"
)

const Namespace = "github_com/azubkokshe/krakend-basicauth"

var ErrNoConfig = errors.New("no config defined for the module")

func ParseConfig(cfg config.ExtraConfig) (Config, error) {
	res := Config{}
	e, ok := cfg[Namespace]
	if !ok {
		return res, ErrNoConfig
	}
	b, err := json.Marshal(e)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(b, &res)
	return res, err
}
