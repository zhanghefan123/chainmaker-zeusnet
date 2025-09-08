package fileop

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines reads all lines of the file.
func ReadLines(path string) (result []string, err error) {
	// 创建文件handle和最终结果
	var file *os.File
	var lines []string
	// 为文件handle赋值
	file, err = os.Open(path)
	// 如果出现错误
	if err != nil {
		return nil, fmt.Errorf("error while open file: %w", err)
	}
	// 最后进行文件的关闭
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()
	// 进行循环的一行一行的读取
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func CreateFile(path string) (err error) {
	var file *os.File
	file, err = os.Create(path)
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()
	if err != nil {
		return fmt.Errorf("error while open file: %w", err)
	}
	return nil
}
