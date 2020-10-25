class Reflector:
  def __init__(self, alphabet, output):
    self.alphabet = alphabet
    self.output = output

  def reflect(self, index):
    reflected = self.output[index]
    for i, c in enumerate(self.alphabet):
      if reflected == c:
        return i
    return -1