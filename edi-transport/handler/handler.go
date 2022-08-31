package gsftp

import (
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/????lab/pkg/log"
	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

var logger *zap.Logger

func init() {
	// nolint:errcheck
	logger, _ = log.NewLogger()
}

// Example Handlers
func (fs *gcsHandler) Fileread(r *sftp.Request) (io.ReaderAt, error) {
	object := fs.bucket.Object(r.Filepath[1:])

	logger.Debug("reading-file", zap.String("name", r.Filepath))

	reader, err := object.NewReader(r.Context())
	if err != nil {
		return nil, err
	}

	return NewReadAtBuffer(reader)
}

func (fs *gcsHandler) Filewrite(r *sftp.Request) (io.WriterAt, error) {
	object := fs.bucket.Object(r.Filepath[1:])

	logger.Debug("writing-file", zap.String("name", r.Filepath))

	writer := object.NewWriter(r.Context())

	return NewWriteAtBuffer(writer, []byte{}), nil
}

func (fs *gcsHandler) Filecmd(r *sftp.Request) error {
	switch r.Method {
	case "Setstat":
		return nil
	case "Rename":
		return fmt.Errorf("not implemented")
	case "Rmdir", "Remove":
		return fmt.Errorf("not implemented")
	case "Mkdir":
		object := fs.bucket.Object(r.Filepath[1:] + "/")

		logger.Debug("creating-directory", zap.String("name", r.Filepath))

		writer := object.NewWriter(r.Context())

		err := writer.Close()
		return err
	case "Symlink":
		return fmt.Errorf("not implemented")
	}
	return nil
}

func (fs *gcsHandler) Filelist(r *sftp.Request) (sftp.ListerAt, error) {
	switch r.Method {
	case "List":
		logger.Debug("listing-directory", zap.String("path", r.Filepath))

		prefix := r.Filepath[1:]
		if prefix != "" {
			prefix += "/"
		}

		objects := fs.bucket.Objects(r.Context(), &storage.Query{
			Delimiter: "/",
			Prefix:    prefix,
		})

		list := []os.FileInfo{}

		for {
			objAttrs, err := objects.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				logger.Error("error-iterating-directory", zap.String("path", r.Filepath), zap.Error(err))

				return nil, err
			}

			// Don't include self.
			if ((prefix != "") && (objAttrs.Prefix == prefix)) || ((objAttrs.Prefix == "") && (objAttrs.Name == prefix)) {
				continue
			}

			list = append(list, &SyntheticFileInfo{
				prefix:  prefix,
				objAttr: objAttrs,
			})
		}

		return listerat(list), nil
	case "Stat":
		if r.Filepath == "/" {
			return listerat([]os.FileInfo{
				&SyntheticFileInfo{
					objAttr: &storage.ObjectAttrs{
						Prefix: "/",
					},
				},
			}), nil
		}

		object := fs.bucket.Object(r.Filepath[1:])

		logger.Debug("getting-file-info", zap.String("name", r.Filepath))

		attrs, err := object.Attrs(r.Context())
		if err == storage.ErrObjectNotExist {
			object := fs.bucket.Object(r.Filepath[1:] + "/")

			logger.Debug("retrieving-file-info", zap.String("path", r.Filepath+"/"))

			attrs, err = object.Attrs(r.Context())
		}
		if err != nil {
			return nil, err
		}

		file := &SyntheticFileInfo{
			objAttr: attrs,
		}
		return listerat([]os.FileInfo{file}), nil
	case "Readlink":
		return nil, fmt.Errorf("not implemented")
	}
	return nil, nil
}
