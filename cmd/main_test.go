package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_run_basic(t *testing.T) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = orig }()

	var buf bytes.Buffer
	buf.Write([]byte("PLACE 0,0,NORTH\n"))
	buf.Write([]byte("MOVE\n"))
	buf.Write([]byte("REPORT\n"))

	err := run(bytes.NewReader(buf.Bytes()))
	require.NoError(t, err)
	w.Close()

	got, _ := io.ReadAll(r)
	require.Equal(t, "0,1,NORTH\n", string(got))
}
