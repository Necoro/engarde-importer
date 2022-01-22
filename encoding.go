package main

import (
	"bytes"
	"io"

	"github.com/gogs/chardet"
	"golang.org/x/text/encoding/ianaindex"
)

// getEncodedReader tries to determine the encoding of the content of `r`.
// It returns a new reader that returns UTF-8 content.
func getEncodedReader(r io.Reader) (io.Reader, error) {
	buf := make([]byte, 128)

	n, err := io.ReadFull(r, buf)
	switch {
	case err == io.ErrUnexpectedEOF:
		buf = buf[:n]
		// as `buf` holds the whole content, we can use it as the underlying reader
		r = bytes.NewReader(buf)
	case err != nil:
		return nil, err
	default:
		// re-append `buf`
		r = io.MultiReader(bytes.NewReader(buf), r)
	}

	res, err := chardet.NewTextDetector().DetectBest(buf)
	if err != nil {
		return nil, err
	}

	enc, err := ianaindex.IANA.Encoding(res.Charset)
	if err != nil {
		return nil, err
	}

	return enc.NewDecoder().Reader(r), nil
}
