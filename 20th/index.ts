import { readFileSync } from "fs";

const testingNr: number = 0;
const dataTxt = (testingNr === 1 ?
	`broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`: testingNr === 2 ? `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
` : readFileSync(__dirname + "/data.txt", "utf-8")
).trim().split("\n");


const HIGH = true;
const LOW = false;

type Signal = typeof HIGH | typeof LOW;

let queue: { from: string, to: string, signal: Signal; }[] = [];

class Module {
	outputs: string[];
	name: string;
	constructor(outputs: string[], name: string) {
		this.name = name;
		this.outputs = outputs;
	}

	input = (signal: Signal, input: string): { high: number, low: number; } => {
		let count = { high: 0, low: 0 };
		this.outputs.forEach(v => {
			if (signal === LOW)
				++count.low;
			else
				++count.high;

			queue.push({ to: v, from: this.name, signal: signal });
		});
		return count;
	};
	connect = (input: string) => { };
}

class DebugModule extends Module {
	constructor(name: string) {
		super([], name);
	}

	input = (signal: Signal, input: string): { high: number, low: number; } => {
		console.log("recieved signal: ", { signal, input });
		return { high: 0, low: 0 };
	};
}

class FlipFlop extends Module {
	state: Signal;
	constructor(outputs: string[], name: string) {
		super(outputs, name);
		this.state = false;
	}

	input = (signal: Signal): { high: number, low: number; } => {
		let count = { high: 0, low: 0 };
		if (signal === LOW) {
			this.state = !this.state;
			this.outputs.forEach(v => {
				if (this.state === LOW)
					++count.low;
				else
					++count.high;
				queue.push({ to: v, from: this.name, signal: this.state });
			});
		}
		return count;
	};

	connect = (input: string) => { };
}

class Conjunction extends Module {
	memory: Map<string, boolean>;
	constructor(outputs: string[], name: string) {
		super(outputs, name);
		this.memory = new Map<string, boolean>();
	}

	input = (signal: boolean, input: string): { high: number, low: number; } => {
		let count = { high: 0, low: 0 };
		this.memory.set(input, signal);
		if ([...this.memory.values()].every(v => v === HIGH))
			this.outputs.forEach(v => {
				++count.low;
				queue.push({ to: v, from: this.name, signal: LOW });
			});
		else
			this.outputs.forEach(v => {
				++count.high;
				queue.push({ to: v, from: this.name, signal: HIGH });
			});
		return count;
	};

	connect = (input: string) => {
		this.memory.set(input, false);
	};
}

const nodes = new Map<string, Module>();

for (let node of dataTxt) {
	const [name, outputs] = node.split(" -> ");
	switch (name[0]) {
		case "%": nodes.set(name.slice(1), new FlipFlop(outputs.split(", "), name.slice(1))); break;
		case "&": nodes.set(name.slice(1), new Conjunction(outputs.split(", "), name.slice(1))); break;
		default: nodes.set(name, new Module(outputs.split(", "), name)); break;
	}
}

for (let node of nodes.keys()) {
	nodes.get(node)!.outputs.forEach(v => {
		if (!nodes.has(v))
			nodes.set(v, new DebugModule(v));
		nodes.get(v)!.connect(node);
	});
}



const pushButton = () => {
	queue.push({ from: "", to: "broadcaster", signal: LOW });

	let signalCount = { low: 1, high: 0 };
	while (queue.length > 0) {
		const instruction = queue.shift()!;
		const res = nodes.get(instruction.to)!.input(instruction.signal, instruction.from);
		signalCount.low += res.low;
		signalCount.high += res.high;
	}
	return signalCount;
};

let count = { low: 0, high: 0 };
for (let i = 0; i < 1000; ++i) {
	const res = pushButton();
	count.low += res.low;
	count.high += res.high;
}

console.log(count);
console.log(count.low * count.high);
