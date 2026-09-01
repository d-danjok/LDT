; LDT Windows Installer Script (NSIS)
; This script creates an installer for LDT that includes the binary and appdata folder

!include "MUI2.nsh"

; Basic Installer Information
Name "LDT"
OutFile "LDT-installer.exe"
InstallDir "$PROGRAMFILES\LDT"
InstallDirRegKey HKCU "Software\LDT" ""

; Request application privileges for Windows Vista and higher
RequestExecutionLevel admin

; MUI Settings
!insertmacro MUI_PAGE_WELCOME
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH

!insertmacro MUI_LANGUAGE "English"

; Installer sections
Section "LDT Application"
  SetOutPath "$INSTDIR"
  
  ; Copy the main executable
  File "LDT.exe"
  
  ; Create and copy appdata folder
  SetOutPath "$INSTDIR\appdata\definitions"
  File "appdata\definitions\LCVersions.json"
  File "appdata\definitions\general.def"
  File "appdata\definitions\Assemblies.json"
  
  ; Store installation folder in registry
  WriteRegStr HKCU "Software\LDT" "" $INSTDIR
  
  ; Create Start Menu shortcut
  SetShellVarContext all
  CreateDirectory "$SMPROGRAMS\LDT"
  CreateShortcut "$SMPROGRAMS\LDT\LDT.lnk" "$INSTDIR\LDT.exe"
  CreateShortcut "$SMPROGRAMS\LDT\Uninstall LDT.lnk" "$INSTDIR\uninstall.exe"
  
  ; Create Desktop shortcut (optional)
  CreateShortcut "$DESKTOP\LDT.lnk" "$INSTDIR\LDT.exe"
  
  ; Create uninstaller
  WriteUninstaller "$INSTDIR\uninstall.exe"
  WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\LDT" "DisplayName" "LDT"
  WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\LDT" "UninstallString" "$INSTDIR\uninstall.exe"
SectionEnd

; Uninstaller section
Section "Uninstall"
  ; Remove files
  Delete "$INSTDIR\LDT.exe"
  Delete "$INSTDIR\appdata\definitions\LCVersions.json"
  Delete "$INSTDIR\appdata\definitions\general.def"
  Delete "$INSTDIR\appdata\definitions\Assemblies.json"
  Delete "$INSTDIR\uninstall.exe"
  
  ; Remove directories
  RMDir "$INSTDIR\appdata\definitions"
  RMDir "$INSTDIR\appdata"
  RMDir "$INSTDIR"
  
  ; Remove Start Menu shortcuts
  SetShellVarContext all
  Delete "$SMPROGRAMS\LDT\LDT.lnk"
  Delete "$SMPROGRAMS\LDT\Uninstall LDT.lnk"
  RMDir "$SMPROGRAMS\LDT"
  
  ; Remove Desktop shortcut
  Delete "$DESKTOP\LDT.lnk"
  
  ; Remove registry entries
  DeleteRegKey HKCU "Software\LDT"
  DeleteRegKey HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\LDT"
SectionEnd
