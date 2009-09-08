package net.sf.lapg.lex;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;

import net.sf.lapg.INotifier;

public class RegexpParser {

	private INotifier err;

	// result
	private int[] symbols;
	private ArrayList<int[]> setpool;

	// temporary variables
	private int[] sym;
	private int[] stack;
	private int[] set;

	public RegexpParser(INotifier err) {
		this.err = err;
		this.sym = new int[LexConstants.MAX_ENTRIES];
		this.stack = new int[LexConstants.MAX_DEEP];
		this.set = new int[LexConstants.SIZE_SYM];

		this.symbols = new int[LexConstants.SIZE_SYM];
		this.setpool = new ArrayList<int[]>();

		Arrays.fill(symbols, 0);
		symbols[(0) / LexConstants.BITS] |= (1 << ((0) % LexConstants.BITS));
		symbols[(1) / LexConstants.BITS] |= (1 << ((1) % LexConstants.BITS));
	}

	private int index;
	private char[] re;
	private String regexp;

	public static int parseHex(String s) {
		int result = 0;
		for(int i = 0; i < s.length(); i++) {
			result <<= 4;
			int c = s.codePointAt(i);
			if(c >= 'a' && c <= 'f') {
				result |= 10 + c - 'a';
			} else if(c >= 'A' && c <= 'F') {
				result |= 10 + c - 'A';
			} else if(c >= '0' && c <= '9') {
				result |= c - '0'; 
			} else {
				throw new NumberFormatException();
			}			
		}
		return result;
	}

	private int escape() {
		index++;
		if (index >= re.length) {
			err.error("lex: \\ found at the end of expression: /" + regexp + "/\n");
			return -1;
		}

		int c = re[index];
		switch (c) {
		case 'a':
			return 7;
		case 'b':
			return '\b';
		case 'f':
			return '\f';
		case 'n':
			return '\n';
		case 'r':
			return '\r';
		case 't':
			return '\t';
		case 'u':
		case 'x': {
			if (index + 4 >= re.length) {
				err.error("lex: unicode symbol is incomplete: /" + regexp + "/\n");
				return -1;
			}
			index += 4;
			try {
				return parseHex(new String(re, index - 3, 4));
			} catch(NumberFormatException ex) {
				err.error("lex: bad unicode symbol: /" + regexp + "/\n");
				return -1;
			}
		}

		default:
			return c;
		}
	}

	private int storeSet() {
		int[] newSet = new int[LexConstants.SIZE_SYM];
		for (int i = 0; i < LexConstants.SIZE_SYM; i++) {
			newSet[i] = set[i];
		}
		int setIndex = setpool.size() * LexConstants.SIZE_SYM;
		setpool.add(newSet);
		return setIndex;
	}

	/**
	 * regexp "tokenizer"
	 */
	private int getnext() {
		int i, e;
		boolean invert = false;

		switch (re[index]) {
		case '(':
			index++;
			return LexConstants.LBR;

		case ')':
			index++;
			return LexConstants.RBR;

		case '|':
			index++;
			return LexConstants.OR;

		case '.':
			index++;
			symbols[('\n') / LexConstants.BITS] |= (1 << (('\n') % LexConstants.BITS));
			return LexConstants.ANY;

		case '[':
			index++;
			invert = index < re.length && re[index] == '^';
			if (invert) {
				index++;
			}

			Arrays.fill(set, 0);

			while (index < re.length && re[index] != ']') {
				i = re[index];
				if (i == '\\') {
					i = escape();
					if (i == -1) {
						return -1;
					}
				}
				set[(i) / LexConstants.BITS] |= (1 << ((i) % LexConstants.BITS));
				index++;
				if (index + 1 < re.length && re[index] == '-' && re[index + 1] != ']') {
					e = re[++index];
					if (e == '\\') {
						e = escape();
						if (e == -1) {
							return -1;
						}
					}

					if (e > i) {
						for (i++; i <= e; i++) {
							set[(i) / LexConstants.BITS] |= (1 << ((i) % LexConstants.BITS));
						}
					} else {
						for (i--; i >= e; i--) {
							set[(i) / LexConstants.BITS] |= (1 << ((i) % LexConstants.BITS));
						}
					}

					index++;
				}
			}

			if (index >= re.length) {
				err.error("lex: enclosing square bracket not found: /" + regexp + "/\n");
				return -1;
			}

			for (i = 0; i < LexConstants.SIZE_SYM; i++) {
				symbols[i] |= set[i];
			}

			if (invert) {
				for (i = 0; i < LexConstants.SIZE_SYM; i++) {
					set[i] = ~set[i];
				}
				set[(0) / LexConstants.BITS] &= ~(1 << ((0) % LexConstants.BITS));
			}

			index++;
			return storeSet();

		case '\\':
			i = escape();
			if (i == -1) {
				return -1;
			}
			symbols[(i) / LexConstants.BITS] |= (1 << ((i) % LexConstants.BITS));
			index++;
			return i | LexConstants.SYM;

		default:
			i = re[index];
			// TODO support i > 256 (encode as utf-8)
			symbols[(i) / LexConstants.BITS] |= (1 << ((i) % LexConstants.BITS));
			index++;
			return i | LexConstants.SYM;
		}
	}

	/**
	 * @return Engine representation of regular expression
	 */
	public int[] compile(int number, String name, String regexp) {

		int length = 0, deep = 1;
		boolean addbrackets = false;

		this.index = 0;
		this.regexp = regexp;
		this.re = this.regexp.toCharArray();

		if (re.length == 0) {
			err.error("lex: regexp for `" + name + "' does not contain symbols\n");
			return null;
		}

		while (index < re.length) {
			if (length > LexConstants.MAX_ENTRIES - 5) {
				err.error("lex: regexp for `" + name + "' is too long: /" + regexp + "/\n");
				return null;
			}

			sym[length] = getnext();

			switch (sym[length]) {
			case LexConstants.LBR:
				stack[deep] = length;
				if (++deep >= LexConstants.MAX_DEEP) {
					err.error("lex: regexp for `" + name + "' is too deep: /" + regexp + "/\n");
					return null;
				}
				break;
			case LexConstants.OR:
				if (deep == 1) {
					addbrackets = true;
				}
				break;
			case LexConstants.RBR:
				if (--deep == 0) {
					err.error("lex: error in `" + name + "', wrong parantheses: /" + regexp + "/\n");
					return null;
				}
				sym[stack[deep]] |= length;
				/* FALLTHROUGH */
			default:
				if (index < re.length && (re[index] == '+' || re[index] == '?' || re[index] == '*')) {
					switch (re[index]) {
					case '+':
						sym[length] |= 1 << 29;
						break;
					case '*':
						sym[length] |= 2 << 29;
						break;
					case '?':
						sym[length] |= 3 << 29;
						break;
					}
					if (re[index] != '?') {
						sym[++length] = LexConstants.SPL;
					}
					index++;
				}
				break;
			case -1:
				return null;
			}
			length++;
		}

		if (deep != 1) {
			err.error("lex: error in `" + name + "', wrong parantheses: /" + regexp + "/\n");
			return null;
		}

		int e;
		if (addbrackets) {
			for (e = 0; e < length; e++) {
				if ((sym[e] & LexConstants.MASK) == LexConstants.LBR) {
					sym[e] = (sym[e] & ~0xffff) | ((sym[e] & 0xffff) + 1);
				}
			}

			sym[length++] = LexConstants.RBR;
		}

		sym[length++] = -1 - number;

		e = 0;
		int[] result = new int[length + (addbrackets ? 1 : 0)];
		if (addbrackets) {
			result[e++] = LexConstants.LBR | (length - 1);
		}
		for (int i = 0; i < length; i++, e++) {
			result[e] = sym[i];
		}

		return result;
	}

	/**
	 * @return array of sets of symbols
	 */
	public int[] getSetpool() {
		int[] result = new int[setpool.size() * LexConstants.SIZE_SYM];
		int e = 0;

		for (Iterator<int[]> it = setpool.iterator(); it.hasNext();) {
			int[] from = it.next();
			for (int i = 0; i < LexConstants.SIZE_SYM; i++) {
				result[e++] = from[i];
			}
		}

		return result;
	}

	/**
	 * @return set of used symbols
	 */
	public int[] getSymbolSet() {
		return symbols;
	}
}
