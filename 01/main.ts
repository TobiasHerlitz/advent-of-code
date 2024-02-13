const patternOne = /(\d)/g; // Part 1
const patternTwo = /(?=([0-9]|one|two|three|four|five|six|seven|eight|nine))/g; // Part 2
const textToDigit = {
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9
}

function parseToNumber(original: string): number {
  const digit = Number(original);
  if (digit) {
    return digit;
  }

  if (original in textToDigit) {
    return textToDigit[original];
  }

  throw new Error(
    `Could not translate digit representaion into number. Got: ${original}`
  )
}

function getCalibrationSum(digitPattern: RegExp, inputRows: string[]) {
  const calibrationValues = inputRows.map((inputRow) => {
    const inputDigits = [...inputRow.matchAll(digitPattern)];
    const first = inputDigits[0][1]
    const last = inputDigits[inputDigits.length - 1][1];

    return Number([parseToNumber(first), parseToNumber(last)].join(''))
  });

  return calibrationValues.reduce((a, b) => a + b);
}

const file = Bun.file('input');
const text = await file.text();
const inputRows = text.split('\n');

console.log(getCalibrationSum(patternOne, inputRows)); // 55607
console.log(getCalibrationSum(patternTwo, inputRows)); // 55291

export {}
