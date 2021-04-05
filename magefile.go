// +build mage

// This is a magefile, and is a "makefile for go".
// See https://magefile.org/
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/carolynvs/magex/shx"
)

// Any commands executed by "must" (as opposed to shx.RunV for example), will stop
// the build immediately when the command fails.
var must = shx.CommandBuilder{StopOnError: true}

// Run go tests
func Test() {
	coverageFile := filepath.Join(getOutputDir(), "coverage.txt")
	must.RunV("go", "test", "-timeout=30s", "-race", "-coverprofile="+coverageFile, "-covermode=atomic", "./cmd/...", "./pkg/...")
}

func Vendor() {
	must.RunV("go", "mod", "vendor")
}

// Check if go.mod matches the contents of vendor/
func VerifyVendor() error {
	output, _ := must.OutputE("git", "status", "--porcelain")
	if output != "" {
		return fmt.Errorf("vendor directory is out-of-date:\n%s", output)
	}
	return nil
}

var outDir = ""

func getOutputDir() string {
	if outDir != "" {
		return outDir
	}
	const sharedVolume = "/shared"
	if _, err := os.Stat(sharedVolume); err == nil {
		return sharedVolume
	}

	return "."
}
