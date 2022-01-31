from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def generator(length):
    value = 0
    fmt_str = '0{}'.format(length)

    while value < 2 ** length:
        yield [True if c == '1' else False for c in f'{value:{fmt_str}b}']
        value += 1

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('target', type=int,
            help='Target value')
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
    containers = list(map(int, args.input.readlines()))

    logger.debug('Read %d containers: %s', len(containers), containers)

    for c in generator(len(containers)):
        if sum(containers[i] if p else 0 for i, p in enumerate(c)) == args.target:
            logger.debug('Found combination: %s', [containers[i] if p else 0 for i, p in enumerate(c)])
            result += 1

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
