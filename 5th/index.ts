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
`.split("\n\n");*/
//destination source length

console.log(data);

const seeds = data[0].split(" ").slice(1).map(v => parseInt(v));
seeds;
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
		[...currentMap.keys()].forEach(v => {
			if (ok)
				return;
			let nums = v.split(" ").map(v => parseInt(v));
			if (nums[0] <= seed && seed <= nums[1]) {
				temp.push(currentMap.get(v)! + (seed - nums[0]));
				ok = true;
			}
		});
		if (!ok)
			temp.push(seed);
	}
	mapped = temp;
}


let min = Infinity;
for (let location of mapped) {
	if (location < min)
		min = location;
}

min;
