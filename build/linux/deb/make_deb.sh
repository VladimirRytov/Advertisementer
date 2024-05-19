#!/bin/env bash

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

trap errorHandler ERR

apt update && apt install build-essential md5deep -y
mkdir -p advertisementer-$1-$2/DEBIAN advertisementer-$1-$2/usr/{bin,share/{applications,doc/advertisementer}}

install -m 0755 build/linux/advertisementer.desktop advertisementer-$1-$2/usr/share/applications/
install -m 0755 linux/advertisementer advertisementer-$1-$2/usr/bin/

for i in 16 32 48 64 128 256;
do 
  mkdir -p advertisementer-$1-$2/usr/share/icons/hicolor/${i}x${i}/apps
  cp assets/icons/*${i}.png advertisementer-$1-$2/usr/share/icons/hicolor/${i}x${i}/apps/advertisementer.png
done

deb_dir=$(pwd)/advertisementer-$1-$2/
md5deep -r advertisementer-$1-$2/usr > advertisementer-$1-$2/DEBIAN/md5sums
sed "s|$deb_dir||g" -i advertisementer-$1-$2/DEBIAN/md5sums
cp build/linux/deb/control advertisementer-$1-$2/DEBIAN
cp build/linux/deb/copyright advertisementer-$1-$2/usr/share/doc/advertisementer
total_size=$(du -sk advertisementer-$1-$2/usr | awk '{ print $1 }')
sed -e 's/${version}/'"$1"'/' -e 's/${release}/'"$2"'/' -e 's/${buildNumber}/'"$3"'/' -e 's/${size}/'"$total_size"'/' -i advertisementer-$1-$2/DEBIAN/control
fakeroot dpkg-deb --build advertisementer-$1-$2