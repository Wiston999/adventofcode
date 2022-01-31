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
    result = 20151125

    match = re.search('Enter the code at row (\d+), column (\d+)', args.input.read())
    row, column = match.groups()
    row, column = int(row), int(column)
    logger.info('Searching for code at row %d and column %d', row, column)

    n = 1
    for i in range(1, column):
        logger.debug('Value at top of column %d: %d', i, n)
        n += i + 1


    for j in range(1, row):
        logger.debug('Value at beggining of row %d: %d', j, n)
        n += j + i

    logger.info('Needed element at position %05d', n)

    for c in range(1, n):
        result = (result * 252533) % 33554393

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
