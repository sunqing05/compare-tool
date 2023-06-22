package pkg

import (
	"os"
	"path/filepath"
	"strings"
)

func MoveDir(srcDir, destDir string) error {
	// 移动
	destFile := getDest(srcDir, destDir)
	_, err := os.Stat(destFile)
	if os.IsNotExist(err) {
		os.MkdirAll(destFile, os.ModePerm)
	} else if err != nil {
		return err
	}
	err = os.Rename(srcDir, filepath.Join(destFile, filepath.Base(srcDir)))
	if err != nil {
		return err
	}

	// 删除源目录
	err = os.RemoveAll(srcDir)
	if err != nil {
		return err
	}

	return nil
}

// 获取目标文件夹路径
func getDest(path, dest string) string {
	rel, _ := filepath.Rel(dest, filepath.Dir(path))
	dirs := strings.Split(rel, string(os.PathSeparator))
	rs := ""
	for i, v := range dirs {
		if i < 1 {
			//跳过第一层
			continue
		}
		rs = filepath.Join(rs, v)
	}
	return filepath.Join(dest, rs)
}
