import logging

logger = logging.getLogger('day23')

class CPU(object):
    def __init__(self):
        self.ptr = 0
        self.a = 0
        self.b = 0
        self.instructions = []

    def __repr__(self):
        return 'CPU<a={} b={} ptr={}>'.format(self.a, self.b, self.ptr)

    @property
    def current(self):
        return self.instructions[self.ptr]

    def step(self):
        # Finished executing
        if self.ptr >= len(self.instructions):
            return False

        logger.debug('Executing instruction: %s', self.current)
        splitted = self.current.split(' ')
        if splitted[0] == 'hlf':
            if splitted[1] == 'a':
                self.a = self.a // 2
            elif splitted[1] == 'b':
                self.b = self.b // 2
        if splitted[0] == 'tpl':
            if splitted[1] == 'a':
                self.a = self.a * 3
            elif splitted[1] == 'b':
                self.b = self.b * 3
        if splitted[0] == 'inc':
            if splitted[1] == 'a':
                self.a += 1
            elif splitted[1] == 'b':
                self.b += 1
        self.ptr += 1
        if splitted[0] in ['jmp', 'jio', 'jie']:
            offset = int(splitted[-1])
            jump = False
            if splitted[0] == 'jmp':
                jump = True
            else:
                register = splitted[1].split(',')[0]
                if register == 'a':
                    jump = ((self.a % 2) == 0 and splitted[0] == 'jie') or (self.a == 1 and splitted[0] == 'jio')
                elif register == 'b':
                    jump = ((self.b % 2) == 0 and splitted[0] == 'jie') or (self.b == 1 and splitted[0] == 'jio')
            if jump:
                self.ptr += offset - 1
        return True
