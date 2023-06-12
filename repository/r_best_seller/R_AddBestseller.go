package r_bestseller

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"toko-buah/config/db"
	m_bestseller "toko-buah/model/m_best_seller"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

type AddBestsellerRepository interface {
	AddBestseller(bestseller *m_bestseller.Bestseller, gambarFile *multipart.FileHeader) (*m_bestseller.Bestseller, error)
}

type addBestsellerRepository struct{}

func NewAddBestsellerRepository() AddBestsellerRepository {
	return &addBestsellerRepository{}
}

// Upload File Dengan Penyimpanan Local
func (r *addBestsellerRepository) AddBestseller(bestseller *m_bestseller.Bestseller, gambarFile *multipart.FileHeader) (*m_bestseller.Bestseller, error) {
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
	basePath := "../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/bestseller/"
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
	bestseller.Gambar = newFileName

	// Simpan bestseller ke database
	err = db.Server().Create(&bestseller).Error
	if err != nil {
		return nil, err
	}

	return bestseller, nil
}

// import (
// 	"io"
// 	"mime/multipart"
// 	"os"
// 	"path/filepath"
// 	"toko-buah/config/db"
// 	m_bestseller "toko-buah/model/m_best_seller"

// 	"github.com/disintegration/imaging"
// 	"github.com/google/uuid"
// 	"github.com/jlaffaye/ftp"
// )

// type AddBestsellerRepository interface {
// 	AddBestseller(bestseller *m_bestseller.Bestseller, gambarFile *multipart.FileHeader) (*m_bestseller.Bestseller, error)
// }

// type addBestsellerRepository struct{}

// func NewAddBestsellerRepository() AddBestsellerRepository {
// 	return &addBestsellerRepository{}
// }

// func (r *addBestsellerRepository) AddBestseller(bestseller *m_bestseller.Bestseller, gambarFile *multipart.FileHeader) (*m_bestseller.Bestseller, error) {
// 	// Generate unique filename using UUID
// 	fileExt := filepath.Ext(gambarFile.Filename)
// 	newFileName := uuid.New().String() + fileExt

// 	// Open the source file
// 	src, err := gambarFile.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer src.Close()

// 	// Create FTP client
// 	client, err := ftp.Dial("ftp.example.com:21") // Replace with your 000webhost FTP server details
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Quit()

// 	// Login to FTP server
// 	err = client.Login("your-ftp-username", "your-ftp-password") // Replace with your 000webhost FTP login credentials
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Change working directory to the desired path on the server
// 	err = client.ChangeDir("public_html/uploads/bestseller") // Replace with the desired directory path on the server
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Determine the destination path on the server
// 	savePath := newFileName

// 	// Create a new file on the server
// 	dst, err := client.Retr(savePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer dst.Close()

// 	// Create a local file to save the retrieved content
// 	localFilePath := filepath.Join("../../../../../../../laragon/www/Toko-Buah-Admin/pages/uploads/bestseller/", newFileName) // Replace with the desired local directory path
// 	localFile, err := os.Create(localFilePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer localFile.Close()

// 	// Copy the contents from the FTP server to the local file
// 	_, err = io.Copy(localFile, dst)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Open the saved image
// 	img, err := imaging.Open(localFilePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Resize the image to 500x500
// 	resizedImg := imaging.Resize(img, 500, 500, imaging.Lanczos)

// 	// Save the resized image to the local file
// 	err = imaging.Save(resizedImg, localFilePath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Menyimpan di database hanya nama file saja
// 	bestseller.Gambar = newFileName

// 	// Simpan bestseller ke database
// 	err = db.Server().Create(&bestseller).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return bestseller, nil
// }
