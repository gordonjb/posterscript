# Posterscript

Posterscript is a CLI written in Go to perform a few tasks related to my Plex media.

## Usage

```bash
Plex related scripts

Usage:
  posterscript [command]

Available Commands:
  check       Validate Plex local posters
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  libib       Filter libib library by tags
  logo        Overlay a logo file over a poster file, for example to add network logos
  tpdbarchive Unpack TPDb archive to Plex format

Flags:
  -h, --help      help for posterscript
  -v, --version   version for posterscript

Use "posterscript [command] --help" for more information about a command.
```

## libib

### Usage

```bash
Reads in a CSV export of a Libib library, and filters entries based on combinations of tags.

Usage:
  posterscript libib FILE [flags]

Flags:
  -e, --exclude TAG   A libib TAG. Can be specified multiple times. If an item in the library has *any* of these TAGs, it will not be returned
  -h, --help          help for libib
  -i, --include TAG   A libib TAG. Can be specified multiple times. Items in the library must have *all* of these TAGs, or they will not be returned
```

#### FILE argument

The path to a Libib library CSV, as exported from [Settings > Libraries > Export Library (.csv)](https://www.libib.com/library/settings-libraries)

## check

### Usage

```bash
Checks Plex Movie and TV library folders, and output when it believes that 'poster.ext', 'fanart.ext' and 'seasonxx.ext' local image files are missing.

Each subdirectory of the starting directory PATH will be treated as a Library folder (e.g. with root PATH 'Images/Plex Posters' which contains
'TV', 'Movie' and 'Other' folders, checks will only be made within the 'Images/Plex Posters/TV/', 'Images/Plex Posters/Movie/' and 'Images/Plex Posters/Other/' folders)

Usage:
  posterscript check PATH [flags]

Flags:
  -e, --exclude DIRECTORY   Ignore this DIRECTORY when scanning the root path. Can be specified multiple times
  -h, --help                help for check
  -a, --show-all            Show every validated item, not just failures
```

The Plex local media assets expected layout is described in their documentation:

- [Local Media Assets – Movies](https://support.plex.tv/articles/200220677-local-media-assets-movies/)
- [Local Media Assets – TV Shows](https://support.plex.tv/articles/200220717-local-media-assets-tv-shows/)

#### PATH argument

The path containing Plex libraries from which to start the check.

## tpdbarchive

### Usage

```bash
Unzip a poster set from The Poster Database into the correct folder structure for Plex to pick them up.

Supports Movie and TV sets.

Usage:
  posterscript tpdbarchive FILE [flags]

Aliases:
  tpdbarchive, archive, tpdb, a

Flags:
  -h, --help   help for tpdbarchive
```

The Plex local media assets expected layout is described in their documentation:

- [Local Media Assets – Movies](https://support.plex.tv/articles/200220677-local-media-assets-movies/)
- [Local Media Assets – TV Shows](https://support.plex.tv/articles/200220717-local-media-assets-tv-shows/)

#### FILE argument

The path to a poster set ZIP, as downloaded from [The Poster Database](https://theposterdb.com/) using the "Download Set Posters" option.

## logo

### Usage

```bash
Overlay a logo file over a poster file, for example to add network logos

Usage:
  posterscript logo FILE [flags]

Flags:
  -h, --help             help for logo
  -l, --logoDir string   Directory containing overlay images (default "logos")
```

Currently, as this is largely for my own use, I have a folder of logo overlays, most of which were extracted from the PSD posted [on r/PlexPosters by u/tscardino](https://www.reddit.com/r/PlexPosters/comments/mdbbp9/cornernetwork_tv_template/), with a few others I added to the template for myself over time. The poster file will be resized to 1000x1500 to match these overlays, which are all transparent 1000x1500 PNGs. I have uploaded my folder of logos [here](https://drive.google.com/drive/folders/1Z6bZDPczbnMgF6YNutdSgw4ZdABp5ypB?usp=share_link)

When run, the script will read all files in logoDir, and allow the user to select one to overlay.

#### FILE argument

The path to a poster image.
