from __future__ import print_function
import argparse
import logging
import sys

import hashlib

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

    data = args.input.read().strip()
    current = 0
    while True:
        md5 = hashlib.md5('{}{}'.format(data, current).encode()).hexdigest()
        if md5.startswith('000000') and md5[6] in '123456789':
            logger.info('Found hash: %s (%s)', md5, '{}{}'.format(data, current))
            break
        current += 1
    result = current


    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
