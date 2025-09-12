data1 = [2, 3, 8, 8, 2, 3]
data2 = [1, 2, 3, 4, 5, 6]
data3 = [2, 2, 2, 2, 2, 2, 2]


def can_form_hexahon(lst):
    if not len(lst) == 6:
        return False
    lst.sort()
    for i in range(0, 6, 2):
        if not lst[i] == lst[i + 1]:
            return False
    return True


print(can_form_hexahon(data1))
print(can_form_hexahon(data2))
print(can_form_hexahon(data3))
