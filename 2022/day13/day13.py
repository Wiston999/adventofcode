import argparse
import logging
import sys

import json

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

def sorted_pairs(l, r):
    logger.debug(f'Comparing {l} to {r}')
    if isinstance(l, int) and isinstance(r, int):
        if l == r:
            return None
        return l < r
    elif isinstance(l, int) and isinstance(r, list):
        return sorted_pairs([l], r)
    elif isinstance(l, list) and isinstance(r, int):
        return sorted_pairs(l, [r])
    else:
        for i, e in enumerate(l):
            if len(r) > i:
                comp = sorted_pairs(l[i], r[i])
                if isinstance(comp, bool):
                    return comp
        if len(l) < len(r):
            return True
        elif len(l) > len(r):
            return False
    return

class Pair(object):
    def __init__(self):
        self.L = []
        self.R = []

    def is_sorted(self):
        return sorted_pairs(self.L, self.R)

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-o', '--output', type=argparse.FileType('w'), default=sys.stdout,
            help='Output file, use - for stdout')
    arg_parser.add_argument('-i', '--input', type=argparse.FileType('r'), default='input.txt',
            help='Input file, use - for stdin')
    arg_parser.add_argument('-l', '--loglevel', type=str.upper, default='info',
            choices=['DEBUG', 'INFO', 'WARNING', 'ERROR'], help='Output file (when new set is created)')
    args = arg_parser.parse_args(argv)

    logger.setLevel(args.loglevel)

    logger.debug('Log level  : %s', args.loglevel)
    logger.debug('Output file: %s', args.output.name)
    logger.debug('Input file : %s', args.input.name)

    p = Pair()
    pairs = []
    packets = []
    for i, l in enumerate(args.input):
        if i % 3 == 2:
            pairs.append(p)
            p = Pair()
        elif i % 3 == 0:
            p.L = json.loads(l)
            packets.append(p.L)
            logger.debug('Parsed L: %s', p.L)
        else:
            p.R = json.loads(l)
            packets.append(p.R)
            logger.debug('Parsed R: %s', p.R)
    pairs.append(p)

    result = 0
    for i, p in enumerate(pairs):
        if p.is_sorted():
            logger.info(f'Pair {i+1} is sorted')
            result += i + 1

    print(f"Result 1 is {result}", file=args.output)

    packets.append([[2]])
    packets.append([[6]])

    n = len(packets)
    for i in range(n):
        for j in range(0, n-i-1):
            if not sorted_pairs(packets[j], packets[j+1]):
                packets[j], packets[j+1] = packets[j+1], packets[j]

    result = 1
    for i, p in enumerate(packets):
        if p == [[2]] or p == [[6]]:
            result *= i + 1
        logger.debug('[%03d] Packet: %s', i+1, p)

    print(f"Result 2 is {result}", file=args.output)

if __name__ == '__main__':
    main()
