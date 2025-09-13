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

def min_assembly_time(parts):
  collector = {}
  for part in parts:
    if part['arrivalDays'] not in collector:
      collector[part['arrivalDays']] = 0
    collector[part['arrivalDays']] += part['assemblyHours']
    last_day = max(collector.keys())
    hrs = 24 * last_day + collector[last_day]
  return(hrs)



res = min_assembly_time([
  {"name": "keycaps", "arrivalDays": 1, "assemblyHours": 2},
  {"name": "switches", "arrivalDays": 2, "assemblyHours": 3},
  {"name": "stabilizers", "arrivalDays": 0, "assemblyHours": 1},
  {"name": "PCB", "arrivalDays": 1, "assemblyHours": 4},
  {"name": "case", "arrivalDays": 3, "assemblyHours": 2}
]) # Expected output: 74

print(res)
