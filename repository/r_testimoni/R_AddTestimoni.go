package r_testimoni

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"toko-buah/config/db"
	m_testimoni "toko-buah/model/m_testimoni"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

type AddTestimoniRepository interface {
	AddTestimoni(testimoni *m_testimoni.Testimoni, gambarFile *multipart.FileHeader) (*m_testimoni.Testimoni, error)
}

type addTestimoniRepository struct{}

func NewAddTestimoniRepository() AddTestimoniRepository {
	return &addTestimoniRepository{}
}

func (r *addTestimoniRepository) AddTestimoni(testimoni *m_testimoni.Testimoni, gambarFile *multipart.FileHeader) (*m_testimoni.Testimoni, error) {
	// Generate unique filename using UUID
	fileExt := filepath.Ext(gambarFile.Filename)
	newFileName := uuid.New().String() + fileExt

	// Open the source file
	src, err := gambarFile.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Determine the destination path
	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/testimoni/"
	savePath := filepath.Join(basePath, newFileName)
	// savePath := "./uploads/" + newFileName
	dst, err := os.Create(savePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(dst, src)
	if err != nil {
		return nil, err
	}

	// Open the saved image
	img, err := imaging.Open(savePath)
	if err != nil {
		return nil, err
	}

	// Resize the image to 500x500
	resizedImg := imaging.Resize(img, 150, 150, imaging.Lanczos)

	// Save the resized image
	err = imaging.Save(resizedImg, savePath)
	if err != nil {
		return nil, err
	}

	// Menyimpan di database hanya nama file saja
	testimoni.Gambar = newFileName

	// Simpan produk ke database
	err = db.Server().Create(&testimoni).Error
	if err != nil {
		return nil, err
	}
	return testimoni, nil
}
