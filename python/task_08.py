# Given a multi-line string and a sequence of Vim navigation commands (h for left, j for down, k for up, and l for right), and starting at the top-left character, write a function that processes the commands and returns the character under the cursor. If the cursor tries to move out of bounds, keep it at the last valid position.

# Example:

# ```
# const string = `Hello, world!
# how are ya?`; // or "Hello, world!\nhow are ya?"
# const commands = 'jlhll';

# getCharAfterVimCommands(string, commands)
# > 'w'

# ```


def get_char_after_vim_commands(text, cmds):
    positions = {"line": 0, "cursor": 0}
    lines = text.split("\n")
    for cmd in cmds:
        move(cmd, positions)
    line = lines[positions["line"]]
    char = line[positions["cursor"]]

    return char


def move(cmd, positions):
    match cmd:
        case "h":
            if positions["cursor"] > 0:
                positions["cursor"] -= 1
        case "l":
            positions["cursor"] += 1
        case "k":
            if positions["line"] > 0:
                positions["line"] -= 1
        case "j":
            positions["line"] += 1


str = """Hello, world!
how are ya?"""
commands = "jlhll"
print(get_char_after_vim_commands(str, commands))  # "w"
