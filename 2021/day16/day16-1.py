import argparse
import logging
import sys

from lib import *

__version__ = '0.1.0'

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

    binary = ''.join(f'{int(c, base=16):0>4b}' for c in args.input.read().strip())

    parser = PacketParser(binary)
    parser.parse()

    all_packets = list(parser.get_all_packets())
    logger.debug('All packets found: %s', all_packets)
    result = sum(p.version for p in all_packets)
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
