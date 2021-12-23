from __future__ import print_function
import argparse
import logging
import sys

from lib import *

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger('day21')

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-g', '--goal', type=int, default=1000,
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

    dice = DeterministicDice()
    players = [
        Player('player1', int(lines[0].split(' ')[-1])),
        Player('player2', int(lines[1].split(' ')[-1])),
    ]
    turn = 0
    while True:
        if any(p.score >= args.goal for p in players):
            logger.info('One of the players reached score %s, P1: %s P2: %s', args.goal, players[0].score, players[1].score)
            break

        steps = [dice.play() for _ in range(3)]
        players[turn].move(sum(steps))
        logger.debug('Turn %s, steps %s, players: %s', turn, steps, players)
        turn = (turn + 1) % 2

    result = min(p.score for p in players) * dice.rolls
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
