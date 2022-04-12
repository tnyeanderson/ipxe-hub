package server

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pin/tftp"
	"github.com/tnyeanderson/pixie/config"
)

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	print("TFTP get: " + filename)
	if !strings.HasPrefix(filename, config.Pixie.Paths.FileServer) {
		return errors.New("Path must begin with " + config.Pixie.Paths.FileServer)
	}
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
	return nil
}

// writeHandler is called when client starts file upload to server
func writeHandler(filename string, wt io.WriterTo) error {
	if !strings.HasPrefix(filename, config.Pixie.Paths.FileServer) {
		return errors.New("Path must begin with " + config.Pixie.Paths.FileServer)
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
	if err := s.SetPortRange(65500, 65515); err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}
