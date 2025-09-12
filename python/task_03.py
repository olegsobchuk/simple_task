durations = [120, 90, 60, 150, 80]
max_duration = 200

# def group_durations(durations, max_duration):
#     if not durations:
#         return []

#     durations.sort()
#     res = []
#     idxs = list(range(len(durations)))
#     idxs.reverse()
#     rm = []
#     while len(idxs) != len(rm) :
#         group = []
#         for idx in idxs:
#             if idx in rm: continue

#             group_sum = sum(group) + durations[idx]
#             if group_sum < max_duration:
#                 group.append(durations[idx])
#                 rm.append(idx)
#             elif group_sum == max_duration:
#                 group.append(durations[idx])
#                 rm.append(idx)
#                 break
#         res.append(group)
#     return res

# def group_durations(durations, max_duration):
#     if not durations:
#         return []

#     durations = sorted(durations, reverse=True)  # Sort in descending order
#     groups = []
#     current_group = []

#     for duration in durations:
#         if sum(current_group) + duration <= max_duration:
#             current_group.append(duration)
#         else:
#             if current_group:  # Only append non-empty groups
#                 groups.append(current_group)
#             current_group = [duration]

#     if current_group:  # Don't forget the last group
#         groups.append(current_group)

#     return groups

def group_durations(durations, max_duration):
    if not durations:
        return []

    durations = sorted(durations, reverse=True)
    groups = []
    used = [False] * len(durations)

    # Try to fill each group as close to max_duration as possible
    while False in used:
        current_group = []
        current_sum = 0

        # Try to fit each unused duration into current group
        for i, duration in enumerate(durations):
            if not used[i] and current_sum + duration <= max_duration:
                current_sum += duration
                current_group.append(duration)
                used[i] = True

        if current_group:
            groups.append(current_group)

    return groups

res = group_durations(durations, 150)
print(res)
