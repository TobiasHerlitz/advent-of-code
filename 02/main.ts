type Turn = {
  green: number;
  red: number;
  blue: number;
}

type Game = {
  id: string;
  turns: Turn[];
}

const RED_CUBE_MAX = 12;
const GREEN_CUBE_MAX = 13;
const BLUE_CUBE_MAX = 14;

// Builds a turn object from string with the shape " 1 green, 4 blue, 13 red"
function buildTurn(turnInput: string): Turn {
  const turn = { green: 0, red: 0, blue: 0 };

  turnInput.split(',').forEach((cubeQuantityInput) => {
    const [qty, color] = cubeQuantityInput.trim().split(' ');
    turn[color] = Number(qty);
  });

  return turn;
}

// Builds a game object from string with the shape "Game 3: 5 green, 2 blue, 8 red; 1 green, 4 blue, 13 red"
function buildGame(gameInput: string): Game {
  const [metaInput, turnsInput] = gameInput.split(':');

  return {
    id: metaInput.split(' ')[1],
    turns: turnsInput.split(';').map(buildTurn)
  };
}

function validateTurn(turn: Turn) {
  if (turn.green > GREEN_CUBE_MAX) return false;
  if (turn.red > RED_CUBE_MAX) return false;
  if (turn.blue > BLUE_CUBE_MAX) return false;

  return true
}

/*
Calculate the sum of all game ID:s that are playable
with a set number of colored cubes
*/
function partOne(games: Game[]) {
  let sumOfIds = 0;
  games.forEach((game) => {
    if (game.turns.every(validateTurn)) {
      sumOfIds += Number(game.id);
    };
  })

  return sumOfIds;
}

/*
Find minimum number of different colored cubes for each game.
Then multiply the three quantities per game and sum for all games
*/
function partTwo(games: Game[]) {
  let sumOfProducts = 0;

  games.forEach((game) => {
    let green = 0;
    let red = 0;
    let blue = 0;

    game.turns.forEach((turn) => {
      if (turn.green > green) {
        green = turn.green;
      }

      if (turn.red > red) {
        red = turn.red;
      }

      if (turn.blue > blue) {
        blue = turn.blue;
      }
    })

    sumOfProducts += green * red * blue;
  })

  return sumOfProducts;
}

const file = Bun.file('input');
const text = await file.text();
const inputRows = text.split('\n');

const games: Game[] = inputRows.map(buildGame)

console.log(partOne(games)); // 2278
console.log(partTwo(games)); // 67953

export {}