from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

class Sheet(object):
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.data = [['.'] * x for _ in range(y)]

    def add_point(self, x, y):
        self.data[y][x] = '#'

    def fold(self, axis, coord):
        logger.debug('Folding on %s-%s', axis, coord)

        max_coord = self.x if axis == 'x' else self.y

        rang = max_coord - coord - 1
        logger.debug('Range of folding is %s from %s', rang, max_coord)

        for i in range(rang):
            new_coord = coord - i - 1
            folded_coord = coord + i + 1
            logger.debug('Coordinate %s will become %s', folded_coord, new_coord)
            for j in range(self.y if axis == 'x' else self.x):
                if axis == 'x' and self.data[j][folded_coord] == '#':
                    self.data[j][new_coord] = self.data[j][folded_coord]
                elif axis == 'y' and self.data[folded_coord][j] == '#':
                    self.data[new_coord][j] = self.data[folded_coord][j]

        if axis == 'x':
            self.x = coord
        else:
            self.y = coord

        logger.debug('New sheet size is %s, %s', self.x, self.y)

    @property
    def points(self):
        return sum(self.data[j][i] == '#' for i in range(self.x) for j in range(self.y))

    def __repr__(self):
        result = ''
        for i in range(self.y):
            for j in range(self.x):
                result += self.data[i][j]
            result += '\n'
        return result

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('--np', action='store_true',
            help='Don\'t print final sheet')
    arg_parser.add_argument('-f', '--folds', type=int, default=1,
            help='Number of folds to be applied')
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
    coords = []
    folds = []

    for l in args.input:
        l = l.strip()
        if l.startswith('fold along'):
            folds.append(l.split(' ')[-1].split('='))
            folds[-1][-1] = int(folds[-1][-1])
        elif len(l) > 0:
            coords.append(list(map(int, l.split(','))))

    x = max(c[0] for c in coords) + 1
    y = max(c[1] for c in coords) + 1

    logger.info('Creating sheet with size: %s, %s', x, y)
    sheet = Sheet(x, y)
    for c in coords:
        sheet.add_point(*c)

    for i in range(min(len(folds), args.folds)):
        logger.info('Applying fold %s', i + 1)
        sheet.fold(*folds[i])

    result = sheet.points

    print ("Result is", result, file=args.output)
    if not args.np:
        print ("Folded sheet is", file=args.output)
        print (sheet, file=args.output)


if __name__ == '__main__':
    main()
