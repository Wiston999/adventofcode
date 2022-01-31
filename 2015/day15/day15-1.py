from __future__ import print_function
import argparse
import logging
import sys

import re
import itertools

__version__ = '0.1.0'

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

class Ingredient(object):
    def __init__(self, capacity, durability, flavor, texture, calories):
        self.capacity = capacity
        self.durability = durability
        self.flavor = flavor
        self.texture = texture
        self.calories = calories

    def quality(self, value):
        return self.capacity * value, self.durability * value, self.flavor * value, self.texture * value, self.calories * value

def combinations(capacity, ingredients):
    return (pair for pair in itertools.product(range(capacity+1), repeat=ingredients) if sum(pair) == capacity)

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

    ingredients = {}
    for l in args.input:
        regex = re.match('(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)', l)
        name, capacity, durability, flavor, texture, calories = regex.groups()
        ingredients[name] = Ingredient(int(capacity), int(durability), int(flavor), int(texture), int(calories))

    ingredient_names = ingredients.keys()
    for combination in combinations(100, len(ingredients)):
        capacity, durability, flavor, texture = 0, 0, 0, 0
        for i, p in enumerate(combination):
            ingredient = ingredients[ingredient_names[i]]
            i_capacity, i_durability, i_flavor, i_texture, _ = ingredient.quality(p)
            capacity += i_capacity
            durability += i_durability
            flavor += i_flavor
            texture += i_texture
        if any(c < 0 for c in [capacity, durability, flavor, texture]):
            continue
        logger.debug('Combination %s gave result: %06d', combination, capacity * durability * flavor * texture)
        if result < (capacity * durability * flavor * texture):
            logger.info('Best combination so far: %s (%s)', combination, capacity * durability * flavor * texture)
            result = capacity * durability * flavor * texture

    print ("Result is", result, file=args.output)


if __name__ == '__main__':
    main()
