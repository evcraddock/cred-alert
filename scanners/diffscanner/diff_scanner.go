package diffscanner

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	"github.com/pivotal-cf/cred-alert/scanners"

	"code.cloudfoundry.org/lager"
)

const MaxLineSize = 1024 * 1024

var fileHeaderRegexp = regexp.MustCompile(`^\+\+\+\s(?:\w\/(.*)|/dev/null)$`)
var hunkHeaderRegexp = regexp.MustCompile(`^@@.*\+(\d+),?\d*\s@@`)
var addedOrExistingLineRegexp = regexp.MustCompile(`^([\s\+])(.*)`)

type DiffScanner struct {
	scanner           *bufio.Scanner
	content           []byte
	currentPath       string
	currentLineNumber int
}

func NewDiffScanner(diff io.Reader) *DiffScanner {
	buf := make([]byte, 0, 4096)
	s := &DiffScanner{
		scanner: bufio.NewScanner(diff),
	}
	s.scanner.Buffer(buf, MaxLineSize)
	return s
}

func (d *DiffScanner) Scan(logger lager.Logger) bool {
	logger = logger.Session("diff-scanner").Session("scan")
	logger.Debug("starting")
	defer logger.Debug("done")

	for d.scanner.Scan() {
		line := d.scanner.Text()

		matches := fileHeaderRegexp.FindStringSubmatch(line)
		if len(matches) == 2 {
			d.currentPath = matches[1]
			continue
		}

		matches = hunkHeaderRegexp.FindStringSubmatch(line)
		if len(matches) == 2 {
			startLine, err := strconv.Atoi(matches[1])
			if err != nil {
				logger.Error("failed", err)
				break
			}
			d.currentLineNumber = startLine - 1
			continue
		}

		matches = addedOrExistingLineRegexp.FindStringSubmatch(line)
		if len(matches) == 3 {
			d.currentLineNumber++
			if matches[1] == "+" {
				d.content = []byte(matches[2])
				return true
			}
		}
	}

	return false
}

func (d *DiffScanner) Line(lager.Logger) *scanners.Line {
	return &scanners.Line{
		Content:    d.content,
		LineNumber: d.currentLineNumber,
		Path:       d.currentPath,
	}
}

func (d *DiffScanner) Err() error {
	return d.scanner.Err()
}
