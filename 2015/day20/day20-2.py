from __future__ import print_function
import argparse
import logging
import sys

from functools import lru_cache, reduce
from collections import defaultdict
import math

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def factors(n, m=1):
    return set(reduce(list.__add__,
                ([i, n//i] for i in range(1, int(n**0.5) + 1) if n % i == 0)))

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('input', type=int,
            help='Input value')
    arg_parser.add_argument('-o', '--output', type=argparse.FileType('w'), default=sys.stdout,
            help='Output file, use - for stdout')
    arg_parser.add_argument('-l', '--loglevel', type=str.upper, default='info',
            choices=['DEBUG', 'INFO', 'WARNING', 'ERROR'], help='Output file (when new set is created)')
    args = arg_parser.parse_args(argv)

    logger.setLevel(args.loglevel)

    logger.debug('Log level: %s', args.loglevel)
    logger.debug('Input: %s', args.input)
    logger.debug('Output file: %s', args.output.name)
    result = 1

    visited_factors = defaultdict(int)
    while True:
        computed = 0
        for f in factors(result):
            if visited_factors[f] < 50:
                computed += 11 * f
            visited_factors[f] += 1
        if logger.isEnabledFor(logging.DEBUG):
            logger.debug('Count for house %03d: %05d', result, computed)
        if computed >= args.input:
            break
        result += 1

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
