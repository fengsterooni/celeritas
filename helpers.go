package celeritas

import "os"

func (c *Celeritas) CreateDirIfNotExist(path string) error {
	const mod = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mod)
		if err != nil {
			return err
		}
	}

	return nil
}
