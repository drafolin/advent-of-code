import * as fs from "fs";

/*const dataTxt =
	`Time:      7  15   30
Distance:  9  40  200`.split("\n");*/
const dataTxt = fs.readFileSync(__dirname + "/data.txt", "utf-8").split("\n");
const times = dataTxt[0].split(" ").filter(v => v !== "").slice(1).map(v => parseInt(v));
const distances = dataTxt[1].split(" ").filter(v => v !== "").slice(1).map(v => parseInt(v));

const data: { time: number; distance: number; }[] = [];


for (let i = 0; i < times.length; ++i) {
	data.push({ time: times[i], distance: distances[i] });
}

let total = 1;

for (let race of data) {
	let count = 0;
	for (let i = 0; i <= race.time; i++) {
		// i = temps appuyé = speed mm/ms
		const runTime = race.time - i;
		const spd = i;
		const distance = spd * runTime;

		if (distance > race.distance) {
			++count;
		}
	}
	total *= count;
}

total;
