/*************************************************************
 * Copyright (c) 2002-2009 Evgeny Gryaznov
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Evgeny Gryaznov - initial API and implementation
 *************************************************************/
package net.sf.lapg.test.cases;

import net.sf.lapg.LexerTables;
import net.sf.lapg.ParserTables;
import net.sf.lapg.common.FormatUtil;

public class OutputUtils {

	public static void print(StringBuffer sb, int[] table, int maxwidth, int leftpadding ) {
		for( int i = 0; i < table.length; i++ ) {
			if( i > 0 ) {
				if( (i%maxwidth) == 0 ) {
					sb.append("\n");
					for( int e = 0; e < leftpadding; e++) {
						sb.append("\t");
					}
				} else {
					sb.append(" ");
				}
			}
			sb.append(table[i]);
			sb.append(",");
		}
	}

	public static void print(StringBuffer sb, short[] table, int maxwidth, int leftpadding ) {
		for( int i = 0; i < table.length; i++ ) {
			if( i > 0 ) {
				if( (i%maxwidth) == 0 ) {
					sb.append("\n");
					for( int e = 0; e < leftpadding; e++) {
						sb.append("\t");
					}
				} else {
					sb.append(" ");
				}
			}
			sb.append(table[i]);
			sb.append(",");
		}
	}

	public static void print(StringBuffer sb, int[][] table, int leftpadding, char startrow, char endrow ) {
		for( int i = 0; i < table.length; i++ ) {
			if( i > 0 ) {
				for( int e = 0; e < leftpadding; e++) {
					sb.append("\t");
				}
			}
			sb.append(startrow);
			sb.append(" ");
			int[] row = table[i];
			for( int e = 0; e < row.length; e++ ) {
				sb.append(row[e]);
				sb.append(", ");
			}
			sb.append(endrow);
			sb.append(",\n");
		}
	}

	public static void printTables(StringBuffer sb, LexerTables lt) {

		sb.append("// tables.txt\n\n");

		sb.append("lapg_char2no = {\n\t");
		print(sb, lt.char2no, 16, 1);
		sb.append("\n}\n\n");

		sb.append("lapg_lexem ["+lt.nstates+","+lt.nchars+"] = {\n\t");
		print(sb, lt.change, 1, '{', '}');
		sb.append("}\n\n");
	}

	public static String toIdentifier(String s, int number) {

		if( s.startsWith("\'") && s.endsWith("\'")) {
			StringBuffer res = new StringBuffer();
			String inner = s.substring(1, s.length()-1);
			for( int i = 0; i < inner.length(); i++ ) {
				int c = inner.charAt(i);
				if( c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' ) {
					res.append((char)c);
				} else {
					String name;
					switch( c ) {
						case '{': name = "LBRACE"; break;
						case '}': name = "RBRACE"; break;
						case '[': name = "LBRACKET"; break;
						case ']': name = "RBRACKET"; break;
						case '(': name = "LROUNDBRACKET"; break;
						case ')': name = "RROUNDBRACKET"; break;
						case '.': name = "DOT"; break;
						case ',': name = "COMMA"; break;
						case ':': name = "COLON"; break;
						case ';': name = "SEMICOLON"; break;
						case '+': name = "PLUS"; break;
						case '-': name = "MINUS"; break;
						case '*': name = "MULT"; break;
						case '/': name = "DIV"; break;
						case '%': name = "PERC"; break;
						case '&': name = "AMP"; break;
						case '|': name = "OR"; break;
						case '^': name = "XOR"; break;
						case '!': name = "EXCL"; break;
						case '~': name = "TILDE"; break;
						case '=': name = "EQ"; break;
						case '<': name = "LESS"; break;
						case '>': name = "GREATER"; break;
						case '?': name = "QUESTMARK"; break;
						default: name = "N" + FormatUtil.asHex(c, 2);break;
					}
					res.append(name);
				}
			}

			return res.toString();
		} else if( s.equals("{}") ) {
			return "_sym" + number;
		} else {
			return s;
		}
	}

	public static void printTables(StringBuffer sb, ParserTables pt) {
		sb.append("lapg_action ["+pt.nstates+"] = {\n\t");
		print(sb, pt.action_index, 16, 1);
		sb.append("\n}\n\n");

		if( pt.nactions > 0 ) {
			sb.append("lapg_lalr ["+pt.nactions+"] = {\n\t");
			print(sb, pt.action_table, 16, 1);
			sb.append("\n}\n\n");
		}

		sb.append("lapg_sym_goto ["+pt.nsyms+"+1] = {\n\t");
		print(sb, pt.sym_goto, 16, 1);
		sb.append("\n}\n\n");

		sb.append("lapg_sym_from ["+pt.sym_goto[pt.nsyms]+"] = {\n\t");
		print(sb, pt.sym_from, 16, 1);
		sb.append("\n}\n\n");

		sb.append("lapg_sym_to ["+pt.sym_goto[pt.nsyms]+"] = {\n\t");
		print(sb, pt.sym_to, 16, 1);
		sb.append("\n}\n\n");

		sb.append("lapg_rlen ["+pt.rules+"] = {\n\t");
		for( int i = 0; i < pt.rules; i++ ) {
			if( i > 0 ) {
				if( (i%16) == 0 ) {
					sb.append("\n\t");
				} else {
					sb.append(" ");
				}
			}
			int e = 0;
			for(; pt.rright[ pt.rindex[i]+e ] >= 0; e++) {
				;
			}
			sb.append(e);
			sb.append(",");
		}
		sb.append("\n}\n\n");

		sb.append("lapg_rlex ["+pt.rules+"] = {\n\t");
		print(sb, pt.rleft, 16, 1);
		sb.append("\n}\n\n");

		sb.append("lapg_syms = {\n");
		for( int i = 0; i < pt.nsyms; i++ ) {
			sb.append("\t\"");
			sb.append(pt.sym[i].getName());
			sb.append("\",\n");
		}
		sb.append("}\n\n");

		sb.append("Tokens = {\n");
		for( int i = 0; i < pt.nsyms; i++ ) {
			sb.append("\t");
			sb.append(toIdentifier(pt.sym[i].getName(), i));
			sb.append(",\n");
		}
		sb.append("}\n\n");

	}
}