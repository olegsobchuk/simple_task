# You are given two sorted arrays, a and b, where a has a large enough size buffer at the end
# to hold b (which can be spaces, zeroes, or nulls). Write a function to merge b into a in
# sorted order.

#     Example:

#     let a = [1, 3, 5, 0, 0, 0];
#     let b = [2, 4, 6];

#     > merge(a, b)
#     > [1, 2, 3, 4, 5, 6]

BUFFER_MARKERS = ['', ' ', 0, nil]

def merge(prior, secondary)
  return [] if prior.empty?

  prior_size = prior.size - 1
  idx = prior.index { |el| BUFFER_MARKERS.include?(el) }
  return [] unless idx

  buffer_size = prior_size - idx
  (prior[0...idx] + secondary[0..buffer_size]).sort
end

puts merge([1, 3, 5, 0, 0, 0], [2, 4, 6]).inspect
