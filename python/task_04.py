"""
Given an array arr representing the positions of monsters along a straight line,
and an integer d representing the minimum safe distance required between any two
monsters, write a function to determine if all monsters are at least d units apart.
If not, return the smallest distance found between any two monsters. If all monsters
are safely spaced, return -1.
"""

monsters = [3, 8, 10, 15]
d = 6


def min_monster_distance(points, dist):
    min_dist = -1
    max_position = len(points)

    if max_position <= 1:
        min_dist

    for n in range(max_position):
        if n == (max_position - 1):
            continue

        points_diff = points[n + 1] - points[n]
        if points_diff < dist:
            if min_dist == -1 or points_diff < min_dist:
                min_dist = points_diff

    return min_dist


print(min_monster_distance(monsters, d))  # => 2
print(min_monster_distance([5, 9, 14, 18], 4))  # => -1
