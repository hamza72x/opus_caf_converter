package caf

import (
	"encoding/binary"
	"errors"
	"io"
)

func (h *FileHeader) Decode(r io.Reader) error {
	err := binary.Read(r, binary.BigEndian, h)
	if err != nil {
		return err
	}
	if h.FileType != stringToChunkType("caff") {
		return errors.New("invalid caff header")
	}
	return nil
}

func (h *FileHeader) Encode(w io.Writer) error {
	err := binary.Write(w, binary.BigEndian, h)
	if err != nil {
		return err
	}
	return nil
}
