import * as fs from "fs";


enum Combo {
	high,
	s2,
	s2x2,
	s3,
	full,
	s4,
	s5,
}

enum Card {
	j,
	c2 = 1,
	c3,
	c4,
	c5,
	c6,
	c7,
	c8,
	c9,
	t,
	q,
	k,
	a
}

const dataTxt = fs.readFileSync(__dirname + "/data.txt", "utf-8").split("\n");
/*const dataTxt = `2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41
`.split("\n");*/

dataTxt.pop();

const strToCard = (card: string): Card => {
	if (card.length !== 1)
		throw new Error("Not a card");

	switch (card) {
		case "2": return Card.c2;
		case "3": return Card.c3;
		case "4": return Card.c4;
		case "5": return Card.c5;
		case "6": return Card.c6;
		case "7": return Card.c7;
		case "8": return Card.c8;
		case "9": return Card.c9;
		case "T": return Card.t;
		case "J": return Card.j;
		case "Q": return Card.q;
		case "K": return Card.k;
		case "A": return Card.a;
		default: throw new Error("Not a card");
	}
};

const cardToStr = (card: Card): string => {
	switch (card) {
		case Card.c2: return "2";
		case Card.c3: return "3";
		case Card.c4: return "4";
		case Card.c5: return "5";
		case Card.c6: return "6";
		case Card.c7: return "7";
		case Card.c8: return "8";
		case Card.c9: return "9";
		case Card.t: return "T";
		case Card.j: return "J";
		case Card.q: return "Q";
		case Card.k: return "K";
		case Card.a: return "A";
	}
};

let data = dataTxt.map(v => {
	let current = v.split(" ");
	return { game: current[0].split("").map(v => strToCard(v)), bid: parseInt(current[1]) };
});

const count = <T>(arr: T[], srch: T) => {
	let count = 0;
	for (let item of arr) {
		if (item === srch)
			++count;
	}
	return count;
};

type CalculatedGame = { combo: Combo, hand: Card[]; };

const calculateHand = (hand: Card[]): CalculatedGame => {

	const counts: Map<Card, number> = new Map();
	let jokers = 0;
	for (let item of hand) {
		if (item !== Card.j)
			counts.set(item, (counts.get(item) ?? 0) + 1);
		else
			++jokers;
	}

	const combos = [...counts.values()];
	const max = Math.max(0, Math.max(...combos));
	if (combos.includes(5) || max + jokers === 5)
		return { combo: Combo.s5, hand };
	else if (combos.includes(4) || max + jokers === 4)
		return { combo: Combo.s4, hand };
	else if (combos.includes(3)) {
		return {
			combo: combos.includes(2) ?
				Combo.full :
				Combo.s3,
			hand
		};
	} else if (max + jokers >= 3) {
		const sortedCombos = combos.sort((a, b) => b - a);
		const remaining = jokers - (3 - sortedCombos[0]);
		if (sortedCombos[1] + remaining >= 2) {
			return {
				combo: Combo.full,
				hand
			};
		} else {
			return {
				combo: Combo.s3,
				hand
			};
		}
	} else if (combos.includes(2) || max + jokers >= 2) {
		return {
			combo: count(combos, 2) === 2 ||
				(combos.includes(2) && jokers >= 1) ||
				(count(combos, 1) === 2 && jokers === 2) ?
				Combo.s2x2 :
				Combo.s2,
			hand
		};
	} else {
		return {
			combo: Combo.high,
			hand
		};
	}
};

/**
 * Compares two calculated games
 * @param g1 first game
 * @param g2 second game
 * @returns {number} Positive if g1 comes after g2, negative if g1 comes before g2, 0 if g1 is equivalent to g2
 */
const compare = (g1: CalculatedGame, g2: CalculatedGame): number => {
	if (g1.combo > g2.combo)
		return 1;
	else if (g1.combo < g2.combo)
		return -1;
	else {
		for (let i = 0; i <= 5; ++i) {
			if (g1.hand[i] > g2.hand[i])
				return 1;
			else if (g1.hand[i] < g2.hand[i])
				return -1;
		}
		return 0;
	}
};

let ordered: { game: Card[], bid: number; }[] = data.sort((g1, g2) => {
	const c1 = calculateHand(g1.game);
	const c2 = calculateHand(g2.game);
	return compare(c1, c2);
});

let sum = 0;
for (let i of ordered.keys()) {
	sum += ordered[i].bid * (i + 1);
}

sum;
