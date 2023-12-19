import { readFileSync } from "fs";

const testing = false;
const [rules, inputs] = (testing ?
	`px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}
`: readFileSync(__dirname + "/data.txt", "utf-8")
).trim().split("\n\n").map(v => v.split("\n"));

type entry = { x: number, m: number, a: number, s: number; };

const entries = inputs.map(v => {
	let entry: entry = { x: 0, m: 0, a: 0, s: 0 };
	v.slice(1, v.length - 1).split(",").forEach(v => {
		const [k, val] = v.split("=");
		entry[k] = parseInt(val);
	});
	return entry;
});

type rule = {
	condition: {
		field: keyof entry,
		condition: "<" | ">",
		value: number;
	},
	action: "R" | "A" | string;
};

type workflow = {
	rules: rule[],
	default: "R" | "A" | string;
};

const flows = new Map<string, workflow>();
for (let flow of rules) {
	const flowRules = flow.split("{")[1].slice(0, flow.split("{")[1].length - 1).split(",");
	let conditions: rule[] = [];
	const d = flowRules.pop()!;
	for (let ruleTxt of flowRules) {
		const [field, condition, value, action] = [...ruleTxt.matchAll(/(.)(<|>)([0-9]*):(.*)/g)][0].slice(1) as any;
		conditions.push({ condition: { field, condition, value: parseInt(value) }, action });
	}


	flows.set(flow.split("{")[0], { rules: conditions, default: d });
}

console.dir(flows);

const applyWorkflow = (input: entry, wf: string): "R" | "A" | string => {
	const flow = flows.get(wf)!;
	for (let rule of flow.rules) {
		if (rule.condition.condition === ">") {
			if (input[rule.condition.field] > rule.condition.value)
				return rule.action;
		} else if (rule.condition.condition === "<") {
			if (input[rule.condition.field] < rule.condition.value)
				return rule.action;
		}
	}
	return flow.default;
};

const runTillEnd = (input: entry, flow: string) => {
	const res = applyWorkflow(input, flow);
	return res === "R" || res === "A" ? res : runTillEnd(input, res);
};

let sum = 0;
for (let input of entries) {
	if (runTillEnd(input, "in") === "A")
		for (let v in input)
			sum += input[v];
}

sum;
