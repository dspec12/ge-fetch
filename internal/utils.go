package internal

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// Unmarhsll JSON into a slice of release types
func getReleases(url string) []release {
	bs := fetchReleaseData(url)
	var releases []release
	err := json.Unmarshal(bs, &releases)
	if err != nil {
		log.Println(err)
	}
	return releases
}

// Fetch Release data from github
func fetchReleaseData(url string) []byte {
	rs, err := http.Get(url + "?per_page=100")
	if err != nil {
		log.Println(err)
	}

	bs, err := io.ReadAll(rs.Body)
	if err != nil {
		log.Println(err)
	}
	return bs
}

// GetHomeDirectory : return the home directory
func getHomeDirectory() string {
	homedir, errHome := homedir.Dir()
	if errHome != nil {
		log.Fatalf("Failed to get home directory %v\n", errHome)
	}
	return homedir
}

// // GetCurrentDirectory : return the current directory
// func getCurrentDirectory() string {
// 	dir, err := os.Getwd() //get current directory
// 	if err != nil {
// 		log.Printf("Failed to get current directory %v\n", err)
// 		os.Exit(1)
// 	}
// 	return dir
// }

func verifyChecksum(filePath, checksum string) bool {
	f, _ := os.Open(filePath)
	defer f.Close()
	h := sha512.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	sha := hex.EncodeToString(h.Sum(nil))
	if checksum != sha {
		log.Fatalln("checksum failed")
		return false
	}
	log.Println("checksum passed")
	return true
}

// Use systems tar program to decompress and extract
func extract(src, dst string) {
	cmd := exec.Command("tar", "-xf", src, "-C", dst)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func userConfirm(msg string) bool {
	var input string

	fmt.Println(msg)

	// Taking input from user
	fmt.Scanln(&input)
	if strings.ToLower(input) == "yes" || input == "y" {
		return true
	}

	fmt.Println("aborting..")
	return false
}
