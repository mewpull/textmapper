language json(go);

lang = "json"

:: lexer

'{': /\{/
'}': /\}/
'[': /\[/
']': /\]/
':': /:/
',': /,/

space: /[\t\r\n ]+/ (space)

hex = /[0-9a-fA-F]/

# TODO
JSONString(string): /"([^"\\]|\\(["\/\\bfnrt]|u{hex}{4}))*"/		{ $$ = l.Text() }
#JSONString: /"([^"\\\x00-\x1f]|\\(["\/\\bfnrt]|u{hex}{4}))*"/

fraction = /\.[0-9]+/
exp = /[eE][+-]?[0-9]+/
JSONNumber: /-?(0|[1-9][0-9]*){fraction}?{exp}?/

id: /[a-zA-Z][a-zA-Z0-9]*/ (class)

'null': /null/
'true': /true/
'false': /false/

error:

:: parser

%input JSONText;

JSONText ::=
	  JSONValue ;

JSONValue (Value) ::=
	  'null'						{ $$ = &Literal{value: "null"} }
	| 'true'
	| 'false'
	| JSONObject
	| JSONArray
	| JSONString
	| JSONNumber
;

JSONObject ::=
	  '{' JSONMemberList? '}' ;

JSONMember (*Field) ::=
	  JSONString ':' JSONValue		{ $$ = &Field{name: $JSONString} } ;

JSONMemberList ::=
	  JSONMember
	| JSONMemberList ',' JSONMember
;

JSONArray ::=
	  '[' JSONElementList? ']' ;

JSONElementList ::=
	  JSONValue
	| JSONElementList ',' JSONValue
;