package fileserver

import (
	"net/http"

	"github.com/Nico3012/webserver"
)

func CreateFileServer(addr string, certPath string, certKeyPath string, rootPath string) error {
	return webserver.CreateWebServer(addr, certPath, certKeyPath, http.FileServer(http.Dir(rootPath)))
}

func CreateWebServerAndCertificate(addr string, caPath string, caKeyPath string, hosts []string, rootPath string) error {
	return webserver.CreateWebServerAndCertificate(addr, caPath, caKeyPath, hosts, http.FileServer(http.Dir(rootPath)))
}
