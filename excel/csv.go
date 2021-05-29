package excel

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"time"

	"github.com/xmseekshine/eutil/filetool"
)

type (
	XLSXFile interface {
		FullName() string
		ShortName() string
		Close() error
		Write(cells []string) error
		WriteAll(records [][]string) error
		Flush()
	}
	defaultXLSXFile struct {
		f         *os.File
		w         *csv.Writer
		fileName  string
		shortName string
		cfg       ExportCfg
	}
	ExportCfg struct {
		ExportTeamPath     string // "./temp"
		ExportFileNameLen  uint   //32
		ExportFileDuration uint   //导出文件最大存活时间 单位秒
		ExportMaxPerPage   uint   //导出分页查询每页最大记录数
	}
)

//导出文件位置
// const (
// 	DefaultExportTeamPath     = "./temp"
// 	DefaultExportFileNameLen  = 32
// 	DefaultExportFileDuration = 600  //导出文件最大存活时间 单位秒
// 	DefaultExportMaxPerPage   = 1000 //导出分页查询每页最大记录数
// )

//NewFileSteamSysTem 保存到系统临时文件
func NewFileSteamSysTem(c ExportCfg, sheet string, headers []string) (XLSXFile, error) {

	tf, err := filetool.CreateTempFile(c.ExportTeamPath)
	if err != nil {
		return nil, err
	}
	fileName := tf.Name()
	// 写入UTF-8 BOM
	if _, err := tf.WriteString("\xEF\xBB\xBF"); err != nil {
		return nil, err
	}
	w := csv.NewWriter(tf)
	csv := &defaultXLSXFile{
		f:         tf,
		w:         w,
		fileName:  fileName,
		shortName: filepath.Base(fileName),
		cfg:       c,
	}
	if err := csv.Write(headers); err != nil {
		_ = tf.Close()
		return nil, err
	}

	return csv, nil
}

//XLSXFile ...
// type XLSXFile struct {
// 	f         *os.File
// 	w         *csv.Writer
// 	fileName  string
// 	shortName string
// }

//FullName ...
func (x *defaultXLSXFile) FullName() string {
	return x.fileName
}

//ShortName ...
func (x *defaultXLSXFile) ShortName() string {
	return x.shortName
}

//Close ...
func (x *defaultXLSXFile) Close() error {
	if x == nil || x.f == nil {
		return nil
	}
	err := x.f.Close()
	if err == nil {
		//TODO设置定时器移除文件
		//x.cfg.ExportFileDuration
		filetool.RemoveFileAfter(x.fileName, time.Duration(x.cfg.ExportFileDuration)*time.Second)
	}
	return err
}

//Write ...
func (x *defaultXLSXFile) Write(cells []string) error {
	if x == nil || x.w == nil {
		return nil
	}
	if err := x.w.Write(cells); err != nil {
		return err
	}
	x.Flush()
	return nil
}

//WriteAll ...
func (x *defaultXLSXFile) WriteAll(records [][]string) error {
	if x == nil || x.w == nil {
		return nil
	}
	return x.w.WriteAll(records)
}

//Flush ...
func (x *defaultXLSXFile) Flush() {
	if x == nil || x.w == nil {
		return
	}
	x.w.Flush()
}
