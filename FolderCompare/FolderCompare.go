// FolderCompare project main.go
// 基于MD5或sha算法对两个文件夹中的文件进行对比，给出文件的差异清单。
// 如果是指定的文本文件，对文本文件的行进行对比。（后期扩展，可以考虑针对不同的文件扩展名采用不同的算法进行比较）

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Fctype string       //设置比较算法  MD5 或 SHA
var Fsource string      //比较的源目录  文件包
var Fdestination string //比较的目标目录  基线目录
var Foutput string      //设置输出的文件名
var Fhelp bool          //是否使用了 -h

func init() {
	flag.BoolVar(&Fhelp, "h", false, "查看帮助信息")
	flag.StringVar(&Fdestination, "d", "", "比较的目标目录（基线目录）,该参数必须设定")
	flag.StringVar(&Foutput, "o", "", "设置文件差异清单输出的文件，默认输出到屏幕")
	flag.StringVar(&Fsource, "s", ".", "比较的源目录 默认为当前目录")
	flag.StringVar(&Fctype, "t", "MD5", "设置比较算法 MD5 或 SHA  默认为MD5")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `FolderCompare (v1.0)
  基于源目录与目标目录中的文件进行比较，给出源目录中不同的文件清单。文件的比较算法基于MD5或SHA算法；
  用法：FolderCompare -t=MD5|SHA -s=源目录 -d=目标目录 -o=输出差异清单文件名
  -d string
        比较的目标目录（基线目录）,该参数必须设定
  -h    查看帮助信息
  -o string
        设置文件差异清单输出的文件，默认输出到屏幕
  -s string
        比较的源目录 默认为当前目录 (default ".")
  -t string
        设置比较算法 MD5 或 SHA  默认为MD5 (default "MD5")
        
  作者：宋京 by 2020年1月
  源码：https://github.com/hasjing/MyTools/FolderCompare
`)
	//	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if Fhelp {
		flag.Usage()
		os.Exit(1)
	}
	if Fdestination == "" {
		fmt.Println("未指定 -d 参数，必须设定比较的目标目录。")
		flag.Usage()
		os.Exit(2)
	}
	fmt.Println("Fctype==", Fctype)
	fmt.Println("Fsource==", Fsource)
	fmt.Println("Fdestination==", Fdestination)
	fmt.Println("Foutput==", Foutput)
	Fdestination = strings.ReplaceAll(Fdestination, "\\", "/")
	Fsource = strings.ReplaceAll(Fsource, "\\", "/")
	//	name := "GenRsaKey/rsa_private_key.pem"
	fmt.Println("Openfile=", Fdestination)
	if contents, err := ioutil.ReadFile(Fdestination); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		//   result := strings.Replace(string(contents),"\n","",1)
		fmt.Println(string(contents))
	}
}
