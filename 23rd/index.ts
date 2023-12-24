import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ?
	`#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#
`: readFileSync(`${__dirname}/data.txt`, "utf-8"))
	.trim().split("\n").map(v => v.split(""));

const map = new Map<string, string>();
for (let y of dataTxt.keys())
	for (let x of dataTxt[y].keys())
		map.set(`${x},${y}`, dataTxt[y][x]);

let startColumn = 0;
while (dataTxt[0][startColumn] === "#")
	startColumn++;

let endColumn = 0;
while (dataTxt[dataTxt.length - 1][endColumn] === "#")
	endColumn++;

type Point = { x: number, y: number; };

const deltas: Point[] = [
	{ x: 0, y: -1 },
	{ x: 0, y: 1 },
	{ x: -1, y: 0 },
	{ x: 1, y: 0 }
];

const directions: Record<string, Point[]> = {
	'^': [{ x: 0, y: -1 }],
	'v': [{ x: 0, y: 1 }],
	'<': [{ x: -1, y: 0 }],
	'>': [{ x: 1, y: 0 }],
	'.': deltas,
};


function getGraph() {
	const points = new Map<string, string>(dataTxt.flatMap((row, rowI) => row.map((char, colI) => [`${colI},${rowI}`, char])));

	const graph = Object.fromEntries(
		[...points].map((point) => [
			point[0],
			{} as Record<string, number>,
		])
	);

	for (const [sx, sy] of [...points.keys()].map(s => s.split(","))) {
		const stack: [number, number, number][] = [[0, +sy, +sx]];
		const seen = new Set<string>([`${sx},${sy}`]);

		while (stack.length) {
			const [n, y, x] = stack.shift()!;

			if (n && points.has(`${x},${y}`)) {
				graph[`${sx},${sy}`][`${x},${y}`] = n;
				continue;
			}

			const tile = points.get(`${x},${y}`)!;

			if (tile === '#') continue;

			for (const { x: dx, y: dy } of (deltas) as Point[]) {
				const nx = x + dx;
				const ny = y + dy;

				if (ny < 0 || ny >= dataTxt.length || nx < 0 || ny >= dataTxt[0].length)
					continue;

				const newTile = dataTxt.at(ny)?.at(nx);

				if (!newTile || newTile === '#') continue;

				if (seen.has(`${nx},${ny}`)) continue;

				stack.push([n + 1, ny, nx]);
				seen.add(`${nx}, ${ny}`);
			}
		}
	}

	return graph;
}

const seen = new Set<string>();
const graph = getGraph();
const dfs = (point: Point, depth: number) => {
	if (point.x === endColumn && point.y === dataTxt.length - 1)
		return 0;


	let max = -Infinity;

	seen.add(`${point.x},${point.y}`);
	const connections = graph[`${point.x},${point.y}`];

	for (let v of Object.keys(connections)) {
		const [x, y] = v.split(",");
		if (!seen.has(v))
			max = Math.max(max, dfs({ x: +x, y: +y }, depth + 1) + connections[v]);
	};

	seen.delete(`${point.x},${point.y}`);

	return max;
};

console.log(dfs({ x: startColumn, y: 0 }, 0));
