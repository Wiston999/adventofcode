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
    arg_parser.add_argument('-d', '--days', type=int, default=80,
            help='Number of days for the simulation')
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

    fishes = list(map(int, args.input.read().split(',')))

    for n in range(args.days):
        logger.info('There are %04d fishes at day %s', len(fishes), n)
        logger.debug(fishes)
        for i in range(len(fishes)):
            fishes[i] = fishes[i] - 1
            if fishes[i] == -1:
                logger.debug('Adding fish for position %s', i)
                fishes[i] = 6
                fishes.append(8)

    print ("Result is", len(fishes), file=args.output)

if __name__ == '__main__':
    main()
