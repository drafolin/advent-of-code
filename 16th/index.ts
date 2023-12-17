import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ?
	`.|...\\....
|.-.\\.....
.....|-...
........|.
..........
.........\\
..../.\\\\..
.-.-/..|..
.|....-|.\\
..//.|....
`: readFileSync(__dirname + "/data.txt", "utf-8")).trim().split("\n").map(v => v.split(""));


type Point = {
	x: number;
	y: number;
};

type Direction = {
	dx: number;
	dy: number;
};

const energised = new Set<string>();

class Beam {
	direction: Direction;
	position: Point;
	constructor(dir: Direction, pos: Point) {
		this.direction = dir;
		this.position = pos;
	}

	public move() {
		let char = dataTxt[this.position.y][this.position.x];
		switch (char) {
			case "/":
				this.direction = {
					dx: -this.direction.dy,
					dy: -this.direction.dx
				};
				break;
			case "\\":
				this.direction = {
					dx: this.direction.dy,
					dy: this.direction.dx
				};
				break;
			case "|":
				if (this.direction.dy !== 0)
					break;
				return [
					new Beam({ dx: 0, dy: 1 }, { x: this.position.x, y: this.position.y + 1 }),
					new Beam({ dx: 0, dy: -1 }, { x: this.position.x, y: this.position.y - 1 })
				];
			case "-":
				if (this.direction.dx !== 0)
					break;
				return [
					new Beam({ dx: 1, dy: 0 }, { x: this.position.x + 1, y: this.position.y }),
					new Beam({ dx: -1, dy: 0 }, { x: this.position.x - 1, y: this.position.y })
				];

		}
		this.position.x += this.direction.dx;
		this.position.y += this.direction.dy;
		return [this];
	}

	public isValid() {
		if (this.position.x < 0 || dataTxt[0].length <= this.position.x ||
			this.position.y < 0 || dataTxt.length <= this.position.y)
			return false;

		const char = dataTxt[this.position.y][this.position.x];
		if ((char === "|" || char === "-") && energised.has(JSON.stringify(this.position)))
			return false;

		return true;
	}
};

let beams: Beam[] = [new Beam({ dx: 1, dy: 0 }, { x: 0, y: 0 })];

while (beams.length > 0) {
	let newBeams: Beam[] = [];
	for (let beam of beams) {
		energised.add(JSON.stringify(beam.position));
		newBeams.push(...beam.move());
	}
	let validBeams: Beam[] = [];
	for (let beam of newBeams) {
		if (beam.isValid())
			validBeams.push(beam);
	}
	beams = validBeams;
}

console.log(energised.size);
