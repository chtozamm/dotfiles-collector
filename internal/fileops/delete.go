package fileops

import (
	"fmt"
	"os"
)

// Delete removes a file with the given path.
func Delete(path string) error {
	fstat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("source %q does not exist", path)
		}
		return fmt.Errorf("stat file: %v", err)
	}

	err = os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("remove %q: %v", fstat.Name(), err)
	}
	return nil
}
