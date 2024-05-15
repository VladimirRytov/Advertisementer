#!/bin/env bash

apt update && apt install curl file -y
curl -L  https://github.com/AppImage/AppImageKit/releases/download/continuous/appimagetool-x86_64.AppImage -o appImage 
chmod +x appImage
./appImage --appimage-extract
mkdir -p Advertisementer.AppDir/usr/{bin,lib,share/applications}
cp linux/advertisementer Advertisementer.AppDir/usr/bin
chmod +x Advertisementer.AppDir/usr/bin/advertisementer

cp squashfs-root/usr/bin/AppRun Advertisementer.AppDir/

install -m 0755 build/linux/advertisementer.desktop Advertisementer.AppDir/
install -m 0755 build/linux/advertisementer.desktop Advertisementer.AppDir/usr/share/applications/

mkdir debs unpacked
apt -o Dir::Cache::archives="$(pwd)/debs" --download-only install libgtk-3-0 -y
find debs/ -name '*.deb' -exec dpkg -x {} unpacked \;

for lib in $(find unpacked/ | grep -e $(ldd /mnt/Clienbt/linux/advertisementer | grep 'not found' | cut -d' ' -f1 | xargs | sed 's/ / -e /g') | sed 's/unpacked\///');
do 
  cp --parents unpacked/$lib Advertisementer.AppDir/
done

for i in 16 32 48 64 128 256;
do 
  mkdir -p Advertisementer.AppDir/usr/share/icons/hicolor/${i}x${i}/apps
  cp assets/icons/*${i}.png Advertisementer.AppDir/usr/share/icons/hicolor/${i}x${i}/apps/advertisementer.png
done
cp Advertisementer.AppDir/usr/share/icons/hicolor/256x256/apps/advertisementer.png Advertisementer.AppDir
cp Advertisementer.AppDir/usr/share/icons/hicolor/256x256/apps/advertisementer.png Advertisementer.AppDir/.DirIcon
cp -r unpacked/usr/share/icons/Adwaita Advertisementer.AppDir/usr/share/icons/

squashfs-root/AppRun Advertisementer.AppDir
err=$?
if [ $err -ne 0 ];
then 
  echo "An error occured while creating AppImage package."
  exit $err
fi