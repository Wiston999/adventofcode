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
    mapa = [[int(e) for e in l.strip()] for l in args.input.readlines()]
    logger.debug('Map is: %s', mapa)

    limit_x = len(mapa)
    limit_y = len(mapa[0])

    for x in range(limit_x):
        for y in range(limit_y):
            element = mapa[x][y]
            surrounding = []
            if (x - 1) >= 0:
                surrounding.append(mapa[x-1][y])
            if (x + 1) < limit_x:
                surrounding.append(mapa[x+1][y])
            if (y - 1) >= 0:
                surrounding.append(mapa[x][y-1])
            if (y + 1) < limit_y:
                surrounding.append(mapa[x][y+1])
            logger.debug('Comparing %s (%s, %s) with %s', element, x, y, surrounding)
            if element < min(surrounding):
                result += element + 1

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
