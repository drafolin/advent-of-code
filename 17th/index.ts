import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ?
	`2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
`: readFileSync(__dirname + "/data.txt", "utf-8")
).trim().split("\n").map(v => v.split(""));

class Direction {
	readonly row: number;
	readonly col: number;

	constructor(row: number, col: number) {
		this.row = row;
		this.col = col;
		Object.freeze(this);
	}

	add(other: Direction) {
		return new Direction(this.row + other.row, this.col + other.col);
	}

	reverse() { return new Direction(-this.row, -this.col); }
	turnLeft() { return new Direction(-this.col, this.row); }
	turnRight() { return new Direction(this.col, -this.row); }

	static up = new Direction(-1, 0);
	static down = new Direction(1, 0);
	static left = new Direction(0, -1);
	static right = new Direction(0, 1);
}

class Position {
	readonly row: number;
	readonly col: number;

	constructor(row: number, col: number) {
		this.row = row;
		this.col = col;
		Object.freeze(this);
	}

	clone() { return new Position(this.row, this.col); }
	up() { return new Position(this.row - 1, this.col); }
	down() { return new Position(this.row + 1, this.col); }
	left() { return new Position(this.row, this.col - 1); }
	right() { return new Position(this.row, this.col + 1); }
	move(dir: Direction) {
		return new Position(this.row + dir.row, this.col + dir.col);
	}
	toString() { return `(${this.col};${this.row})`; }
}

let allPoints: [string, string][] = [];
dataTxt.forEach((v, y) => v.forEach((w, x) => allPoints.push([new Position(x, y).toString(), w])));
const map = new Map<string, string>(allPoints);
type Gnome = { pos: Position, dir: Direction, dist: number; };


const seen = new Set<string>();
let totalHeatLoss = 0;
const queue: Gnome[][] = [[{ dir: Direction.right, dist: 0, pos: new Position(0, 0) }, { dir: Direction.down, dist: 0, pos: new Position(0, 0) }]];

const tryMove = (baseGnome: Gnome, dir: Direction) => {
	const newGnome: Gnome = {
		pos: baseGnome.pos.move(dir),
		dir,
		dist: baseGnome.dir === dir ? baseGnome.dist + 1 : 1
	};

	if (newGnome.dist > 10)
		return;

	if (!map.has(newGnome.pos.toString()))
		return;

	const id = [newGnome.pos.col, newGnome.pos.row, newGnome.dir.col, newGnome.dir.row, newGnome.dist].join();

	if (seen.has(id))
		return;

	seen.add(id);


	const newHeatLoss = totalHeatLoss + parseInt(map.get(newGnome.pos.toString())!);
	queue[newHeatLoss] ??= [];
	queue[newHeatLoss].push(newGnome);
};

const checkStep = () => {
	for (const gnome of (queue[totalHeatLoss] ?? [])) {
		if (gnome.pos.col === dataTxt[0].length - 1 && gnome.pos.row === dataTxt.length - 1 && gnome.dist >= 4)
			return;
		tryMove(gnome, gnome.dir);
		if (gnome.dist >= 4) {
			tryMove(gnome, gnome.dir.turnLeft());
			tryMove(gnome, gnome.dir.turnRight());
		}
	}
	totalHeatLoss++;
	checkStep();
};
checkStep();
console.log(totalHeatLoss);
