package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/integrii/flaggy"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

var supportTypes map[string]subtitle = map[string]subtitle{
	".srt": &srtType{},
	".ass": &assType{},
}

var (
	inputPath  string
	shift      int
	gFrom      int
	gTo        int
	fromStr    string
	toStr      string
	fromRegexp string
	toRegexp   string
	dry        bool
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
	flaggy.String(&fromRegexp, "sr", "startRegexp", "start from regular expression")
	flaggy.String(&toRegexp, "er", "endRegexp", "end at regular expression")
	flaggy.Bool(&dry, "d", "dry", "dry run")
	flaggy.SetVersion(version)
	flaggy.Parse()
}

func parseFromTo() error {
	re := regexp.MustCompile(`(-?\d+):(\d\d):(\d\d)`)

	str2ms := func(str string) (ms int, err error) {
		matches := re.FindStringSubmatch(str)
		if matches == nil {
			return 0, fmt.Errorf("invalid time format, expected HH:MM:SS, example: 0:23:45")
		}

		hours, _ := strconv.Atoi(matches[1])
		minutes, _ := strconv.Atoi(matches[2])
		seconds, _ := strconv.Atoi(matches[3])

		sign := 1
		if hours < 0 {
			sign = -1
		}
		ms = (hours*3600 + minutes*60 + seconds) * 1000 * sign
		return
	}

	if fromStr != "" && fromRegexp != "" {
		return fmt.Errorf("multiple start conditions are not allowed")
	}

	if toStr != "" && toRegexp != "" {
		return fmt.Errorf("multiple end conditions are not allowed")
	}

	if fromStr != "" {
		ms, err := str2ms(fromStr)
		if err != nil {
			return err
		} else {
			gFrom = ms
		}
	}

	if toStr != "" {
		ms, err := str2ms(toStr)
		if err != nil {
			return err
		} else {
			gTo = ms
		}
	}

	if gTo != 0 && gFrom != 0 {
		if gFrom > gTo {
			return fmt.Errorf("end must be greater than start")
		}
	}

	if fromRegexp != "" {
		re, err := regexp.Compile(fromRegexp)
		if err != nil {
			return err
		}
		fromRegexp = re.String()
	}

	if toRegexp != "" {
		re, err := regexp.Compile(toRegexp)
		if err != nil {
			return err
		}
		toRegexp = re.String()
	}

	return nil
}

func timeInclude(t int, from int, to int) bool {
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

var fileUpdated = 0

func main() {
	if shift == 0 {
		fmt.Println("shift 0ms means nothing to do.")
		return
	}
	if err := parseFromTo(); err != nil {
		fmt.Println(err)
		return
	}

	inputPath, _ = filepath.Abs(inputPath)
	_, err := os.Stat(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	filepath.Walk(inputPath, walker)

	if dry {
		fmt.Println("\n[Info] Dry run, no file changes.")
	} else {
		fmt.Printf("\n[Info] %d subtitle files updated.\n", fileUpdated)
	}
}

func walker(realPath string, f os.FileInfo, err error) error {
	ext := filepath.Ext(f.Name())
	if f.Name()[0] == '.' {
		return filepath.SkipDir
	}
	ext = strings.ToLower(ext)

	if subs, ok := supportTypes[ext]; ok {
		subFile, err := os.ReadFile(realPath)
		fmt.Println("\n" + filepath.Base(realPath))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// Detect the encoding
		encoding, name, certain := charset.DetermineEncoding(subFile, "")
		if !certain {
			fmt.Printf("Warning: uncertain encoding detected for %s, assuming %s\n", realPath, name)
		}
		// Transcode to UTF-8
		utf8Bytes, _, err := transform.Bytes(encoding.NewDecoder(), subFile)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		subs.setContent(utf8Bytes)
		newLines := subtitleShift(subs, shift)
		if !dry {
			err = os.WriteFile(realPath, []byte(strings.Join(newLines, "\n")), 0644)
			if err != nil {
				fmt.Println("[ERROR] "+filepath.Base(realPath), err)
			} else {
				fileUpdated++
			}
		}
	}
	return nil
}

type subtitle interface {
	setContent(content []byte)
	getContent() []byte
	match2Ms(match []string) int
	time2Str(totalMs int) string
	re() *regexp.Regexp
}

type srtType struct {
	content []byte
}

func (s *srtType) setContent(content []byte) {
	s.content = content
}

func (s srtType) getContent() []byte {
	return s.content
}

func (s srtType) re() *regexp.Regexp {
	return regexp.MustCompile(`(-?\d+):(\d\d):(\d\d)\,(\d{1,3})`)
}

func (s srtType) match2Ms(match []string) int {
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
	return totalMs
}

func (s srtType) time2Str(totalMs int) string {
	sign := 1
	if totalMs < 0 {
		sign = -1
	}
	totalMs *= sign
	return fmt.Sprintf("%02d:%02d:%02d,%03d",
		totalMs/3600000*sign,
		(totalMs/60000)%60,
		(totalMs/1000)%60,
		totalMs%1000)
}

type assType struct {
	content []byte
}

func (s *assType) setContent(content []byte) {
	s.content = content
}

func (s assType) getContent() []byte {
	return s.content
}

func (s assType) re() *regexp.Regexp {
	return regexp.MustCompile(`(-?\d+):(\d\d):(\d\d)\.(\d{1,3})`)
}

func (s assType) match2Ms(match []string) int {

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
	return totalMs
}

func (s assType) time2Str(totalMs int) string {
	sign := 1
	if totalMs < 0 {
		sign = -1
	}
	totalMs *= sign
	return fmt.Sprintf("%d:%02d:%02d.%02d",
		totalMs/3600000*sign,
		(totalMs/60000)%60,
		(totalMs/1000)%60,
		totalMs%1000/10)
}

func subtitleShift(s subtitle, shift int) (newLines []string) {
	lines := strings.Split(string(s.getContent()), "\n")
	newLines = make([]string, 0)

	// if regexp is not empty, then use regexp to match time
	var reFrom, reTo *regexp.Regexp
	if fromRegexp != "" {
		reFrom = regexp.MustCompile(fromRegexp)
	}
	if toRegexp != "" {
		reTo = regexp.MustCompile(toRegexp)
	}
	from := gFrom
	to := gTo
	if fromRegexp != "" || toRegexp != "" {
		for _, v := range lines {
			matches := s.re().FindAllStringSubmatch(v, 2)
			if matches != nil {
				ms := s.match2Ms(matches[1])
				if fromRegexp != "" && from == 0 {
					if reFrom.MatchString(v) {
						from = ms
					}
				}
				if toRegexp != "" && to == 0 {
					if reTo.MatchString(v) {
						to = ms
					}
				}
			}
		}
	}

	linesShifted := 0
	for _, v := range lines {
		matches := s.re().FindAllStringSubmatch(v, 2)
		for _, match := range matches {
			totalMs := s.match2Ms(match)
			if !timeInclude(totalMs, from, to) {
				continue
			}
			totalMs += shift
			newTime := s.time2Str(totalMs)
			v = strings.Replace(v, match[0], newTime, 1)
			linesShifted++
		}
		newLines = append(newLines, v)
	}
	var _from string
	var _to string
	if from != 0 {
		_from = s.time2Str(from)
	} else {
		_from = "start"
	}
	if to != 0 {
		_to = s.time2Str(to)
	} else {
		_to = "end"
	}
	fmt.Printf("From %s to %s, %d lines shifted %dms\n", _from, _to, linesShifted, shift)
	return
}
