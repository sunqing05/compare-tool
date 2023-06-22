package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"compare-tool/conf"
	"compare-tool/pkg"
)

func main() {
	fmt.Printf("开始执行... \n")
	//读配置
	conf := conf.Config()

	//读data目录
	root := conf.DataDir
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("读取数据文件夹 '%s' 错误：%v \n", root, err)
		return
	}
	//重建diff文件夹
	for _, entry := range entries {
		//只处理文件夹
		if entry.IsDir() {
			//每组文件夹处理开始
			dataRoot := filepath.Join(root, entry.Name())
			//先删掉目录下的diff文件夹
			diffFolder := filepath.Join(dataRoot, conf.DiffFolderName)
			os.RemoveAll(diffFolder)
			//读取目录
			des, err := os.ReadDir(dataRoot)
			if err != nil {
				fmt.Printf("读取数据文件夹 '%s' 错误:%v \n", dataRoot, err)
				return
			}
			baseRoot := ""
			others := []string{}
			for _, de := range des {
				//找到base文件夹
				if !de.IsDir() {
					continue
				}
				p := filepath.Join(dataRoot, de.Name())
				if strings.Contains(de.Name(), conf.BaseFolderName) {
					baseRoot = p
				} else {
					others = append(others, p)
				}
			}
			if baseRoot == "" {
				fmt.Printf("数据文件夹 '%s' 下未找到含关键字 '%s' 的文件夹 \n", dataRoot, conf.BaseFolderName)
				return
			}

			//重建diff文件夹
			if err := os.MkdirAll(diffFolder, os.ModePerm); err != nil {
				fmt.Printf("创建文件夹 '%s' 失败 \n", diffFolder)
				return
			}
			//比较
			compare(dataRoot, baseRoot, others, diffFolder)
		}
	}
	fmt.Sprintln("执行完成")

	//等待手动结束终端
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("结束请输入: exit")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				return
			}
		} else {
			fmt.Printf("读取输入发生错误:%v,请手动关闭终端", scanner.Err())
		}
	}
}

func compare(root string, base string, others []string, diffFolder string) {
	//读base文件目录
	//文件夹下的文件名
	ffMap := make(map[string][]string)
	ens, err := os.ReadDir(base)
	if err != nil {
		fmt.Printf("读取数据文件夹 '%s' 错误:%v \n", base, err)
		return
	}
	for _, en := range ens {
		//这层也只处理文件夹
		if en.IsDir() {
			ffMap[en.Name()] = []string{}
			cbName := filepath.Join(base, en.Name())
			cens, err := os.ReadDir(cbName)
			if err != nil {
				fmt.Printf("读取数据文件夹 '%s' 错误:%v \n", cbName, err)
				return
			}
			for _, cen := range cens {
				ffMap[en.Name()] = append(ffMap[en.Name()], getFileName(cen.Name()))
			}
		}
	}
	//开始处理
	for _, o := range others {
		enos, err := os.ReadDir(o)
		if err != nil {
			fmt.Printf("读取数据文件夹 '%s' 错误:%v \n", o, err)
			return
		}
		for _, eno := range enos {
			if eno.IsDir() {
				nroot := filepath.Join(o, eno.Name())
				if _, ok := ffMap[eno.Name()]; !ok {
					if err := pkg.MoveDir(nroot, diffFolder); err != nil {
						fmt.Printf("移动文件夹 '%s' 错误:%v \n", nroot, err)
						return
					}
					continue
				}
				//否则比较里面的文件名
				cenos, err := os.ReadDir(nroot)
				if err != nil {
					fmt.Printf("读取数据文件夹 '%s' 错误:%v \n", nroot, err)
					return
				}
				for _, ceno := range cenos {
					//比较文件名是否在base文件夹中
					f := 0
					for _, i := range ffMap[eno.Name()] {
						if getFileName(ceno.Name()) == i {
							f = 1
						}
					}
					if f == 0 {
						froot := filepath.Join(nroot, ceno.Name())
						if err := pkg.MoveDir(froot, diffFolder); err != nil {
							fmt.Printf("移动文件夹 '%s' 错误:%v \n", froot, err)
							return
						}
					}
				}
			}
		}
	}
}

// 获取文件的文件名
func getFileName(path string) string {
	fileName := filepath.Base(path)
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
