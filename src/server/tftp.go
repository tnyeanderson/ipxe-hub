package server

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/pin/tftp"
	"github.com/tnyeanderson/pixie/api"
	"github.com/tnyeanderson/pixie/config"
)

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	fmt.Printf("TFTP get: %s\n", filename)
	conf := config.Config{}
	if err := conf.Load(api.ConfigPath); err != nil {
		return errors.New("Failed to load config")
	}
	staticRoot := conf.StaticRoot
	if filename == "pixie.kpxe" {
		// For compatibility reasons, allow loading pixie.kpxe from the root path
		filename = path.Join(staticRoot, filename)
	}
	if !strings.HasPrefix(filename, staticRoot) {
		return errors.New("Path must begin with " + staticRoot)
	}
	// TODO: This should add the FileServer prefix, skip the above check
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes sent\n", n)
	//queries.LogLastAccessed(filename)
	return nil
}

// writeHandler is called when client starts file upload to server
func writeHandler(filename string, wt io.WriterTo) error {
	conf := config.Config{}
	if err := conf.Load(api.ConfigPath); err != nil {
		return errors.New("Failed to load config")
	}
	staticRoot := conf.StaticRoot
	if !strings.HasPrefix(filename, staticRoot) {
		return errors.New("Path must begin with " + staticRoot)
	}
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := wt.WriteTo(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes received\n", n)
	return nil
}

func ListenTFTP() {
	s := tftp.NewServer(readHandler, writeHandler)
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}
