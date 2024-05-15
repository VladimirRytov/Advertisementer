Name:           advertisementer
Version:        ${version}
Release:        ${release}%{?dist}
Summary:        An advertisements managment package.
ExclusiveArch:  x86_64
License:        MIT

Requires:       gtk3

%description
An advertisements managment package.
Build number ${buildNumber}.

%install
mkdir -p %{buildroot}/%{_bindir}
install -m 0755 %{name} %{buildroot}/%{_bindir}

mkdir -p %{buildroot}/%{_datadir}/applications/
install -m 0755 advertisementer.desktop %{buildroot}/%{_datadir}/applications/

for i in 16 32 48 64 128 256;
do 
  mkdir -p %{buildroot}/%{_datadir}/icons/hicolor/${i}x${i}/apps/
  cp *${i}.png %{buildroot}/%{_datadir}/icons/hicolor/${i}x${i}/apps/advertisementer.png
done

%files
%{_bindir}/advertisementer
%{_datadir}/applications/advertisementer.desktop
%{_datadir}/icons/hicolor/16x16/apps/advertisementer.png
%{_datadir}/icons/hicolor/32x32/apps/advertisementer.png
%{_datadir}/icons/hicolor/48x48/apps/advertisementer.png
%{_datadir}/icons/hicolor/64x64/apps/advertisementer.png
%{_datadir}/icons/hicolor/128x128/apps/advertisementer.png
%{_datadir}/icons/hicolor/256x256/apps/advertisementer.png

%changelog
* Sun May 05 2024 Super User
- 
