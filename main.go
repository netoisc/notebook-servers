package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/3Blades/core-server"
	_ "github.com/lib/pq"
)

var (
	databaseURL        string
	redisURL           string
	hashID             string
	resourceDir        string
	port               int
	hubServerSecretKey string
	serverStateKey     string
)

func main() {
	flag.StringVar(&databaseURL, "db", "", "Database url")
	flag.StringVar(&redisURL, "redis", "redis://localhost:6379/0", "Redis url")
	flag.StringVar(&hashID, "hashid", "", "Workspace hash id")
	flag.StringVar(&resourceDir, "resource", "", "Resource dir")
	flag.IntVar(&port, "port", -1, "Notebook port")
	flag.StringVar(&hubServerSecretKey, "secret", "", "Secret key")
	flag.Parse()
	serverStateKey = "server_state_" + hashID
	serverID, err := core.DecodeHashID(hubServerSecretKey, hashID)
	if err != nil {
		go core.SetStatus(redisURL, "Error", serverStateKey)
		log.Fatal(err)
	}
	core.StartScript(databaseURL, serverID)
	err = core.CreateSSHTunnels(databaseURL, resourceDir, serverID)
	if err != nil {
		go core.SetStatus(redisURL, "Error", serverStateKey)
		log.Fatal(err)
	}
	err = runNotebook()
	if err != nil {
		log.Fatal(err)
	}
}

func runNotebook() error {
	args := flag.Args()
	log.Println(args)
	command := args[0]
	cargs := args[1:]
	err := os.Chdir(resourceDir)
	if err != nil {
		go core.SetStatus(redisURL, "Error", serverStateKey)
		return err
	}
	cmd := exec.Command(command, cargs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go core.SetStatus(redisURL, "Running", serverStateKey)
	return cmd.Run()
}
