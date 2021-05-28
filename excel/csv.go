package excel

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

//XLSXFile ...
type XLSXFile struct {
	f         *os.File
	w         *csv.Writer
	fileName  string
	shortName string
}

//FullName ...
func (x *XLSXFile) FullName() string {
	return x.fileName
}

//ShortName ...
func (x *XLSXFile) ShortName() string {
	return x.shortName
}

//Close ...
func (x *XLSXFile) Close() error {
	if x == nil || x.f == nil {
		return nil
	}
	err := x.f.Close()
	if err == nil {
		//TODO设置定时器移除文件
		//filetool.RemoveFileAfter(x.fileName, define.ExportFileDuration*time.Second)
	}
	return err
}

//Write ...
func (x *XLSXFile) Write(cells []string) error {
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
func (x *XLSXFile) WriteAll(records [][]string) error {
	if x == nil || x.w == nil {
		return nil
	}
	return x.w.WriteAll(records)
}

//Flush ...
func (x *XLSXFile) Flush() {
	if x == nil || x.w == nil {
		return
	}
	x.w.Flush()
}

//NewFileSteamSysTem 保存到系统临时文件
func NewFileSteamSysTem(sheet string, headers []string) (*XLSXFile, error) {
	tf, err := filetool.CreateTempFile(define.ExportTeamPath)
	if err != nil {
		return nil, err
	}
	fileName := tf.Name()
	// 写入UTF-8 BOM
	if _, err := tf.WriteString("\xEF\xBB\xBF"); err != nil {
		return nil, err
	}
	w := csv.NewWriter(tf)
	csv := &XLSXFile{
		f:         tf,
		w:         w,
		fileName:  fileName,
		shortName: filepath.Base(fileName),
	}
	if err := csv.Write(headers); err != nil {
		_ = tf.Close()
		return nil, err
	}

	return csv, nil
}
