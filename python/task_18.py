# Given a string str and an array of positive integers widths, write a function that splits the
# string into lines, each with the exact number of characters as specified by the corresponding width.
# Return an array of the substrings. Use the last width for any remaining characters if the array is
# shorter than needed.

# Example:

# const str = "Supercalifragilisticexpialidocious";
# const widths = [5, 9, 4];

# > splitByWidths(str, widths);
# > ['Super', 'califragi', 'list', 'icex', 'pial', 'idoc', 'ious']


def split_by_widths(str, numbers):
    str_length = len(str)
    result = []
    start_indx = 0
    for number in numbers:
        result.append(str[start_indx : start_indx + number])
        start_indx = start_indx + number

    while start_indx < str_length:
        result.append(str[start_indx : start_indx + numbers[-1]])
        start_indx += numbers[-1]

    return result


print(split_by_widths("Supercalifragilisticexpialidocious", [5, 9, 4]))
# ['Super', 'califragi', 'list', 'icex', 'pial', 'idoc', 'ious']
