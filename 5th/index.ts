import * as fs from "fs";

const data = fs.readFileSync(__dirname + "/data.txt", "utf-8").split("\n\n");
/*const data = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`.split("\n\n").map(v => v.trim());
//destination source length*/
/*
console.log(data);

const seedsData = data[0].split(" ").slice(1).map(v => parseInt(v));
let min = Infinity;

for (let i = 0; i < seedsData.length; i += 2) {
	let seeds: number[] = [];
	console.log(i);
	console.log(seedsData[i + 1]);
	let k = 0;
	for (let j = 0; j < seedsData[i + 1]; ++j) {
		if (j > k * 1000 + 1000) {
			k = j / 1000;
			console.log(Math.floor(j / 1000) * 1000);
		}
		seeds = [seedsData[i] + j];
		const maps = data.slice(1).map(v => v.trim().split("\n").slice(1).map(v => v.split(" ").map(v => parseInt(v))));
		let mapped = seeds;
		for (let map of maps) {
			const currentMap = new Map<string, number>();

			for (let line of map) {
				let destination = line[0];
				let source = line[1];
				currentMap.set(`${source} ${source + line[2] - 1}`, destination);
			}

			let temp: number[] = [];
			for (let seed of mapped) {
				let ok = false;
				for (let v of [...currentMap.keys()]) {
					if (ok)
						break;
					let nums = v.split(" ").map(v => parseInt(v));
					if (nums[0] <= seed && seed <= nums[1]) {
						temp.push(currentMap.get(v)! + (seed - nums[0]));
						ok = true;
						break;
					}
				};
				if (!ok)
					temp.push(seed);
			}
			mapped = temp;
		}


		for (let location of mapped) {
			if (location < min)
				min = location;
		}
	}
}

min;
*/

let location = 45000000;
const seedsData = data[0].split(":")[1].trim().split(" ");
seedsData;
const mapData = data.slice(1);
mapData;
let min = Infinity;
let nextLog = 45000000;

let maxSeed = 0;
for (let i = 0; i < seedsData.length; i += 2) {
	maxSeed = Math.max(parseInt(seedsData[i]) + parseInt(seedsData[i + 1]), maxSeed);
}

console.log(maxSeed);

do {
	let seed = location;
	for (let i = mapData.length - 1; i >= 0; i--) {
		const map = mapData[i].split("\n").slice(1);
		for (let entry of map) {
			entry;
			const entryArray = entry.split(" ");
			if (seed >= parseInt(entryArray[0]) && (parseInt(entryArray[0]) + parseInt(entryArray[2]) - 1 >= seed)) {
				seed = seed - parseInt(entryArray[0]) + parseInt(entryArray[1]);
				break;
			}
		}
	}

	for (let seedI = 0; seedI < seedsData.length; seedI += 2) {
		if (seed >= parseInt(seedsData[seedI]) && parseInt(seedsData[seedI + 1]) + parseInt(seedsData[seedI]) >= seed) {
			min = Math.min(location, min);
			console.log("FOUND:", min);
			break;
		}
	}

	--location;
	if (nextLog >= location) {
		console.log(location);
		nextLog = location - 1000000;
	}
} while (location > 0);
console.log("done!");
console.log(min);
