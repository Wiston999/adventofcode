from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

class Line(object):
    def __init__(self, data):
        parts = data.split('->')
        self.start_x, self.start_y = map(int, parts[0].split(','))
        self.end_x, self.end_y = map(int, parts[1].split(','))

    def is_horizontal(self):
        return self.start_x == self.end_x

    def is_vertical(self):
        return self.start_y == self.end_y

    @property
    def points(self):
        if self.start_x == self.end_x:
            start = min([self.start_y, self.end_y])
            end = max([self.start_y, self.end_y])
            for y in range(start, end + 1):
                yield (self.start_x, y)
        if self.start_y == self.end_y:
            start = min([self.start_x, self.end_x])
            end = max([self.start_x, self.end_x])
            for x in range(start, end + 1):
                yield (x, self.start_y)

    def __repr__(self):
        return '({}, {}) -> ({}, {})'.format(
            self.start_x,
            self.start_y,
            self.end_x,
            self.end_y,
        )

class Map(object):
    def __init__(self, size):
        self.data = [[0] * size for _ in range(size)]

    def mark_line(self, line):
        for point in line.points:
            logger.debug('Marking point %s', point)
            self.data[point[0]][point[1]] += 1

    def intersections(self):
        return sum(sum(x > 1 for x in row) for row in self.data)

    def __repr__(self):
        return '\n'.join(''.join(map(str, row)) for row in self.data)

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
    lines = []
    map_size = 0
    for l in args.input:
        line = Line(l)
        if line.is_horizontal() or line.is_vertical():
            logger.debug('Line %s used', line)
            map_size = max([map_size, line.start_x, line.end_x, line.start_y, line.end_y])
            lines.append(line)

    logger.info('Generating map of size %s', map_size)
    m = Map(map_size + 1)

    for l in lines:
        logger.info('Marking line %s', l)
        m.mark_line(l)

    logger.debug('Map after marks is: %s', m)
    result = m.intersections()
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
