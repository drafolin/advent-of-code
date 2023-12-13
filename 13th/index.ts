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

const countDiff = (s1: string, s2: string): number => {
	let count = 0;
	let a1 = s1.split("");
	let a2 = s2.split("");

	for (let i in a1) {
		if (a1[i] !== a2[i])
			++count;
	}

	console.log(a1, a2, count);

	return count;
};

const findMirror = (arr: string[]): number | null => {
	for (let lineI = 1; lineI < arr.length; ++lineI) {
		let diffCount = 0;
		if (lineI > arr.length / 2) {
			for (let i = 0; i < arr.length - lineI; ++i) {
				diffCount += countDiff(arr[lineI - i - 1], arr[lineI + i]);
				if (diffCount > 1)
					break;
			}
		} else {
			for (let i = 0; i < lineI; ++i) {
				diffCount += countDiff(arr[lineI - i - 1], arr[lineI + i]);
				if (diffCount > 1)
					break;
			}
		}
		console.log(lineI, diffCount);
		if (diffCount === 1)
			return lineI;
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
