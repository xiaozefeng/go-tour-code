package improve

import (
	"bufio"
	"io"
)

func CountLines1(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

func CountLines2(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	lines:= 0
	for scanner.Scan() {
		lines ++
	}
	return lines, scanner.Err()
}

