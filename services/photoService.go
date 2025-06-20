package services

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/su-andrey/kr_aip/config"
	"github.com/su-andrey/kr_aip/database"
	"github.com/su-andrey/kr_aip/models"
)

func UploadToCloudinary(file multipart.File, fileHeader *multipart.FileHeader, cfg config.Config) (string, error) {
	cld, _ := cloudinary.NewFromParams(cfg.CloudinaryCloudName, cfg.CloudinaryAPIKey, cfg.CloudinaryAPISecret)
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}

func DeletePhotoFromClodinary(cfg config.Config, url string) error {
	cld, _ := cloudinary.NewFromParams(cfg.CloudinaryCloudName, cfg.CloudinaryAPIKey, cfg.CloudinaryAPISecret)
	parts := strings.Split(url, "/")
	publicIDWithExt := parts[len(parts)-1]
	publicID := strings.TrimSuffix(publicIDWithExt, path.Ext(publicIDWithExt))

	_, err := cld.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: publicID})
	if err != nil {
		return err
	}

	return nil
}

func GetPhotos(ctx context.Context, optCondition ...Condition) ([]models.Photo, error) {
	whereStatement := ""
	args := []any{}
	if len(optCondition) > 0 {
		whereStatement = fmt.Sprintf(" WHERE %s %s $1", optCondition[0].Name, optCondition[0].Operator)
		args = append(args, optCondition[0].Value)
	}

	var photos []models.Photo

	rows, err := database.DB.Query(ctx,
		`SELECT id, post_id, url FROM photos`+whereStatement, args...)
	if err != nil {
		return photos, errors.New("error getting posts from DB")
	}
	defer rows.Close()

	for rows.Next() {
		var photo models.Photo
		err := rows.Scan(&photo.ID, &photo.PostID, &photo.Url)
		if err != nil {
			return photos, errors.New("error processing data")
		}

		photos = append(photos, photo)
	}

	return photos, nil
}

func UploadPostPhotos(ctx context.Context, cfg config.Config, postID string, files []*multipart.FileHeader) ([]string, error) {
	var urls []string

	for _, fileHeader := range files {
		if fileHeader.Size > int64(cfg.MaxImageWeight) {
			return urls, errors.New("file size extends limits")
		}

		file, err := fileHeader.Open()
		if err != nil {
			return urls, errors.New("error opening fileheader")
		}

		buf := make([]byte, 512)
		_, err = file.Read(buf)
		if err != nil {
			file.Close()
			return urls, errors.New("error opening filereader")
		}

		fileType := http.DetectContentType(buf)
		if !cfg.AllowedImagesTypes[fileType] {
			file.Close()
			return urls, errors.New("unsupported file type: " + fileType)
		}
		file.Seek(0, 0)

		tx, err := database.DB.Begin(ctx)
		if err != nil {
			return urls, errors.New("error starting transaction")
		}
		defer tx.Rollback(ctx)

		url, err := UploadToCloudinary(file, fileHeader, cfg)
		file.Close()

		if err != nil {
			return urls, errors.New("error uploading to cloud")
		}
		urls = append(urls, url)

		_, err = tx.Exec(ctx, "INSERT INTO photos (post_id, url) VALUES ($1, $2)", postID, url)
		if err != nil {
			return urls, errors.New("error inserting url to DB")
		}
		err = tx.Commit(ctx)
		if err != nil {
			return urls, errors.New("error commiting changes")
		}
	}

	return urls, nil
}

func DeletePhoto(ctx context.Context, id string) error {
	var url string
	err := database.DB.QueryRow(ctx, `
		SELECT url FROM photos WHERE id = $1`, id).Scan(&url)

	if err != nil {
		return errors.New("photo not found")
	}

	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return errors.New("error starting transaction")
	}
	defer tx.Rollback(ctx)

	cfg := config.LoadConfig()
	err = DeletePhotoFromClodinary(cfg, url)
	if err != nil {
		return errors.New("error deleting photo from cloudinary")
	}

	_, err = tx.Exec(ctx, `
		DELETE FROM photos WHERE id = $1`, id)
	if err != nil {
		return errors.New("error deleting from DB")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.New("error commiting transaction")
	}

	return nil
}
