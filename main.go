package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/integrii/flaggy"
)

var (
	inputPath   string
	shift       int
	from        int
	to          int
	fromStr     string
	toStr       string
	fileShifted int
)

var version = "1.0.0"

func init() {
	flaggy.SetName("ass-shifter")
	flaggy.SetDescription("ASS subtitle shifter")
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.DefaultParser.AdditionalHelpPrepend = "https://github.com/Nigh/subtitle-ass-shifter"
	flaggy.AddPositionalValue(&inputPath, "path", 1, true, "the subtitle path to shift")
	flaggy.Int(&shift, "t", "shift", "shift ms")
	flaggy.String(&fromStr, "s", "start", "start from HH:MM:SS")
	flaggy.String(&toStr, "e", "end", "end at HH:MM:SS")
	flaggy.SetVersion(version)
	flaggy.Parse()
}

func parseFromTo() error {
	re := regexp.MustCompile(`(-?\d+):(\d\d):(\d\d)`)

	if fromStr != "" {
		matches := re.FindStringSubmatch(fromStr)
		if matches == nil {
			return fmt.Errorf("invalid start time format, expected HH:MM:SS, example: 0:23:45")
		}

		hours, _ := strconv.Atoi(matches[1])
		minutes, _ := strconv.Atoi(matches[2])
		seconds, _ := strconv.Atoi(matches[3])

		sign := 1
		if hours < 0 {
			sign = -1
		}
		from = (hours*3600 + minutes*60 + seconds) * 1000 * sign
	}

	if toStr != "" {
		matches := re.FindStringSubmatch(toStr)

		if matches == nil {
			return fmt.Errorf("invalid end time format, expected HH:MM:SS, example: 1:32:54")
		}
		hours, _ := strconv.Atoi(matches[1])
		minutes, _ := strconv.Atoi(matches[2])
		seconds, _ := strconv.Atoi(matches[3])

		sign := 1
		if hours < 0 {
			sign = -1
		}
		to = (hours*3600 + minutes*60 + seconds) * 1000 * sign
	}

	return nil
}

func timeInclude(t int) bool {
	if from == 0 && to == 0 {
		return true
	}
	if from != 0 && t < from {
		return false
	}
	if to != 0 && t > to {
		return false
	}
	return true
}

func main() {
	if shift == 0 {
		fmt.Println("shift 0ms means nothing to do.")
		return
	}
	if err := parseFromTo(); err != nil {
		fmt.Println(err)
		return
	}
	if to != 0 && from != 0 {
		if from > to {
			fmt.Println("end must be greater than start")
			return
		}
	}
	inputPath, _ = filepath.Abs(inputPath)
	_, err := os.Stat(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	filepath.Walk(inputPath, walker)

	if fileShifted > 0 {
		fmt.Print("Total " + strconv.Itoa(fileShifted) + " files shifted " + strconv.Itoa(shift) + "ms ")
		if from != 0 || to != 0 {
			if from != 0 {
				fmt.Print("from " + fromStr)
			} else {
				fmt.Print("from start")
			}
			if to != 0 {
				fmt.Print(" to " + toStr)
			} else {
				fmt.Print(" to end")
			}
		}
		fmt.Println()
	}
}

func srtShift(realPath string) {
	srtFile, err := os.ReadFile(realPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(srtFile), "\n")
	// 00:01:36,649
	re := regexp.MustCompile(`(-?\d+):(\d\d):(\d\d)\,(\d{1,3})`)
	newLines := make([]string, 0)
	for _, v := range lines {
		matches := re.FindAllStringSubmatch(v, -1)
		for _, match := range matches {
			hours, _ := strconv.Atoi(match[1])
			minutes, _ := strconv.Atoi(match[2])
			seconds, _ := strconv.Atoi(match[3])
			milliseconds, _ := strconv.Atoi(match[4])

			for i := 0; i < 3-len(match[4]); i++ {
				milliseconds *= 10
			}

			sign := 1
			if hours < 0 {
				sign = -1
			}
			minutes *= sign
			seconds *= sign
			milliseconds *= sign

			totalMs := (hours*3600+minutes*60+seconds)*1000 + milliseconds
			if !timeInclude(totalMs) {
				continue
			}
			totalMs += shift

			if totalMs < 0 {
				sign = -1
			} else {
				sign = 1
			}
			totalMs *= sign

			newTime := fmt.Sprintf("%02d:%02d:%02d,%03d",
				totalMs/3600000*sign,
				(totalMs/60000)%60,
				(totalMs/1000)%60,
				totalMs%1000)

			v = strings.Replace(v, match[0], newTime, 1)
		}
		newLines = append(newLines, v)
	}

	err = os.WriteFile(realPath, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Println("[ERROR] "+filepath.Base(realPath), err)
		return
	}
	fmt.Println("[SUCCESS] " + filepath.Base(realPath))
	fileShifted++
}

func assShift(realPath string) {
	assFile, err := os.ReadFile(realPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(assFile), "\n")
	// 0:00:10.19
	re := regexp.MustCompile(`(-?\d+):(\d\d):(\d\d)\.(\d{1,3})`)
	newLines := make([]string, 0)
	for _, v := range lines {
		matches := re.FindAllStringSubmatch(v, -1)
		for _, match := range matches {
			hours, _ := strconv.Atoi(match[1])
			minutes, _ := strconv.Atoi(match[2])
			seconds, _ := strconv.Atoi(match[3])
			milliseconds, _ := strconv.Atoi(match[4])

			for i := 0; i < 3-len(match[4]); i++ {
				milliseconds *= 10
			}

			sign := 1
			if hours < 0 {
				sign = -1
			}
			minutes *= sign
			seconds *= sign
			milliseconds *= sign

			totalMs := (hours*3600+minutes*60+seconds)*1000 + milliseconds
			if !timeInclude(totalMs) {
				continue
			}
			totalMs += shift

			if totalMs < 0 {
				sign = -1
			} else {
				sign = 1
			}
			totalMs *= sign

			newTime := fmt.Sprintf("%d:%02d:%02d.%02d",
				totalMs/3600000*sign,
				(totalMs/60000)%60,
				(totalMs/1000)%60,
				totalMs%1000/10)

			v = strings.Replace(v, match[0], newTime, 1)
		}
		newLines = append(newLines, v)
	}

	err = os.WriteFile(realPath, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Println("[ERROR] "+filepath.Base(realPath), err)
		return
	}
	fmt.Println("[SUCCESS] " + filepath.Base(realPath))
	fileShifted++
}

func walker(realPath string, f os.FileInfo, err error) error {
	ext := filepath.Ext(f.Name())
	if f.Name()[0] == '.' {
		return filepath.SkipDir
	}
	switch strings.ToLower(ext) {
	case ".srt":
		srtShift(realPath)
	case ".ass":
		assShift(realPath)
	default:
		return nil
	}
	return nil
}
