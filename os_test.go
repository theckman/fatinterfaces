package fatinterfaces

import (
	"os"
	"testing"
)

func TestOsFile(t *testing.T) {
	var _ OsFile = (*os.File)(nil)
}
