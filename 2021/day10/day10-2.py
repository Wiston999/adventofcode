from __future__ import print_function
import argparse
import logging
import sys

import statistics

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

def score_line(chars):
    result = 0
    scores = {'(': 1, '[': 2, '{': 3, '<': 4}
    for c in reversed(chars):
        result = (result * 5) + scores[c]
    return result

def check_line(line):
    stack = []

    for c in line:
        if c in '([{<':
            stack.append(c)
        elif c in ')]}>':
            opening_c = stack.pop(-1)
            if opening_c == '(' and c != ')':
                return c, stack
            if opening_c == '[' and c != ']':
                return c, stack
            if opening_c == '{' and c != '}':
                return c, stack
            if opening_c == '<' and c != '>':
                return c, stack

    return None, stack

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
    line_scores = []

    lines = args.input.readlines()

    for i, l in enumerate(lines):
        logger.info('Processing line %02d / %02d: %s', i + 1, len(lines), l)

        error_c, pending_c = check_line(l)
        if error_c is not None:
            logger.debug('Found error character %s, ignoring line', error_c)
            continue

        if len(pending_c) > 0:
            line_scores.append(score_line(pending_c))
            logger.info('Line is incomplete, pending matches are: %s, score is %06d', pending_c, line_scores[-1])

    logger.info('Incomplete lines scores: %s', line_scores)
    result = statistics.median(line_scores)

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
