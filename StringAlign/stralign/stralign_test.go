package stralign

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var validChars []int

func init() {
	rand.Seed(time.Now().UnixNano())

	for i := 32; i < 64; i++ {
		if i != 34 && i != 39 {
			validChars = append(validChars, i)
		}
	}
}

func TestLjust(t *testing.T) {
	strInput := randomString()
	widthInput := randomWidth()
	fillInput := randomFill()

	strOutput, err := Ljust(strInput, widthInput, fillInput)
	require.NoError(t, err)
	require.NotEmpty(t, strOutput)

	//get output from Python
	strPythonOutput, err := executePython("ljust", strInput, widthInput, fillInput)
	require.NoError(t, err)
	require.NotEmpty(t, strPythonOutput)

	require.Equal(t, strOutput, strPythonOutput)
}

func TestRjust(t *testing.T) {
	strInput := randomString()
	widthInput := randomWidth()
	fillInput := randomFill()

	strOutput, err := Rjust(strInput, widthInput, fillInput)
	require.NoError(t, err)
	require.NotEmpty(t, strOutput)

	//get output from Python
	strPythonOutput, err := executePython("rjust", strInput, widthInput, fillInput)
	require.NoError(t, err)
	require.NotEmpty(t, strPythonOutput)

	require.Equal(t, strOutput, strPythonOutput)
}

func TestCenter(t *testing.T) {
	strInput := randomString()
	widthInput := randomWidth()
	fillInput := randomFill()

	strOutput, err := Center(strInput, widthInput, fillInput)
	require.NoError(t, err)
	require.NotEmpty(t, strOutput)

	//get output from Python
	strPythonOutput, err := executePython("center", strInput, widthInput, fillInput)
	require.NoError(t, err)
	require.NotEmpty(t, strPythonOutput)

	require.Equal(t, strOutput, strPythonOutput)
}

func randomString() string {
	var sb strings.Builder
	k := len(alphabet)

	numWords := 1 + rand.Intn(5)
	numCharPerWord := 1 + rand.Intn(5)
	for wc := 0; wc < numWords; wc++ {
		for i := 0; i < numCharPerWord; i++ {
			c := alphabet[rand.Intn(k)]
			sb.WriteByte(c)
		}
		sb.WriteByte(' ')
	}
	return sb.String()
}

func randomWidth() int32 {
	return rand.Int31n(50)
}

func randomFill() string {
	k := len(validChars)
	return string(rune(validChars[rand.Intn(k)]))
}

func executePython(funcName string, strInput string, widthInput int32, fillInput string) (string, error) {
	pyProg := fmt.Sprintf("print(\"%s\".%s(%d, \"%s\"))", strInput, funcName, widthInput, fillInput)
	fmt.Println(pyProg)
	cmd := exec.Command("python3", "-c", pyProg)

	var out, stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	fmt.Printf("output buffer: %v", out)
	return out.String()[:out.Len()-1], nil
}
