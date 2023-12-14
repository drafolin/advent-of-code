import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ?
	`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`: readFileSync(__dirname + "/data.txt", "utf-8")).trim().split("\n");

let cols: string[][] = [];
for (let i = 0; i < dataTxt[0].split("").length; ++i) {
	cols[i] = [];
	dataTxt.forEach(v => {
		cols[i].push(v[i]);
	});
}

console.log(cols);

while (cols.some(v => v.join("").match(/\.O/))) {
	cols.forEach((v, colI) => {
		for (let i = 1; i < v.length; ++i) {
			if (v[i] === "O" && v[i - 1] === ".") {
				cols[colI][i] = ".";
				cols[colI][i - 1] = "O";
			}
		}
	});
}

cols;

let totalLoad = 0;
for (let i = 0; i < cols[0].length; ++i) {
	const load = cols[0].length - i;
	cols.forEach(v => {
		if (v[i] === "O")
			totalLoad += load;
	});
}

totalLoad;
