from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

def find_basin(mapa, x, y, limit_x, limit_y, used):
    element = mapa[x][y]
    lower_locations = []
    if (x - 1) >= 0 and mapa[x-1][y] > element and mapa[x-1][y] != 9 and (x-1, y) not in used:
        lower_locations.append((x-1, y))
    if (y - 1) >= 0 and mapa[x][y-1] > element and mapa[x][y-1] != 9 and (x, y-1) not in used:
        lower_locations.append((x, y-1))
    if (x + 1) < limit_x and mapa[x+1][y] > element and mapa[x+1][y] != 9 and (x+1, y) not in used:
        lower_locations.append((x+1, y))
    if (y + 1) < limit_y and mapa[x][y+1] > element and mapa[x][y+1] != 9 and (x, y+1) not in used:
        lower_locations.append((x, y+1))

    used.update(lower_locations)
    logger.debug('Found lower positions from %s (%s, %s): %s', element, x, y, lower_locations)
    return [(x, y, element)] + sum([find_basin(mapa, nx, ny, limit_x, limit_y, used) for nx, ny in lower_locations], [])

def print_basin(basin):
    min_x = min(x for x, _, _ in basin)
    max_x = max(x for x, _, _ in basin)
    min_y = min(y for _, y, _ in basin)
    max_y = max(y for _, y, _ in basin)

    text = ''
    for x in range(min_x, max_x + 1):
        for y in range(min_y, max_y + 1):
            element = [b[2] for b in basin if b[0] == x and b[1] == y]
            if len(element) > 0:
                text += str(element[0]) + ' '
            else:
                text += '  '
        text += '\n'

    return text

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
    mapa = [[int(e) for e in l.strip()] for l in args.input.readlines()]
    logger.debug('Map is: %s', mapa)

    limit_x = len(mapa)
    limit_y = len(mapa[0])

    basins = []
    lower_points = []

    for x in range(limit_x):
        for y in range(limit_y):
            element = mapa[x][y]
            surrounding = []
            if (x - 1) >= 0:
                surrounding.append(mapa[x-1][y])
            if (x + 1) < limit_x:
                surrounding.append(mapa[x+1][y])
            if (y - 1) >= 0:
                surrounding.append(mapa[x][y-1])
            if (y + 1) < limit_y:
                surrounding.append(mapa[x][y+1])
            if element < min(surrounding):
                lower_points.append((x, y))

    logger.info('Found %s lower points', len(lower_points))
    for x, y in lower_points:
        element = mapa[x][y]
        basin = find_basin(mapa, x, y, limit_x, limit_y, set())
        logger.debug('Found basin: %s', basin)
        print("Basin found:", file=args.output)
        print(print_basin(basin), file=args.output)
        basins.append(len(basin))
        logger.info('Found basin of size: %04d from (%s, %s)', basins[-1], x, y)

    basins = sorted(basins, reverse=True)
    logger.info('3 biggest basins size: %s', basins[:3])
    result = basins[0]*basins[1]*basins[2]

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
