package database

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	. "project/tables"

	"gorm.io/gorm"
)

const DB_TEMPLATE_PATH = "db_media/template.anki2"
const MEDIA_TEMPLATE_PATH = "db_media/media.template.db2"
const MAIN_DB = "Anki2/User 1/collection.anki2"

func ExportDB(decks string, name string) (err error) {
	apkgExport, err := os.Create(name + ".apkg")
	if err != nil {
		fmt.Println("Error creating .apkg file:", err)
		return
	}
	defer apkgExport.Close()

	zipWriter := zip.NewWriter(apkgExport)
	if err != nil {
		panic(err)
	}

	defer zipWriter.Close()

	CreateEmptyTemplate(name)
	//defer os.Remove(name + ".anki2")

	exportDB, err := OpenDB(name + ".anki2")
	if err != nil {
		log.Fatal(err)
		return
	}
	source, err := OpenDB(MAIN_DB)
	if err != nil {
		log.Fatal(err)
		return
	}
	CopyDatabase(source, exportDB)
	ZipFiles(zipWriter, name)

	return
}

func CopyDatabase(source *gorm.DB, target *gorm.DB) error {
	if err := ImportDB[Deck](source, target, "decks"); err != nil {
		fmt.Println("Error importing decks:", err)
		log.Fatal(err)
		return err
	}
	if err := ImportDB[Card](source, target, "cards"); err != nil {
		fmt.Println("Error importing decks:", err)
		log.Fatal(err)
		return err
	}
	if err := ImportDB[Note](source, target, "notes"); err != nil {
		fmt.Println("Error importing decks:", err)
		log.Fatal(err)
		return err
	}
	return nil
}

func CreateEmptyTemplate(name string) (err error) {
	cmd := exec.Command("cp", DB_TEMPLATE_PATH, name+".anki2")
	if err = cmd.Run(); err != nil {
		fmt.Println("Error creating template", err)
		return
	}
	return
}

func ZipFiles(zipWriter *zip.Writer, name string) (err error) {
	files := []string{name + ".anki2", MEDIA_TEMPLATE_PATH}
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			fmt.Println("Error adding file to zip:", err)
			return
		}
	}
	return
}

func AddFileToZip(zipWriter *zip.Writer, file string) (err error) {
	fileToZip, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileToZip.Close()

	writer, err := zipWriter.Create(filepath.Base(file))
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	return
}
