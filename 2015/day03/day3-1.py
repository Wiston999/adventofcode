from __future__ import print_function
import argparse
import logging
import sys

import collections

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

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

    visits = collections.defaultdict(int)
    current = (0, 0)
    visits[current] += 1
    for d in args.input.read():
        if d == 'v':
            current = (current[0], current[1] - 1)
        if d == '^':
            current = (current[0], current[1] + 1)
        if d == '<':
            current = (current[0] - 1, current[1])
        if d == '>':
            current = (current[0] + 1, current[1])
        visits[current] += 1

    result = len(visits)
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()