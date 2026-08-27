# LDT (Lethal Download Tools)

A desktop helper for installing and managing Lethal Company mod setups, including full assembly downloads and Thunderstore mod installs with dependency resolution.

## Overview

LDT is built to simplify mod installation for older and legacy Lethal Company versions. Instead of manually tracking compatibility and dependency chains, the app helps you:

- install a ready-made mod assembly from Steam
- install individual mods from Thunderstore by URL
- choose a target LC version and automatically fetch compatible package versions
- create a new game instance or work with an existing installation

The project is built as a Wails desktop app and currently targets Windows usage for the main workflow.

## Features

- Full assembly install flow using Steam depot downloads and QR-based Steam login
- Individual mod installation from Thunderstore links
- Dependency-aware package installation for the selected LC version
- Support for a fresh LC install or an existing game folder
- Automatic choice of compatible mod versions based on LC version
- Automatic download and extraction of DepotDownloader when needed

## Installation modes

When the app starts, it asks which mode you want to use:

*  Complete assembly installation 
    > installs a complete mod assembly from a remote source and requires Steam authentication via QR code
*  Individual mod installation
   > installs mods from Thunderstore by link to a new or existing LC install


### Complete assembly installation

This mode installs a pre-configured mod assembly using Steam depot downloads via [DepotDownloader](https://github.com/SteamRE/DepotDownloader).

1. Choose the Steam folder to use.
  - By default, the app points to `C:\Program Files (x86)\Steam`.
  - If Steam is elsewhere, choose the correct folder manually.
2. Select the assembly from the available list.
3. Log in to Steam when prompted with the QR code in the Steam mobile app.
4. Wait for the download to finish and the install folder will open automatically.

> Tip: make sure Steam is running when launching the game after installation.

### Individual mod installation

This mode installs one or more mods from Thunderstore and resolves dependencies for your chosen LC version.

1. Choose the installation target:
  - create a new LC instance
  - install into an existing installation
2. Pick the LC version you are targeting.
3. Enter a Thunderstore package URL in this format:

```text
https://thunderstore.io/c/lethal-company/p/[author]/[packageName]/
```

Example:

```text
https://thunderstore.io/c/lethal-company/p/BepInEx/BepInExPack/
```

4. The app resolves the latest compatible version, downloads dependencies, and installs everything into the correct folder.
5. After each install, you can continue with another mod or stop.

## Requirements

- Windows (primary supported platform)
- Steam installed and available on the system
- Internet connection
- [.NET 8.0+](https://dotnet.microsoft.com/download/dotnet/8.0)
- A Steam account with Lethal Company ownership for full assembly installation

## Build and run

This project uses Go and Wails.

```bash
go mod download
wails build
```

For local development with the frontend:

```bash
wails dev
```

## Project purpose

The original motivation was to make older Lethal Company mod setups easier to install and debug, especially when dependencies and version compatibility become complicated. LDT is meant to reduce that complexity and make mod installation safer and more repeatable.
 
