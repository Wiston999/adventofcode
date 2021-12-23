from __future__ import print_function
import argparse
import logging
import sys

from lib import *
import itertools
from functools import lru_cache
import collections

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger('day21')

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-g', '--goal', type=int, default=21,
            help='Goal score')
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

    lines = args.input.readlines()

    dice = DeterministicDice(max_value=3)
    turn = 0
    players = [
        Player('player1', int(lines[0].split(' ')[-1])),
        Player('player2', int(lines[1].split(' ')[-1])),
    ]

    states = {(players[0].position, 0, players[1].position, 0): 1}
    counts = collections.Counter(
        sum(l) for l in list(itertools.product(range(1, 4), repeat=3))
    )

    player1_wins, player2_wins = 0, 0
    while states:
        current_states = collections.defaultdict(int)
        logger.debug('Current States: %s', states)
        for roll_value, roll_count in counts.items():
            for (p1_position, p1_score, p2_position, p2_score), state_count in states.items():
                p1_position = ((p1_position + roll_value - 1) % 10) + 1
                p1_score += p1_position
                if p1_score >= args.goal:
                    player1_wins += state_count * roll_count
                else:
                    current_states[(p1_position, p1_score, p2_position, p2_score)] += roll_count * state_count
        states = current_states.copy()
        current_states = collections.defaultdict(int)
        for roll_value, roll_count in counts.items():
            for (p1_position, p1_score, p2_position, p2_score), state_count in states.items():
                p2_position = ((p2_position + roll_value - 1) % 10) + 1
                p2_score += p2_position
                if p2_score >= args.goal:
                    player2_wins += state_count * roll_count
                else:
                    current_states[(p1_position, p1_score, p2_position, p2_score)] += roll_count * state_count
        states = current_states.copy()

    result = max(player1_wins, player2_wins)
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
