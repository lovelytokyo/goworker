#!/usr/bin/env python

"""Parallel Executor

Usage:
  ./multiprocess <process> <count>

Options:
  <process>    process.
  <count>      thumbnail count.
"""

from docopt import docopt
from multiprocessing import Pool
import os

def function(hoge):
    cmd = './generator ./test.mp4 2 126 73 4 thumbnails_%s.jpg' % hoge
    os.system(cmd)
    return hoge

def multi(process, count):
    p = Pool(process)
    result = p.map(function, range(count))
    return result

def main():
    arguments = docopt(__doc__, version='0.0.2')
    data = multi(int(arguments['<process>']), int(arguments['<count>']))
    for i in data:
        print i

main()