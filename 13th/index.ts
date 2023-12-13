import { readFileSync } from "fs";

const testing = false;

const dataTxt = (testing ? `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
` : readFileSync(__dirname + "/data.txt", "utf-8")).trim().split("\n\n");

const findMirror = (arr: string[]): number | null => {
	for (let lineI = 1; lineI < arr.length; ++lineI) {
		if (arr[lineI - 1] === arr[lineI]) {
			let isValid = true;
			if (lineI > arr.length / 2) {
				for (let i = 0; i < arr.length - lineI; ++i) {
					if (arr[lineI + i] !== arr[lineI - i - 1])
						isValid = false;
				}
			} else {
				for (let i = 0; i < lineI; ++i) {
					if (arr[lineI + i] !== arr[lineI - i - 1])
						isValid = false;
				}
			}
			if (isValid)
				return lineI;
		}
	}
	return null;
};

let count = 0;
for (let pattern of dataTxt) {
	const lines = pattern.split("\n");
	const lineRes = findMirror(lines);

	if (lineRes) {
		count += lineRes * 100;
		continue;
	}

	const columns: string[] = [];

	for (let x = 0; x < lines[0].length; ++x) {
		lines.forEach((_, y) => {
			if (columns[x] === undefined)
				columns[x] = "";
			columns[x] += lines[y][x];
		});
	}

	count += findMirror(columns)!;
}

count;
