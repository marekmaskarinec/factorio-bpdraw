#!/usr/bin/python
from pprint import pprint
import json
import sys

json = json.loads(open("entities.json", "r").read())

if len(sys.argv) == 2:
    pprint(json[sys.argv[1]])
else:
    print("usage: ./print.py entity-name\navailable options:\n")
    print(" ".join(json.keys()))
