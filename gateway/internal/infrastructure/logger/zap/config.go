package zap

import "errors"

type Config struct {
	DevMode   bool
	Directory string
}

func (c Config) Validate() error {
	if c.Directory == "" {
		return errors.New("directory is empty")
	}

	return nil
}
