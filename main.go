package main

import (
	"encoding/csv"
	"log"
	"os"
	"os/exec"

	"github.com/schollz/progressbar/v3"
)

var FilePath = envOrPanic("CSV_ABS_FILE")
var TargetCloneDirectory = envOrPanic("OUTPUT_ABS_DIR")

func envOrPanic(k string) string {
	v, present := os.LookupEnv(k)
	if !present {
		panic(k + "variable is not set.")
	}
	return v
}

func parseRepositoryList() [][]string {
	f, err := os.Open(FilePath)
	if err != nil {
		log.Fatal("Can't open file at " + FilePath + ".")
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV at " + FilePath + ".")
	}

	return records
}

func main() {
	repositories := parseRepositoryList()
	bar := progressbar.Default(int64(len(repositories)))
	_ = bar.RenderBlank()
	for _, repo := range repositories {

		url, commitId, targetFolderName := repo[0], repo[1], repo[2]
		targetAbsolutePath := TargetCloneDirectory + "/" + targetFolderName
		log.Println(url, commitId, targetFolderName)
		cmd := exec.Command(
			"git",
			"clone",
			"--depth",
			"1",
			"--revision",
			commitId,
			url,
			targetAbsolutePath)

		err := cmd.Run()
		if err != nil {
			log.Printf(err.Error())
			log.Printf("Failed to clone %s at %s\n", url, commitId)
		}
		_ = bar.Add(1)
	}
}
