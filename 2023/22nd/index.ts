import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ? `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9
` : readFileSync(__dirname + "/data.txt", "utf-8"))
	.trim()
	.split("\n");

class NumRange {
	start: number;
	end: number;

	constructor(start: number, end: number) {
		this.start = start;
		this.end = end;
	}

	intersect = (other: NumRange) => {
		const maxStart = Math.max(this.start, other.start);
		const minEnd = Math.min(this.end, other.end);

		if (maxStart > minEnd)
			return null;

		return new NumRange(maxStart, minEnd);
	};
}

class Brick {
	supports: Brick[] = [];
	supportedBy: Brick[] = [];
	x: NumRange;
	y: NumRange;
	z: NumRange;

	constructor(x: NumRange, y: NumRange, z: NumRange) {
		this.x = x;
		this.y = y;
		this.z = z;
	}

	addOnTop = (other: Brick) => {
		this.supports.push(other);
	};

	addSupport = (other: Brick) => {
		this.supportedBy.push(other);
	};

	toString = () =>
		`[${this.x.start}-${this.x.end},${this.y.start}-${this.y.end},${this.z.start}-${this.z.end}]`;
}

let brickStack: Brick[] = [];

let fallingBricks = dataTxt.map(v => {
	const description = v.split("~").map(w => w.split(",").map(Number));
	return new Brick(
		new NumRange(description[0][0], description[1][0]),
		new NumRange(description[0][1], description[1][1]),
		new NumRange(description[0][2], description[1][2])
	);
}).sort((a, b) => a.z.start - b.z.start);

let fallenBricks: Brick[] = [];
fallingBricks.forEach(brick => {
	var fallingBrick = brick;
	var supported = false;
	while (!supported) {
		let supportingBricks = fallenBricks.filter((v) =>
			v.x.intersect(fallingBrick.x) !== null &&
			v.y.intersect(fallingBrick.y) !== null &&
			v.z.end === fallingBrick.z.start - 1
		);
		supported = supportingBricks.length !== 0 || fallingBrick.z.start === 0;
		if (supported) {
			supportingBricks.forEach(v => {
				v.addOnTop(fallingBrick);
				fallingBrick.addSupport(v);
			});
			fallenBricks.push(fallingBrick);
		} else {
			let nextZ = 0;
			let bricksUnder = fallenBricks
				.filter(v =>
					v.z.end < (fallingBrick.z.start - 1)
				);
			if (bricksUnder.length > 0)
				nextZ = bricksUnder.reduce((acc, cv) => Math.max(cv.z.end, acc), 0) + 1;
			let brickHeight = fallingBrick.z.end - fallingBrick.z.start;
			fallingBrick = new Brick(
				fallingBrick.x, fallingBrick.y, new NumRange(nextZ, nextZ + brickHeight)
			);
		}
	}
});

brickStack = fallenBricks;
console.log(brickStack.filter(brickToTake => !brickToTake.supports.some(v => v.supportedBy.length === 1)).length);
console.log(brickStack
	.filter(brickToTake => brickToTake.supports.some(v => v.supportedBy.length === 1))
	.reduce((sum, cv) => {
		let collapsingStack = brickStack.filter(v => v.toString() !== cv.toString());
		let goneBricks = [cv];
		let notSupported: Brick[] = [];
		do {
			notSupported = collapsingStack
				.filter(v => v.supportedBy.length > 0 &&
					v.supportedBy.every(v => goneBricks.includes(v))
				);
			collapsingStack = collapsingStack.filter(v => !notSupported.includes(v));
			goneBricks.push(...notSupported);
		} while (notSupported.length !== 0);
		return goneBricks.length - 1 + sum;
	}, 0));
