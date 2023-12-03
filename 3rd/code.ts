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
