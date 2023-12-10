import re

data_file = open("./data.txt", "r")
data = data_file.read()
data_file.close()
data_lines = data.split("\n")

line_pattern = r'.*: ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) \| ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d) ([0-9 ]\d)'

class Card:
  @staticmethod
  def from_str(st):
    match = re.search(line_pattern, st)
    return Card(
      set((
        match.group(1),
        match.group(2),
        match.group(3),
        match.group(4),
        match.group(5),
        match.group(6),
        match.group(7),
        match.group(8),
        match.group(9),
        match.group(10),
      )),
      set((
        match.group(11),
        match.group(12),
        match.group(13),
        match.group(14),
        match.group(15),
        match.group(16),
        match.group(17),
        match.group(18),
        match.group(19),
        match.group(20),
        match.group(21),
        match.group(22),
        match.group(23),
        match.group(24),
        match.group(25),
        match.group(26),
        match.group(27),
        match.group(28),
        match.group(29),
        match.group(30),
        match.group(31),
        match.group(32),
        match.group(33),
        match.group(34),
        match.group(35),
      )),
    )

  def __init__(self, winning, have):
    self.winning = winning
    self.have = have

  def __str__(self):
    return f"{self.id} {self.winning} {self.have}"
  
  def winning_numbers(self):
    r = set()
    for have_number in self.have:
      if have_number in self.winning:
        r.add(have_number)
    return r

all_cards = []
for line in data_lines:
  all_cards.append(Card.from_str(line))

total = 0
for card in all_cards:
  winning_numbers = card.winning_numbers()
  if len(winning_numbers) > 0:
    total += 2 ** (len(winning_numbers) - 1)
print(f"Part 1: {total}")

resolved = []
for card in all_cards:
  resolved.append([card])
for card_number, card_set in enumerate(resolved):
  for card in card_set:
    winning_numbers = len(card.winning_numbers())
    for clone_card_number in list(range(card_number + 1, min(card_number + 1 + winning_numbers, len(resolved)))):
      resolved[clone_card_number].append(resolved[clone_card_number][0])

total_cards = 0
for card_set in resolved:
  total_cards += len(card_set)

print(f"Part 2: {total_cards}")
