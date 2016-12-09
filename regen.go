// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var goPkgOptRe = regexp.MustCompile(`(?m)^option go_package = (.*);`)

func main() {
	goOutDir := flag.String("go_out", "", "go_out argument to pass to protoc-gen-go")
	pkgPrefix := flag.String("pkg_prefix", "", "only include proto files with go_package starting with this prefix")
	flag.Parse()

	if *goOutDir == "" {
		log.Fatal("need go_out flag")
	}

	pkgFiles := make(map[string][]string)
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() || !strings.HasSuffix(path, ".proto") {
			return nil
		}
		return register(path, pkgFiles, *pkgPrefix)
	}
	for _, root := range flag.Args() {
		if err := filepath.Walk(root, walkFn); err != nil {
			log.Fatal(err)
		}
	}
	for _, fnames := range pkgFiles {
		if err := protoc(*goOutDir, flag.Args(), fnames); err != nil {
			log.Fatal(err)
		}
	}
}

func register(fname string, pkgFiles map[string][]string, pkgPrefix string) error {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	var pkgName string
	if match := goPkgOptRe.FindSubmatch(content); len(match) > 0 {
		pn, err := strconv.Unquote(string(match[1]))
		if err != nil {
			return err
		}
		pkgName = pn
	}
	if p := strings.IndexRune(pkgName, ';'); p > 0 {
		pkgName = pkgName[:p]
	}
	if strings.HasPrefix(pkgName, pkgPrefix) {
		pkgFiles[pkgName] = append(pkgFiles[pkgName], fname)
	}
	return nil
}

func protoc(goOut string, includes, fnames []string) error {
	args := []string{"--go_out=plugins=grpc:" + goOut}
	for _, inc := range includes {
		args = append(args, "-I"+inc)
	}
	args = append(args, fnames...)
	return exec.Command("protoc", args...).Run()
}
