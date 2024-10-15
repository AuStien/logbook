package book

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadDay(file *os.File) (Day, error) {
	var day Day

	fi, err := file.Stat()
	if err != nil {
		return Day{}, err
	}

	b := make([]byte, fi.Size())
	_, err = file.Read(b)
	if err != nil {
		return Day{}, err
	}

	lines := strings.Split(string(b), "\n")

	if !strings.HasPrefix(lines[0], "#") {
		return Day{}, errors.New("missing day header")
	}

	header := strings.Split(lines[0], " ")
	// Expect the following format
	// # Wednesday 24/07/2024
	if len(header) != 3 {
		return Day{}, fmt.Errorf("day header %q invalid", lines[0])
	}
	date, err := time.Parse("02/01/2006", header[2])
	if err != nil {
		return Day{}, fmt.Errorf("failed to parse %q", header[2])
	}

	var entries []Entry
	from := -1
	to := -1
	entryLines := lines[1:]
	for i, line := range entryLines {
		if strings.HasPrefix(line, "##") {
			if from == -1 {
				from = i
			} else {
				to = i - 1
			}
		} else if i == len(entryLines[1:])-1 {
			to = i
		}

		if from != -1 && to != -1 {
			entry, err := readEntry(entryLines[from : to+1])
			if err != nil {
				return Day{}, err
			}
			entries = append(entries, entry)

			from = -1
			to = -1
			if strings.HasPrefix(line, "##") {
				from = i
			}
		}
	}

	day.Date = date
	day.Entries = entries
	return day, nil
}

func readEntry(lines []string) (Entry, error) {
	var entry Entry

	if !strings.HasPrefix(lines[0], "##") {
		return Entry{}, fmt.Errorf("expected entry header, got %q", lines[0])
	}
	header := strings.Split(lines[0], " ")
	if len(header) != 2 {
		return Entry{}, fmt.Errorf("invalid entry header %q", lines[0])
	}
	time, err := time.Parse("15:04", header[1])
	if err != nil {
		return Entry{}, fmt.Errorf("failed to parse %q", header[1])
	}

	// Remove last line if empty
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	text := ""
	if len(lines) > 1 {
		text = strings.Join(lines[1:], "\n")
	}

	entry.Time = time
	entry.Text = text
	return entry, nil
}
