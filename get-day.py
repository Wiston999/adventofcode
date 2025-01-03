from __future__ import print_function
import argparse
import logging
import sys
import os, os.path

from datetime import datetime
from urllib import request
import re
import subprocess

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

GO_RUN_SH='''#!/bin/bash
go run day{:02d}.go utils.go $@
'''

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

def go_workspace_tune(pwd):
    actions = [
        'go mod init github.com/Wiston999/adventofcode/{}'.format(pwd),
        'go get github.com/sirupsen/logrus',
        'go get github.com/urfave/cli/v2',
        'go get github.com/oleiade/lane/v2',
        'go mod tidy',
        'chmod a+x run.sh',
    ]

    rcs = []
    for action in actions:
        rc = subprocess.run(action, shell=True, cwd='./{}'.format(pwd)).returncode
        if rc != 0:
            logger.warning('Action "%s" failed executing', action)
        rcs.append(rc == 0)
    return all(rcs)

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-y', '--year', type=int, default=datetime.today().year,
            help='Year of advent of code')
    arg_parser.add_argument('-m', '--language', choices=['python', 'go'], default='go',
            help='Language skel to be used')
    arg_parser.add_argument('day', type=int, nargs='?', default=datetime.today().day,
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

    langs = {
        'python': {
            'skel': 'py-skel.py',
            'extension': 'py',
            'workspace_tune': lambda x: True
        },
        'go': {
            'skel': 'go-skel.go',
            'extension': 'go',
            'workspace_tune': go_workspace_tune
        },
    }

    print ("Creating folder for day:", args.day, file=args.output)
    try:
        os.mkdir('{}/day{:02d}'.format(args.year, args.day))
    except FileExistsError:
        logger.warning('Folder %s/day%02d already exists, ignoring', args.year, args.day)

    print ("Fetching README for day", args.day, file=args.output)
    readme = fetch_readme(args.day, args.year)

    open(os.path.join(str(args.year), 'day{:02d}'.format(args.day), 'README.md'), 'w').write(html2mk(readme))


    print ("Creating input files", file=args.output)
    for f in ['input.txt', 'input.test.1.txt']:
        open(os.path.join(str(args.year), 'day{:02d}'.format(args.day), f), mode='a').close()

    print ("Copying {} skeletons".format(args.language), file=args.output)

    if args.language == 'python':
        for f in '12':
            f = os.path.join(
                str(args.year),
                'day{:02d}'.format(args.day),
                'day{:02d}-{}.{}'.format(args.day, f, langs[args.language]['extension'])
            )
            if not os.path.exists(f):
                cp(langs[args.language]['skel'], f)
            else:
                logger.warning('%s already exists', f)
    elif args.language == 'go':
        f = os.path.join(
            str(args.year),
            'day{:02d}'.format(args.day),
            'day{:02d}.{}'.format(args.day, langs[args.language]['extension'])
        )
        if not os.path.exists(f):
            cp(langs[args.language]['skel'], f)
        else:
            logger.warning('%s already exists', f)

        f = os.path.join(
            str(args.year),
            'day{:02d}'.format(args.day),
            'utils.go',
        )

        if not os.path.exists(f):
            cp('utils.go', f)
        else:
            logger.warning('%s already exists', f)

        f = os.path.join(
            str(args.year),
            'day{:02d}'.format(args.day),
            'run.sh',
        )

        if not os.path.exists(f):
            open(f, 'w').write(GO_RUN_SH.format(args.day))
        else:
            logger.warning('%s already exists', f)

    setup_success = langs[args.language]['workspace_tune'](os.path.join(
        str(args.year),
        'day{:02d}'.format(args.day)
    ))

    if not setup_success:
        logger.warning('Failed setting up workspace')

    print ("Finished setting up day, visit https://adventofcode.com/{}/day/{}/input to get your input".format(
        args.year,
        args.day,
        ), file=args.output)


if __name__ == '__main__':
    main()
