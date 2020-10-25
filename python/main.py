import sys
import json
import argparse
from enigma import Enigma

def main():
  print("""\n
  \t███████╗███╗   ██╗██╗ ██████╗ ███╗   ███╗ █████╗ 
	██╔════╝████╗  ██║██║██╔════╝ ████╗ ████║██╔══██╗
	█████╗  ██╔██╗ ██║██║██║  ███╗██╔████╔██║███████║
	██╔══╝  ██║╚██╗██║██║██║   ██║██║╚██╔╝██║██╔══██║
	███████╗██║ ╚████║██║╚██████╔╝██║ ╚═╝ ██║██║  ██║
	╚══════╝╚═╝  ╚═══╝╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝
  """)

  args = get_args()
  config = read_config_file(args.config)

  enigma = Enigma(
    config['rotors'],
    config['alphabet'],
    config['reflector'],
    config['plugboard']
  )

  user_input = ""
  output = ""
  while True:
    print("------------")
    user_input = input("plaintext  > ")
    try:
      output = enigma.encrypt(user_input)
    except Exception as e:
      print(e)
      enigma = Enigma(
        config['rotors'],
        config['alphabet'],
        config['reflector'],
        config['plugboard']
      )
      continue
    print("output  > " + output)
    print("------------")

def read_config_file(path):
  with open(path, 'r') as myfile:
    data = myfile.read()
  config = json.loads(data)
  return config

def get_args():
    parser = argparse.ArgumentParser(description='')
    parser.add_argument('--config', type=str, default='',
                        help='Specify your enigma configuration file')
    args = parser.parse_args()

    if not args.config:
        parser.print_help()
        exit(1)
    return args

if __name__ == '__main__':
  main()