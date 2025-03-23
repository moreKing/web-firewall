package vars

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"path"
)

var (
	rootPath string
	crtFile  string
	keyFile  string
)

func init() {
	getwd, err := os.Getwd()
	if err != nil {
		g.Log().Fatal(context.Background(), "getwd err", err)
		return
	}
	rootPath = getwd

	crtFile = g.Cfg().MustGet(context.Background(), "server.httpsCertPath", path.Join(rootPath, "manifest/config/server.crt")).String()
	keyFile = g.Cfg().MustGet(context.Background(), "server.httpsKeyPath", path.Join(rootPath, "manifest/config/server.key")).String()

}

func GetRootPath() string {
	return rootPath
}

// GetCrtFile https crt file
func GetCrtFile() string {

	return crtFile
}

// GetKeyFile https key file
func GetKeyFile() string {
	return keyFile
}
