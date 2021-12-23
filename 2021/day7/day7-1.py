from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
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
    positions = list(map(int, args.input.read().split(',')))

    logger.info('Read %s crabs', len(positions))
    max_position = max(positions)
    min_cost = sys.maxsize
    min_position = -1

    for i in range(max_position):
        cost = sum(abs(c - i) for c in positions)
        logger.debug('Computed cost at position %04d: %d', i, cost)

        if cost < min_cost:
            logger.info('Found new best position at %s with cost %s', i, cost)
            min_cost = cost
            min_position = i

    result = min_cost
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
