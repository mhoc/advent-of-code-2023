import std/nre
import strutils

type GamePull = object 
  red: int
  green: int
  blue: int

type
  Game = object
    id: int
    pulls: seq[GamePull]

var games = newSeq[Game]()

let entireFile = readFile("data.txt")
for line in splitLines(entireFile):
  let splitColon = split(line, ":")
  let gameId = parseInt(find(splitColon[0], re"Game (\d+)").get.captures[0])
  let pulls = split(splitColon[1], ";")
  var game = Game(
    id: gameId, 
    pulls: newSeq[GamePull](),
  )
  for pull in pulls:
    var gamepull = GamePull()
    for colorPull in split(pull, ","):
      let captures = find(colorPull, re"(\d+) (red|green|blue)").get.captures
      let count = captures[0]
      let color = captures[1]
      if color == "red":
        gamepull.red = parseInt(count)
      elif color == "green":
        gamepull.green = parseInt(count)
      elif color == "blue":
        gamepull.blue = parseInt(count)
    game.pulls.add(gamepull)
  games.add(game)

let allowedRed = 12
let allowedGreen = 13
let allowedBlue = 14

var gameIdSum = 0

for game in games:
  var gameAllowed = true
  for pull in game.pulls:
    if pull.red > allowedRed:
      gameAllowed = false
    if pull.green > allowedGreen:
      gameAllowed = false 
    if pull.blue > allowedBlue:
      gameALlowed = false
  if gameAllowed:
    gameIdSum += game.id

echo "Part 1: " & intToStr(gameIdSum)

var totalPower = 0
for game in games:
  var maxRed = 0
  var maxBlue = 0
  var maxGreen = 0
  for pull in game.pulls:
    if pull.red > maxRed:
      maxRed = pull.red
    if pull.green > maxGreen:
      maxGreen = pull.green
    if pull.blue > maxBlue:
      maxBlue = pull.blue  
  totalPower += (maxRed * maxBlue * maxGreen)

echo "Part 2: " & intToStr(totalPower)