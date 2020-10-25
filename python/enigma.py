from python.rotor import Rotor 
from python.plugboard import Plugboard 

class Enigma:
  def __init__(self, rotors, alphabet, reflector, plugboard):
    plugboardPairs = plugboard.replace(" ", "").split(",")
    plugboardIn = ""
    plugboardOut = ""

    for pair in plugboardPairs:
      plugboardIn += pair[0]
      plugboardOut += pair[1]

    builtRotors = []
    for r in rotors:
      builtRotors.append(Rotor(
        r.start,
        r.alphabet,
        r.wiring,
        r.ringSetting,
        r.turnover
      ))
    
    self.rotors = builtRotors
    self.alphabet = alphabet
    self.reflector = reflector
    self.plugboard = Plugboard(plugboardIn, plugboardOut)
  
  def encrypt(self, plaintext):
    rotors = self.rotors
    rotorCount = len(self.rotors)
    cipher = ""

    for letter in plaintext:
      index = -1
      nextLetter = self.plugboard.map(letter)

      for j, c in self.alphabet:
        if c == nextLetter:
          index = j
          break
      
      if index == -1:
        return "Invalid input. Character '" + letter + "' not in alphabet '" + self.alphabet + "'"

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
        index = rotors[rlIndex].encode_rl()
        rlIndex -= 1

      # reflect the signal
      index = self.reflector.reflect(index)

      for rotor in rotors:
        index = rotor.encode_lr(index)

      # map the signal back through the plugboard
      nextLetter = self.plugboard.map(alphabet[index])

      for i, l in self.alphabet:
        if l == nextLetter:
          index = i
      
      cipher += alphabet[index]
    return cipher
