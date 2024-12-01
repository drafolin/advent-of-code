import { log } from "console";
import { readFileSync } from "fs";

const input = readFileSync(__dirname + "/data.txt", "utf-8").split("\n");
const cols: string[][] = input.map(v => v.split(""));

const isTouching = (line: number, start: number, end: number): boolean => {
	//check previous line
	if (line - 1 > 0) {
		for (let i = start - 1; i <= end + 1; ++i) {
			if (cols[line - 1][i] === undefined)
				continue;
			if (!/[\d\.]/.test(cols[line - 1][i])) {
				return true;
			}
		}
	}

	// check this line
	for (let i = start - 1; i <= end + 1; ++i) {
		if (cols[line][i] === undefined)
			continue;
		if (!/[\d\.]/.test(cols[line][i])) {
			return true;
		}
	}

	// check next line
	//check previous line
	if (line + 1 < cols.length) {
		for (let i = start - 1; i <= end + 1; ++i) {
			if (cols[line + 1][i] === undefined)
				continue;
			if (!/[\d\.]/.test(cols[line + 1][i])) {
				return true;
			}
		}
	}

	return false;
};

const boundNumber = (line: number, col: number): { start: number, end: number, val: number; } => {
	let start = col;
	while (cols[line][start] !== undefined &&
		/\d/.test(cols[line][start])) {
		start--;
	}

	start++;

	let end = col;
	while (cols[line][start] !== undefined &&
		/\d/.test(cols[line][end])) {
		end++;
	}

	end--;

	let value = "";

	for (let i = start; i <= end; ++i) {
		value += cols[line][i];
	}
	return { start, end, val: parseInt(value) };
};

const gearRatio = (line: number, col: number): number => {
	let firstNumber: {
		val: number,
		start: number,
		line: number,
		end: number;
	} | undefined = undefined;
	let secondNumber: {
		val: number,
		start: number,
		line: number,
		end: number;
	} | undefined | undefined = undefined;

	//check previous line
	if (line - 1 >= 0) {
		for (let i = col - 1; i <= col + 1; ++i) {
			if (cols[line - 1][i] === undefined)
				continue;
			if (/\d/.test(cols[line - 1][i])) {
				if (!secondNumber) {
					console.log(firstNumber);
					if (!firstNumber) {
						firstNumber = { ...boundNumber(line - 1, i), line: line - 1 };
						continue;
					} else if (firstNumber.line !== line - 1 || (firstNumber.start > i || i > firstNumber.end)) {
						secondNumber = { ...boundNumber(line - 1, i), line: line - 1 };
						continue;
					} else {
						continue;
					}
				} else {
					if ((line + 1 === firstNumber?.line &&
						firstNumber.start <= i && i <= firstNumber.end) ||
						(line + 1 === secondNumber?.line &&
							secondNumber.start <= i && i <= secondNumber.end))
						continue;
					return 0;
				}
			}
		}
	};

	// check this line
	for (let i = col - 1; i <= col + 1; ++i) {
		if (cols[line][i] === undefined)
			continue;
		if (/\d/.test(cols[line][i])) {
			if (!secondNumber) {
				if (!firstNumber) {
					firstNumber = { ...boundNumber(line, i), line: line };
					continue;
				} else if (firstNumber.line !== line || firstNumber.start > i || i > firstNumber.end) {
					secondNumber = { ...boundNumber(line, i), line: line };
					continue;
				} else {
					continue;
				}
			} else {
				if ((line + 1 === firstNumber?.line &&
					firstNumber.start <= i && i <= firstNumber.end) ||
					(line + 1 === secondNumber?.line &&
						secondNumber.start <= i && i <= secondNumber.end))
					continue;
				return 0;
			}
		}
	}

	// check next line
	if (line + 1 < cols.length) {
		for (let i = col - 1; i <= col + 1; ++i) {
			if (cols[line + 1][i] === undefined)
				continue;
			if (/\d/.test(cols[line + 1][i])) {
				if (!secondNumber) {
					if (!firstNumber) {
						firstNumber = { ...boundNumber(line + 1, i), line: line + 1 };
						continue;
					} else if (firstNumber.line !== line + 1 || firstNumber.start > i || i > firstNumber.end) {
						secondNumber = { ...boundNumber(line + 1, i), line: line + 1 };
						continue;
					} else {
						continue;
					}
				} else {
					if ((line + 1 === firstNumber?.line &&
						firstNumber.start <= i && i <= firstNumber.end) ||
						(line + 1 === secondNumber?.line &&
							secondNumber.start <= i && i <= secondNumber.end))
						continue;
					return 0;
				}
			}
		}
	}


	console.log(line, col);
	console.log(firstNumber);
	console.log(secondNumber);
	return !firstNumber || !secondNumber ? 0 : firstNumber.val * secondNumber.val;
};

let sum = 0;

cols.forEach((_, i) => {
	let isInNumber = false;
	let wasInNumber = false;
	let currentNumber = "";
	cols[i].forEach((_, j) => {
		if (cols[i][j] <= "9" && cols[i][j] >= "0") {
			wasInNumber = true;
			isInNumber = true;
			currentNumber += cols[i][j];
		} else {
			isInNumber = false;
		}

		if (!isInNumber && wasInNumber) {
			wasInNumber = false;

			// check for adjascent
			let end = j - 1;
			let start = j - currentNumber.length;
			if (isTouching(i, start, end)) {
				sum += parseInt(currentNumber);
			}
			log(currentNumber);
			currentNumber = "";
		}
	});
	if (wasInNumber) {
		wasInNumber = false;

		// check for adjascent
		let j = cols[i].length;
		let end = j - 1;
		let start = j - currentNumber.length;
		if (isTouching(i, start, end)) {
			sum += parseInt(currentNumber);
		}
		log(currentNumber);
		currentNumber = "";
	}
});

sum;

let gearSum = 0;

cols.forEach((_, i) => {
	cols[i].forEach((_, j) => {
		if (cols[i][j] === "*") {
			gearSum += gearRatio(i, j);
		}
	});
});

gearSum;
