package chunk

import (
	"os"

	"github.com/yowcow/xsort/types"
)

type Chunk struct {
	file string
}

func (c *Chunk) createFile(bytes types.Bytes) error {
	w, err := os.OpenFile(c.file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer w.Close()

	for _, b := range bytes {
		_, err = w.Write(b)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
