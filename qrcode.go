package qrcode

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type Encoder interface {
	// Encode кодирует текст в qr код
	Encode(text string, size int) ([]byte, error)
}

type encoder struct {
}

func NewEncoder() Encoder {
	return &encoder{}
}

// Encode кодирует сообщение в qr код картинку
func (e *encoder) Encode(text string, size int) ([]byte, error) {
	if size <= 0 {
		return nil, errors.New("size is <= 0")
	}

	qrCode, err := qrcode.New(text)
	if err != nil {
		return nil, err
	}

	buffer := closerByteBuffer{}

	writer := standard.NewWithWriter(
		&buffer,
		standard.WithQRWidth(uint8(size)),
		standard.WithBorderWidth(0),
	)
	if err != nil {
		return nil, err
	}

	err = qrCode.Save(writer)
	if err != nil {
		return nil, fmt.Errorf("qr code save: %v", err)
	}

	return buffer.Bytes(), nil
}

// closerByteBuffer закрывающийся байтовый буфер
type closerByteBuffer struct {
	bytes.Buffer
}

func (wc *closerByteBuffer) Close() error {
	return nil
}
