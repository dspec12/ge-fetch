package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

type release struct {
	TagName   string  `json:"tag_name"`
	Name      string  `json:"name"`
	AssetsURL string  `json:"assets_url"`
	HTMLURL   string  `json:"html_url"`
	Published string  `json:"published_at"`
	Assets    []asset `json:"assets"`
}

type asset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"browser_download_url"`
}

// Download Asset
func (r release) download(program string) {
	// Set asset URL and name vars
	var url, filePath string
	for _, a := range r.Assets {
		if strings.HasSuffix(a.Name, ".tar.xz") || strings.HasSuffix(a.Name, ".tar.gz") {
			url = a.DownloadURL
			filePath = fmt.Sprintf("%s/%s/%s", getHomeDirectory(), program, a.Name)
		}
	}

	// Confirm
	fmt.Printf("%s will be installed to '%s'\n", r.TagName, fmt.Sprintf("%s/%s/", getHomeDirectory(), program))
	if !userConfirm("Do you want to proceed? [y/N]") {
		return
	}

	// Download asset tarball
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fileHandle, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup
	defer func() {
		fileHandle.Close()
		err = os.Remove(filePath)
		if err != nil {
			log.Printf("Failed to cleanup '%s'\n%n", filePath, err)
		}
	}()

	// Download progress bar
	bar := progressbar.DefaultBytes(resp.ContentLength, "downloading")

	io.Copy(io.MultiWriter(fileHandle, bar), resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Verify Checksum
	verifyChecksum(filePath, r.getChecksum())

	// Extract to location
	dst := fmt.Sprintf("%s/%s", getHomeDirectory(), program)
	log.Println("extracting to ", dst)
	extract(filePath, dst)
}

func (r release) getChecksum() string {
	for _, a := range r.Assets {
		if strings.HasSuffix(a.Name, ".sha512sum") {
			resp, err := http.Get(a.DownloadURL)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			bs, _ := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			return strings.Split(string(bs), " ")[0]
		}
	}
	return ""
}
