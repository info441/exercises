package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

type errorReader struct{}

func (er errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test read error")
}

func TestSign(t *testing.T) {
	cases := []struct {
		name           string
		input          io.Reader
		key            []byte
		expectedOutput []byte
		expectError    bool
	}{
		{
			"Valid Data",
			strings.NewReader("test"),
			[]byte("secret"),
			[]byte{3, 41, 160, 107, 98, 205, 22, 179, 62, 182, 121, 43, 232, 198, 11, 21, 141, 137, 162, 238, 58, 135, 111, 206, 154, 136, 30, 187, 72, 140, 9, 20},
			false,
		},
		{
			"Empty Data",
			strings.NewReader(""),
			[]byte("secret"),
			[]byte{249, 230, 110, 23, 155, 103, 71, 174, 84, 16, 143, 130, 248, 173, 232, 179, 194, 93, 118, 253, 48, 175, 222, 108, 57, 88, 34, 197, 48, 25, 97, 105},
			false,
		},
		{
			"Error from Reader",
			errorReader{},
			[]byte("secret"),
			nil,
			true,
		},
	}
	
	for _, c := range cases {
		output, err := Sign(c.input, c.key)
		if c.expectError && err == nil {
			t.Errorf("case %s: should have received error but didn't get one",
				c.name)
		}
		if !c.expectError && err != nil {
			t.Errorf("case %s: unexpected error: %v", c.name, err)
		}
		if !bytes.Equal(output, c.expectedOutput) {
			t.Errorf("case %s: incorrect output: expected %v but got %v",
				c.name, c.expectedOutput, output)
		}
	}
}

func TestValidate(t *testing.T) {
	cases := []struct {
		name           string
		input          io.Reader
		key            []byte
		signature      []byte
		expectedOutput bool
		expectError    bool
	}{
		{
			"Valid Signature",
			strings.NewReader("test"),
			[]byte("secret"),
			[]byte{3, 41, 160, 107, 98, 205, 22, 179, 62, 182, 121, 43, 232, 198, 11, 21, 141, 137, 162, 238, 58, 135, 111, 206, 154, 136, 30, 187, 72, 140, 9, 20},
			true,
			false,
		},
		{
			"Invalid Signature but Valid Key",
			strings.NewReader("test"),
			[]byte("secret"),
			[]byte{4, 41, 160, 107, 98, 205, 22, 179, 62, 182, 121, 43, 232, 198, 11, 21, 141, 137, 162, 238, 58, 135, 111, 206, 154, 136, 30, 187, 72, 140, 9, 20},
			false,
			false,
		},
		{
			"Valid Signature but Invalid Key",
			strings.NewReader("test"),
			[]byte("secretx"),
			[]byte{3, 41, 160, 107, 98, 205, 22, 179, 62, 182, 121, 43, 232, 198, 11, 21, 141, 137, 162, 238, 58, 135, 111, 206, 154, 136, 30, 187, 72, 140, 9, 20},
			false,
			false,
		},
		{
			"Invalid Signature and Invalid Key",
			strings.NewReader("test"),
			[]byte("secretx"),
			[]byte{3, 42, 160, 107, 98, 205, 22, 179, 62, 182, 121, 43, 232, 198, 11, 21, 141, 137, 162, 238, 58, 135, 111, 206, 154, 136, 30, 187, 72, 140, 9, 20},
			false,
			false,
		},
		{
			"Error from Reader",
			errorReader{},
			[]byte("secret"),
			[]byte{3, 41, 160, 107, 98, 205, 22, 179, 62, 182, 121, 43, 232, 198, 11, 21, 141, 137, 162, 238, 58, 135, 111, 206, 154, 136, 30, 187, 72, 140, 9, 20},
			false,
			true,
		},
	}
	
	for _, c := range cases {
		output, err := Verify(c.input, c.key, c.signature)
		if c.expectError && err == nil {
			t.Errorf("case %s: should have received error but didn't get one",
				c.name)
		}
		if !c.expectError && err != nil {
			t.Errorf("case %s: unexpected error: %v", c.name, err)
		}
		if output != c.expectedOutput {
			t.Errorf("case %s: incorrect output: expected %v but got %v",
				c.name, c.expectedOutput, output)
		}
	}
}