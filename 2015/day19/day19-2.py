from __future__ import print_function
import argparse
import logging
import sys

import re
import random

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def search(target, replacements):
    t = target
    shuffles = 0
    i = 0
    while t != 'e':
        prev = t
        for dst, src in replacements:
            if src in t:
                i += t.count(src)
                t = t.replace(src, dst)
                logger.debug('Got new target (%03d) %s using %s --> %s', i, t, src, dst)
        if prev == t:
            random.shuffle(replacements)
            t = target
            shuffles += 1
            i = 0
            logger.info('(%05d) Dead end with %s', shuffles, t)

    return i


def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-i', '--input', type=argparse.FileType('r'), default=sys.stdin,
            help='Intput file, use - for stdin')
    arg_parser.add_argument('-o', '--output', type=argparse.FileType('w'), default=sys.stdout,
            help='Output file, use - for stdout')
    arg_parser.add_argument('-l', '--loglevel', type=str.upper, default='info',
            choices=['DEBUG', 'INFO', 'WARNING', 'ERROR'], help='Output file (when new set is created)')
    args = arg_parser.parse_args(argv)

    logger.setLevel(args.loglevel)

    logger.debug('Log level: %s', args.loglevel)
    logger.debug('Input file: %s', args.input.name)
    logger.debug('Output file: %s', args.output.name)
    result = 0

    replacements = []
    molecule = ''
    for l in args.input:
        l = l.strip()
        if '=>' in l:
            replacements.append(list(map(lambda x:x.strip(), l.split('=>'))))
        elif len(l) > 0:
            molecule = l


    logger.info('Read molecule: %s', molecule)
    logger.info('Read %02d replacements', len(replacements))
    logger.debug(replacements)

    result = search(molecule, replacements)

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
