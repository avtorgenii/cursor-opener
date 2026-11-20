package launcher

import "os/exec"

func Open(path string) error {
	return exec.Command("cursor", path).Start()
}
