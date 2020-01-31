package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/millerjem/konvoyctl/cmd"
)

var konvoyClusterPath = os.Getenv("HOME") + "/.konvoy/clusters"

func main() {
	cmd.Execute()
	//listClusters()
	//initCluster("bah")
	//listClusters()
}

func normalizeClusterName(name string) string {
	return strings.ToLower(name)
}

func listClusters() {

	files, err := ioutil.ReadDir(konvoyClusterPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func mkClusterDir(clusterName string) {
	os.MkdirAll(konvoyClusterPath+"/"+normalizeClusterName(clusterName), os.ModePerm)
}

func initCluster(clusterName string) {
	mkClusterDir(clusterName)
	cmd := exec.Command("konvoy", "init", "--cluster-name", normalizeClusterName(clusterName))
	location := path.Join(os.Getenv("HOME"), "/.konvoy/clusters/"+normalizeClusterName(clusterName))
	cmd.Dir = location
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}

func provisionCLuster(clusterName string) {}
