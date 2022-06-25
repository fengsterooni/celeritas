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

func (c *Celeritas) CreateFileIfNotExist(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	return nil
}
