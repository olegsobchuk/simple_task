"""
Write a generator function createLaundryItem() that returns an object
representing a laundry item. This object should have a method nextCycle()
which, when called, advances the item through a series of laundry cycles
in order: "soak", "wash", "rinse", "spin", and "dry". After the final cycle,
subsequent calls to nextCycle() should return "done".
"""


class Laundry:
    CYCLES = ("init", "soak", "wash", "rinse", "spin", "dry", "done")

    def __init__(self):
        self.state = 0
        self.last_state = len(self.CYCLES) - 1

    def in_last_state(self):
        return self.state == self.last_state

    def next_cycle(self):
        if self.in_last_state():
            return self.CYCLES[self.state]

        self.state += 1
        return self.CYCLES[self.state]


def create_laundry_item():
    return Laundry()


towel = create_laundry_item()

print(towel.next_cycle())  # => "soak"
print(towel.next_cycle())  # => "wash"
print(towel.next_cycle())  # => "rinse"
print(towel.next_cycle())  # => "spin"
print(towel.next_cycle())  # => "dry"
print(towel.next_cycle())  # => "done"
print(towel.next_cycle())  # => "done"
