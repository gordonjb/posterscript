# Posterscript

This repo contains a couple of Python scripts related to my Plex media.

## dbscript

dbscript reads in a CSV export of a [Libib library](https://www.libib.com/library/home), and filters entries based on combinations of tags.

### Usage

```bash
python dbscript.py
```

The config file is [dbscript.cfg](./dbscript.cfg). This is read by [configparser](https://docs.python.org/3/library/configparser.html).

#### [files]

##### path

The complete path to the Libib library CSV, as exported from [Settings > Libraries > Export Library (.csv)](https://www.libib.com/library/settings-libraries)

#### [tags]

##### include_list

Comma separated list of tags. Items in the library must have *all* of these tags, or they will not be returned.

##### exclude_list

Comma separated list of tags. If an item in the library has *any* of these tags, it will not be returned.

## posterscript

posterscript checks Plex Movie and TV library folders, and outputs when it believes that `poster.ext`, `fanart.ext` and `seasonxx.ext` local image files are missing. The starting directory, and any directories it contains, will not be checked for assets (e.g. with root `path` `Images/Plex Posters` which contains `TV`, `Movie` and `Other` folders, checks will only be made within the `Images/Plex Posters/TV/*`, `Images/Plex Posters/Movie/*` and `Images/Plex Posters/Other/*` folders)

### Usage

```bash
python posterscript.py
```

The config file is [posterscript.cfg](./posterscript.cfg). This is read by [configparser](https://docs.python.org/3/library/configparser.html).

#### [files]

##### path

The complete path to the root directory containing the Library folders.

##### exclude_dir_list

Comma separated list of directory names to ignore.
