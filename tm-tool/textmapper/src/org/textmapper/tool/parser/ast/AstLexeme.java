/**
 * Copyright 2002-2012 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textmapper.tool.parser.ast;

import org.textmapper.tool.parser.LapgTree.TextSource;

public class AstLexeme extends AstNode implements AstLexerPart {

	private final AstIdentifier name;
	private final String type;
	private final AstRegexp regexp;
	private final AstReference transition;
	private final AstLexemAttrs attrs;
	private final int priority;
	private final AstCode code;

	public AstLexeme(AstIdentifier name, String type, AstRegexp regexp, AstReference transition,
					 Integer priority, AstLexemAttrs attrs, AstCode code, TextSource source, int offset,
					 int endoffset) {
		super(source, offset, endoffset);
		this.name = name;
		this.type = type;
		this.regexp = regexp;
		this.transition = transition;
		this.attrs = attrs;
		this.priority = priority != null ? priority : 0;
		this.code = code;
	}

	public AstIdentifier getName() {
		return name;
	}

	public String getType() {
		return type;
	}

	public AstRegexp getRegexp() {
		return regexp;
	}

	public AstReference getTransition() {
		return transition;
	}

	public int getPriority() {
		return priority;
	}

	public AstLexemAttrs getAttrs() {
		return attrs;
	}

	public AstCode getCode() {
		return code;
	}

	@Override
	public void accept(AbstractVisitor v) {
		if (!v.visit(this)) {
			return;
		}
		if (name != null) {
			name.accept(v);
		}
		if (regexp != null) {
			regexp.accept(v);
		}
		if (transition != null) {
			transition.accept(v);
		}
		if (attrs != null) {
			attrs.accept(v);
		}
		if (code != null) {
			code.accept(v);
		}
	}
}