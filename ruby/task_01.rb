# Find the last non-repeating character in a given string. If all characters repeat, return an empty string.

# Example:

# ```
# > nonRepeat('candy canes do taste yummy')
# > 'u'

# ```

def non_repeat(str)
  counted = str.split('').tally
  counted.keep_if { |_, v| v == 1 }.to_a.last&.first || ''
end

puts non_repeat('candy canes do taste yummy')
