import logging

logger = logging.getLogger('day21')

class DeterministicDice(object):
    def __init__(self, start=1, max_value=100):
        self.value = start - 1
        self.rolls = 0
        self.max_value = max_value

    def play(self):
        value = self.value + 1
        self.rolls += 1
        self.value = (self.value + 1 ) % self.max_value
        return value

class Player(object):
    def __init__(self, name, position=1, score=0):
        self.name = name
        self._position = position - 1
        self.score = score

    def copy(self):
        other = Player(self.name)
        other.position = self.position
        other.score = self.score
        return other

    @property
    def position(self):
        return self._position + 1

    @position.setter
    def position(self, value):
        self._position = value - 1

    def move(self, steps):
        self._position = (self._position + steps) % 10
        self.score += self.position

    def __repr__(self):
        return 'Player<name %s - score %03d> @ %s' % (self.name, self.score, self.position)

