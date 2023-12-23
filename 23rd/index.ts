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

class Solver {
	pos: Point;
	/**
	 * Not including current pos.
	 * Contains the solver path, in string formatted "x,y"
	*/
	path: string[];
	constructor(pos: Point, path: string[]) {
		this.pos = pos;
		this.path = path;
	}

	tryMove = () => {
		let res: Solver[] = [];
		[
			{ x: this.pos.x - 1, y: this.pos.y },
			{ x: this.pos.x + 1, y: this.pos.y },
			{ x: this.pos.x, y: this.pos.y - 1 },
			{ x: this.pos.x, y: this.pos.y + 1 }
		]
			.filter(v => !this.path.includes(`${v.x},${v.y}`))
			.forEach(v => {
				if (map.has(`${v.x},${v.y}`)) {
					switch (map.get(`${v.x},${v.y}`)) {
						case "#": return;
						case ">":
							if (this.pos.x >= v.x)
								return;
							break;
						case "^":
							if (this.pos.y <= v.y)
								return;
							break;
						case "<":
							if (this.pos.x <= v.x)
								return;
							break;
						case "v":
							if (this.pos.y >= v.y)
								return;
							break;
					}
					res.push(new Solver(v, [...this.path, `${this.pos.x},${this.pos.y}`]));
				}
			});
		return res;
	};

	get isFinished() {
		return this.pos.x === endColumn && (this.pos.y === (dataTxt.length - 1));
	}
}

let concurrents = [new Solver(
	{ x: startColumn, y: 0 },
	[]
)];

let lastPath: Solver | undefined;

while (concurrents.length > 0) {
	const pathsToMove = concurrents;
	const newPaths: Solver[] = [];
	pathsToMove.forEach(v => {
		v.tryMove().forEach(w => {
			if (w.isFinished) {
				lastPath = w;
			}
			else
				newPaths.push(w);
		});
	});
	concurrents = newPaths;
}

console.log(lastPath?.path.length);
