import logging
import operator
from functools import reduce

logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s')
logger = logging.getLogger(__file__)

class PacketParser(object):
    def __init__(self, data):
        self.data = data
        self.packets = []

    def parse(self):
        remaining = self.data
        while len(remaining) >= 11: # Min packet size
            logger.debug('Parsing: %s', remaining)
            p = Packet(remaining)
            if p.type == 4:
                logger.debug('Parsed as literal packet')
                p = LiteralPacket(remaining)
            else:
                logger.debug('Parsed as operator packet')
                p = OperatorPacket(remaining)

            remaining = p.parse()
            self.packets.append(p)

            logger.info('Processed packet: %s', p)
        logger.info('Processed stream of packets')

    def get_all_packets(self):
        pending = self.packets
        while len(pending) > 0:
            p = pending.pop(0)
            if isinstance(p, OperatorPacket):
                pending.extend(p.subpackets)
            yield p

class Packet(object):
    def __init__(self, data):
        self.raw_data = data
        self._value = None

    @property
    def version(self):
        return int(self.raw_data[:3], base=2)

    @property
    def type(self):
        return int(self.raw_data[3:6], base=2)

    @property
    def raw_info(self):
        return self.raw_data[6:]

    def __repr__(self):
        return 'Packet<version: {} - type: {}>'.format(self.version, self.type)

class LiteralPacket(Packet):
    def parse(self):
        str_value = ''
        position = 6 # discard headers
        while True:
            str_value += self.raw_data[position + 1:position + 5]
            position += 5

            if self.raw_data[position - 5] == '0': # If it was last digit
                break


        self._value = int(str_value, base=2)
        remaining = self.raw_data[position:]
        self.raw_data = self.raw_data[:position]
        return remaining

    @property
    def value(self):
        return self._value

    def __repr__(self):
        return 'LiteralPacket<version: {} - type: {} - value: {}>'.format(
            self.version,
            self.type,
            self.value,
        )

class OperatorPacket(Packet):
    def __init__(self, data):
        super().__init__(data)
        self.subpackets = []

    @property
    def length_type(self):
        return self.raw_data[6]

    @property
    def raw_info(self):
        return self.raw_data[6:]

    @property
    def operator(self):
        if self.type == 0:
            return '+'
        if self.type == 1:
            return '*'
        if self.type == 2:
            return 'min'
        if self.type == 3:
            return 'max'
        if self.type == 5:
            return '>'
        if self.type == 6:
            return '<'
        if self.type == 7:
            return '=='

    def parse(self):
        logger.debug('Parsing OperatorPacket with length_type: %s', self.length_type)
        if self.length_type == '0':
            next_bytes = int(self.raw_data[8:22], base=2) # 23 == 8+15
            logger.debug('Operator subpacket to be parsed with length: %03d (%s)', next_bytes, self.raw_data[8:22])
            parser = PacketParser(self.raw_data[22:22 + next_bytes])
            parser.parse()
            self.subpackets.extend(parser.packets)
            remaining = self.raw_data[22 + next_bytes:]
            self.raw_data = self.raw_data[:22 + next_bytes]
        else:
            next_packets = int(self.raw_data[8:18], base=2)
            logger.debug('Operator subpacket to be parsed with %03d (%s) packets contained', next_packets, self.raw_data[8:18])
            remaining = self.raw_data[18:]
            for i in range(next_packets):
                p = Packet(remaining)
                if p.type == 4:
                    p = LiteralPacket(remaining)
                else:
                    p = OperatorPacket(remaining)

                remaining = p.parse()
                logger.debug('Parsed subpacket %s of %s: %s', i + 1, self, p)
                self.subpackets.append(p)

        return remaining

    def __repr__(self):
        return 'OperatorPacket<version: {} - type: {} - length_type: {} - subpackets: {}>'.format(
            self.version,
            self.type,
            self.length_type,
            len(self.subpackets),
        )

    @property
    def value(self):
        value = None
        subp_values = [v.value for v in self.subpackets]
        logger.debug('Operating %s on values: %s', self.operator, subp_values)
        if self.operator == '+':
            value = sum(subp_values)
        if self.operator == '*':
            value = reduce(operator.mul, subp_values, 1)
        if self.operator == 'max':
            value = max(subp_values)
        if self.operator == 'min':
            value = min(subp_values)
        if self.operator == '<':
            value = int(subp_values[0] < subp_values[1])
        if self.operator == '>':
            value = int(subp_values[0] > subp_values[1])
        if self.operator == '==':
            value = int(subp_values[0] == subp_values[1])

        return value

