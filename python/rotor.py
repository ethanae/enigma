
class Rotor:
  def __init__(self, start, alphabet, wiring, ringSetting, turnover):
    # determine the shift for the ring setting
    shift = 0
    for c in alphabet:
      if c == ringSetting:
        break
      shift += 1

    shiftedWiring = ""
    alphabetLen = len(alphabet)
    firstAlphaChar = ord(alphabet[0])
    dotPosition = 0

    # shift all the wiring based on the shift distance
    for i, c in enumerate(wiring):
      if c == "A":
        dotPosition = (i + shift) % alphabetLen
      index = (ord(c) - firstAlphaChar + shift) % alphabetLen
      shiftedChar = alphabet[index]
      shiftedWiring += shiftedChar
    
    # shift the wiring to match the dot position on the rotor
    while shiftedWiring[dotPosition] != ringSetting:
        shiftedWiring = shiftedWiring[1:] + shiftedWiring[0]

    self.input = alphabet
    self.output = shiftedWiring
    self.turnover = turnover

    while self.input[0] != start:
      self.rotate()

  def encode_rl(self, index):
    outChar = self.output[index]
    for i, c in enumerate(self.input):
      if c == outChar:
        return i
    return -1

  def encode_lr(self, index):
    inChar = self.input[index]
    for i, c in enumerate(self.output):
      if c == inChar:
        return i
    return -1

  def rotate(self):
    self.input = self.input[1:] + self.input[0]
    self.output = self.output[1:] + self.output[0]
    return self.input[0] == self.turnover

  def ratchet_engaged(self):
    window = ord(self.turnover)
    return self.input[0] == chr(window - 1)
  

  