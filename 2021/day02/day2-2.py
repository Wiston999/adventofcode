from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
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

    result = {'x': 0, 'y': 0}
    aim = 0
    for l in args.input:
        action, steps = l.split(' ')
        if action == 'forward':
            result['x'] += int(steps)
            result['y'] += int(steps) * aim
        elif action == 'up':
            aim -= int(steps)
        elif action == 'down':
            aim += int(steps)
        else:
            logger.warning('Unknown action: %s', action)

    logger.debug('Partial result is: %s', result)
    print ("Result is", result['x'] * result['y'], file=args.output)

if __name__ == '__main__':
    main()
