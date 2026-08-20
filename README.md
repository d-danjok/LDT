LCAD (Lethal Company Assembly Downloader)
==

Convenient tool for downloading and installing mod assemblies for legacy Lethal Company versions

## Contents

* [Motivation](#motivation)
* [Concept](#concept)
* [Usage](#usage)
  * [CLI](#cli)
  * [GUI](#gui)

## Motivation

Idea for this project came after 2 days of trying to make Wesley's Moons mod work, the complexity of this was in fact 
that currently Wesley's Moons does not support v80 Lethal Company. And because of how much dependencies this mod has 
it turned into nightmare when debugging. So in order to simplify life for my friends I came up with an idea to write 
a simple program which would allow to easily install mod assemblies.

## Concept

Brief idea was already mentioned in [motivation](#motivation), but here it will be explained detailed.

Main idea is to provide convenient tool for downloading and sharing mod assemblies, potentially installing older 
versions of mods independently.

## Usage

NOTE:
Currently, program is only implemented with CLI, plans are to add GUI launch mode which will allow much more functionality
to be implemented

### CLI

1. On start program will ask you to confirm if you want to use default Steam folder location (C:\Program Files (x86)\Steam)
   ```terminaloutput
    Do you want to use default Steam folder location (C:\Program Files (x86)\Steam) [y/n] 
    :
    ```
   if you have Steam installed else where you just have to deny, and then you will have an ability to locate Steam 
   folder manually in popped up window.
   ![](/docs/resorces/img/selectSteamFolderPopUp.png)

2. After you locate Steam folder list of mod assemblies will be displayed
   ```terminaloutput
    Select assembly you want to install by entering corresponding number

    0:      v73 Non modded
    1:      v73 Wesley's Basic
    ...
    n:      vXX ...
    
    :
    ```
    Then you can choose an exact assembly by typing its number, after that download will begin

3. After this program will ask for logging in to Steam using QR code, this is necessary because versions are downloaded
    directly from steam, using [DepotDownloader](https://github.com/SteamRE/DepotDownloader)
   
4. When download is complete folder with LC assembly will pop up, and you can start playing  
    TIP: do not forget to launch steam before starting game, without this game won't work correctly

### GUI

To be introduced
