package mail_services

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"path/filepath"
	"strings"

	"billionmail-core/api/mail_services/v1"
)

func (c *ControllerV1) GetConfigFile(ctx context.Context, req *v1.GetConfigFileReq) (res *v1.GetConfigFileRes, err error) {
	res = &v1.GetConfigFileRes{}

	// get service config file pathe config file path
	var configPath string
	switch req.ServiceType {
	case v1.ServiceTypePostfix:
		configPath = v1.ServiceType_Postfix
	case v1.ServiceTypeDovecot:
		configPath = v1.ServiceType_Dovecot
	case v1.ServiceTypeRspamd:
		configPath = v1.ServiceType_Rspamd
	default:
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid service type")
	}

	// check if the file exists
	if !gfile.Exists(configPath) {
		return nil, gerror.NewCode(gcode.CodeNotFound, "config file not found: "+configPath)
	}

	// get file info
	fileInfo, err := os.Stat(configPath)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "get file info failed: "+err.Error())
	}

	// check if the file format is supported for online editing
	ext := filepath.Ext(configPath)
	if ext != "" {
		ext = ext[1:] // remove the leading dot
	}

	unsupportedExts := []string{"gz", "zip", "rar", "exe", "db", "pdf", "doc", "xls", "docx",
		"xlsx", "ppt", "pptx", "7z", "bz2", "png", "gif", "jpg", "jpeg", "bmp",
		"icon", "ico", "pyc", "class", "so", "pyd"}

	for _, unsupportedExt := range unsupportedExts {
		if strings.EqualFold(ext, unsupportedExt) {
			return nil, gerror.NewCode(gcode.CodeInvalidParameter, "unsupported file format")
		}
	}

	// file size limit (3MB)
	maxSize := int64(3 * 1024 * 1024)
	fileSize := fileInfo.Size()

	// fill response data
	res.Data.ConfigPath = configPath
	res.Data.LastModified = fileInfo.ModTime().Unix()
	res.Data.Size = fileSize

	// read file content
	if fileSize > maxSize {
		// file too large, only read the last part
		content, err := tailFile(configPath, 10000)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, "read file failed: "+err.Error())
		}

		res.Data.Content = content
		res.Data.OnlyRead = true
		res.Data.Encoding = "utf-8"
	} else {
		// read file normally
		content := gfile.GetContents(configPath)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, "read file failed: "+err.Error())
		}

		res.Data.Content = content
		res.Data.Encoding = "utf-8" // default assume utf-8 encoding
		res.Data.OnlyRead = false
	}

	res.SetSuccess("config file get success")
	return res, nil
}

// tailFile read the last few lines of a file
func tailFile(filename string, maxLines int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// get file size
	info, err := file.Stat()
	if err != nil {
		return "", err
	}
	fileSize := info.Size()

	// decide the start position
	var startPos int64
	if fileSize > int64(maxLines*100) { // assume each line is 100 bytes
		startPos = fileSize - int64(maxLines*100)
	} else {
		startPos = 0
	}

	// move to the start position
	_, err = file.Seek(startPos, 0)
	if err != nil {
		return "", err
	}

	// read content
	buffer := make([]byte, fileSize-startPos)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	content := string(buffer)

	// if not read from the beginning of the file, add a prompt message
	if startPos > 0 {
		return "... (file too large, only show the last part) ...\n" + content, nil
	}

	return content, nil
}
