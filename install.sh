#!/bin/bash



mkdir ~/Documents/goki

cp ./db_media/collection.anki2 ./db_media/template.anki2 ./db_media/media.template.db2 ~/Documents/goki/

go install fyne.io/fyne/v2/cmd/fyne@latest

cd src

sudo ~/go/bin/fyne package -os darwin -icon ../Icon.png && ~/go/bin/fyne install -icon ../Icon.png