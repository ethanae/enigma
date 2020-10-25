from rotor import Rotor 
from plugboard import Plugboard 
from reflector import Reflector 

class Enigma:
  def __init__(self, rotors, alphabet, reflector, plugboard):
    plugboardIn = ""
    plugboardOut = ""

    if plugboard != "":
      plugboardPairs = plugboard.replace(" ", "").split(",")
      for pair in plugboardPairs:
        plugboardIn += pair[0]
        plugboardOut += pair[1]

    builtRotors = []
    for r in rotors:
      builtRotors.append(Rotor(
        r['start'],
        alphabet,
        r['wiring'],
        r['ring'],
        r['turnoverLetter']
      ))
    
    self.rotors = builtRotors
    self.alphabet = alphabet
    self.reflector = Reflector(alphabet, reflector)
    self.plugboard = Plugboard(plugboardIn, plugboardOut)

  def handle_notch_rotations(self, rotorToTurn):
    if rotorToTurn != -1:
      if self.rotors[rotorToTurn].rotate():
          self.handle_notch_rotations(rotorToTurn - 1)
  
  def encrypt(self, plaintext):
    rotors = self.rotors
    rotorCount = len(self.rotors)
    cipher = ""

    for letter in plaintext:
      index = -1
      nextLetter = self.plugboard.map(letter)

      for j, c in enumerate(self.alphabet):
        if c == nextLetter:
          index = j
          break
      
      if index == -1:
        raise Exception("Invalid input. Character '" + letter + "' not in alphabet '" + self.alphabet + "'")

      # traverse rotors in reverse - right-to-left
      # also check whether a rotation of one rotator casued its neighbour to rotate
      didRotateNextRotors = False
      rlIndex = rotorCount - 1
      while rlIndex > -1:
        if rlIndex == rotorCount - 1:
          if rotors[rlIndex].rotate() == True:
            didRotateNextRotors = True
            self.handle_notch_rotations(rlIndex -1)
        elif not didRotateNextRotors and rotors[rlIndex].ratchet_engaged():
          self.handle_notch_rotations(rlIndex -1)
        index = rotors[rlIndex].encode_rl(index)
        rlIndex -= 1

      # reflect the signal
      index = self.reflector.reflect(index)

      for rotor in rotors:
        index = rotor.encode_lr(index)

      # map the signal back through the plugboard
      nextLetter = self.plugboard.map(self.alphabet[index])

      for i, l in enumerate(self.alphabet):
        if l == nextLetter:
          index = i
      
      cipher += self.alphabet[index]
    return cipher
