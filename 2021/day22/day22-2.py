from __future__ import print_function
import argparse
import logging
import sys

from functools import lru_cache
import collections
import re

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

class Map(object):
    def __init__(self):
        self.cuboids = []

    def add(self, c):
        self.cuboids.append(c)

    def area(self, cuboids):
        if len(cuboids) > 0:
            logger.debug('Computing %s: %s', len(cuboids), cuboids[0])
            current = cuboids[0]
            if current.action == 'off':
                logger.debug('OFF %s', current)
                return self.area(cuboids[1:])

            this = current.area
            rest = self.area(cuboids[1:])
            intersect = self.area(self.intersections(current, cuboids[1:]))
            logger.debug('Partial value: %03d + %03d - %03d = %s', this, rest, intersect, this + rest - intersect)
            return this + rest - intersect
        else:
            return 0

    def intersections(self, cuboid, checks):
        intersections = []
        for c in checks:
            intersect = cuboid.intersection(c)
            if intersect.x != (0, 0) or intersect.y != (0, 0) or intersect.z != (0, 0):
                intersections.append(intersect)

        logger.debug('Returning %s intersections: %s', len(intersections), intersections)
        return intersections

class Cuboid(object):
    def __init__(self, x_range, y_range, z_range, action):
        self.x = x_range
        self.y = y_range
        self.z = z_range
        self.action = action

    def __repr__(self):
        return 'Cube<({}, {}) x ({}, {}) x ({}, {}) -> {}>'.format(
            self.x[0], self.x[1],
            self.y[0], self.y[1],
            self.z[0], self.z[1],
            self.action
        )

    @property
    @lru_cache(maxsize=None)
    def area(self):
        return (1 + self.x[1] - self.x[0]) * (1 + self.y[1] - self.y[0]) * (1 + self.z[1] - self.z[0])

    def intersection(self, other):
        min_x = self.x[0] if self.x[0] > other.x[0] else other.x[0]
        min_y = self.y[0] if self.y[0] > other.y[0] else other.y[0]
        min_z = self.z[0] if self.z[0] > other.z[0] else other.z[0]
        max_x = self.x[1] if self.x[1] < other.x[1] else other.x[1]
        max_y = self.y[1] if self.y[1] < other.y[1] else other.y[1]
        max_z = self.z[1] if self.z[1] < other.z[1] else other.z[1]

        intersection = max(0, max_x - min_x + 1) * max(0, max_y - min_y + 1) * max(0, max_z - min_z + 1)
        if intersection > 0:
            new_c = Cuboid((min_x, max_x), (min_y, max_y), (min_z, max_z), None)
        else:
            new_c = Cuboid((0, 0), (0, 0), (0, 0), None)

        return new_c

def count(cuboids):
    if len(cuboids):
        current = cuboids.pop(0)
    else:
        return 0

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
    planes = {
        'x': collections.defaultdict(int),
        'y': collections.defaultdict(int),
        'z': collections.defaultdict(int)
    }
    m = Map()
    for i, l in enumerate(args.input):
        regex = re.search('(?P<action>on|off) x=(?P<min_x>-?\d+)\.\.(?P<max_x>-?\d+),y=(?P<min_y>-?\d+)\.\.(?P<max_y>-?\d+),z=(?P<min_z>-?\d+)\.\.(?P<max_z>-?\d+)', l.strip())
        captures = regex.groupdict()
        logger.debug('Captured: %s', captures)
        action = captures['action']
        min_x, max_x = int(captures['min_x']), int(captures['max_x'])
        min_y, max_y = int(captures['min_y']), int(captures['max_y'])
        min_z, max_z = int(captures['min_z']), int(captures['max_z'])

        c = Cuboid((min_x, max_x), (min_y, max_y), (min_z, max_z), action)
        m.add(c)
        logger.debug('Read cuboid area: %s', c.area)

    result = m.area(m.cuboids)
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
