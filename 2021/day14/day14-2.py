from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig()
logger = logging.getLogger(__file__)

def find_patterns(s, pattern):
    result = []
    position = 0
    while True:
        position = s.find(pattern, position)
        if position == -1:
            break
        logger.debug('Found %s at %04d in %s', pattern, position, s)
        result.append(position)
        position += 1
    return result

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('-s', '--steps', type=int, default=10,
            help='Number of steps')
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

    tmpl = ''
    insertion_pairs = {}

    unique_chars = set()

    for l in args.input:
        l = l.strip()
        if '->' in l:
            k, v = l.split(' -> ')
            insertion_pairs[k] = v
            unique_chars.add(v)
        elif len(l) > 0:
            tmpl = l

    unique_chars.update(tmpl)

    logger.info('Read polymer: %s', tmpl)
    logger.info('Read insertion rules: %s', insertion_pairs)

    c_count = {c: tmpl.count(c) for c in unique_chars}

    pairs_count = {p: tmpl.count(p) for p in insertion_pairs.keys()}

    for i in range(args.steps):
        logger.info('Step %03d of %03d', i + 1, args.steps)
        logger.debug('Chars count: %s', c_count)
        logger.debug('Pairs count: %s', pairs_count)

        tmp_count = pairs_count.copy()
        for p, t in insertion_pairs.items():
            if tmp_count[p] > 0:
                c_count[t] += tmp_count[p]
                pairs_count[p] -= tmp_count[p]
                new_pair = '{}{}'.format(p[0], t)
                if new_pair in insertion_pairs:
                    pairs_count[new_pair] += tmp_count[p]

                new_pair = '{}{}'.format(t, p[1])
                if new_pair in insertion_pairs:
                    pairs_count[new_pair] += tmp_count[p]

    counters_values = sorted(c_count.values())
    result = counters_values[-1] - counters_values[0]
    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
