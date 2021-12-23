from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

def score(error):
    if error == ')':
        return 3
    elif error == ']':
        return 57
    elif error == '}':
        return 1197
    elif error == '>':
        return 25137
    else:
        logger.warning('Unknown error code: %s', error)
        return 0

def check_line(line):
    stack = []

    for c in line:
        if c in '([{<':
            stack.append(c)
        elif c in ')]}>':
            opening_c = stack.pop(-1)
            if opening_c == '(' and c != ')':
                return c
            if opening_c == '[' and c != ']':
                return c
            if opening_c == '{' and c != '}':
                return c
            if opening_c == '<' and c != '>':
                return c

    return None

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

    lines = args.input.readlines()

    for i, l in enumerate(lines):
        logger.info('Processing line %02d / %02d: %s', i + 1, len(lines), l)

        error_c = check_line(l)
        if error_c is not None:
            result += score(error_c)
            logger.debug('Found error character %s. Score is %04d', error_c, result)

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
