from __future__ import print_function
import argparse
import logging
import sys

import re

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

    lights = [[0] * 1000 for _ in range(1000)]

    for action in args.input:
        regex = re.match('([\w ]+) (\d+,\d+) through (\d+,\d+)', action.strip())
        change = regex.group(1)
        x_start, y_start = map(int, regex.group(2).split(','))
        x_end, y_end = map(int, regex.group(3).split(','))

        for x in range(x_start, x_end + 1):
            for y in range(y_start, y_end + 1):
                if change == 'toggle':
                    lights[x][y] += 2
                elif change == 'turn on':
                    lights[x][y] += 1
                elif change == 'turn off':
                    lights[x][y] -= 1 if lights[x][y] > 0 else 0

    result = sum(sum(l) for l in lights)
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
