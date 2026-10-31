package input

import (
	"bufio"
	"io"
	"io/fs"
)

// ReadInputSingle opens a file and allows specification of a worker function for each line.
// Returns error if file can't be opened or if scanning fails.
// fs.FS is used per the techniques learned for testing in Learn Go with Tests
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files#file-system-abstractions-introduced-in-go-1.16
func ReadInputSingle(fileSys fs.FS, path string, worker func(lineNum int, line string)) error {
	f, err := fileSys.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return readFile(f, worker)
}

func ReadInputDir(path string, worker func(lineNum int, line string)) error {
	return nil // stubbed out for now
}

func readFile(r io.Reader, worker func(lineNum int, line string)) error {
	scanner := bufio.NewScanner(r)
	for lineNum := 1; scanner.Scan(); lineNum++ {
		worker(lineNum, scanner.Text())
	}
	return scanner.Err()
}