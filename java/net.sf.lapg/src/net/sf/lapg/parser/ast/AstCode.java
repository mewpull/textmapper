package net.sf.lapg.parser.ast;

import net.sf.lapg.parser.TextSource;

public class AstCode extends Node {

	public AstCode(TextSource source, int offset, int endoffset) {
		super(source, offset, endoffset);
	}
}
