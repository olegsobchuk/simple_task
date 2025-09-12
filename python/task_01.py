# Find the last non-repeating character in a given string. If all characters repeat, return an empty string.

# Example:

# > nonRepeat('candy canes do taste yummy')
# > 'u'

def non_repeat(str):
  counter = dict()
  char_list = list(str)
  for char in char_list:
    if char not in counter: counter[char] = 1
    else: counter[char] += 1

  single_chars = [k for k, v in counter.items() if v == 1]
  return(single_chars[-1] if single_chars else '')

print(non_repeat("candy canes do taste yummy"))
