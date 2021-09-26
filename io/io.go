package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

// ReadFileByIOUtil 读取文件最快
func ReadFileByIOUtil() {
	data, err := ioutil.ReadFile("F:\\repo\\go\\README.md")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

// ReadWholeFile 读取整个文件
func ReadWholeFile() {
	file, err := os.Open("F:\\repo\\go\\README.md")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

// ReadFileByByte 读取文件, 按字节缓冲区
func ReadFileByByte() {
	file, err := os.Open("F:\\repo\\go\\README.md")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)
	bufReader := bufio.NewReader(file)
	data := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := bufReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
	}
	fmt.Println(string(data))
}

// ReadFileByLine 读取文件, 按行
func ReadFileByLine() {
	file, err := os.Open("F:\\repo\\go\\README.md")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)
	stat, err := file.Stat()
	size := stat.Size()
	fmt.Println("file size: ", size, "byte")
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("file read completed!")
				break
			} else {
				fmt.Println("file read error!", err)
				return
			}
		}
		line = strings.TrimSpace(line)
		fmt.Println(line)
	}
}

// WriteFileByIOUtil 每次写入覆盖原文件, 文件不存在则创建
func WriteFileByIOUtil() {
	data := []byte("第一行\n第二行\n")
	err := ioutil.WriteFile("F:\\repo\\go\\iotest.txt", data, 0644)
	if err != nil {
		panic(err)
	}
}

// WriteFileByOS
func WriteFileByOS() {
	filePath := "F:\\repo\\go\\iotest.txt"
	var file *os.File
	// 文件不存在时创建文件
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("file not exists!")
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Println("fail to create file!")
			return
		}
	}

	// 文件存在, 以追加的模式写入数据
	file, err := os.OpenFile(filePath, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error!")
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error")
			return
		}
	}(file)
	data := "追加\n追加\n"
	n, err := io.WriteString(file, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("writes ", n, "bytes")
}

func WriteFile() {
	filePath := "F:\\repo\\go\\iotest.txt"
	var file *os.File
	// 文件不存在时创建文件
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("file not exists!")
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Println("fail to create file!")
			return
		}
	}

	// 文件存在, 以追加的模式写入数据
	file, err := os.OpenFile(filePath, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error!")
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error")
			return
		}
	}(file)
	data := "追加\n追加\n"
	n, err := file.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("writes ", n, "bytes")
}

func WriteFileByBufIO() {
	filePath := "F:\\repo\\go\\iotest.txt"
	var file *os.File
	// 文件不存在时创建文件
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("file not exists!")
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Println("fail to create file!")
			return
		}
	}

	// 文件存在, 以追加的模式写入数据
	file, err := os.OpenFile(filePath, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error!")
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error")
			return
		}
	}(file)
	data := "追加一行\n追加一行\n"
	w := bufio.NewWriter(file)
	n, err := w.WriteString(data)
	if err != nil {
		panic(err)
	}
	err1 := w.Flush()
	if err1 != nil {
		return
	}
	fmt.Println("writes %d bytes", n)
}

func ReadFromStdin() {
	var data string
  _, err := fmt.Scanln(&data)
  if err != nil {
    return
  }
	fmt.Println(data)
}

func main() {
	//data, err := ReadFrom(os.Stdin, 11)

	// 从普通文件读取，其中 file 是 os.File 的实例
	//data, err := ReadFrom(file, 9)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(data)
	//ReadFromFile()

	//ReadFileByByte()

	//ReadFileByLine()

	//ReadFileByIOUtil()

	//WriteFileByIOUtil()

	//WriteFileByOS()

	//WriteFile()

	//WriteFileByBufIO()

  ReadFromStdin()

}
