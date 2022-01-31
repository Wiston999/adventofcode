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

    visits = [collections.defaultdict(int), collections.defaultdict(int)]
    currents = [(0, 0), (0, 0)]
    visits[0][currents[0]] += 1
    visits[1][currents[1]] += 1
    for i, d in enumerate(args.input.read()):
        current = currents[i % 2]
        if d == 'v':
            current = (current[0], current[1] - 1)
        if d == '^':
            current = (current[0], current[1] + 1)
        if d == '<':
            current = (current[0] - 1, current[1])
        if d == '>':
            current = (current[0] + 1, current[1])
        currents[i % 2] = current
        visits[i % 2][current] += 1

    result = len(set(visits[0].keys()) | set(visits[1].keys()))
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
