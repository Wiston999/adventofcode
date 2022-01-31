from __future__ import print_function
import argparse
import logging
import sys

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

def password_gen(current='aaaaaaaa'):
    current = list(map(ord, current))
    while not all(c == 122 for c in current): # 122 == ord('z')
        yield ''.join(chr(c) for c in current)
        for i in range(len(current) - 1, 0, -1):
            current[i] += 1
            if current[i] > 122:
                current[i] = 97 # 97 == ord('a')
            else:
                break

def is_valid_password(value):
    if 'i' in value or 'o' in value or 'l' in value:
        return False

    len_value = len(value)
    pairs = 0
    stairs = 0
    for i in range(len_value - 1):
        j = i
        while j < len_value and value[i] == value[j]:
            j += 1
        if j == (i + 2):
            pairs += 1
        j = i + 1
        while j < (len_value - 1) and ord(value[j]) == (ord(value[j + 1]) - 1):
            j += 1
        if (j - i) >= 3:
            stairs += 1

    logger.debug('Found %02d pairs and %02d stairs in %s', pairs, stairs, value)
    if pairs < 2:
        return False

    if stairs < 1:
        return False

    return True

def main(argv=None):
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument('-v', '--version', action='version',
            version='%%(prog)s v%s' % __version__)
    arg_parser.add_argument('input', type=str,
            help='Input value')
    arg_parser.add_argument('-o', '--output', type=argparse.FileType('w'), default=sys.stdout,
            help='Output file, use - for stdout')
    arg_parser.add_argument('-l', '--loglevel', type=str.upper, default='info',
            choices=['DEBUG', 'INFO', 'WARNING', 'ERROR'], help='Output file (when new set is created)')
    args = arg_parser.parse_args(argv)

    logger.setLevel(args.loglevel)

    logger.debug('Log level: %s', args.loglevel)
    logger.debug('Input: %s', args.input)
    logger.debug('Output file: %s', args.output.name)
    result = 0

    for i, current in enumerate(password_gen(args.input)):
        logger.debug('Generated %04d passwords', i + 1)
        if is_valid_password(current):
            result = current
            break

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
