package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type mp4 struct {
	Path     string
	Basename string
	Ext      string
}

type info struct {
	Path     string
	Basename string
	Ext      string
}

type PartTime struct {
	Time int `json:"Time"`
}

type result struct {
	Type                 int        `json:"Type"`
	GroupKey             string     `json:"GroupKey"`
	File                 string     `json:"File"`
	IsSelected           bool       `json:"IsSelected"`
	IsLastHit            bool       `json:"IsLastHit"`
	DownloadState        string     `json:"DownloadState"`
	Aid                  int        `json:"Aid"`
	Cid                  int        `json:"Cid"`
	SeasonId             int        `json:"SeasonId"`
	EpisodeId            int        `json:"EpisodeId"`
	Title                string     `json:"Title"`
	Uploader             string     `json:"Uploader"`
	Description          string     `json:"Description"`
	CoverURL             string     `json:"CoverURL"`
	Tag                  string     `json:"Tag"`
	From                 string     `json:"From"`
	PartNo               string     `json:"PartNo"`
	PartName             string     `json:"PartName"`
	Format               int        `json:"Format"`
	TotalParts           int        `json:"TotalParts"`
	DownloadTimeRelative int        `json:"DownloadTimeRelative"`
	CreateDate           string     `json:"CreateDate"`
	TotalTime            string     `json:"TotalTime"`
	TotalTimeString      string     `json:"TotalTimeString"`
	PartTime             []PartTime `json:"PartTime"`
	TotalSizeString      string     `json:"TotalSizeString"`
	IsSinglePart         bool       `json:"IsSinglePart"`
	IsDownloaded         bool       `json:"IsDownloaded"`
	HasDanmaku           bool       `json:"HasDanmaku"`
	FontSize6            float32    `json:"FontSize6"`
	FontSize8            float32    `json:"FontSize8"`
	FontSize10           float32    `json:"FontSize10"`
	FontSize12           float32    `json:"FontSize12"`
	FontSize14           float32    `json:"FontSize14"`
	FontSize16           float32    `json:"FontSize16"`
	FontSize18           float32    `json:"FontSize18"`
	FontSize20           float32    `json:"FontSize20"`
	FontSize22           float32    `json:"FontSize22"`
	FontSize24           float32    `json:"FontSize24"`
	FontSize30           float32    `json:"FontSize30"`
	FontSize34           float32    `json:"FontSize34"`
}

func main() {
	var Finfo os.FileInfo
	var Mp4 mp4
	var Info info
	files := Scandir(".")

	for _, file := range files {
		Finfo, _ = os.Stat(file)
		if Finfo.IsDir() {
			FMp4 := Scandir(file)
			for _, fl := range FMp4 {
				if strings.EqualFold(path.Ext(fl), ".info") {
					Info.Path, Info.Basename = filepath.Split(fl)
					Info.Ext = path.Ext(fl)
				} else {
					Mp4.Path, Mp4.Basename = filepath.Split(fl)
					Mp4.Ext = path.Ext(fl)
				}
			}
			Nfname := OpenJson(Info.Path + Info.Basename)
			fmt.Println(Mp4.Path+Mp4.Basename, " ==> ", filepath.Base(file)+`-`+Nfname+Mp4.Ext)
			err := os.Rename(Mp4.Path+Mp4.Basename, filepath.Base(file)+`-`+Nfname+Mp4.Ext)
			if err != nil {
				panic(err)
			}
		}
	}
}

func OpenJson(fileName string) string {
	var data result

	// 读取JSON文件内容 返回字节切片
	bytes, _ := ioutil.ReadFile(fileName)
	//	fmt.Println("*** data.json content: ***")

	// 打印时需要转为字符串
	//	fmt.Println(string(bytes))

	// 将字节切片映射到指定结构上
	json.Unmarshal(bytes, &data)

	//	fmt.Println("*** unmarshal result: ***")

	// 打印对象结构
	//	fmt.Println(data.PartName)
	return data.PartName
}

func Scandir(dirName string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}
	var fileList []string
	for _, file := range files {
		fileList = append(fileList, dirName+string(os.PathSeparator)+file.Name())
	}
	return fileList
}
