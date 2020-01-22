// GetHDInfo.go project main.go
package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

func main() {
	fmt.Println("操作系统：", runtime.GOOS)
	fmt.Println("CPU逻辑内核数:", runtime.NumCPU())
	i, id := getCpuId()
	fmt.Println("CPU序列号:", i, id)
	l, mac := getNicMac()
	fmt.Println("Mac 地址:", l, mac)
}

//  返回CPU数量与CPU 序列号
func getCpuId() (int, []string) {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		panic("CPU信息采集失败!")
	}
	str := string(out)
	//fmt.Println(str)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\n")
	str = reg.ReplaceAllString(str, "")
	strs := strings.Fields(str)
	sort.Strings(strs[1:])
	return (len(strs) - 1), strs[1:]
}

//  返回在线的网卡数量与Mac地址 wmic nic where NetEnabled=TRUE get MACAddress
func getNicMac() (int, []string) {
	//cmd := exec.Command("wmic", "nic where NetEnabled=TRUE get MACAddress")
	cmd := exec.Command("wmic", "nic", "where", "NetEnabled=TRUE", "get", "MACAddress")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		panic("网卡信息采集失败!")
	}
	str := string(out)
	//	fmt.Println(str)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\n")
	str = reg.ReplaceAllString(str, "")
	strs := strings.Fields(str)
	sort.Strings(strs[1:])
	return (len(strs) - 1), strs[1:]
}
