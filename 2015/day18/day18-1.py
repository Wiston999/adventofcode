from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def print_grid(grid, max_x, max_y):
    r = ''
    for x in range(max_x):
        for y in range(max_y):
            r += '#' if grid[(x, y)] else '.'
        r += '\n'
    return r

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('steps', type=int,
            help='Number of steps')
    arg_parser.add_argument('-i', '--input', type=argparse.FileType('r'), default=sys.stdin,
            help='Intput file, use - for stdin')
    arg_parser.add_argument('-p', '--print', action='store_true',
            help='Print grid in each iteration')
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
    grid = {(x, y): c == '#' for x, l in enumerate(args.input) for y, c in enumerate(l)}
    max_x = min(100, max(x for x, _ in grid) + 1)
    max_y = min(100, max(y for _, y in grid) + 1)

    logger.info('Read grid of %s x %s', max_x, max_y)

    for s in range(args.steps):
        logger.info('Current step is %04d: %06d', s + 1, sum(grid.values()))
        if args.print:
            logger.debug(print_grid(grid, max_x, max_y))
        tmp_grid = grid.copy()
        for (x, y), v in tmp_grid.items():
            on = sum(tmp_grid[(xx, yy)] for xx in range(max(0, x - 1), min(max_x, x + 2))
                    for yy in range(max(0, y - 1), min(max_y, y + 2)) if xx != x or yy != y)
            logger.debug('(%s, %s) {%s} [(%s, %s) -> (%s, %s)] neighbors on: %02d',
                    x, y, v,
                    max(0, x - 1), max(0, y - 1),
                    min(max_x, x + 2), min(max_y, y + 2),
                    on)
            if v:
                if on not in [2, 3]:
                    grid[(x, y)] = False
            else:
                if on == 3:
                    grid[(x, y)] = True

    result = sum(grid.values())
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
