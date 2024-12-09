const parseInput = (text: string) => {
  const leftList: number[] = [];
  const rightList: number[] = [];

  const inputRows = text.split('\n');
  inputRows.forEach((inputRow) => {
    const inputLocations = inputRow.split('   ');
    if (inputRow === "") {
      return;
    }
    leftList.push(Number(inputLocations[0]));
    rightList.push(Number(inputLocations[1]));
  })

  if (leftList.length !== rightList.length) {
    throw new Error('Lists are not of equal length')
  }
  return [leftList, rightList];
}

const calculateTotalDistance = (leftList: number[], rightList: number[]) => {
  leftList.sort();
  rightList.sort();

  let totalDistance = 0;
  for (let i = 0; i < rightList.length; i++) {
    const distance = Math.abs(leftList[i] - rightList[i])
    totalDistance += distance;
  }

  return totalDistance;
}

const calculateSimilarityScore = (leftList: number[], rightList: number[]) => {
  let similarityScore = 0;
  leftList.forEach((leftListItem) => {
    rightList.forEach((rightListItem) => {
      if (leftListItem !== rightListItem) {
        return;
      };
      similarityScore += leftListItem;
    })
  })

  return similarityScore
}

const file = Bun.file('input');
const text = await file.text();

const [leftList, rightList] = parseInput(text)

console.log(`Part 1 - Total distance is: ${calculateTotalDistance(leftList, rightList)}`)
console.log(`Part 2 - Similarity score is: ${calculateSimilarityScore(leftList, rightList)}`)

export {}
