package SearchAnyPath_test

import (
	"testing"

	SearchPath "SearchPath/internal"

	"github.com/stretchr/testify/assert"
)

func TestGetPath(t *testing.T) {
	dirName := "SearchPath"
	want := "home/SearchPath"

	assert.Equal(t, want, SearchPath.AnyProjectPath(dirName))
}
