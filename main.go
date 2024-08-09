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
	inputPath string
	shift     int32
)

var version = "0.0.0-dev"

func init() {
	flaggy.SetName("ass-shifter")
	flaggy.SetDescription("ASS subtitle shifter")
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.DefaultParser.AdditionalHelpPrepend = "https://github.com/Nigh/subtitle-ass-shifter"
	flaggy.AddPositionalValue(&inputPath, "path", 1, true, "the subtitle path to shift")
	flaggy.Int32(&shift, "t", "shift", "shift ms")
	flaggy.SetVersion(version)
	flaggy.Parse()
}
func main() {
	if shift == 0 {
		fmt.Println("0ms shift means nothing to do.")
		return
	}
	inputPath, _ = filepath.Abs(inputPath)
	_, err := os.Stat(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	filepath.Walk(inputPath, walker)
}

func assShift(realPath string, shift int32) {
	assFile, err := os.ReadFile(realPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(assFile), "\n")
	// 0:00:10.19
	re := regexp.MustCompile(`(-?\d+):(\d\d):(\d\d)\.(\d\d)`)
	newLines := make([]string, 0)
	for _, v := range lines {
		matches := re.FindAllStringSubmatch(v, -1)
		for _, match := range matches {
			hours, _ := strconv.Atoi(match[1])
			minutes, _ := strconv.Atoi(match[2])
			seconds, _ := strconv.Atoi(match[3])
			milliseconds, _ := strconv.Atoi(match[4])

			sign := 1
			if hours < 0 {
				sign = -1
			}
			minutes *= sign
			seconds *= sign
			milliseconds *= sign

			totalMs := (hours*3600+minutes*60+seconds)*1000 + milliseconds*10
			totalMs += int(shift)

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

	// savePath := filepath.Dir(realPath) + "\\shifted_" + filepath.Base(realPath)
	// fmt.Println(savePath)
	err = os.WriteFile(realPath, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Println("[ERROR] "+filepath.Base(realPath), err)
		return
	}
	fmt.Println("[SUCCESS] Shifted " + strconv.Itoa(int(shift)) + "ms -> " + filepath.Base(realPath))
}

func walker(realPath string, f os.FileInfo, err error) error {
	ext := filepath.Ext(f.Name())
	if f.Name()[0] == '.' {
		return filepath.SkipDir
	}
	if strings.ToLower(ext) != ".ass" {
		return nil
	}
	// fmt.Println(realPath)
	assShift(realPath, shift)
	return nil
}
