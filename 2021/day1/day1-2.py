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

    lines = args.input.readlines()
    num_lines = len(lines)
    prev_sum = None
    result = 0
    for i, l in enumerate(lines):
        if i <= (num_lines - 3):
            cur_sum = sum(map(int, lines[i:i+3]))
            logger.debug('Sum %s from %s', cur_sum, lines[i:i+3])
            if prev_sum is not None and cur_sum > prev_sum:
                logger.debug('Sum increased: %s > %s', cur_sum, prev_sum)
                result += 1
            prev_sum = cur_sum

    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
