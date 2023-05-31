package r_newarrival

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"toko-buah/config/db"
	m_newarrival "toko-buah/model/m_new_arrival"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

type AddNewArrivalRepository interface {
	AddNewArrival(newarrival *m_newarrival.NewArrival, gambarFile *multipart.FileHeader) (*m_newarrival.NewArrival, error)
}

type addNewArrivalRepository struct{}

func NewAddNewArrivalRepository() AddNewArrivalRepository {
	return &addNewArrivalRepository{}
}

func (r *addNewArrivalRepository) AddNewArrival(newarrival *m_newarrival.NewArrival, gambarFile *multipart.FileHeader) (*m_newarrival.NewArrival, error) {
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
	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/newarrival/"
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
	resizedImg := imaging.Resize(img, 500, 500, imaging.Lanczos)

	// Save the resized image
	err = imaging.Save(resizedImg, savePath)
	if err != nil {
		return nil, err
	}

	// Menyimpan di database hanya nama file saja
	newarrival.Gambar = newFileName

	// Simpan produk ke database
	err = db.Server().Create(&newarrival).Error
	if err != nil {
		return nil, err
	}
	return newarrival, nil
}
