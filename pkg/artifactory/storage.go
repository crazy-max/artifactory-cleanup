package artifactory

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// LastModified holds last modified response
type LastModified struct {
	URI          string `json:"uri"`
	LastModified string `json:"lastModified"`
}

// FileStats holds file stats response
type FileStats struct {
	URI                  string `json:"uri"`
	DownloadCount        int64  `json:"downloadCount"`
	LastDownloaded       int64  `json:"lastDownloaded"`
	RemoteDownloadCount  int64  `json:"remoteDownloadCount"`
	RemoteLastDownloaded int64  `json:"remoteLastDownloaded"`
}

// FileList holds file list response
type FileList struct {
	URI     string    `json:"uri"`
	Created time.Time `json:"created"`
	Files   []struct {
		URI          string    `json:"uri"`
		Size         int64     `json:"size"`
		LastModified time.Time `json:"lastModified"`
		Folder       bool      `json:"folder"`
		Sha1         string    `json:"sha1"`
		Sha2         string    `json:"sha2"`
	} `json:"files"`
}

// LastDownloaded returns the last downloaded time of an artifact
func (c *Client) LastDownloaded(repo string, path string) (time.Time, error) {
	fs, err := c.FileStats(repo, path)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, fs.LastDownloaded*int64(time.Millisecond)), nil
}

// FileStats returns the file stats of an artifact
func (c *Client) FileStats(repo string, path string) (FileStats, error) {
	var fs FileStats

	resp, err := c.restCli.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&fs).
		Get(fmt.Sprintf("%s/api/storage/%s/%s?stats", c.url, repo, path))

	log.Trace().Interface("file_stats", fs).Msg("Artifactory artifact file stats response")

	if err != nil {
		return fs, err
	} else if resp.IsError() {
		return fs, errors.New(resp.Status())
	}

	return fs, err
}

// FileList gets a flat listing of the files and folders within a folder
func (c *Client) FileList(repo string, path string) (FileList, error) {
	var fl FileList

	resp, err := c.restCli.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&fl).
		Get(fmt.Sprintf("%s/api/storage/%s/%s?list&deep=1", c.url, repo, path))

	log.Trace().Interface("file_list", fl).Msg("Artifactory file list response")

	if err != nil {
		return fl, err
	} else if resp.IsError() {
		return fl, errors.New(resp.Status())
	}

	return fl, nil
}

// DeleteItem deletes a file or a folder on Artifactory
func (c *Client) DeleteItem(repo string, path string) error {
	resp, err := c.restCli.R().
		Delete(fmt.Sprintf("%s/%s/%s", c.url, repo, path))

	if err != nil {
		return err
	} else if resp.IsError() {
		return errors.New(resp.Status())
	}

	return nil
}
