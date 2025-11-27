# For an array of numbers, generate an array where for every element, all neighboring
# elements are added to itself, and return the sum of that array.

# Examples:

# ```text
# []               -> 0
# [1]              -> 1
# [1, 4]           -> 10 // (1+4 + 4+1)
# [1, 4, 7]        -> 28 // (1+4 + 4+1+7 + 7+4)
# [1, 4, 7, 10]    -> 55
# [-1, -2, -3]     -> -14
# [0.1, 0.2, 0.3]  -> 1.4
# [1,-20,300,-4000,50000,-600000,7000000] -> 12338842

from decimal import Decimal


def neighboring_sum(arry: list) -> Decimal:
    d_arry: list[Decimal] = [Decimal(str(x)) for x in arry]
    res: Decimal = Decimal("0.0")
    length: int = len(d_arry)
    for idx, num in enumerate(d_arry):
        sum: int = num
        if idx > 0:
            sum += d_arry[idx - 1]
        if idx < length - 1:
            sum += d_arry[idx + 1]
        res += sum
    return res


print(neighboring_sum([]))  # 0
print(neighboring_sum([1]))  # 1
print(neighboring_sum([1, 4]))  # 10 // (1+4 + 4+1)
print(neighboring_sum([1, 4, 7]))  # 28 // (1+4 + 4+1+7 + 7+4)
print(neighboring_sum([1, 4, 7, 10]))  # 55
print(neighboring_sum([-1, -2, -3]))  # -14
print(neighboring_sum([0.1, 0.2, 0.3]))  # 1.4
print(neighboring_sum([1, -20, 300, -4000, 50000, -600000, 7000000]))  # 12338842
