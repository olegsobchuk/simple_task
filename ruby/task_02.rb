#   You’re assembling a custom mechanical keyboard. Each required part has a delivery time in days and an assembly time in hours. You can only assemble a part after it arrives, and you can only work on one part at a time. Given an array of parts where each part is { name, arrivalDays, assemblyHours }, return the minimum total hours needed to finish assembling all parts, starting from hour 0.

# Example:

# ```
# minAssemblyTime([
#   { name: "keycaps", arrival_days: 1, assembly_hours: 2 },
#   { name: "switches", arrival_days: 2, assembly_hours: 3 },
#   { name: "stabilizers", arrival_days: 0, assembly_hours: 1 },
#   { name: "PCB", arrival_days: 1, assembly_hours: 4 },
#   { name: "case", arrival_days: 3, assembly_hours: 2 }
# ])

# > 74

# ```

def min_assembly_time(parts)
  rtturn 0 if parts.empty?

  parts.sort_by! { |part| part[:arrival_days] }
  parts.last[:arrival_days] * 24 + parts.last[:assembly_hours]
end

parts = [
  { name: 'keycaps', arrival_days: 1, assembly_hours: 2 },
  { name: 'switches', arrival_days: 2, assembly_hours: 3 },
  { name: 'stabilizers', arrival_days: 0, assembly_hours: 1 },
  { name: 'PCB', arrival_days: 1, assembly_hours: 4 },
  { name: 'case', arrival_days: 3, assembly_hours: 2 }
]

puts min_assembly_time(parts)
