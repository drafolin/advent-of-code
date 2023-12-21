import { readFileSync } from "fs";

const testing = true;
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
...........`: readFileSync(__dirname + "/data.txt", "utf-8")).split("\n").map(v => v.split(""));
