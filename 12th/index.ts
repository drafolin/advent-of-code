import { readFileSync } from "fs";

const testing = true;

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