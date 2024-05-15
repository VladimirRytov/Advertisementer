#!/bin/env bash

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

trap errorHandler ERR

dnf install -y rpmdevtools rpmlint
HOME=$(pwd)
rpmdev-setuptree

for i in 16 32 48 64 128 256;
do 
  cp assets/icons/*${i}.png ~/rpmbuild/BUILD/advertisementer${i}.png
done

cp build/linux/advertisementer.desktop ~/rpmbuild/BUILD
cp linux/advertisementer ~/rpmbuild/BUILD
cp build/linux/rpm/advertisementer.spec ~/rpmbuild/SPECS/
sed -e 's/${version}/'"$1"'/' -e 's/${release}/'"$2"'/' -e 's/${buildNumber}/'"$3"'/' -i ~/rpmbuild/SPECS/advertisementer.spec
rpmlint ~/rpmbuild/SPECS/advertisementer.spec
rpmbuild -bb ~/rpmbuild/SPECS/advertisementer.spec