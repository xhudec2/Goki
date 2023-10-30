package database

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const DB_TEMPLATE_PATH = "db_media/template.anki2"
const MEDIA_TEMPLATE_PATH = "db_media/media.template.db2"
const MAIN_DB = "Anki2/User 1/collection.anki2"

func Export_db(decks string, name string) (err error) {
	apkg_export, err := os.Create(name + ".apkg")
	if err != nil {
		fmt.Println("Error creating .apkg file:", err)
		return
	}
	defer apkg_export.Close()

	zip_writer := zip.NewWriter(apkg_export)
	if err != nil {
		panic(err)
	}

	defer zip_writer.Close()

	write_data(decks, zip_writer, name)

	return
}

func Create_empty_template(name string) (err error) {
	cmd := exec.Command("cp", DB_TEMPLATE_PATH, name+".anki2")
	if err = cmd.Run(); err != nil {
		fmt.Println("Error creating template", err)
		return
	}
	return
}

func write_data(decks string, zip_writer *zip.Writer, name string) (err error) {
	collection, err := os.Open(DB_TEMPLATE_PATH)
	if err != nil {
		panic(err)
	}
	defer collection.Close()

	temp_writer, err := zip_writer.Create(name + ".anki2")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(temp_writer, collection); err != nil {
		panic(err)
	}

	media, err := os.Open(MEDIA_TEMPLATE_PATH)
	if err != nil {
		panic(err)
	}
	defer media.Close()

	temp_writer, err = zip_writer.Create(name + ".media.db2")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(temp_writer, media); err != nil {
		panic(err)
	}
	return
}
