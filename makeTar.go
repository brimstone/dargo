package dargo

import (
	"archive/tar"
	"bytes"
	"io/ioutil"
	"os"
)

// makeTar builds a tar ready to upload to a Docker engine including a
// Dockerfile and a copy of the running binary
func makeTar() (*bytes.Reader, error) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new tar archive.
	tw := tar.NewWriter(buf)

	self, err := ioutil.ReadFile(os.Args[0])
	if err != nil {
		return nil, err
	}
	Dockerfile := "FROM scratch\r\nADD app /app\r\nENTRYPOINT [\"/app\"]"

	if err := tw.WriteHeader(&tar.Header{
		Name: "Dockerfile",
		Mode: 0600,
		Size: int64(len(Dockerfile)),
	}); err != nil {
		return nil, err
	}
	if _, err := tw.Write([]byte(Dockerfile)); err != nil {
		return nil, err
	}
	if err := tw.WriteHeader(&tar.Header{
		Name: "app",
		Mode: 0755,
		Size: int64(len(self)),
	}); err != nil {
		return nil, err
	}
	if _, err := tw.Write(self); err != nil {
		return nil, err
	}

	// Make sure to check the error on Close.
	if err := tw.Close(); err != nil {
		return nil, err
	}

	// Open the tar archive for reading.
	return bytes.NewReader(buf.Bytes()), nil
}
