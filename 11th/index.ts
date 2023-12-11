import { readFileSync } from "fs";

const dataTxt = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`.split("\n");
//const dataTxt = readFileSync(__dirname + "/data.txt", "utf-8").split("\n");


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
