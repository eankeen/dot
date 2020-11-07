package e2e

import (
	"log"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/eankeen/dotty/actions"
	"github.com/eankeen/dotty/config"
	"github.com/eankeen/dotty/internal/util"
)

func run() {

}

func _filename() string {
	_, _filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Panicln("runtime.Caller not OK")
	}

	return _filename
}

func _dirname() string {
	return filepath.Dir(_filename())
}

func do(dotfilesDir string, srcDir string, destDir string) {
	actions.Apply(dotfilesDir, srcDir, destDir)

	time.Sleep(time.Millisecond * 500)

	actions.Unapply(dotfilesDir, srcDir, destDir)
}

func TestFull(t *testing.T) {
	testDir := filepath.Join(filepath.Dir(_dirname()), "testdata")
	test1 := filepath.Join(testDir, "test1")

	dotfilesDir := test1
	dottyCfg := config.DottyCfg(dotfilesDir)

	srcDir := util.Src(dotfilesDir, dottyCfg, "user")
	destDir := util.Dest(dotfilesDir, dottyCfg, "user")

	do(dotfilesDir, srcDir, destDir)

	t.Log("thing")
}