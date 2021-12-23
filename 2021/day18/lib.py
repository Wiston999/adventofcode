import logging
import math
import json
import re

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger('day18')

class NumberStr(object):
    def __init__(self, value):
        self.value = value

    def __repr__(self):
        return self.value

    def __add__(self, other):
        n = NumberStr(json.dumps([json.loads(self.value), json.loads(other.value)], separators=(',', ':')))
        n.reduce()
        return n

    def reduce(self):
        while True:
            while True:
                if not self.explode():
                    break
            if not self.split():
                break

    def split(self):
        logger.debug('Splitting %s', self)
        regexp = re.search('(\d\d)', self.value)
        if regexp:
            value = int(regexp.group(1))
            new_value = '[{},{}]'.format(math.floor(value / 2), math.ceil(value / 2))
            self.value = self.value[:regexp.start(1)] + new_value + self.value[regexp.end(1):]
        logger.debug('Splitted %s', self)
        return regexp is not None

    def explode(self, depth=5):
        logger.debug('Exploding %s', self)
        exploded = False
        stack = list()
        for i, c in enumerate(self.value):
            if c == '[':
                stack.append(i)
            if c == ']':
                stack.pop(-1)

            if len(stack) >= depth and re.match('^\[\d+,\d+\]', self.value[stack[-1]:]):
                start = stack[-1]
                end = self.value.index(']', start)
                logger.debug('Found explode point at %s to %s', start, end)
                values = list(map(int, self.value[start + 1:end].split(',')))
                next_value_re = re.search('(\d+)', self.value[end:])
                if next_value_re is not None:
                    next_value_start = next_value_re.start(1)
                    next_value_end = next_value_re.end(1)
                    next_value = int(next_value_re.group(1)) + values[-1]
                    self.value = self.value[:end+next_value_start] + str(next_value) + self.value[end+next_value_end:]
                    logger.debug('Exploded to left: %s', self.value)

                prev_values = list(re.finditer('(\d+)', self.value[:start]))
                if prev_values:
                    prev_start = prev_values[-1].start()
                    prev_end = prev_values[-1].end()
                    prev_value = int(self.value[prev_start:prev_end]) + values[0]
                    self.value = self.value[:prev_start] + str(prev_value) + self.value[prev_end:]
                    start += len(str(prev_value)) - (prev_end-prev_start)
                    end += len(str(prev_value)) - (prev_end-prev_start)
                    logger.debug('Exploded to right: %s', self.value)

                self.value = self.value[:start] + '0' + self.value[end+1:]
                exploded = True
                break
        logger.debug('Exploded %s', self)
        return exploded

    @property
    def magnitude(self):
        return Number.parse(json.loads(self.value)).magnitude

class Number(object):
    @classmethod
    def parse(cls, value):
        if isinstance(value[0], list):
            v0 = cls.parse(value[0])
        else:
            v0 = value[0]

        if isinstance(value[1], list):
            v1 = cls.parse(value[1])
        else:
            v1 = value[1]

        return cls([v0, v1])


    def __init__(self, value):
        self.value = value

    def __repr__(self):
        return '[{}, {}]'.format(self.first, self.second)

    @property
    def first(self):
        return self.value[0]

    @first.setter
    def first(self, value):
        self.value[0] = value

    @property
    def second(self):
        return self.value[1]

    @second.setter
    def second(self, value):
        self.value[1] = value

    def copy(self):
        return Number([
            self.first.copy() if isinstance(self.first, Number) else self.first,
            self.second.copy() if isinstance(self.second, Number) else self.second,
        ])

    def ladd(self, value):
        if isinstance(self.first, Number):
            return self.first.ladd(value)
        else:
            self.first += value
            return True
        return False

    def radd(self, value):
        if isinstance(self.second, Number):
            return self.second.radd(value)
        else:
            self.second += value
            return True
        return False

    def split(self):
        change = False
        if isinstance(self.first, int) and self.first >= 10:
            self.first = Number([
                math.floor(self.first / 2),
                math.ceil(self.first / 2),
            ])
            change = True
            logger.debug('Splitted %s @ %s', self.first, self)
        elif isinstance(self.first, Number):
            change = self.first.split()
        if isinstance(self.second, int) and self.second >= 10:
            self.second = Number([
                math.floor(self.second / 2),
                math.ceil(self.second / 2),
            ])
            change = True
            logger.debug('Splitted %s @ %s', self.second, self)
        elif isinstance(self.second, Number):
            change = self.second.split()
        return change

    def explode(self, level=4):
        change = False
        left, right = None, None
        if level == 0:
            logger.debug('Found explode level 0: %s', self)
            change = True
            left, right = self.first, self.second
        else:
            if isinstance(self.first, Number):
                change, left, right = self.first.explode(level - 1)
                logger.debug('Exploded %s: (%s, %s, %s)', self.first, change, left, right)
                if left is not None and right is not None:
                    self.first = 0
                    logger.debug('Set self.first to 0: %s', self)
                if right is not None:
                    if isinstance(self.second, int):
                        logger.debug('Adding %s to %s: %s', right, self.second, self)
                        self.second += right
                    else:
                        logger.debug('Right added %s to %s: %s', right, self.second, self)
                        self.second.ladd(right)
                    right = None
            if isinstance(self.second, Number):
                change, left, right = self.second.explode(level - 1)
                logger.debug('Exploded %s: (%s, %s, %s)', self.second, change, left, right)
                if left is not None and right is not None:
                    self.second = 0
                    logger.debug('Set self.second to 0: %s', self)
                if left is not None:
                    if isinstance(self.second, int):
                        logger.debug('Adding %s to %s: %s', left, self.second, self)
                        self.second += left
                    else:
                        logger.debug('Left added %s to %s: %s', left, self.second, self)
                        self.second.radd(left)
                    left = None
        return change, left, right

    def reduce(self):
        while True:
            while True:
                change, _, _ = self.explode()
                if not change:
                    break
            change = self.split()
            if not change:
                break

    def __eq__(self, other):
        if isinstance(other, Number):
            return self.first == other.first and self.second == self.second
        return False

    def __add__(self, other):
        n = Number.parse([self.value, other.value])
        n.reduce()
        return n

    def __mul__(self, other):
        if isinstance(other, Number):
            return self.magnitude * other.magnitude
        else:
            return self.magnitude * other

    def __rmul__(self, other):
        return other * self

    @property
    def magnitude(self):
        return self.first * 3 + self.second * 2
