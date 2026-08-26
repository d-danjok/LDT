LDT (Lethal Download Tools)
==
 
Convenient tool for downloading and installing mod assemblies for legacy [Lethal Company](https://store.steampowered.com/app/1966720/Lethal_Company/) versions
 
## Contents
 
* [Motivation](#motivation)
* [Concept](#concept)
* [Usage](#usage)
  * [Complete Assembly Installation](#complete-assembly-installation)
  * [Individual Mod Installation](#individual-mod-installation)
* [Requirements](#requirements)
## Motivation
 
Idea for this project came after 2 days of trying to make Wesley's Moons mod work, the complexity of this was in fact
that currently Wesley's Moons does not support v80 Lethal Company. And because of how much dependencies this mod has
it turned into nightmare when debugging. So in order to simplify life for my friends I came up with an idea to write
a simple program which would allow to easily install mod assemblies.
 
## Concept
 
Brief idea was already mentioned in [motivation](#motivation), but here it will be explained in detail.
 
Main idea is to provide a convenient tool for downloading and sharing mod assemblies, with support for installing mods
compatible with older versions of the game independently. LCAD automatically resolves and installs mod dependencies,
downloading the correct version of each mod for your target LC version.
 
 
## Usage
 
On launch you will be asked to select an installation mode:
 
```
Select installation type by entering corresponding number
 
0:  Complete assembly installation
      (installs complete mod assembly with mods from remote source, requires steam authorisation using QR code)
1:  Individual mod installation
      (installing mods from Thunderstore by link, can be installed as new game instance or into existing installation folder)
 
:
```
 
---
 
### Complete Assembly Installation
 
This mode downloads a full pre-configured mod assembly directly from Steam using [DepotDownloader](https://github.com/SteamRE/DepotDownloader).
 
1. **Locate Steam folder** — on start the program will ask you to confirm the default Steam folder location (`C:\Program Files (x86)\Steam`)
```
   Do you want to use default Steam folder location (C:\Program Files (x86)\Steam) [confirm/deny]
   :
```
 
   If Steam is installed elsewhere, deny and locate the folder manually in the pop-up window.
 
   ![Steam folder selection pop-up](/docs/resorces/img/selectSteamFolderPopUp.png)
 
2. **Select assembly** — a list of available mod assemblies will be displayed
```
   Select assembly you want to install by entering corresponding number
 
   0:    v73 Non modded
   1:    v73 Wesley's Basic
   ...
   n:    vXX ...
 
   :
```
 
3. **Steam login via QR code** — because game versions are downloaded directly from Steam, you will be prompted to log in by scanning a QR code with the Steam mobile app
4. **Done** — when the download is complete, the installation folder will open automatically. You can then launch the game.
   > **Tip:** Make sure Steam is running before launching the game, otherwise it won't work correctly.
---
 
### Individual Mod Installation
 
This mode installs mods from Thunderstore by link, automatically resolving and installing all dependencies at the correct version for your target LC version.
 
1. **Choose installation target** — select whether to create a fresh LC instance or install into an existing one
```
   Select installation type by entering corresponding number
 
   0:  Install new LC instance
   1:  Install into existing LC installation
 
   :
```
 
   **New LC instance:**
   - Select the LC version to download (v50, v56, v73)
   - Enter a name for the installation
   - Steam QR login will be required to download the game files
   
   **Existing LC installation:**
   - Locate your existing LC installation folder in the pop-up
   - If BepInEx mods are already present, choose to install alongside or erase existing mods
   - Select which LC version you have installed (used to pick compatible mod versions)
2. **Install mods** — enter a Thunderstore link in the following format:
```
   https://thunderstore.io/c/lethal-company/p/[author]/[packageName]/
```
 
   Example:
```
   https://thunderstore.io/c/lethal-company/p/BepInEx/BepInExPack/
```
 
   LCAD will automatically:
   - Find the latest mod version compatible with your LC version
   - Download the mod and all its dependencies
   - Install everything into the correct location
3. **Install more mods** — after each successful install, you will be asked if you want to install another mod. Type `abort` or `a` at any time to stop.
4. **Done** — the installation folder will open when you're finished.
---
 
## Requirements
 
- **Windows** (only supported platform)
- **Steam** installed (default or custom path)
- **Internet connection**
- [DepotDownloader](https://github.com/SteamRE/DepotDownloader) — downloaded automatically on first run, requires [.NET 8.0+](https://dotnet.microsoft.com/download/dotnet/8.0)
- A **Steam account** with Lethal Company owned (required for complete assembly installation)
 
