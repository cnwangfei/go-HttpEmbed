package httpEmbed

import (
	"embed"
	"io/fs"
	"path"
	"path/filepath"
)

type Fs struct {
	*embed.FS        // 静态资源
	Path      string // 设置embed文件到静态资源的相对路径，也就是embed注释里的路径
}

func (f Fs) Open(name string) (fs.File, error) {
	// 拼接相对路径
	fullName := filepath.Join(f.Path, filepath.FromSlash(path.Clean("/"+name)))
	// 修改http路径中的\为/
	fullName = filepath.ToSlash(fullName)
	//fmt.Println(fullName)
	fs, err := f.FS.Open(fullName)
	return fs, err
}
