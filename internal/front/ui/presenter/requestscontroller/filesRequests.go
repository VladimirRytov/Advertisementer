package requestscontroller

import (
	"context"
	"net/url"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
)

func (r *RequestsHandler) AllFiles(ctx context.Context) {
	go r.fileManager.AllFiles(ctx)
}

func (r *RequestsHandler) NewFile(ctx context.Context, file *presenter.File) error {
	go r.fileManager.NewFile(ctx, r.converter.FileToDto(file))
	return nil
}

func (r *RequestsHandler) FileByName(name string) error {
	go r.fileManager.FileByName(name)
	return nil
}

func (r *RequestsHandler) RemoveFile(name string) error {
	go r.fileManager.RemoveFileByName(name)
	return nil
}

func (r *RequestsHandler) FileMiniatureByName(name string) error {
	go r.fileManager.FileMiniatureByName(r.b64.ToBase64URLString([]byte(name)), "miniature")
	return nil
}

func (r *RequestsHandler) LargeFileByName(name string) error {
	go r.fileManager.FileMiniatureByName(r.b64.ToBase64URLString([]byte(name)), "large")
	return nil
}

func (r *RequestsHandler) UploadFiles(ctx context.Context, file string) {
	go r.fileManager.NewFileMultipart(ctx, file)
}

func (r *RequestsHandler) GetFileURI(fileName string) (string, error) {
	connectionInfo := r.req.ConnectionInfo()
	filePath, err := url.JoinPath(connectionInfo["adress"], connectionInfo["apiPath"], "files", r.b64.ToBase64URLString([]byte(fileName)))
	if err != nil {
		return "", err
	}
	urlPath, err := url.Parse(filePath)
	if err != nil {
		return "", err
	}
	urlPath.RawQuery = "token=" + connectionInfo["token"]
	return urlPath.String(), err
}
