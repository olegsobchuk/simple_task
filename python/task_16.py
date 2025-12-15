# You're building a tool that tracks component edits and groups them into a changelog.
# Given an array of edit actions, each with a timestamp and a component name, return an array
# of grouped changelog entries. Edits to the same component within a 10-minute window should be
# merged into one changelog entry, showing the component name and the range of timestamps affected.

# Example:

# const edits = [
#   { timestamp: "2025-10-06T08:00:00Z", component: "Header" },
#   { timestamp: "2025-10-06T08:05:00Z", component: "Header" },
#   { timestamp: "2025-10-06T08:20:00Z", component: "Header" },
#   { timestamp: "2025-10-06T08:07:00Z", component: "Footer" },
#   { timestamp: "2025-10-06T08:15:00Z", component: "Footer" },
# ];

# > groupChangelogEdits(edits)
# > [
#     {
#         "component": "Footer",
#         "start": "2025-10-06T08:07:00Z",
#         "end": "2025-10-06T08:15:00Z"
#     },
#     {
#         "component": "Header",
#         "start": "2025-10-06T08:00:00Z",
#         "end": "2025-10-06T08:05:00Z"
#     },
#     {
#         "component": "Header",
#         "start": "2025-10-06T08:20:00Z",
#         "end": "2025-10-06T08:20:00Z"
#     }
# ]

from datetime import datetime, timedelta

TIME_FORMAT = "%Y-%m-%dT%H:%M:%SZ"
TIME_DELTA = timedelta(minutes=10)


def group_changelog_edits(components):
    result = []
    for element in components:
        existing = next(
            (
                el
                for el in result
                if el.get("component") == element.get("component")
                and abs(
                    datetime.strptime(el.get("start"), TIME_FORMAT)
                    - datetime.strptime(element.get("timestamp"), TIME_FORMAT)
                )
                <= TIME_DELTA
            ),
            None,
        )
        if not existing:
            result.append(
                {
                    "component": element.get("component"),
                    "start": element.get("timestamp"),
                    "end": element.get("timestamp"),
                }
            )
        else:
            existing["end"] = element.get("timestamp")
    return result


components = [
    {"timestamp": "2025-10-06T08:00:00Z", "component": "Header"},
    {"timestamp": "2025-10-06T08:05:00Z", "component": "Header"},
    {"timestamp": "2025-10-06T08:20:00Z", "component": "Header"},
    {"timestamp": "2025-10-06T08:07:00Z", "component": "Footer"},
    {"timestamp": "2025-10-06T08:15:00Z", "component": "Footer"},
]
print(group_changelog_edits(components))
