import { readFileSync } from "fs";

/*const dataTxt = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`.split("\n");*/
const dataTxt = readFileSync(__dirname + "/data.txt", "utf-8").split("\n");

type Point = { x: number, y: number; };

dataTxt.pop();
const dataTmp = dataTxt.map(v => v.split(""));
let dataRowTmp: string[][] = [];
for (let line of dataTmp) {
	if (!line.includes("#"))
		dataRowTmp.push(line.map(v => v.replace(".", "-")));
	else
		dataRowTmp.push(line);
}
let data: string[][] = [];
for (let i in dataRowTmp[0]) {
	if (dataRowTmp.every(v => v[i] === "." || v[i] === "-"))
		dataRowTmp.forEach((v, j) => {
			if (!data[j])
				data[j] = [];
			if (dataRowTmp[j][i] === ".")
				data[j].push("|");
			else
				data[j].push("+");
		});
	else
		dataRowTmp.forEach((v, j) => {
			if (!data[j])
				data[j] = [];
			data[j].push(v[i]);
		});
}

let galaxies: Point[] = [];

data.forEach((v, i) => {
	v.forEach((c, j) => {
		if (c === "#")
			galaxies.push({ x: j, y: i });
	});
});

console.log(data.map(v => v.join("")).join("\n"));

const EXPANSION_FACTOR = 1_000_000;

let sum = 0;
for (let i = 0; i < galaxies.length - 1; ++i) {
	for (let j = i + 1; j < galaxies.length; ++j) {
		const galaxy1 = galaxies[i];
		const galaxy2 = galaxies[j];

		const minX = Math.min(galaxy1.x, galaxy2.x);
		const maxX = Math.max(galaxy1.x, galaxy2.x);
		const minY = Math.min(galaxy1.y, galaxy2.y);
		const maxY = Math.max(galaxy1.y, galaxy2.y);

		let totalX = 0;
		let totalY = 0;

		for (let x = minX + 1; x <= maxX; ++x) {
			if (data[minY][x] === "|")
				totalX += EXPANSION_FACTOR;
			else
				totalX += 1;
		}

		for (let y = minY + 1; y <= maxY; ++y) {
			if (data[y][maxX] === "-")
				totalY += EXPANSION_FACTOR;
			else
				totalY += 1;
		}

		sum += totalX + totalY;
	}
}

sum;
