import { readFileSync } from "fs";
import { add, memoize } from "lodash";

const testing = false;
const dataTxt = (testing ? `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`: readFileSync(__dirname + "/data.txt", "utf-8")).split("\n").map(v => v.split(""));


type Point = { x: number, y: number; };

const map = new Map<string, string>();
const stepsNeeded = testing ? 6 : 64;
const startLine = dataTxt.findIndex(v => v.includes("S"));
const startPos = { x: dataTxt[startLine].findIndex(v => v === "S"), y: startLine };
dataTxt[startPos.y][startPos.x] = ".";
let reachablePositions: Set<string> = new Set([JSON.stringify(startPos)]);

for (let line in dataTxt)
	for (let char in dataTxt[line])
		map.set(JSON.stringify({ x: +char, y: +line }), dataTxt[line][char]);

const addStep = memoize((reachablePositions: Set<string>) => {
	let newPositions: Set<string> = new Set();
	for (let pos of reachablePositions) {
		let { x, y } = JSON.parse(pos) as Point;
		[
			{ x, y: y - 1 },
			{ x, y: y + 1 },
			{ x: x - 1, y },
			{ x: x + 1, y }
		].forEach(v => {
			if (map.get(JSON.stringify(v)) === ".")
				newPositions.add(JSON.stringify(v));
		});
	}
	return newPositions;
}, (reachablePositions: Set<string>) => [...reachablePositions.values()].sort().join(","));

for (let i = 0; i < stepsNeeded; ++i) {
	reachablePositions = addStep(reachablePositions);
}

console.log(reachablePositions.size);
