	"github.com/twpayne/go-vfs/v5"
	"github.com/twpayne/go-vfs/v5/vfst"
				config := newTestConfig(t, fileSystem, withStdout(&stdout))
				assert.NoError(t, config.execute(append([]string{"diff"}, tc.args...)))