from __future__ import print_function
import argparse
import logging
import sys

import collections
import re

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

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
    limits_x, limits_y, limits_z = (-50, 50), (-50, 50), (-50, 50)
    cube = collections.defaultdict(int)
    for l in args.input:
        regex = re.search('(?P<action>on|off) x=(?P<min_x>-?\d+)\.\.(?P<max_x>-?\d+),y=(?P<min_y>-?\d+)\.\.(?P<max_y>-?\d+),z=(?P<min_z>-?\d+)\.\.(?P<max_z>-?\d+)', l.strip())
        captures = regex.groupdict()
        logger.debug('Captured: %s', captures)
        action = captures['action']
        min_x, max_x = int(captures['min_x']), int(captures['max_x'])
        min_y, max_y = int(captures['min_y']), int(captures['max_y'])
        min_z, max_z = int(captures['min_z']), int(captures['max_z'])

        min_x = max(limits_x[0], min_x)
        min_y = max(limits_y[0], min_y)
        min_z = max(limits_z[0], min_z)
        max_x = min(limits_x[1], max_x)
        max_y = min(limits_y[1], max_y)
        max_z = min(limits_z[1], max_z)
        for x in range(min_x, max_x+1):
            for y in range(min_y, max_y+1):
                for z in range(min_z, max_z+1):
                    if min_x <= x <= max_x and min_y <= y <= max_y and min_z <= z <= max_z:
                        cube[(x, y, z)] = 1 if action == 'on' else 0

    result = sum(cube.values())
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
