#!/bin/bash

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

trap  errorHandler ERR

iconFolder=$1/share/icons

libs=(libatk-1.0-0.dll libbz2-1.dll 
libcairo-2.dll libcairo-gobject-2.dll libepoxy-0.dll 
libexpat-1.dll libffi-8.dll libfontconfig-1.dll libfreetype-6.dll 
libgcc_s_seh-1.dll libgdk-3-0.dll libgdk_pixbuf-2.0-0.dll 
libgio-2.0-0.dll libglib-2.0-0.dll libgmodule-2.0-0.dll libgobject-2.0-0.dll 
libgtk-3-0.dll libharfbuzz-0.dll libintl-8.dll  libjpeg-62.dll libpango-1.0-0.dll 
libpangocairo-1.0-0.dll libpangoft2-1.0-0.dll libpangowin32-1.0-0.dll libpcre2-8-0.dll 
libpixman-1-0.dll libpng16-16.dll libstdc++-6.dll libwinpthread-1.dll 
zlib1.dll iconv.dll libtiff-5.dll libfribidi-0.dll libssp-0.dll)

create_icon() {
go install github.com/tc-hib/go-winres@latest
~/go/bin/go-winres init
cp ./assets/icons/*.png ./winres
cp ./scripts/winres-template.json ./winres/winres.json
sed -i 's/${VERSION}'/"$3"'/' ./winres/winres.json
~/go/bin/go-winres make
}

converter() {
    dest=$(realpath -s $2/16x16/$(basename $(dirname $1)))
    mkdir -p $dest
    gtk-encode-symbolic-svg -o $dest $1 16x16
}

set -o errtrace
trap "exit 1" ERR
create_icon
mv ./*.syso ./cmd/advertisementer/
CC=x86_64-w64-mingw32-gcc PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig GOOS=windows GOARCH=amd64 \
go build -C ./cmd/advertisementer/ -v -ldflags "-w -s -linkmode=external -H windowsgui -X main.version=$2" -o $1

echo "====> copy libraries <===="
for lib in ${libs[@]};
do
  strip --strip-debug /usr/x86_64-w64-mingw32/sys-root/mingw/bin/$lib -o $1/$lib
done

mkdir -p $1/share/{glib-2.0,icons}
cp -r /usr/share/glib-2.0/schemas/ $1/share/glib-2.0/
cp -r /usr/x86_64-w64-mingw32/sys-root/mingw/share/icons/* $iconFolder

echo "====> convert icons <===="
for icons in $(find $iconFolder -maxdepth 1 -mindepth 1 -type d);
do
  for file in $(find $icons -name *.svg);
  do 
     converter ${file} $icons
  done
done

cp /usr/x86_64-w64-mingw32/sys-root/mingw/bin/gdbus.exe $1
echo "====> building for windows is done <===="