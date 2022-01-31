from __future__ import print_function
import argparse
import logging
import sys

import re
import itertools
__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def compute_happiness(graph, sit):
    result = 0
    for i in range(len(sit)):
        result += graph[sit[i]][sit[(i+1) % len(sit)]]
        result += graph[sit[i]][sit[(i-1) % len(sit)]]

    return result


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

    graph = {}
    people = set()
    for l in args.input:
        regex = re.match('(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)', l)
        src, gain, weight, dst = regex.groups()
        weight = int(weight) if gain == 'gain' else -int(weight)
        people.update([src, dst])
        if src not in graph:
            graph[src] = {}
        graph[src][dst] = weight

    for sit in itertools.permutations(people):
        happiness = compute_happiness(graph, sit)
        if happiness > result:
            result = happiness
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
