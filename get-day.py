from __future__ import print_function
import argparse
import logging
import sys
import os, os.path

from urllib import request
import re

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

def fetch_readme(day, year):
    url = 'https://adventofcode.com/{}/day/{}'.format(year, day)
    logger.debug('Fetching %s', url)
    try:
        response = request.urlopen(url)
        logger.debug('Got response %s', response)
        return response.read().decode('utf-8')
    except:
        logger.error('Error fetching url %s', url)

    return

def cp(source, dest):
    with open(source, 'rb') as src, open(dest, 'wb') as dst: dst.write(src.read())

def html2mk(text):
    text = re.sub('.*?<article .+?>(.*?)</article>.*', r'\1', text, flags=re.M | re.S)
    logger.debug('First cleaned text: %s', text)

    text = re.sub('<h2>(.*?)</h2>', r'# \1\n\n', text, flags=re.M | re.S)
    logger.debug('Replaced title: %s', text)

    text = re.sub('<a href="(.*?)".*?>(.*?)</a>', r'[\1](\2)', text, flags=re.M | re.S)
    logger.debug('Replaced anchors: %s', text)

    text = re.sub('<p>(.*?)</p>', r'\1\n\n', text, flags=re.M | re.S)
    logger.debug('Replaced paragraphs: %s', text)

    text = re.sub('<em>(.*?)</em>', r'<em><b>\1</b></em>', text, flags=re.M | re.S)
    logger.debug('Re-emphasized text: %s', text)

    return text

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-y', '--year', type=int, default=2021,
            help='Year of advent of code')
    arg_parser.add_argument('day', type=int,
            help='Day to be fetched and prepared')
    arg_parser.add_argument('-o', '--output', type=argparse.FileType('w'), default=sys.stdout,
            help='Output file, use - for stdout')
    arg_parser.add_argument('-l', '--loglevel', type=str.upper, default='info',
            choices=['DEBUG', 'INFO', 'WARNING', 'ERROR'], help='Output file (when new set is created)')
    args = arg_parser.parse_args(argv)

    logger.setLevel(args.loglevel)

    logger.debug('Log level: %s', args.loglevel)
    logger.debug('Day: %s', args.day)
    logger.debug('Year: %s', args.year)
    logger.debug('Output file: %s', args.output.name)

    print ("Creating folder for day:", args.day, file=args.output)
    try:
        os.mkdir('{}/day{}'.format(args.year, args.day))
    except FileExistsError:
        logger.warning('Folder %s/day%s already exists, ignoring', args.year, args.day)

    print ("Fetching README for day", args.day, file=args.output)
    readme = fetch_readme(args.day, args.year)

    open(os.path.join(str(args.year), 'day{}'.format(args.day), 'README.md'), 'w').write(html2mk(readme))

    print ("Copying python skeletons", file=args.output)

    for f in '12':
        f = os.path.join(str(args.year), 'day{}'.format(args.day), 'day{}-{}.py'.format(args.day, f))
        if not os.path.exists(f):
            cp('py-skel.py', f)
        else:
            logger.warning('%s already exists', f)

    print ("Finished setting up day, visit https://adventofcode.com/{}/day/{}/input to get your input".format(
        args.year,
        args.day,
        ), file=args.output)


if __name__ == '__main__':
    main()
