import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ?
	`19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
`: readFileSync(`${__dirname}/data.txt`, "utf-8"))
	.trim()
	.split("\n")
	.map(v =>
		v
			.split("@")
			.map(w =>
				w
					.trim()
					.split(", ")
			)
	);

console.clear();

type Point2 = {
	x: number,
	y: number,
};

type Point3 = Point2 & {
	z: number;
};

const p3tos = (point: Point3) => `${point.x},${point.y},${point.z}`;
const stop3 = (s: string) => {
	const [x, y, z] = s.split(",");
	return { x: +x, y: +y, z: +z };
};

const hailStones = new Map<string, Point3>();

for (let stone of dataTxt) {
	const [[sPx, sPy, sPz], [sDx, sDy, sDz]] = stone;
	const [[px, py, pz], [dx, dy, dz]] = [[+sPx, +sPy, +sPz], [+sDx, +sDy, +sDz]];

	hailStones.set(p3tos({ x: px, y: py, z: pz }), { x: dx, y: dy, z: dz });
}

const availableStones = [...hailStones.keys()];

const zoneToCheck: { from: Point2, to: Point2; } = testing ? {
	from: {
		x: 7,
		y: 7,
	},
	to: {
		x: 27,
		y: 27
	}
} : {
	from: {
		x: 200000000000000,
		y: 200000000000000,
	},
	to: {
		x: 400000000000000,
		y: 400000000000000
	}
};

let collisionCount = 0;

for (let stoneI in availableStones) {
	const iPosS = availableStones[stoneI];
	const iPos = stop3(iPosS);
	const { x: iDx, y: iDy } = hailStones.get(iPosS)!;
	for (let stoneJ = +stoneI + 1; stoneJ < availableStones.length; ++stoneJ) {
		const jPosS = availableStones[stoneJ];
		const jPos = stop3(jPosS);
		const { x: jDx, y: jDy } = hailStones.get(jPosS)!;
		let dx = jPos.x - iPos.x;
		let dy = jPos.y - iPos.y;
		let det = jDx * iDy - jDy * iDx;
		let u = (dy * jDx - dx * jDy) / det;
		let v = (dy * iDx - dx * iDy) / det;
		let result = { x: iPos.x + iDx * u, y: iPos.y + iDy * u };
		console.log(result);
		if (u > 0 && v > 0 &&
			zoneToCheck.from.x <= result.x &&
			result.x <= zoneToCheck.to.x &&
			zoneToCheck.from.y <= result.y &&
			result.y <= zoneToCheck.to.y)
			++collisionCount;
	}
}

console.log(collisionCount);
