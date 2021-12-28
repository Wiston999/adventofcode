from __future__ import print_function
import argparse
import logging
import sys

import collections

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

class Map(object):
    def __init__(self, data):
        self.data = {(x, y): v for x, l in enumerate(data) for y, v in enumerate(l.strip())}
        self.max_x = max(x for x, _ in self.data) + 1
        self.max_y = max(y for _, y in self.data) + 1

    def __repr__(self):
        return '\n'.join(''.join(self.data[(x, y)] for y in range(self.max_y)) for x in range(self.max_x))

    def move(self, herd):
        moved = False
        new_data = self.data.copy()
        for x in range(self.max_x):
            for y in range(self.max_y):
                v = self.data[(x, y)]
                if v == herd:
                    if herd == '>' and self.data[(x, (y+1) % self.max_y)] == '.':
                        new_data[(x, y)] = '.'
                        new_data[(x, (y+1) % self.max_y)] = v
                        moved = True
                    if herd == 'v' and self.data[((x+1) % self.max_x, y)] == '.':
                        new_data[(x, y)] = '.'
                        new_data[((x+1) % self.max_x, y)] = v
                        moved = True
        self.data = new_data
        return moved

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
    mapa = Map(args.input.readlines())
    while True:
        result += 1
        moved_left = mapa.move('>')
        logger.debug('Current map >:\n%s', mapa)
        moved_down = mapa.move('v')
        logger.debug('Current map v:\n%s', mapa)
        if not moved_down and not moved_left:
            break
        logger.info('Moved %03d times', result)

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
