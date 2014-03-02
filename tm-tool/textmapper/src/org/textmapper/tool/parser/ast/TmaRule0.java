/**
 * Copyright 2002-2014 Evgeny Gryaznov
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

import org.textmapper.tool.parser.TMTree.TextSource;

import java.util.List;

public class TmaRule0 extends TmaNode {

	private final List<ITmaRhsPart> list;
	private final TmaRhsPrefix prefix;
	private final TmaRhsSuffix suffix;
	private final TmaSyntaxProblem error;

	public TmaRule0(TmaRhsPrefix prefix, List<ITmaRhsPart> list, TmaRhsSuffix attr, TextSource source,
					int offset, int endoffset) {
		super(source, offset, endoffset);
		this.list = list;
		this.suffix = attr;
		this.prefix = prefix;
		this.error = null;
	}

	public TmaRule0(TmaSyntaxProblem err) {
		super(err.getSource(), err.getOffset(), err.getEndoffset());
		this.list = null;
		this.suffix = null;
		this.prefix = null;
		this.error = err;
	}

	public boolean hasSyntaxError() {
		return error != null;
	}

	public List<ITmaRhsPart> getList() {
		return list;
	}

	public TmaRhsPrefix getPrefix() {
		return prefix;
	}

	public TmaRhsSuffix getSuffix() {
		return suffix;
	}

	@Deprecated
	public String getAlias() {
		// TODO use getPrefix()
		return prefix != null && prefix.getName() != null ? prefix.getName().getID() : null;
	}

	@Deprecated
	public TmaAnnotations getAnnotations() {
		// TODO -use getPrefix()
		return prefix != null ? prefix.getAnnotations() : null;
	}

	public void accept(TmaVisitor v) {
		if (error != null) {
			v.visit(error);
			return;
		}
		if (!v.visit(this)) {
			return;
		}
		if (prefix != null) {
			prefix.accept(v);
		}
		if (list != null) {
			for (ITmaRhsPart part : list) {
				part.accept(v);
			}
		}
		if (suffix != null) {
			suffix.accept(v);
		}
	}
}
