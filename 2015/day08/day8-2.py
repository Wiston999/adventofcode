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

    code_len = 0
    literal_len = 0
    for l in args.input:
        l = l.strip()
        literal_len += len(l)
        code_l = l.replace('\\', '\\\\').replace('"', '\\"')
        tmp_count = len(code_l) + 2
        code_len += tmp_count
        logger.debug('Tmp count: (%s) --> (%s) %03d %03d', l, code_l, len(l), tmp_count)
    logger.info('Literal count is %03d', literal_len)
    logger.info('Code count is %03d', code_len)
    result = code_len - literal_len
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()