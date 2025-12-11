# Turn an array of integers into a nested array. You can think of this as the opposite of
# flattening an array!

# Examples:

# nestArray([1, 2, 3, 4])
# > [1, [2, [3, [4]]]]

# nestArray([1])
# > [1]


def nest_array(arr):
    nested = []
    for indx in range(len(arr) - 1, -1, -1):
        if indx == len(arr) - 1:
            nested = [arr[indx]]
        else:
            nested = [arr[indx], nested]
    return nested


print(nest_array([1, 2, 3, 4]))  # [1, [2, [3, [4]]]]
