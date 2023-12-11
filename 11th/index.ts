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
		dataRowTmp.push(line);
	dataRowTmp.push(line);
}
let data: string[][] = [];
for (let i in dataRowTmp[0]) {
	if (dataRowTmp.every(v => v[i] === "."))
		dataRowTmp.forEach((v, j) => {
			if (!data[j])
				data[j] = [];
			data[j].push(v[i]);
		});
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

let sum = 0;
for (let i = 0; i < galaxies.length - 1; ++i) {
	for (let j = i + 1; j < galaxies.length; ++j) {
		sum += Math.abs(galaxies[i].x - galaxies[j].x) + Math.abs(galaxies[i].y - galaxies[j].y);
	}
}

sum;
