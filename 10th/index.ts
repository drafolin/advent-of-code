import { readFileSync } from "fs";

/*const dataTxt = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`.split("\n");*/
const dataTxt = readFileSync(__dirname + "/data.txt", "utf-8").split("\n");

dataTxt.pop();
const data = dataTxt.map(v => v.split(""));

type Position = [number, number];

const startingRow = data.find(v => v.includes("S"))!;
const startingColIndex = startingRow.indexOf("S");
const startingRowIndex = data.indexOf(startingRow);

const startingPos: Position = [startingRowIndex, startingColIndex];

let currentPos: Position = startingPos;
let lastPos = currentPos;

const arrayEquals = <T>(a1: T[], a2: T[]) =>
	a1.map((v, i) => v === a2[i]).every(v => v);

const connecting = (pos: Position): [Position, Position] | null => {
	const letter = data[pos[0]][pos[1]];
	switch (letter) {
		case "F":
			return [
				[pos[0], pos[1] + 1],
				[pos[0] + 1, pos[1]]
			];
		case "J":
			return [
				[pos[0], pos[1] - 1],
				[pos[0] - 1, pos[1]]
			];
		case "7":
			return [
				[pos[0], pos[1] - 1],
				[pos[0] + 1, pos[1]]
			];
		case "L":
			return [
				[pos[0], pos[1] + 1],
				[pos[0] - 1, pos[1]]
			];
		case "-":
			return [
				[pos[0], pos[1] - 1],
				[pos[0], pos[1] + 1]
			];
		case "|":
			return [
				[pos[0] - 1, pos[1]],
				[pos[0] + 1, pos[1]]
			];
		case "S":
			let found: Position[] = [];
			if (connecting([pos[0] - 1, pos[1]])?.some(v => arrayEquals(v, pos)))
				found.push([pos[0] - 1, pos[1]]);
			if (connecting([pos[0] + 1, pos[1]])?.some(v => arrayEquals(v, pos)))
				found.push([pos[0] + 1, pos[1]]);
			if (connecting([pos[0], pos[1] - 1])?.some(v => arrayEquals(v, pos)))
				found.push([pos[0], pos[1] - 1]);
			if (connecting([pos[0], pos[1] + 1])?.some(v => arrayEquals(v, pos)))
				found.push([pos[0], pos[1] + 1]);
			return [found[0], found[1]];
		default:
			return null;
	}
};

let previous = JSON.parse(JSON.stringify(currentPos)) as Position;
currentPos = connecting(startingPos)![0];
let iterations = 1;
while (!arrayEquals(currentPos, startingPos)) {
	console.log(startingPos, currentPos);
	const connectingPos = connecting(currentPos)!;
	let temp: Position = connectingPos[arrayEquals(connectingPos[0], previous) ? 1 : 0];
	previous = JSON.parse(JSON.stringify(currentPos)) as Position;
	currentPos = temp;
	++iterations;
};

const res = Math.ceil(iterations / 2.0);
res;
