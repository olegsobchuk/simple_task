# Simple Task

List of tasks:

1. Find the last non-repeating character in a given string. If all characters repeat, return an empty string.

Example:

```
> nonRepeat('candy canes do taste yummy')
> 'u'

```

[Golang](./golang/task_01.go) | [Python](./python/task_01.py) | [Ruby](./ruby/task_01.rb)

2. You’re assembling a custom mechanical keyboard. Each required part has a delivery time in days and an assembly time in hours. You can only assemble a part after it arrives, and you can only work on one part at a time. Given an array of parts where each part is { name, arrivalDays, assemblyHours }, return the minimum total hours needed to finish assembling all parts, starting from hour 0.

Example:

```
minAssemblyTime([
  { name: "keycaps", arrivalDays: 1, assemblyHours: 2 },
  { name: "switches", arrivalDays: 2, assemblyHours: 3 },
  { name: "stabilizers", arrivalDays: 0, assemblyHours: 1 },
  { name: "PCB", arrivalDays: 1, assemblyHours: 4 },
  { name: "case", arrivalDays: 3, assemblyHours: 2 }
])

> 74

```

[Golang TBD](Golang) | [Python](./python/task_02.py) | [Ruby TBD]()

3. Given an array of audio file durations, write a function to group the files into playlists such that each playlist's total duration does not exceed a given limit maxDuration. Return an array of playlists, where each playlist is an array of file durations. Try to minimize the number of playlists.

Example:

```
const files = [120, 90, 60, 150, 80];
const maxDuration = 200;

groupAudioFiles(files, maxDuration)
> [[150], [120,80], [90,60]]

groupAudioFiles(files, 160)
> [[150], [120], [90,60], [80]]

```

[Golang TBD](Golang) | [Python](./python/task_03.py) | [Ruby TBD]()

4. Given an array arr representing the positions of monsters along a straight line, and an integer d representing the minimum safe distance required between any two monsters, write a function to determine if all monsters are at least d units apart. If not, return the smallest distance found between any two monsters. If all monsters are safely spaced, return -1.

Examples:

```
let monsters = [3, 8, 10, 15];
let d = 6;
minMonsterDistance(bees, d)
> 2

minMonsterDistance([5, 9, 14, 18], 4)
> -1

```

[Golang TBD](Golang) | [Python](./python/task_04.py) | [Ruby TBD]()

5. Write a generator function createLaundryItem() that returns an object representing a laundry item. This object should have a method nextCycle() which, when called, advances the item through a series of laundry cycles in order: "soak", "wash", "rinse", "spin", and "dry". After the final cycle, subsequent calls to nextCycle() should return "done".

Example:

```
let towel = createLaundryItem();

console.log(towel.nextCycle()); // "soak"
console.log(towel.nextCycle()); // "wash"
console.log(towel.nextCycle()); // "rinse"
console.log(towel.nextCycle()); // "spin"
console.log(towel.nextCycle()); // "dry"
console.log(towel.nextCycle()); // "done"
console.log(towel.nextCycle()); // "done"

```

[Golang TBD](Golang) | [Python](./python/task_05.py) | [Ruby TBD]()

6. Given an array of order objects for a restaurant, each with a table number and a list of ordered items, write a function that returns an object mapping each table number to a summary of how many times each item was ordered at that table. Extra credit: Could you go so far as to make this a restaurant management game?

Example:

```
const orders = [
  { table: 1, items: ["burger", "fries"] },
  { table: 2, items: ["burger", "burger", "fries"] },
  { table: 1, items: ["salad"] },
  { table: 2, items: ["fries"] }
];

> orderSummary(orders)
{
  1: { burger: 1, fries: 1, salad: 1 },
  2: { burger: 2, fries: 2 }
}
// or, string output format:
"Table 1 ordered 1 burger, 1 fries, and 1 salad. Table 2 ordered 2 burgers and 2 fries."

```

[Golang TBD](Golang) | [Python](./python/task_06.py) | [Ruby TBD]()

7. Given an array of side lengths, write a function to determine they can form a hexagon with three side-length pairs (as in, three pairs of equal sides needed). Return true if possible.

Examples:

```
canFormHexagon([2, 3, 8, 8, 2, 3])
> true

canFormHexagon([1, 2, 3, 4, 5, 6])
> false

canFormHexagon([2, 2, 2, 2, 2, 2, 2])
> false

```

[Golang TBD](Golang) | [Python](./python/task_07.py) | [Ruby TBD]()

8. Given a multi-line string and a sequence of Vim navigation commands (h for left, j for down, k for up, and l for right), and starting at the top-left character, write a function that processes the commands and returns the character under the cursor. If the cursor tries to move out of bounds, keep it at the last valid position.

Example:

```
const string = `Hello, world!
how are ya?`; // or "Hello, world!\nhow are ya?"
const commands = 'jlhll';

getCharAfterVimCommands(string, commands)
> 'w'

```

[Golang TBD](Golang) | [Python TBD]() | [Ruby TBD]()

9. Given an array of strings representing a sequence of traffic light states ("red" for stop, "green" for go, "yellow" for slow), write a function that returns true if the sequence could represent a valid state machine for a standard traffic light. The only valid transitions are: red to green, green to yellow, and yellow to red.

Example:

```
> isValidTrafficSequence(["red", "green", "yellow", "red", "green"])
> true

> isValidTrafficSequence(["red", "yellow", "green"]);
> false

> isValidTrafficSequence([])
> true

```

[Golang TBD](Golang) | [Python TBD]() | [Ruby TBD]()

10. Imagine a simplified version of the game Battleship played on a 2D grid. The grid represents the sea, and each cell can either be empty (.) or contain a part of a ship (X). Ships are placed horizontally or vertically, and there are no adjacent ships. Given a grid, count the number of battleships in it. Extra credit: can you make a layout generator for the game given these rules?

Example:

```
const ships = [
  ['X', 'X', '.', 'X'],
  ['.', '.', '.', 'X'],
  ['.', '.', '.', 'X'],
  ['.', '.', '.', '.'],
];

numberOfShips(ships)
> 2

```

[Golang TBD](Golang) | [Python TBD]() | [Ruby TBD]()

11. For an array of numbers, generate an array where for every element, all neighboring elements are added to itself, and return the sum of that array.

Examples:

```
[]               -> 0
[1]              -> 1
[1, 4]           -> 10 // (1+4 + 4+1)
[1, 4, 7]        -> 28 // (1+4 + 4+1+7 + 7+4)
[1, 4, 7, 10]    -> 55
[-1, -2, -3]     -> -14
[0.1, 0.2, 0.3]  -> 1.4
[1,-20,300,-4000,50000,-600000,7000000] -> 12338842

```

[Golang TBD](Golang) | [Python TBD]() | [Ruby TBD]()

12. Given an array of strings representing the names of monarchs and their ordinal numbers, write a function that returns the list of names sorted first by name and then by their ordinal value (in Roman numerals), in ascending order.

Example:

```
> sortMonarchs(["Louis IX", "Louis VIII", "Philip II", "Philip I"])
> ["Louis VIII", "Louis IX", "Philip I", "Philip II"]

> sortMonarchs(["George VI", "George V", "Elizabeth II", "Edward VIII"])
> ["Edward VIII", "Elizabeth II", "George V", "George VI"]

```

[Golang TBD](Golang) | [Python TBD]() | [Ruby TBD]()

13. Given an array of fireworks representing a series going off, write a function to find the "grand finale" of the show! A grand finale is defined as the longest subarray where the average size is at least 5, the minimum velocity is 3, and the difference between the min and max height is no more than 10. Return the starting index of the grand finale.

Example:

```
const fireworks = [
  {height: 10, size: 6, velocity: 4},
  {height: 13, size: 3, velocity: 2},
  {height: 17, size: 6, velocity: 3},
  {height: 21, size: 8, velocity: 4},
  {height: 19, size: 5, velocity: 3},
  {height: 18, size: 4, velocity: 4}
];

> grandFinaleStart(fireworks)
> 2

```

[Golang TBD](Golang) | [Python TBD]() | [Ruby TBD]()

14. Turn an array of integers into a nested array. You can think of this as the opposite of flattening an array!

Examples:

```
nestArray([1, 2, 3, 4])
> [1, [2, [3, [4]]]]

nestArray([1])
> [1]

```

[Golang TBD](Golang) | [Python TBD]() | [Ruby TBD]()

15. You are given an array of arrays, where each inner array represents the runs scored by each team in an inning of a baseball game: [[home1, away1], [home2, away2], ...]. Write a function that returns an object with the total runs for each team, which innings each team led, and who won the game.


Example:

```
const innings = [[1, 0], [2, 2], [0, 3], [4, 1]];

> analyzeBaseballGame(innings)
> {
    homeTotal: 7,
    awayTotal: 6,
    homeLedInnings: [1, 2, 4],
    awayLedInnings: [3],
    winner: "home"
  }

```

16.

```

```

17.

```

```
