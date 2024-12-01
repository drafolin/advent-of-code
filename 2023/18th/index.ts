import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ?
	`R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
` :
	readFileSync(__dirname + "/data.txt", "utf-8"))
	.trim()
	.split("\n")
	.map(l => l.split(" "))
	.map<[string, number]>(([_1, _2, c]) => {
		const direction = (l: string) => {
			switch (l) {
				case "0": return "R";
				case "1": return "D";
				case "2": return "L";
				case "3": return "U";
				default: return "";
			}
		};
		let color = c.slice(2, c.length - 1);
		const dist = parseInt(color.slice(0, color.length - 1), 16);
		return [direction(color[color.length - 1]), dist];
	});

const dirs = {
	U: [-1, 0],
	D: [1, 0],
	L: [0, -1],
	R: [0, 1]
};

let currentRow = 0;
let currentCol = 0;
let area = 0;
let perimeter = 0;

for (let i = 0; i < dataTxt.length; i++) {
	const dir = dirs[dataTxt[i][0]];
	const dist = dataTxt[i][1];
	let previousRow = currentRow;
	let previousCol = currentCol;
	currentRow += dir[0] * dist;
	currentCol += dir[1] * dist;
	area += previousCol * currentRow - currentCol * previousRow;
	perimeter += dist;
}
console.log(Math.abs(area / 2) + perimeter / 2 + 1);
