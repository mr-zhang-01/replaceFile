package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"strings"
	"strconv"
)

var (
	lineStr string
	mainFile string
	replaceFile string
)

func init(){
	//fmt.Println("-mn:", "输入行数，例子:-mn=5,10")
	//fmt.Println("-hostName:", "输入行数，例子:/home/mySelf/item/go/replaceFile/src/hosts")
	//fmt.Println("-replaceHostName:", "输入行数，例子:/home/mySelf/item/go/replaceFile/src/hosts_zpf")

	flag.StringVar(&lineStr, "mn", "", "输入行数，例子:-mn=5,10")
	flag.StringVar(&mainFile, "hostName", "", "输入hosts文件的路径")
	flag.StringVar(&replaceFile, "replaceHostName", "", "输入替换hosts文件的路径")
}

func main(){
	flag.Parse()

	if lineStr == ""{
		fmt.Println("mn 不能为空")
		return
	}

	mnData := strings.Split(lineStr, ",")

	startLine, err := strconv.Atoi(mnData[0])
	if err != nil{
		fmt.Println("解析错误1", err)
		return
	}

	endLine, err := strconv.Atoi(mnData[1])
	if err != nil{
		fmt.Println("解析错误2", err)
		return
	}

	if startLine > endLine{
		fmt.Println("起始行不能大于终止行")
		return
	}

	if mainFile == ""{
		fmt.Println("hostNmae 不能为空")
		return
	}

	if replaceFile == ""{
		fmt.Println("replaceHostName 不能为空")
		return
	}
	bContentAllMain := readFileLine(mainFile, startLine, endLine)
	//log.Println(bContentAllMain)

	bContentAllReplace := readFileLine(replaceFile, 0, 0)
	//log.Println(bContentAllReplace)

	f, err := os.OpenFile(mainFile, os.O_TRUNC|os.O_RDWR, 0755)
	//f, err := os.OpenFile(mainFile, os.O_RDWR, 0755)
	if err != nil{
		fmt.Println("文件打开失败")
	}

	for k, val := range bContentAllReplace{
		if k == 0{
			continue
		}
		if k + 1 >= 5 && k + 1 <= 10{
			for _, val1 := range bContentAllMain{
				f.WriteString(val1)
				f.WriteString("\n")
			}
		}

		f.WriteString(val)
		f.WriteString("\n")
		//fmt.Println(k, val, f)
	}
}

//读取文件
func readFileLine(name string, start, offset int) ([]string){

	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()

	if err != nil{
		fmt.Println("文件打开失败")
		return nil
	}

	bContentAll := make([]string, 1)
	bContent := make([]byte, 1)

	str := ""
	i := 1
	for {
		_, err := f.Read(bContent)
		if err != nil{
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}

		if string(bContent) == "\n"{
			if i >= start{
				if offset == 0{
					bContentAll = append(bContentAll, str)
				}else{
					if i <= offset{
						bContentAll = append(bContentAll, str)
					}
				}
			}

			i++
			str = ""
		}else{
			str += string(bContent)
		}
	}

	return bContentAll
}