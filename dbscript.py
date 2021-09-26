from os import walk
from pathlib import Path
import csv
import configparser


parser = configparser.ConfigParser()
parser.read("dbscript.cfg")

include_list = parser.get("tags", "include_list").split(',')
exclude_list = parser.get("tags", "exclude_list").split(',')
path = Path(parser.get("files", "path"))

with open(path, encoding='utf8') as csv_file:
    csv_reader = csv.DictReader(csv_file, delimiter=',')
    line_count = 0
    for row in csv_reader:
        if line_count == 0:
            print(f'Items matching tag sets: include {include_list}, exclude {exclude_list}')
            line_count += 1
        row_tags = row["tags"].split(",")
        matches_include_list = all(elem in row_tags for elem in include_list)
        matches_exclude_list = any(elem in row_tags for elem in exclude_list)
        if matches_include_list and not matches_exclude_list:
            print(f'{row["title"]}, tags: {row_tags}')
        line_count += 1
    print(f'Processed {line_count} lines.')