type Report = number[];

const parseInput = (text: string) => {
  const reports: Report[] = [];

  text.split('\n').forEach((inputRow) => {
    if (inputRow === "") {
      return;
    }

    const report = inputRow.split(' ').map((level) => Number(level));
    reports.push(report);
  })

  return reports;
}

const validateReport = (report: Report, canFailOnce: boolean) => {
  let shouldIncrement: boolean | undefined;

  for (let i = 1; i < report.length; i++) {
    const diff = report[i] - report[i - 1];
    if (Math.abs(diff) === 0 || Math.abs(diff) > 3) {
      if (canFailOnce) {
        canFailOnce = false
        continue;
      }
      return false;
    }
    if (shouldIncrement === true && diff < 0) {
      if (canFailOnce) {
        canFailOnce = false
        continue;
      }
      return false;
    }
    if (shouldIncrement === false && diff > 0) {
      if (canFailOnce) {
        canFailOnce = false
        continue;
      }
      return false;
    }

    if (shouldIncrement === undefined) {
      shouldIncrement = diff > 0;
    }
  }

  return true;
}

const countSafeReports = (reports: Report[], problemDampener: boolean = false) => {
  return reports.filter((report) => validateReport(report, problemDampener)).length
}

const file = Bun.file('input');
const text = await file.text();

const reports = parseInput(text);
console.log(`Part 1 - Number of safe reports: ${countSafeReports(reports)}`)
console.log(`Part 2 - Number of safe reports with problem dampener: ${countSafeReports(reports, true)}`) // 687

export {};
