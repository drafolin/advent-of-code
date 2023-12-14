import { readFileSync } from "fs";

const testing = false;
const dataTxt = (testing ? `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`: readFileSync(__dirname + "/data.txt", "utf-8")).split("\n");


