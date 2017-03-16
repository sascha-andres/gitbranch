package app

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// BranchInformation contains data about a branch
type BranchInformation struct {
	Key     string
	Value   string
	Enabled bool
}

// GetBranches connects to remote git repo and returns all branches
func GetBranches(remoteRepository string) ([]BranchInformation, error) {
	dir, err := ioutil.TempDir("", "gitbranch")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	if _, err := execCommand("git init", dir); err != nil {
		return nil, err
	}

	if _, err := execCommand("git remote add origin "+remoteRepository, dir); err != nil {
		return nil, err
	}

	var branches []string
	if branches, err = execCommand("git ls-remote | awk '{print $2}' | grep refs/heads | cut -c 12-", dir); err != nil {
		return nil, err
	}
	var result []BranchInformation

	for _, branch := range branches {
		result = append(result, BranchInformation{Enabled: true, Key: branch, Value: branch})
	}

	return result, nil
}

func execCommand(command, directory string) ([]string, error) {
	var cmd *exec.Cmd
	cmd = exec.Command("bash", "-c", command)
	cmd.Dir = directory
	commandOutput, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(commandOutput))
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if nil != scanner.Err() {
		return nil, err
	}
	return result, nil
}