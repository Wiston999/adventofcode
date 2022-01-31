from __future__ import print_function
import argparse
import logging
import sys

import collections
import re

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

class Aunt(object):
    def __init__(self):
        self.data = {}

    def __getitem__(self, k):
        return self.data.get(k, None)

    def __setitem__(self, k, v):
        self.data[k] = v

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

    aunts = []
    for l in args.input:
        regex = re.match('Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)', l)
        _, k1, v1, k2, v2, k3, v3 = regex.groups()
        aunt = Aunt()
        aunt[k1] = int(v1)
        aunt[k2] = int(v2)
        aunt[k3] = int(v3)
        aunts.append(aunt)

    targets = {
        'children': 3,
        'cats': 7,
        'samoyeds': 2,
        'pomeranians': 3,
        'akitas': 0,
        'vizslas': 0,
        'goldfish': 5,
        'trees': 3,
        'cars': 2,
        'perfumes': 1
    }

    possible = set(range(1, 501))
    for i, a in enumerate(aunts):
        for k, v in targets.items():
            logger.debug('Testing %s with %s = %s (%s)', i + 1, k, v, a[k])
            if a[k] is not None:
                if k in ['cats', 'trees'] and a[k] <= v:
                    logger.info('Removed %s as %s (<=) = %s (%s) is not valid', i + 1, k, v, a[k])
                    possible.remove(i + 1)
                    break
                elif k in ['pomeranians', 'goldfish'] and a[k] >= v:
                    logger.info('Removed %s as %s (>=) = %s (%s) is not valid', i + 1, k, v, a[k])
                    possible.remove(i + 1)
                    break
                elif k not in ['cats', 'trees', 'pomeranians', 'goldfish'] and a[k] != v:
                    logger.info('Removed %s as %s = %s (%s) is not found', i + 1, k, v, a[k])
                    possible.remove(i + 1)
                    break

    result = possible.pop()
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
