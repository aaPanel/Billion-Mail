package mail_services

import (
	"billionmail-core/internal/service/public"
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"

	"billionmail-core/api/mail_services/v1"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

func (c *ControllerV1) SaveConfigFile(ctx context.Context, req *v1.SaveConfigFileReq) (res *v1.SaveConfigFileRes, err error) {
	res = &v1.SaveConfigFileRes{}

	var configPath string
	//var serviceNameForCheck string
	switch req.ServiceType {
	case v1.ServiceTypePostfix:
		configPath = v1.ServiceType_Postfix
		//serviceNameForCheck = "Postfix"
	case v1.ServiceTypeDovecot:
		configPath = v1.ServiceType_Dovecot
		//serviceNameForCheck = "Dovecot"
	case v1.ServiceTypeRspamd:
		configPath = v1.ServiceType_Rspamd
		//serviceNameForCheck = "Rspamd"
	default:
		res.SetError(gerror.New(public.LangCtx(ctx, "Invalid service type")))
		return nil, err
	}

	dirPath := filepath.Dir(configPath)
	if !gfile.Exists(dirPath) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			res.SetError(gerror.New(public.LangCtx(ctx, "Failed to create directory: {}", err.Error())))
			return nil, err
		}
	}

	var backupPath string
	backupMade := false
	if gfile.Exists(configPath) {
		backupPath = configPath + ".bak." + time.Now().Format("20060102150405")
		if err := gfile.Copy(configPath, backupPath); err != nil {
			g.Log().Warningf(ctx, "Failed to back up configuration file %s: %v", configPath, err)
		} else {
			g.Log().Infof(ctx, "A configuration file backup has been created: %s", backupPath)
			backupMade = true
		}
	}

	contentBytes := []byte(req.Content)
	if err := os.WriteFile(configPath, contentBytes, 0644); err != nil {
		if backupMade {
			err = gfile.Move(backupPath, configPath)
		}
		res.SetError(gerror.New(public.LangCtx(ctx, "Failed to save file: {}", err.Error())))
		return nil, err
	}

	res.SetSuccess(public.LangCtx(ctx, "The configuration file was saved successfully"))
	return
}

func getDockerCheckDetails(ctx context.Context, serviceType v1.ServiceType) (containerName string, commandAndArgs []string, execUser string, serviceDisplayName string) {
	execUser = "root"

	switch serviceType {
	case v1.ServiceTypePostfix:
		containerName = "billionmail-postfix-billionmail-1"
		commandAndArgs = []string{"postfix", "check"}
		serviceDisplayName = "Postfix"
		return
	case v1.ServiceTypeDovecot:
		containerName = "billionmail-dovecot-billionmail-1"
		commandAndArgs = []string{"doveconf", "-n"}
		serviceDisplayName = "Dovecot"
		return
	case v1.ServiceTypeRspamd:
		containerName = "billionmail-rspamd-billionmail-1"
		commandAndArgs = []string{"rspamd", "configtest", "--insecure"}
		serviceDisplayName = "Rspamd"
		return
	default:
		return "", nil, "", "Unknown service"
	}
}

func extractCheckError(outputStr string, defaultError string) string {
	if len(outputStr) == 0 {
		return defaultError
	}
	lines := strings.Split(outputStr, "\n")
	for _, line := range lines {
		lowerLine := strings.ToLower(line)
		if strings.Contains(lowerLine, "error") ||
			strings.Contains(lowerLine, "fatal") ||
			strings.Contains(lowerLine, "failed") {
			if len(line) > 256 {
				return line[:256] + "..."
			}
			return line
		}
	}
	if len(lines) > 0 && len(lines[0]) > 0 {
		if len(lines[0]) > 256 {
			return lines[0][:256] + "..."
		}
		return lines[0]
	}
	return defaultError
}
