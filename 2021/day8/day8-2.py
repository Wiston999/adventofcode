from __future__ import print_function
import argparse
import logging
import sys

from itertools import permutations

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

class SegmentMap(object):
    NUMBERS = [
        [0, 1, 2, 4, 5, 6],
        [2, 5],
        [0, 2, 3, 4, 6],
        [0, 2, 3, 5, 6],
        [1, 2, 3, 5],
        [0, 1, 3, 5, 6],
        [0, 1, 3, 4, 5, 6],
        [0, 2, 5],
        [0, 1, 2, 3, 4, 5, 6],
        [0, 1, 2, 3, 5, 6],
    ]

    def __init__(self, mapping):
        self.mapping = mapping

    def __repr__(self):
        return '[{}]'.format(', '.join(self.mapping))

    def get_number(self, data):
        data_mapped = sorted([self.mapping.index(d) for d in data])
        logger.debug('Mapped data to %s from %s', data_mapped, data)
        if data_mapped in self.NUMBERS:
            return self.NUMBERS.index(data_mapped)
        return None

def guess_mapping(data):
    for mapping in permutations(['a', 'b', 'c', 'd', 'e', 'f', 'g']):
        sm = SegmentMap(mapping)

        if all(sm.get_number(d) is not None for d in data):
            return sm
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
    data_lines = args.input.readlines()
    logger.debug('Read data: %s', data_lines)

    for l in data_lines:
        data_input = l.split('|')[0].strip().split(' ')
        data_output = l.split('|')[-1].strip().split(' ')
        segment_map = guess_mapping(data_input)

        if segment_map is None:
            logger.critical('No segment map found for %s', l)
            sys.exit(1)

        result += int(''.join(str(segment_map.get_number(d)) for d in data_output))
        logger.info('Found segment map %s for input %s, partial result is %05d', segment_map, data_input, result)

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
