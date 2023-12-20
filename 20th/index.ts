import { readFileSync } from "fs";

const testing = true;
const dataTxt = (testing ?
	`broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`: readFileSync(__dirname + "/data.txt", "utf-8")
).trim().split("\n");
