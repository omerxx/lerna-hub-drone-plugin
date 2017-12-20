package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os/exec"
    "os"
    "os/exec"
)


func main() {
    lerna := make(map[string]string)
    tempReader, _ := ioutil.ReadFile("lerna.json")
    json.Unmarshal(tempReader, &ec2Instances)
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
