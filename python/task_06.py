"""
Given an array of order objects for a restaurant, each with a table number and a list of ordered items, write a function that returns an object mapping each table number to a summary of how many times each item was ordered at that table. Extra credit: Could you go so far as to make this a restaurant management game?

Example:

orders = [
  { "table": 1, "items": ["burger", "fries"] },
  { "table": 2, "items": ["burger", "burger", "fries"] },
  { "table": 1, "items": ["salad"] },
  { "table": 2, "items": ["fries"] }
];

> order_summary(orders)
{
  1: { "burger": 1, "fries": 1, "salad": 1 },
  2: { "burger": 2, "fries": 2 }
}
// or, string output format:
"Table 1 ordered 1 burger, 1 fries, and 1 salad. Table 2 ordered 2 burgers and 2 fries."
"""

orders = [
    {"table": 1, "items": ["burger", "fries"]},
    {"table": 2, "items": ["burger", "burger", "fries"]},
    {"table": 1, "items": ["salad"]},
    {"table": 2, "items": ["fries"]},
]


def order_summary(orders):
    pass
    tables = set(map(lambda el: el["table"], orders))
    res = {}
    for order in orders:
        items = order["items"]
        if not order["table"] in res:
            res[order["table"]] = items
        else:
            res[order["table"]] += items
    transformer = lambda items: dict((x, items.count(x)) for x in set(items))
    return {table: transformer(items) for table, items in res.items()}


print(order_summary(orders))  # =>
"""
    {
        1: { "burger": 1, "fries": 1, "salad": 1 },
        2: { "burger": 2, "fries": 2 }
    }

"""

"""

"""
