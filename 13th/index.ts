import { readFileSync } from "fs";

const testing = false;

const dataTxt = (testing ? `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
` : readFileSync(__dirname + "/data.txt", "utf-8")).trim().split("\n\n");
