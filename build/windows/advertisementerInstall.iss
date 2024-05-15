; Script generated by the Inno Setup Script Wizard.
; SEE THE DOCUMENTATION FOR DETAILS ON CREATING INNO SETUP SCRIPT FILES!

#define MyAppName "Advertisementer"
#define MyAppVersion "{appversion}"
#define MyAppExeName "advertisementer.exe"

[Setup]
; NOTE: The value of AppId uniquely identifies this application. Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{832F97D2-D7B6-4AC3-9B24-F66131E748A2}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
VersionInfoVersion={#MyAppVersion}
AppPublisher=Vladimir Rytov
AppPublisherURL=https://github.com/VladimirRytov/Advertisementer
;AppVerName={#MyAppName} {#MyAppVersion}
DefaultDirName={autopf64}\{#MyAppName}
DisableProgramGroupPage=yes
LicenseFile=LICENSE
; Uncomment the following line to run in non administrative install mode (install for current user only.)
;PrivilegesRequired=lowest
PrivilegesRequiredOverridesAllowed=dialog
OutputBaseFilename=advertisementerSetup
Compression=lzma
SolidCompression=yes
WizardStyle=modern

[Languages]
Name: "russian"; MessagesFile: "compiler:Languages\Russian.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Files]
Source: "windows\{#MyAppExeName}"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\share\*"; DestDir: "{app}/share"; Flags: ignoreversion recursesubdirs createallsubdirs
Source: "windows\gdbus.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\iconv.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libatk-1.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libbz2-1.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libcairo-2.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libcairo-gobject-2.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libepoxy-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libexpat-1.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libffi-8.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libfontconfig-1.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libfreetype-6.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libfribidi-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libgcc_s_seh-1.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libgdk_pixbuf-2.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libgdk-3-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libgio-2.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libglib-2.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libgmodule-2.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libgobject-2.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libgtk-3-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libharfbuzz-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libintl-8.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libjpeg-62.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libpango-1.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libpangocairo-1.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libpangoft2-1.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libpangowin32-1.0-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libpcre2-8-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libpixman-1-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libpng16-16.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libssp-0.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libstdc++-6.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libtiff-5.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\libwinpthread-1.dll"; DestDir: "{app}"; Flags: ignoreversion
Source: "windows\zlib1.dll"; DestDir: "{app}"; Flags: ignoreversion
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Icons]
Name: "{autoprograms}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"
Name: "{autodesktop}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon

[Run]
Filename: "{app}\{#MyAppExeName}"; Description: "{cm:LaunchProgram,{#StringChange(MyAppName, '&', '&&')}}"; Flags: nowait postinstall skipifsilent

