import * as fs from "fs";

/*const dataTxt =
	`Time:      7  15   30
Distance:  9  40  200`.split("\n");*/
const dataTxt = fs.readFileSync(__dirname + "/data.txt", "utf-8").split("\n");
const times = parseInt(dataTxt[0].split(" ").filter(v => v !== "").slice(1).join(""));
const distances = parseInt(dataTxt[1].split(" ").filter(v => v !== "").slice(1).join(""));

const data: { time: number; distance: number; } = { time: times, distance: distances };

let count = 0;
for (let i = 0; i <= data.time; i++) {
	// i = temps appuyé = speed mm/ms
	const runTime = data.time - i;
	const spd = i;
	const distance = spd * runTime;

	if (distance > data.distance) {
		++count;
	}
}

count;
