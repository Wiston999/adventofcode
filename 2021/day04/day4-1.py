from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

class Board(object):
    def __init__(self, data):
        self.data = []
        for l in data:
            self.data.append(list(map(int, l.split())))

    def __repr__(self):
        return '\n'.join(' '.join(map(str, l)) for l in self.data)

    def mark(self, n):
        for i, row in enumerate(self.data):
            self.data[i] = [-1 if item == n else item for item in row]

    def winner(self):
        any_row = any(all(item == -1 for item in r) for r in self.data)
        any_column = any(all(row[c] == -1 for row in self.data) for c in range(len(self.data[0])))
        return any_row or any_column

    def unmarked_sum(self):
        return sum(sum(filter(lambda x: x != -1, row)) for row in self.data)

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

    lines = list(args.input.readlines())
    game = list(map(int, lines[0].split(',')))

    logger.info('Game read: %s', game)

    boards = []

    tmp_board = []
    for l in lines[2:]:
        l = l.strip()
        if len(l) == 0:
            boards.append(Board(tmp_board))
            logger.debug('Read board %s from lines %s', boards[-1], tmp_board)
            tmp_board = []
            continue
        tmp_board.append(l)

    print ('Starting bingo using', len(boards), 'boards', file=args.output)

    winner_board = None
    for n in game:
        logger.info('Playing number %s', n)
        for board in boards:
            board.mark(n)
            if board.winner():
                winner_board = board
                break

        if winner_board is not None:
            break

    logger.info('Winner board is %s', winner_board)
    logger.info('Won at number: %s', n)
    result = n * winner_board.unmarked_sum()
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
