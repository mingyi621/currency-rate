package models

import (
	"os"
)

type LogFile struct {
	os.File
}

func NewLogFileWithOsFile(f *os.File) *LogFile {
	return &LogFile{*f}
}

func (f *LogFile) GetLastLine() (string, error) {
	fileInfo, err := f.Stat()
	if err != nil {
		return "", err
	}
	str := ""
	size := fileInfo.Size()
	buf := make([]byte, 1)
	started := false
	for i := size - 1; i >= 0; i-- {
		f.ReadAt(buf, i)
		if buf[0] == '\n' {
			if started {
				break
			}
			started = true
			continue
		}
		str += string(buf[0])
	}
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result, nil
}

type FileNames []string

func GetHistoryFileNames() FileNames {
	files, err := os.ReadDir("history/")
	if err != nil {
		return nil
	}
	result := FileNames{}
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result
}

func (f FileNames) GetLatestFileName() string {
	if len(f) == 0 {
		return ""
	}
	return f[len(f)-1]
}
