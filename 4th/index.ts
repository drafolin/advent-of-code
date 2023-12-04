import * as fs from "fs";

const data = fs.readFileSync(__dirname + "/data.txt", "utf-8").split("\n");
data.pop();

let points = 0;

for (let line of data) {
	let linePoints = 0;
	let input = line.split(":")[1];
	let winning = input.split("|")[0]
		.split(" ")
		.filter(v => v !== "")
		.map(v => parseInt(v));

	console.log(winning);

	let numbers = input.split("|")[1]
		.split(" ")
		.filter(v => v !== "")
		.map(v => parseInt(v));

	console.log(numbers);

	for (let number of numbers) {
		if (winning.includes(number))
			linePoints = linePoints === 0 ? 1 : linePoints * 2;
	}
	points += linePoints;
}

points;
