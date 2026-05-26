# Given he current position of a knight as [row, col] in an 8x8 chess board represented as a 2D
#     array, write a function to return all valid moves the knight can make. Extra credit: Do this for
#     every chess piece!

#     Example:

#     knightMoves([4, 4])
#     > [[2, 3], [2, 5], [3, 2], [3, 6], [5, 2], [5, 6], [6, 3], [6, 5]]

#     knightMoves([0, 0])
#     > [[1, 2], [2, 1]]

#     knightMoves([1, 2])
#     > [[0, 0], [0, 4], [2, 0], [2, 4], [3, 1], [3, 3]]

def knight_moves(position)
  moves = [[-2, -1], [-2, 1], [-1, -2], [-1, 2], [1, -2], [1, 2], [2, -1], [2, 1]]

  res = []
  moves.each do |move|
    action = move.zip(position).map(&:sum)
    res.append(action) unless action.any?(&:negative?)
  end

  res
end

puts(knight_moves([4, 4]).inspect) # > [[2, 3], [2, 5], [3, 2], [3, 6], [5, 2], [5, 6], [6, 3], [6, 5]]
puts(knight_moves([0, 0]).inspect) # > [[1, 2], [2, 1]]
puts(knight_moves([1, 2]).inspect) # > [[0, 0], [0, 4], [2, 0], [2, 4], [3, 1], [3, 3]]
