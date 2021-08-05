package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"io"
	"os"
	"strings"
)

func main() {
	opener := func() (io.ReadCloser, error) {
		return Open("python-3.9-alpine.tar.gz")
	}

	img, err := tarball.Image(opener, nil)
	if err != nil {
		fmt.Println("error")
	}
	layers, err := img.Layers()
	if err != nil {
		fmt.Println("error")

	}
	for _, layer := range layers {
		reader, err := layer.Uncompressed()
		if err != nil {
			fmt.Println("error")
		}
		layerDigest ,err := layer.DiffID()
		/*size,err :=  partial.UncompressedSize(layer)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(layerDigest,size)*/
		layerSize, err := UncompressedLayerSize(reader)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(layerDigest,layerSize)
	}

}

func Open(filePath string) (io.ReadCloser, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var r io.ReadCloser
	if strings.HasSuffix(filePath, ".gz") {
		r, err = gzip.NewReader(f)
		if err != nil {
			return nil, err
		}
	} else {
		r = f
	}
	return r, nil
}

func UncompressedLayerSize(r io.Reader) (int64, error) {
	var unCompSize int64
	tf := tar.NewReader(r)
	for {
		hdr, err := tf.Next()

		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		if hdr.Name == "usr/local/bin/"{
			os.Open("usr/local/bin/")
		}
		unCompSize = unCompSize + hdr.Size
	}
	return unCompSize, nil
}
