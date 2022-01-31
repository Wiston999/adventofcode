from __future__ import print_function
import argparse
import logging
import sys

import re
import math

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

class Reindeer(object):
    def __init__(self, name, speed, span, rest):
        self.name = name
        self.speed = speed
        self.span = span
        self.rest = rest

    @property
    def turn(self):
        return self.span + self.rest

    def distance(self, seconds):
        d = (seconds // self.turn) * self.speed * self.span
        d += min(self.span, seconds % self.turn) * self.speed
        return d

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('seconds', type=int,
            help='Number of seconds to compute')
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

    reindeers = {}
    for l in args.input:
        regex = re.match('(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.', l)
        name, speed, span, rest = regex.groups()
        reindeers[name] = Reindeer(name, int(speed), int(span), int(rest))
        logger.debug('%s will fly %s KMs', name, reindeers[name].distance(args.seconds))

    result = max(r.distance(args.seconds) for r in reindeers.values())
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
