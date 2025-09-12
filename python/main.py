# You’re assembling a custom mechanical keyboard. Each required part has a
# delivery time in days and an assembly time in hours. You can only assemble
# a part after it arrives, and you can only work on one part at a time.
# Given an array of parts where each part is { name, arrivalDays, assemblyHours },
# return the minimum total hours needed to finish assembling all parts,
# starting from hour 0.

# Example:

# minAssemblyTime([
#   { name: "keycaps", arrivalDays: 1, assemblyHours: 2 },
#   { name: "switches", arrivalDays: 2, assemblyHours: 3 },
#   { name: "stabilizers", arrivalDays: 0, assemblyHours: 1 },
#   { name: "PCB", arrivalDays: 1, assemblyHours: 4 },
#   { name: "case", arrivalDays: 3, assemblyHours: 2 }
# ])

# > 74

import heapq

def minAssemblyTime(parts):
    # Convert arrival days to arrival hours
    for part in parts:
        part["arrivalHours"] = part["arrivalDays"] * 24

    # Sort parts by arrival time
    parts.sort(key=lambda p: p["arrivalHours"])

    current_time = 0
    i = 0
    n = len(parts)
    heap = []  # Min-heap based on assembly time

    while i < n or heap:
        # Push all parts that have arrived by current_time into the heap
        while i < n and parts[i]["arrivalHours"] <= current_time:
            heapq.heappush(heap, (parts[i]["assemblyHours"], parts[i]["name"], parts[i]["arrivalHours"]))
            i += 1

        if heap:
            # Take part with smallest assembly time
            assembly_time, name, arrival_time = heapq.heappop(heap)
            current_time += assembly_time
        else:
            # Wait for the next part to arrive
            if i < n:
                current_time = parts[i]["arrivalHours"]

    return current_time


print(minAssemblyTime([
  { "name": "keycaps", "arrivalDays": 1, "assemblyHours": 2 },
  { "name": "switches", "arrivalDays": 2, "assemblyHours": 3 },
  { "name": "stabilizers", "arrivalDays": 0, "assemblyHours": 1 },
  { "name": "PCB", "arrivalDays": 1, "assemblyHours": 4 },
  { "name": "case", "arrivalDays": 3, "assemblyHours": 2 }
]))
