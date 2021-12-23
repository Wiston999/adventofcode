from __future__ import print_function
import argparse
import logging
import sys

from queue import PriorityQueue as pq
from functools import lru_cache

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def neighbors(x, y):
    return [
        (x+1, y),
        (x, y+1),
        (x-1, y),
        (x, y-1),
    ]

def construct_path(came_from, current):
    logger.debug('Reconstructing path from %s: %s', current, came_from)
    path = [current]
    while current in came_from:
        current = came_from[current]
        path.insert(0, current)

    return path

@lru_cache(maxsize=None)
def cost(score, current, goal):
    return abs(current[0] - goal[0]) * 1.0 + abs(current[1] - goal[1]) * 1.0 + score * 0.2

def heuristic(mapa, start, goal):
    return abs(goal[0] - start[0]) + abs(goal[1] - start[1])

def a_star(mapa, start, goal):
    open_set = pq()
    open_set.put((0, start))
    came_from = {}

    g_score = {}
    g_score[start] = 0

    f_score = {}
    f_score[start] = heuristic(mapa, start, goal)

    while not open_set.empty():
        current = open_set.get()[1]
        if current == goal:
            logger.debug('GScore was: %s', g_score)
            logger.debug('FScore was: %s', f_score)
            logger.debug('Path scores was: %s', came_from)
            return construct_path(came_from, current)

        neig = [n for n in neighbors(*current) if n in mapa]
        for n in neig:
            tentative = g_score[current] + mapa[n]
            logger.debug('Tentative to %s from %s: %s', n, current, tentative)
            if tentative < g_score.get(n, sys.maxsize):
                logger.debug('Adding %s to candidates list', n)
                came_from[n] = current
                g_score[n] = tentative
                f_score[n] = tentative + heuristic(mapa, n, goal)

                if n not in open_set.queue:
                    open_set.put((f_score[n], n))


def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('--np', action='store_true',
            help='Don\'t print final path')
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

    mapa = {(x, y): int(v) for x, l in enumerate(args.input) for y, v in enumerate(l.strip())}
    path = {}
    logger.info('Read map')

    start_x, start_y = 0, 0

    path = a_star(mapa, (0, 0), (max(x for x, _ in mapa.keys()), max(y for _, y in mapa.keys())))

    logger.info('Path found: %s', path)
    result = sum(mapa[p] for p in path[1:]) # First point is not used

    if not args.np:
        print ("Path followed:", file=args.output)
        for x in range(max(x for x, _ in mapa.keys()) + 1):
            for y in range(max(y for _, y in mapa.keys()) + 1):
                if (x, y) in path:
                    print (mapa[(x, y)], "", end="", file=args.output)
                else:
                    print ('. ', end="", file=args.output)
            print ('', file=args.output)

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
