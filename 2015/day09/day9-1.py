from __future__ import print_function
import argparse
import logging
import sys

import itertools
import re

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def get_cost(graph, path):
    result = 0
    for i, city in enumerate(path[:-1]):
        result += graph[city][path[i+1]]
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
    result = sys.maxsize

    graph = {}
    places = set()
    for l in args.input:
        src, dst, weight = re.findall('(\w+) to (\w+) = (\d+)', l)[0]
        weight = int(weight)
        if src not in graph:
            graph[src] = {}
        if dst not in graph:
            graph[dst] = {}
        graph[src][dst] = weight
        graph[dst][src] = weight
        places.add(src)
        places.add(dst)

    for path in itertools.permutations(places):
        cost = get_cost(graph, path)
        if cost < result:
            result = cost
    logger.debug('Graph read: %s', graph)
    logger.debug('Places read: %s', places)

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
