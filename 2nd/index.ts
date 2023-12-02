import { log } from "console";
import { readFileSync } from "fs";

const input = readFileSync(__dirname + "/data.txt", "utf-8").split("\n");

// remove last (empty) line;
input.pop();

const MAX_RED = 12;
const MAX_GREEN = 13;
const MAX_BLUE = 14;

let sum = 0;

for (let game of input) {
	let gameId = game.split(" ")[1];
	gameId = gameId.slice(0, gameId.length - 1);
	log("id:", gameId);
	let gameData = game.split(":")[1].trim();
	let sets = gameData.split("; ");
	let blue = 0, green = 0, red = 0;
	for (let set of sets) {
		let colors = set.split(", ");
		for (let color of colors) {
			let [countStr, colorName] = color.split(" ");
			let count = parseInt(countStr);
			switch (colorName) {
				case "green": green = Math.max(green, count); break;
				case "blue": blue = Math.max(blue, count); break;
				case "red": red = Math.max(red, count); break;
			}
		}
	}
	console.log(blue, green, red);
	if (blue <= MAX_BLUE &&
		green <= MAX_GREEN &&
		red <= MAX_RED)
		sum += parseInt(gameId);
}

sum;
