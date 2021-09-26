from os import walk
from pathlib import Path
import configparser


parser = configparser.ConfigParser()
parser.read("posterscript.cfg")

exclude_dir_list = parser.get("files", "exclude_dir_list").split(',')
path = Path(parser.get("files", "path"))

root_dirs = []

for (dirpath, dirnames, filenames) in walk(path):
    dirnames[:] = [d for d in dirnames if d not in exclude_dir_list]
    if root_dirs == []:
        root_dirs = dirnames

    folder = path / dirpath
    
    if folder == path or folder.stem in root_dirs:
        continue

    if folder.name.startswith("Season ") or folder.name.startswith("Specials"):
        if folder.name.startswith("Season "):
            season_poster_name = folder.name.replace("Season ", "season").lower()
        elif folder.name.startswith("Specials"):
            season_poster_name = "season-specials-poster"
        contains_poster = False
        for name in filenames:
            file = folder / name
            if file.stem == season_poster_name:
                contains_poster = True
        if contains_poster:
            pass
        else:
            print(dirpath)
            print("Missing Season Poster")
    else:
        contains_art = False
        contains_poster = False
        for name in filenames:
            file = folder / name
            if file.stem == "poster":
                contains_poster = True
            elif file.stem == "fanart":
                contains_art = True
        if contains_art and contains_poster:
            pass
        else:
            print(dirpath)
            if not contains_art:
                print("Missing Fanart")
            if not contains_poster:
                print("Missing Poster")