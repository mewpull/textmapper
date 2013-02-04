/**
 * Copyright 2002-2013 Evgeny Gryaznov
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
package org.textmapper.lapg.api.builder;

import org.textmapper.lapg.api.SourceElement;
import org.textmapper.lapg.api.ast.*;

public interface AstBuilder {

	AstType rawType(String type, SourceElement origin);

	AstList list(AstType inner, boolean nonEmpty, SourceElement origin);

	AstField addField(String name, AstType type, boolean nullable, AstClass container, SourceElement origin);

	AstClass addClass(String name, AstClass container, SourceElement origin);

	AstEnum addEnum(String name, SourceElement origin);

	AstEnumMember addMember(String name, AstEnum container, SourceElement origin);

	void addExtends(AstClass cl, AstClass baseClass);

	AstModel create(SourceElement origin);
}