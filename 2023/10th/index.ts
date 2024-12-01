import { readFileSync } from "fs";

/*const dataTxt = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`.split("\n");*/
const dataTxt = readFileSync(__dirname + "/data.txt", "utf-8").split("\n");

dataTxt.pop();
const data = dataTxt.map(v => v.split(""));

enum direction {
	down,
	neutral,
	up
}
type Position = { x: number, y: number; };
type Node = { pos: Position, direction: direction; };

const startingRow = data.find(v => v.includes("S"))!;
const startingColIndex = startingRow.indexOf("S");
const startingRowIndex = data.indexOf(startingRow);

const startingPos: Position = { x: startingRowIndex, y: startingColIndex };

const arrayEquals = <T>(a1: T[], a2: T[]) =>
	a1.map((v, i) => v === a2[i]).every(v => v);

const posEquals = (p1: Position, p2: Position) =>
	p1.x === p2.x && p1.y === p2.y;

const connecting = (pos: Position): [Position, Position] | null => {
	const letter = data[pos.y][pos.x];
	switch (letter) {
		case "F":
			return [
				{ x: pos.x, y: pos.y + 1 },
				{ x: pos.x + 1, y: pos.y }
			];
		case "J":
			return [
				{ x: pos.x, y: pos.y - 1 },
				{ x: pos.x - 1, y: pos.y }
			];
		case "7":
			return [
				{ x: pos.x, y: pos.y + 1 },
				{ x: pos.x - 1, y: pos.y }
			];
		case "L":
			return [
				{ x: pos.x, y: pos.y - 1 },
				{ x: pos.x + 1, y: pos.y }
			];
		case "-":
			return [
				{ x: pos.x + 1, y: pos.y },
				{ x: pos.x - 1, y: pos.y }
			];
		case "|":
			return [
				{ x: pos.x, y: pos.y + 1 },
				{ x: pos.x, y: pos.y - 1 }
			];
		case "S":
			let found: Position[] = [];
			[
				{ x: pos.x - 1, y: pos.y },
				{ x: pos.x + 1, y: pos.y },
				{ x: pos.x, y: pos.y - 1 },
				{ x: pos.x, y: pos.y + 1 }
			].forEach(v => {
				if (connecting(v)?.some(w => posEquals(w, pos)))
					found.push(v);
			});
			return [found[0], found[1]];
		default:
			return null;
	}
};

const last = <T>(a: T[]) => {
	return a[a.length - 1];
};

let startingNode: Node = { pos: startingPos, direction: direction.neutral };
const next: Node = { pos: connecting(startingPos)![0], direction: direction.neutral };

let loop: Node[] = [startingNode, next];
while (!posEquals(last(loop).pos, loop[0].pos)) {
	const connectingPos = connecting(last(loop).pos)!;
	const newPos = (connectingPos[posEquals(connectingPos[0], loop[loop.length - 2].pos) ? 1 : 0]);
	const newNode: Node = { pos: newPos, direction: direction.neutral };
	if (newPos.y > last(loop).pos.y)
		loop[loop.length - 1].direction = direction.down;
	else if (loop[loop.length - 2].pos.y > last(loop).pos.y)
		loop[loop.length - 1].direction = direction.down;

	loop.push(newNode);
};

loop.pop();

if (loop[1].pos.y > loop[0].pos.y)
	loop[0].direction = direction.down;
if (last(loop).pos.y > loop[0].pos.y)
	loop[0].direction = direction.up;

if (loop[0].pos.y > last(loop).pos.y)
	loop[loop.length - 1].direction = direction.down;
if (loop[loop.length - 2].pos.y > last(loop).pos.y)
	loop[loop.length - 1].direction = direction.up;

const size = Math.ceil(loop.length / 2.0);
size;

let containedCount = 0;
data.forEach((_, y) => {
	let isInside = false;
	data[y].forEach((_, x) => {
		const currentNode = loop.find(v =>
			posEquals(v.pos, { x, y }));
		if (currentNode && currentNode.direction !== direction.neutral)
			isInside = !isInside;
		else if (!currentNode && isInside)
			++containedCount;
	});
});

containedCount;
