from __future__ import print_function
import argparse
import logging
import sys

from collections import defaultdict
import itertools

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def print_image(image, min_x, min_y, max_x, max_y):
    logger.debug('Printing image with borders (%s, %s) x (%s, %s)',
        min_x, max_x, min_y, max_y
    )
    result = '\n'.join(
        ''.join('#' if image[(x, y)] else '.' for y in range(min_y, max_y))
        for x in range(min_x, max_x)
    )
    return result

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('--np', action='store_true',
            help='Don\'t print intermediate images')
    arg_parser.add_argument('-s', '--steps', type=int, default=2,
            help='Number of steps to be applied')
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
    iea = lines[0].strip()

    image = defaultdict(lambda: 0, {(x, y): int(d == '#') for x, l in enumerate(lines[2:]) for y, d in enumerate(l.strip())})
    output = defaultdict(lambda: 0)
    min_x, min_y = -5, -5
    max_x, max_y = max(x for x, _ in image.keys()) + 5, max(y for _, y in image.keys()) + 5

    logger.info('Read Image Enhance algorithm of size: %s', len(iea))
    logger.info('Read image %s x %s', max_x + 1, max_y + 1)

    for i in range(args.steps):
        logger.info('Applying iteration %s', i + 1)
        for x, y in itertools.product(range(min_x, max_x), range(min_y, max_y)):
            value = int(''.join(str(image[(i, j)]) for i in range(x-1, x+2) for j in range(y-1, y+2)), base=2)
            new_value = int(iea[value] == '#')
            logger.debug('Enhanced value for (%03d, %03d) = %s --> %s', x, y, value, new_value)
            output[(x, y)] = new_value

        image = defaultdict(lambda: (i + 1) % 2 if iea[0] == '#' else 0, output.copy())
        if not args.np:
            logger.debug('Dict content: %s', image)
            print(print_image(image, min_x - 5, min_y - 5, max_x + 5, max_y + 5), file=args.output)
        logger.info(
            'Current count: %s',
            sum(image[(x, y)] for (x, y) in itertools.product(range(min_x - 4, max_x + 4), range(min_y - 4, max_y + 4)))
        )
        output = defaultdict(lambda: 0)
        max_x, max_y = max_x + 1, max_y + 1
        min_x, min_y = min_x - 1, min_y - 1

    logger.debug('Dict content: %s', image)
    # Not 100% how, but it only works if the image is printed using print_image
    # NB2: Adding logger.info current cound seems to make all work, this might be due to dynamic behavior of default dict
    result = sum(image[(x, y)] for (x, y) in itertools.product(range(min_x - 4, max_x + 4), range(min_y - 4, max_y + 4)))
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
