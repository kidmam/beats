// build +windows

package eventlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventLogReadFlags(t *testing.T) {
	assert.Equal(t, 0x0001, EVENTLOG_SEQUENTIAL_READ)
	assert.Equal(t, 0x0002, EVENTLOG_SEEK_READ)
	assert.Equal(t, 0x0004, EVENTLOG_FORWARDS_READ)
	assert.Equal(t, 0x0008, EVENTLOG_BACKWARDS_READ)
}

func TestLoadLibraryExFlags(t *testing.T) {
	assert.Equal(t, 0x00000001, DONT_RESOLVE_DLL_REFERENCES)
	assert.Equal(t, 0x00000010, LOAD_IGNORE_CODE_AUTHZ_LEVEL)
	assert.Equal(t, 0x00000002, LOAD_LIBRARY_AS_DATAFILE)
	assert.Equal(t, 0x00000040, LOAD_LIBRARY_AS_DATAFILE_EXCLUSIVE)
	assert.Equal(t, 0x00000020, LOAD_LIBRARY_AS_IMAGE_RESOURCE)
	assert.Equal(t, 0x00000200, LOAD_LIBRARY_SEARCH_APPLICATION_DIR)
	assert.Equal(t, 0x00001000, LOAD_LIBRARY_SEARCH_DEFAULT_DIRS)
	assert.Equal(t, 0x00000100, LOAD_LIBRARY_SEARCH_DLL_LOAD_DIR)
	assert.Equal(t, 0x00000800, LOAD_LIBRARY_SEARCH_SYSTEM32)
	assert.Equal(t, 0x00000400, LOAD_LIBRARY_SEARCH_USER_DIRS)
	assert.Equal(t, 0x00000008, LOAD_WITH_ALTERED_SEARCH_PATH)
}
