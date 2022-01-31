from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def count(data):
    previous = data[0]
    current_count = 1
    result = ''
    for i in range(1, len(data)):
        if previous == data[i]:
            current_count += 1
        else:
            result += '{}{}'.format(current_count, previous)
            previous = data[i]
            current_count = 1
    result += '{}{}'.format(current_count, previous)
    return result

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-n', '--iterations', type=int, default=40,
            help='Number of iterations')
    arg_parser.add_argument('input', type=str,
            help='Input value')
    arg_parser.add_argument('-o', '--output', type=argparse.FileType('w'), default=sys.stdout,
            help='Output file, use - for stdout')
    arg_parser.add_argument('-l', '--loglevel', type=str.upper, default='info',
            choices=['DEBUG', 'INFO', 'WARNING', 'ERROR'], help='Output file (when new set is created)')
    args = arg_parser.parse_args(argv)

    logger.setLevel(args.loglevel)

    logger.debug('Log level: %s', args.loglevel)
    logger.debug('Input value: %s', args.input)
    logger.debug('Output file: %s', args.output.name)
    result = args.input

    for i in range(args.iterations):
        result = count(result)
        logger.debug('Iteration %02d: %s', i + 1, result)

    result = len(result)
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
