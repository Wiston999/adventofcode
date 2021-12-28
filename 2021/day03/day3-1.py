from __future__ import print_function
import argparse
import logging
import sys

from statistics import mode, StatisticsError

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

def find_rate(lines, rate, position = 0):
    if len(lines) == 1:
        return int(lines[0], 2)

    try:
        moda = mode(l[position] for l in lines)
    except StatisticsError:
        moda = '1'

    # If rate is not positive, revert mode
    if rate != '+' and not all(l[position] == moda for l in lines):
        moda = '0' if moda == '1' else '1'

    logger.info('Recursion with position %s and %s lines, found mode %s', position, len(lines), moda)
    logger.debug('Lines: %s', lines)

    filtered_lines = [l for l in lines if l[position] == moda]

    return find_rate(filtered_lines, rate, position + 1)

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

    lines = list(args.input.readlines())
    line_len = len(lines[0]) - 1 # Remove line break char
    logger.debug('Read %s lines', len(lines))
    logger.debug('Lines lenght is %s', line_len)

    o2 = find_rate(lines, '+')
    co2 = find_rate(lines, '-')

    logger.debug('Found O2:  %s', o2)
    logger.debug('Found CO2: %s', co2)

    result = o2 * co2
    print ("Result is", result, file=args.output)

if __name__ == '__main__':
    main()
