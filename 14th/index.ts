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

const spinMatrix = <T>(m: T[][], cw: boolean) => {
	let cols: T[][] = [];
	if (cw)
		for (let vI = m.length - 1; vI >= 0; --vI) {
			let v = m[vI];
			for (let i = 0; i < v.length; ++i) {
				if (!cols[i])
					cols[i] = [];
				cols[i].push(v[i]);
			}
		}
	else
		for (let v of m) {
			for (let i = 0; i < v.length; ++i) {
				if (!cols[i])
					cols[i] = [];
				cols[i].push(v[v.length - 1 - i]);
			}
		}
	return cols;
};

const moveToLeft = (m: string[][]) => {
	let res: string[][] = [];
	m.forEach((v, colI) => {
		res.push(
			v.join("")
				.split("#")
				.map(v => {
					const oCount = v.split("").reduce((pv, cv) => cv === "O" ? (pv ?? 0) + 1 : pv ?? 0, 0);
					return "O".repeat(oCount) + ".".repeat(v.length - oCount);
				})
				.join("#")
				.split(""));
		for (let i = 1; i < v.length; ++i) {
			if (v[i] === "O" && v[i - 1] === ".") {
				grid[colI][i] = ".";
				grid[colI][i - 1] = "O";
			}
		}
	});
	return res;
};

const spinCycle = (m: string[][]) => {
	let newM = m;
	newM = spinMatrix(newM, false);
	for (let i = 0; i < 4; ++i) {
		newM = moveToLeft(newM);
		newM = spinMatrix(newM, true);
	}
	newM = spinMatrix(newM, true);
	return newM;
};

const identity = (m: string[][]) =>
	m.map(v => v.join("")).join("\n");


let grid = dataTxt.map(v => v.split(""));

let seen = new Map<string, number>();
seen.set(identity(grid), 0);

let cycle_start = 0;
let cycle_len = 0;

for (let i = 0; i < 1000000000; ++i) {
	grid = spinCycle(grid);

	if (seen.has(identity(grid))) {
		const prev = seen.get(identity(grid))!;
		cycle_start = prev;
		cycle_len = i - prev + 1;
		break;
	}
	seen.set(identity(grid), i + 1);
}

for (let i = 0; i < (1000000000 - cycle_start) % cycle_len; ++i) {
	grid = spinCycle(grid);
}

let totalLoad = 0;
for (let i of grid.keys()) {
	const load = grid[0].length - i;
	grid[i].forEach(v => {
		if (v === "O")
			totalLoad += load;
	});
}

totalLoad;
