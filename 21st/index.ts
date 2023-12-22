import { readFileSync } from "fs";

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
...........
`: readFileSync(__dirname + "/data.txt", "utf-8")).trim().split(/\r?\n/gm);

const prog = () => {
	const map = new Map<string, number>();
	let start = "";
	for (let y = 0; y < dataTxt.length; y++) {
		for (let x = 0; x < dataTxt[0].length; x++) {
			const char = dataTxt[y][x];
			if (char !== "#") {
				map.set(`${x}, ${y}`, 0);
			}
			if (char === 'S')
				start = `${x}, ${y}`;
		}
	}

	const width = dataTxt[0].length;
	const height = dataTxt.length;
	const toVisit = new Map([[start, 0]]);
	const needed = 26501365;
	const afterNeeded = needed + 1;
	const twoWidth = width * 2;
	const modulo = afterNeeded % twoWidth;
	let good = 0;
	let onestep = 0;
	const validationRounds = 2;

	for (const value of toVisit) {
		const [point, step] = value;
		if (onestep < step && step > (width * 2)) {
			const uv = [...new Set(map.values())]
				.filter(x => x !== 0)
				.sort();
			const groups = [...map.values()]
				.filter(x => x !== 0)
				.reduce((acc, cur) => {
					acc.set(cur, (acc.get(cur) ?? 0) + 1);
					return acc;
				}, new Map<number, number>());

			const groupKeys = [...groups.keys()];
			const neededSeed = Math.floor(2 * afterNeeded / twoWidth) - (modulo === width ? 1 : 0);

			if (uv.length === 2) {
				const seed = Math.floor(2 * step / twoWidth);

				if (seed > validationRounds && (step % twoWidth === modulo)) {
					console.log(groups[groupKeys[0]] * (neededSeed ** 2) + groups[groupKeys[1]] * (neededSeed ** 2 - neededSeed));
					return;
				}
			} else if (uv.length === 3) {
				const seed = Math.floor(2 * step / twoWidth) - (step % twoWidth === width ? 1 : 0);
				if (seed > validationRounds && (step % twoWidth === modulo)) {
					console.log(groups.get(+groupKeys[0])! * (neededSeed ** 2) + groups.get(+groupKeys[1])! * (neededSeed ** 2 + neededSeed) + groups.get(+groupKeys[2])! * (neededSeed ** 2 + neededSeed * 2 + 1));
					return;
				}
			} else {
				throw new Error('cant solve');
			}
			onestep = step;
		}

		const [curX, curY] = point.split(', ').map(Number);
		if (step % 2 === 1) {
			good++;
			const realX = mod(curX, width);
			const realY = mod(curY, height);
			map.set(`${realX}, ${realY}`, map.get(`${realX}, ${realY}`)! + 1);
		}

		[
			{ x: 1, y: 0 },
			{ x: -1, y: 0 },
			{ x: 0, y: 1 },
			{ x: 0, y: -1 }
		].forEach(({ x, y }) => {
			const nextX = curX + x;
			const nextY = curY + y;
			const nextPoint = `${nextX}, ${nextY}`;
			const realX = mod(nextX, width);
			const realY = mod(nextY, height);
			if (map.has(`${realX}, ${realY}`) && !toVisit.has(nextPoint))
				toVisit.set(nextPoint, step + 1);
		});
	};

	console.log(good);
};

function mod(n: number, m: number) {
	return ((n % m) + m) % m;
}

prog();
