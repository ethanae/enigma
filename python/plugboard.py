class Plugboard:
  def __init__(self, alphabet, output):
    self.input = alphabet
    self.output = output

  def map(self, letter):
    for i, c in enumerate(self.input):
      if c == letter:
        return self.output[i]
      if self.output[i] == letter:
        return c
    return letter