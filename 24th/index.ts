// USE NODE.JS FOR Z3. BUN IS NOT WORKING.

import { readFileSync } from "fs";
import { init } from "z3-solver";

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
(async () => {
	const { Context } = await init();
	const Z3 = Context("main");

	const x = Z3.Real.const("x");
	const y = Z3.Real.const("y");
	const z = Z3.Real.const("z");

	const dx = Z3.Real.const("dx");
	const dy = Z3.Real.const("dy");
	const dz = Z3.Real.const("dz");

	const solver = new Z3.Solver();

	for (let i in availableStones) {
		const posS = availableStones[i];
		const diff = hailStones.get(posS)!;
		const pos = stop3(posS);

		const t = Z3.Real.const(`t${i}`);

		solver.add(t.ge(0));
		solver.add(x.add(dx.mul(t)).eq(t.mul(diff.x).add(pos.x)));
		solver.add(y.add(dy.mul(t)).eq(t.mul(diff.y).add(pos.y)));
		solver.add(z.add(dz.mul(t)).eq(t.mul(diff.z).add(pos.z)));
	}
	console.log("Solving...");
	const isSat = await solver.check();
	console.log("Solved.");

	if (isSat !== "sat") console.error("Couldn't find a solution");

	const model = solver.model();
	const rx = Number(model.eval(x));
	const ry = Number(model.eval(y));
	const rz = Number(model.eval(z));

	console.log(rx + ry + rz);
})();
