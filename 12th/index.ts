import { readFileSync } from "fs";
import * as _ from "lodash";

console.time();

const testing = false;

const dataTxt = (testing ?
	`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
` :
	readFileSync(__dirname + "/data.txt", "utf-8")).split("\n");

dataTxt.pop();

const getChoices: (data: string, springCount: number[]) => number = _.memoize((data: string, springCount: number[]) => {
	if (data.length === 0)
		return (springCount.length === 0) ? 1 : 0;

	if (data[0] === ".")
		return getChoices(data.slice(1), springCount);

	if (data[0] === "?") {
		let withPound = getChoices("#" + data.slice(1), springCount);
		let withDot = getChoices("." + data.slice(1), springCount);
		return withPound + withDot;
	}

	if (data[0] === "#") {
		if (data.match(RegExp(`^[\\?#]{${springCount[0]}}(\\?|\\.|$)`)))
			return getChoices(data.slice(springCount[0] + 1), springCount.slice(1));
		else
			return 0;
	}
	return NaN;
}, (data, springCount) => JSON.stringify({ data, springCount }));

let sum = 0;
for (let line of dataTxt) {
	const data = line.split(" ")[0];
	const springCount = line.split(" ")[1].split(",").map(v => parseInt(v));
	let dataArr: string[] = [];
	let countArr: number[] = [];

	for (let i = 0; i < 5; ++i) {
		dataArr.push(data);
		countArr.push(...springCount);
	}

	sum += getChoices(dataArr.join("?"), countArr);
};

console.log(sum);
console.timeEnd();
