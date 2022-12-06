package workspace

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func SetupWorkspace(repo string, branch string, dir string, debug bool) Workspace {
	color.Magenta("Setting up workspace for %s", repo)
	workspace := Workspace{}

	if branch == "" {
		branch = "main"
	}

	cloneCmd := exec.Command("git", "clone", "--depth=1", "-b", branch, repo, dir)

	if debug {
		cloneCmd.Stdout = os.Stdout
		cloneCmd.Stderr = os.Stderr
	}

	color.Magenta("Cloning %s", repo)
	cloneCmd.Run()

	workspace.Directory = fmt.Sprintf("%s/%s", dir, strings.Split(repo, "/")[1])
	workspace.Endpoint = "http://localhost:5002/interactions"

	yarnCmd := exec.Command("yarn", "install", "--immutable")
	yarnCmd.Dir = workspace.Directory

	if debug {
		yarnCmd.Stdout = os.Stdout
		yarnCmd.Stderr = os.Stderr
	}

	color.Magenta("installing dependencies")
	yarnCmd.Run()

	buildCmd := exec.Command("yarn", "build", "--filter=@disploy/framework-example")
	buildCmd.Dir = workspace.Directory
	if debug {
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
	}

	color.Magenta("building @disploy/framework-example")
	buildCmd.Run()

	startCmd := exec.Command("yarn", "workspace", "@disploy/framework-example", "test-server")
	startCmd.Env = os.Environ()
	startCmd.Dir = workspace.Directory

	stdout, _ := startCmd.StdoutPipe()

	color.Magenta("starting test server")
	startCmd.Start()

	reader := io.ReadCloser(stdout)
	buf := make([]byte, 1024)

	for {
		n, _ := reader.Read(buf)

		if debug {
			color.Cyan(string(buf[:n]))
		}

		if n == 0 {
			color.Red("failed to start test server")
			break
		}

		if strings.Contains(string(buf[:n]), "Server Ready!") {
			color.Green("test server ready")
			break
		}
	}

	return workspace
}
