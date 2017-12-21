package main

import (
    "io/ioutil"
    "encoding/json"
    "os"
    "os/exec"
    "syscall"
)


func main() {
    lerna := make(map[string]string)
    tempReader, _ := ioutil.ReadFile("lerna.json")
    json.Unmarshal(tempReader, &lerna )
    version := lerna["version"]

    binary, lookErr := exec.LookPath("lerna")
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"lerna", "publish", "--force-publisj=*", "--repo-version=" + version, "--skip-npm", "--skip-git", "--yes", "--exact"}

    execErr := syscall.Exec(binary, args, os.Environ())
    if execErr != nil {
        panic(execErr)
    }
}
