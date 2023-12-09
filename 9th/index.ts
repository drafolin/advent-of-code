import * as fs from "fs";

/*const dataTxt = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`.split("\n");*/

const dataTxt = fs.readFileSync(__dirname + "/data.txt", "utf-8").split("\n");

dataTxt.pop();

const data = dataTxt.map(v => v.split(" ").map(v => parseInt(v)));
let sum = 0;
for (let line of data) {
	let steps: number[][] = [];
	steps.push(line);
	while (!steps[steps.length - 1].every(v => v === 0)) {
		let currentLine = steps[steps.length - 1];
		let nextLine: typeof currentLine = [];

		for (let i = 1; i < currentLine.length; i++) {
			nextLine.push(currentLine[i] - currentLine[i - 1]);
		}
		steps.push(nextLine);
	};

	steps[steps.length - 1].push(0);

	for (let i = steps.length - 2; i >= 0; --i) {
		const currentLine = steps[i];
		steps[i].push(currentLine[currentLine.length - 1] + steps[i + 1][steps[i + 1].length - 1]);
	}
	sum += steps[0][steps[0].length - 1];
}

sum;
