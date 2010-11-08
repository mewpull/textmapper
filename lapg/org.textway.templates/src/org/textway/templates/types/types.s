#  syntax: template types definition grammar
#
#  Lapg (Lexer and Parser Generator)
#  Copyright 2002-2010 Evgeny Gryaznov
# 
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

.lang        "java" 
.prefix      "Types"
.package	 "org.textway.templates.types"
.lexemend    "on"
.gentree	 "on"
.genast		 "on"
.astprefix	 "Ast"
.positions   "line,offset"
.endpositions "offset"

# Vocabulary

[0]

identifier(String): /[a-zA-Z_][a-zA-Z_0-9]*|'([^\n\\']|\\.)*'/ -1
			{ $lexem = current(); break; }

scon(String):	/"([^\n\\"]|\\.)*"/		{ $lexem = unescape(current(), 1, token.length()-1); break; }
icon(Integer):	/-?[0-9]+/				{ $lexem = Integer.parseInt(current()); break; }
bcon(Boolean):  /true|false/			{ $lexem = current().equals("true"); break; }

_skip:         /[\n\t\r ]+/    		{ return false; }
_skip:  /#.*/

'..':    /\.\./
'*':    /\*/
';':    /;/
',':	/,/
':':    /:/
'=':    /=/
'{':	/\{/
'}':	/\}/
'(':	/\(/
')':	/\)/
'[':	/\[/
']':	/\]/

Lclass:		 /class/
Lextends:	 /extends/
Lnew:		 /new/

Lint:		 /int/
Lbool:		 /bool/
Lstring:	 /string/

Lset:		 /set/
Lchoice:	 /choice/


# Grammar

input ::=
	declarations ;

declarations ::=
	declarations type_declaration | type_declaration ;

type_declaration ::=
	Lclass name=identifier extends=extends_clauseopt '{' feature_declarations '}'
;

extends_clause ::=
	Lextends @pass identifier ;

##### DECLARATIONS

feature_declarations ::=
	feature_declarations feature_declaration | feature_declaration ;

feature_declaration ::=
	type name=identifier modifiersopt defaultvalopt ';' ;

defaultval ::=
	'=' @pass expression ;

modifiers ::=
	'[' @pass constraints ']' ;

constraints ::=
	constraints ';' constraint | constraint ;

constraint ::=
	string_constraint | multiplicity ;

string_constraint ::=
	kind=Lset ':' strings
  | kind=Lchoice ':' strings
  | identifier
;

strings ::=
	strings ',' string | string ;
	
string ::=
	identifier | scon ;	

multiplicity ::=
	lo=icon
  | lo=icon '..' hasNoUpperBound='*'
  | lo=icon '..' hi=icon
;

##### TYPES

type ::=
	kind=Lint
  | kind=Lstring
  | kind=Lbool
  | identifier
  | identifier isReference='*'
;

##### EXPRESSIONS

expression ::=
	structural_expression | literal_expression ;

literal_expression ::=
	  scon
	| icon
	| bcon
;

structural_expression ::=
	  Lnew identifier '(' map_entries ')'
	| '[' expression_list ']'
;

expression_list ::=
	expression
	| expression_list ',' expression
;

map_entries ::=
	  identifier ':' expression
	| map_entries ',' identifier ':' expression
;

##################################################################################

%%

${template java_lexer.lexercode}
private String unescape(String s, int start, int end) {
	StringBuilder sb = new StringBuilder();
	end = Math.min(end, s.length());
	for(int i = start; i < end; i++) {
		char c = s.charAt(i);
		if(c == '\\') {
			if(++i == end) {
				break;
			}
			c = s.charAt(i);
			if(c == 'u' || c == 'x') {
				// FIXME process unicode
			} else if(c == 'n') {
				sb.append('\n');
			} else if(c == 'r') {
				sb.append('\r');
			} else if(c == 't') {
				sb.append('\t');
			} else {
				sb.append(c);
			}
		} else {
			sb.append(c);
		}
	} 
	return sb.toString();
}
${end}