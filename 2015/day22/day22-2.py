from __future__ import print_function
import argparse
import logging
import sys

import itertools

from utils import *
import utils

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger('day22')

BEST = sys.maxsize

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

    me = Player('me', 50, 500)
    enemy = Player('boss')

    for l in args.input:
        if 'Hit Points' in l:
            enemy.hit_points = int(l.split(' ')[-1])
        if 'Damage' in l:
            enemy.damage = int(l.split(' ')[-1])

    find_best_game(me, enemy, SPELLS, [], hard_mode=True)
    result = utils.BEST
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
