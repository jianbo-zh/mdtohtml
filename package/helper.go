package main

import (
	"bytes"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
	"unicode"
)

// 检查错误
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 获取markdown文件的元数据
func getMdMd(filePath string) (map[string][]string, error) {

	metaMaps := make(map[string][]string)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, 500)
	n, err := file.Read(buf)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return metaMaps, nil
		}
		return nil, err
	}

	if bytes.HasPrefix(buf[:n], []byte("---")) {
		if endIdx := bytes.Index(buf[3:n], []byte("---")); endIdx > 0 {
			lis := bytes.Split(bytes.TrimSpace(buf[3:endIdx+3]), []byte("\n"))
			for _, li := range lis {
				if i := bytes.Index(li, []byte(":")); i > 0 {
					key := string(bytes.ToLower(li[:i]))
					val := li[i+1:]

					uvMap := make(map[string]bool)
					vli := bytes.Split(val, []byte(","))
					for _, v := range vli {
						if t := bytes.TrimSpace(v); len(t) > 0 {
							st := string(bytes.ToLower(t))
							if _, ok := uvMap[st]; !ok {
								metaMaps[key] = append(metaMaps[key], st)
								uvMap[st] = true
							}
						}
					}
				}
			}
		}
	}

	return metaMaps, nil
}

// 获取markdown文件的内容
func getMdCnt(filePath string) ([]byte, error) {
	mdCnt, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	mdCnt = bytes.TrimLeftFunc(mdCnt, unicode.IsSpace)

	if bytes.HasPrefix(mdCnt, []byte("---")) {
		if index := bytes.Index(mdCnt[3:], []byte("---")); index > 0 {
			mdCnt = mdCnt[index+6:]
		}
	}

	return mdCnt, nil
}

// 获取命令行选项
func getCliOpts() (string, string, bool, error) {
	var srcDir, dstDir string
	var isLocal bool

	flag.StringVar(&srcDir, "src-dir", "./", "the markdown file path")
	flag.StringVar(&dstDir, "dst-dir", "./html", "where to put the generated html")
	flag.BoolVar(&isLocal, "local", false, "the generated page is local access or network access?")
	flag.Parse()

	if _, err := os.Stat(srcDir); err != nil {
		return srcDir, dstDir, false, err
	}

	return srcDir, dstDir, isLocal, nil
}

// 查看文件时间
func getFileTime(filename string) (time.Time, error) {
	finfo, err := os.Stat(filename)
	if err != nil {
		return time.Now(), err
	}

	return finfo.ModTime(), nil
}
