// generated by Textmapper; DO NOT EDIT

package tm

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	Identifier NodeType = iota + 1
	IntegerLiteral
	StringLiteral
	BooleanLiteral
	Pattern
	Name // (Identifier)+
	Command
	SyntaxProblem
	File          // Header imports=(Import)* options=(Option)* lexer=LexerSection parser=ParserSection?
	Header        // name=Name target=Name?
	LexerSection  // (LexerPart)+
	ParserSection // (GrammarPart)+
	Import        // alias=Identifier? path=StringLiteral
	KeyValue      // key=Identifier value=Expression
	Symref        // name=Identifier args=SymrefArgs?
	RawType
	NamedPattern         // name=Identifier Pattern
	StartConditionsScope // StartConditions (LexerPart)+
	StartConditions      // (Stateref)*
	Lexeme               // StartConditions? name=Identifier RawType? Pattern? priority=IntegerLiteral? attrs=LexemeAttrs? Command?
	LexemeAttrs          // LexemeAttribute
	LexemeAttribute
	DirectiveBrackets // opening=Symref closing=Symref
	InclusiveStates   // states=(LexerState)+
	ExclusiveStates   // states=(LexerState)+
	Stateref          // name=Identifier
	LexerState        // name=Identifier
	Nonterm           // Annotations? name=Identifier params=NontermParams? NontermType? ReportClause? (Rule0)+
	SubType           // reference=Symref
	InterfaceType
	ClassType // Implements?
	VoidType
	Implements // (Symref)+
	Assoc
	ParamModifier
	TemplateParam      // modifier=ParamModifier? ParamType name=Identifier ParamValue?
	DirectivePrio      // Assoc symbols=References
	DirectiveInput     // inputRefs=(Inputref)+
	DirectiveInterface // ids=(Identifier)+
	DirectiveAssert    // RhsSet
	DirectiveSet       // name=Identifier RhsSet
	Inputref           // reference=Symref
	References         // References? Symref
	Rule               // Predicate? (RhsPart)* RhsSuffix? ReportClause?
	Predicate          // PredicateExpression
	RhsSuffix          // Symref
	ReportClause       // action=Identifier kind=Identifier? ReportAs?
	ReportAs           // Identifier
	RhsLookahead       // predicates=(LookaheadPredicate)+
	LookaheadPredicate // Symref
	StateMarker        // name=Identifier
	RhsAnnotated       // Annotations inner=RhsPart
	RhsAssignment      // id=Identifier inner=RhsPart
	RhsPlusAssignment  // id=Identifier inner=RhsPart
	RhsOptional        // inner=RhsPart
	RhsCast            // inner=RhsPart target=Symref
	RhsAsLiteral       // inner=RhsPart Literal
	ListSeparator      // separator_=References
	RhsSymbol          // reference=Symref
	RhsNested          // (Rule0)+
	RhsPlusList        // ruleParts=(RhsPart)+ ListSeparator
	RhsStarList        // ruleParts=(RhsPart)+ ListSeparator
	RhsQuantifier      // inner=RhsPart
	RhsIgnored         // (Rule0)+
	RhsSet             // expr=SetExpression
	SetSymbol          // operator=Identifier? symbol=Symref
	SetCompound        // inner=SetExpression
	SetComplement      // inner=SetExpression
	SetOr              // left=SetExpression right=SetExpression
	SetAnd             // left=SetExpression right=SetExpression
	Annotations        // (Annotation)+
	AnnotationImpl     // name=Identifier Expression?
	NontermParams      // list=(NontermParam)+
	InlineParameter    // param_type=Identifier name=Identifier ParamValue?
	ParamRef           // Identifier
	SymrefArgs         // arg_list=(Argument)*
	ArgumentImpl       // name=ParamRef val=ParamValue?
	ArgumentTrue       // name=ParamRef
	ArgumentFalse      // name=ParamRef
	ParamType
	PredicateNot   // ParamRef
	PredicateEq    // ParamRef Literal
	PredicateNotEq // ParamRef Literal
	PredicateAnd   // left=PredicateExpression right=PredicateExpression
	PredicateOr    // left=PredicateExpression right=PredicateExpression
	Array          // (Expression)*
	InvalidToken
	MultilineComment
	Comment
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
	"Identifier",
	"IntegerLiteral",
	"StringLiteral",
	"BooleanLiteral",
	"Pattern",
	"Name",
	"Command",
	"SyntaxProblem",
	"File",
	"Header",
	"LexerSection",
	"ParserSection",
	"Import",
	"KeyValue",
	"Symref",
	"RawType",
	"NamedPattern",
	"StartConditionsScope",
	"StartConditions",
	"Lexeme",
	"LexemeAttrs",
	"LexemeAttribute",
	"DirectiveBrackets",
	"InclusiveStates",
	"ExclusiveStates",
	"Stateref",
	"LexerState",
	"Nonterm",
	"SubType",
	"InterfaceType",
	"ClassType",
	"VoidType",
	"Implements",
	"Assoc",
	"ParamModifier",
	"TemplateParam",
	"DirectivePrio",
	"DirectiveInput",
	"DirectiveInterface",
	"DirectiveAssert",
	"DirectiveSet",
	"Inputref",
	"References",
	"Rule",
	"Predicate",
	"RhsSuffix",
	"ReportClause",
	"ReportAs",
	"RhsLookahead",
	"LookaheadPredicate",
	"StateMarker",
	"RhsAnnotated",
	"RhsAssignment",
	"RhsPlusAssignment",
	"RhsOptional",
	"RhsCast",
	"RhsAsLiteral",
	"ListSeparator",
	"RhsSymbol",
	"RhsNested",
	"RhsPlusList",
	"RhsStarList",
	"RhsQuantifier",
	"RhsIgnored",
	"RhsSet",
	"SetSymbol",
	"SetCompound",
	"SetComplement",
	"SetOr",
	"SetAnd",
	"Annotations",
	"AnnotationImpl",
	"NontermParams",
	"InlineParameter",
	"ParamRef",
	"SymrefArgs",
	"ArgumentImpl",
	"ArgumentTrue",
	"ArgumentFalse",
	"ParamType",
	"PredicateNot",
	"PredicateEq",
	"PredicateNotEq",
	"PredicateAnd",
	"PredicateOr",
	"Array",
	"InvalidToken",
	"MultilineComment",
	"Comment",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}

var Annotation = []NodeType{
	AnnotationImpl,
	SyntaxProblem,
}

var Argument = []NodeType{
	ArgumentFalse,
	ArgumentImpl,
	ArgumentTrue,
}

var Expression = []NodeType{
	Array,
	BooleanLiteral,
	IntegerLiteral,
	StringLiteral,
	Symref,
	SyntaxProblem,
}

var GrammarPart = []NodeType{
	DirectiveAssert,
	DirectiveInput,
	DirectiveInterface,
	DirectivePrio,
	DirectiveSet,
	Nonterm,
	SyntaxProblem,
	TemplateParam,
}

var LexerPart = []NodeType{
	DirectiveBrackets,
	ExclusiveStates,
	InclusiveStates,
	Lexeme,
	NamedPattern,
	StartConditionsScope,
	SyntaxProblem,
}

var Literal = []NodeType{
	BooleanLiteral,
	IntegerLiteral,
	StringLiteral,
}

var NontermParam = []NodeType{
	InlineParameter,
	ParamRef,
}

var NontermType = []NodeType{
	ClassType,
	InterfaceType,
	RawType,
	SubType,
	VoidType,
}

var Option = []NodeType{
	KeyValue,
	SyntaxProblem,
}

var ParamValue = []NodeType{
	BooleanLiteral,
	IntegerLiteral,
	ParamRef,
	StringLiteral,
}

var PredicateExpression = []NodeType{
	ParamRef,
	PredicateAnd,
	PredicateEq,
	PredicateNot,
	PredicateNotEq,
	PredicateOr,
}

var RhsPart = []NodeType{
	Command,
	RhsAnnotated,
	RhsAsLiteral,
	RhsAssignment,
	RhsCast,
	RhsIgnored,
	RhsLookahead,
	RhsNested,
	RhsOptional,
	RhsPlusAssignment,
	RhsPlusList,
	RhsQuantifier,
	RhsSet,
	RhsStarList,
	RhsSymbol,
	StateMarker,
	SyntaxProblem,
}

var Rule0 = []NodeType{
	Rule,
	SyntaxProblem,
}

var SetExpression = []NodeType{
	SetAnd,
	SetComplement,
	SetCompound,
	SetOr,
	SetSymbol,
}

var ruleNodeType = [...]NodeType{
	Identifier,           // identifier : ID
	Identifier,           // identifier : 'brackets'
	Identifier,           // identifier : 'inline'
	Identifier,           // identifier : 'prec'
	Identifier,           // identifier : 'shift'
	Identifier,           // identifier : 'returns'
	Identifier,           // identifier : 'input'
	Identifier,           // identifier : 'left'
	Identifier,           // identifier : 'right'
	Identifier,           // identifier : 'nonassoc'
	Identifier,           // identifier : 'generate'
	Identifier,           // identifier : 'assert'
	Identifier,           // identifier : 'empty'
	Identifier,           // identifier : 'nonempty'
	Identifier,           // identifier : 'global'
	Identifier,           // identifier : 'explicit'
	Identifier,           // identifier : 'lookahead'
	Identifier,           // identifier : 'param'
	Identifier,           // identifier : 'flag'
	Identifier,           // identifier : 'no-eoi'
	Identifier,           // identifier : 's'
	Identifier,           // identifier : 'x'
	Identifier,           // identifier : 'soft'
	Identifier,           // identifier : 'class'
	Identifier,           // identifier : 'interface'
	Identifier,           // identifier : 'void'
	Identifier,           // identifier : 'space'
	Identifier,           // identifier : 'layout'
	Identifier,           // identifier : 'language'
	Identifier,           // identifier : 'lalr'
	Identifier,           // identifier : 'lexer'
	Identifier,           // identifier : 'parser'
	Identifier,           // identifier_Kw : ID
	Identifier,           // identifier_Kw : 'brackets'
	Identifier,           // identifier_Kw : 'inline'
	Identifier,           // identifier_Kw : 'prec'
	Identifier,           // identifier_Kw : 'shift'
	Identifier,           // identifier_Kw : 'returns'
	Identifier,           // identifier_Kw : 'input'
	Identifier,           // identifier_Kw : 'left'
	Identifier,           // identifier_Kw : 'right'
	Identifier,           // identifier_Kw : 'nonassoc'
	Identifier,           // identifier_Kw : 'generate'
	Identifier,           // identifier_Kw : 'assert'
	Identifier,           // identifier_Kw : 'empty'
	Identifier,           // identifier_Kw : 'nonempty'
	Identifier,           // identifier_Kw : 'global'
	Identifier,           // identifier_Kw : 'explicit'
	Identifier,           // identifier_Kw : 'lookahead'
	Identifier,           // identifier_Kw : 'param'
	Identifier,           // identifier_Kw : 'flag'
	Identifier,           // identifier_Kw : 'no-eoi'
	Identifier,           // identifier_Kw : 's'
	Identifier,           // identifier_Kw : 'x'
	Identifier,           // identifier_Kw : 'soft'
	Identifier,           // identifier_Kw : 'class'
	Identifier,           // identifier_Kw : 'interface'
	Identifier,           // identifier_Kw : 'void'
	Identifier,           // identifier_Kw : 'space'
	Identifier,           // identifier_Kw : 'layout'
	Identifier,           // identifier_Kw : 'language'
	Identifier,           // identifier_Kw : 'lalr'
	Identifier,           // identifier_Kw : 'lexer'
	Identifier,           // identifier_Kw : 'parser'
	Identifier,           // identifier_Kw : 'true'
	Identifier,           // identifier_Kw : 'false'
	Identifier,           // identifier_Kw : 'separator'
	Identifier,           // identifier_Kw : 'as'
	Identifier,           // identifier_Kw : 'import'
	Identifier,           // identifier_Kw : 'set'
	IntegerLiteral,       // integer_literal : icon
	StringLiteral,        // string_literal : scon
	BooleanLiteral,       // boolean_literal : 'true'
	BooleanLiteral,       // boolean_literal : 'false'
	0,                    // literal : string_literal
	0,                    // literal : integer_literal
	0,                    // literal : boolean_literal
	Pattern,              // pattern : regexp
	0,                    // qualified_name : identifier
	0,                    // qualified_name : qualified_name '.' identifier_Kw
	Name,                 // name : qualified_name
	Command,              // command : code
	SyntaxProblem,        // syntax_problem : error
	File,                 // file : header import__optlist option_optlist lexer_section parser_section
	File,                 // file : header import__optlist option_optlist lexer_section
	0,                    // import__optlist : import__optlist import_
	0,                    // import__optlist :
	0,                    // option_optlist : option_optlist option
	0,                    // option_optlist :
	Header,               // header : 'language' name '(' name ')' ';'
	Header,               // header : 'language' name ';'
	LexerSection,         // lexer_section : '::' 'lexer' lexer_parts
	ParserSection,        // parser_section : '::' 'parser' grammar_parts
	Import,               // import_ : 'import' identifier string_literal ';'
	Import,               // import_ : 'import' string_literal ';'
	KeyValue,             // option : identifier '=' expression
	0,                    // option : syntax_problem
	Symref,               // symref : identifier
	Symref,               // symref_Args : identifier symref_args
	Symref,               // symref_Args : identifier
	RawType,              // rawType : code
	0,                    // lexer_parts : lexer_part
	0,                    // lexer_parts : lexer_parts lexer_part_OrSyntaxError
	0,                    // lexer_part : named_pattern
	0,                    // lexer_part : lexeme
	0,                    // lexer_part : lexer_directive
	0,                    // lexer_part : start_conditions_scope
	0,                    // lexer_part_OrSyntaxError : named_pattern
	0,                    // lexer_part_OrSyntaxError : lexeme
	0,                    // lexer_part_OrSyntaxError : lexer_directive
	0,                    // lexer_part_OrSyntaxError : start_conditions_scope
	0,                    // lexer_part_OrSyntaxError : syntax_problem
	NamedPattern,         // named_pattern : identifier '=' pattern
	StartConditionsScope, // start_conditions_scope : start_conditions '{' .recoveryScope lexer_parts '}'
	StartConditions,      // start_conditions : '<' '*' '>'
	StartConditions,      // start_conditions : '<' stateref_list_Comma_separated '>'
	0,                    // stateref_list_Comma_separated : stateref_list_Comma_separated ',' stateref
	0,                    // stateref_list_Comma_separated : stateref
	Lexeme,               // lexeme : start_conditions identifier rawTypeopt ':' pattern integer_literalopt lexeme_attrsopt commandopt
	Lexeme,               // lexeme : start_conditions identifier rawTypeopt ':'
	Lexeme,               // lexeme : identifier rawTypeopt ':' pattern integer_literalopt lexeme_attrsopt commandopt
	Lexeme,               // lexeme : identifier rawTypeopt ':'
	LexemeAttrs,          // lexeme_attrs : '(' lexeme_attribute ')'
	LexemeAttribute,      // lexeme_attribute : 'soft'
	LexemeAttribute,      // lexeme_attribute : 'class'
	LexemeAttribute,      // lexeme_attribute : 'space'
	LexemeAttribute,      // lexeme_attribute : 'layout'
	DirectiveBrackets,    // lexer_directive : '%' 'brackets' symref symref ';'
	InclusiveStates,      // lexer_directive : '%' 's' lexer_state_list_Comma_separated ';'
	ExclusiveStates,      // lexer_directive : '%' 'x' lexer_state_list_Comma_separated ';'
	0,                    // lexer_state_list_Comma_separated : lexer_state_list_Comma_separated ',' lexer_state
	0,                    // lexer_state_list_Comma_separated : lexer_state
	Stateref,             // stateref : identifier
	LexerState,           // lexer_state : identifier
	0,                    // grammar_parts : grammar_part
	0,                    // grammar_parts : grammar_parts grammar_part_OrSyntaxError
	0,                    // grammar_part : nonterm
	0,                    // grammar_part : template_param
	0,                    // grammar_part : directive
	0,                    // grammar_part_OrSyntaxError : nonterm
	0,                    // grammar_part_OrSyntaxError : template_param
	0,                    // grammar_part_OrSyntaxError : directive
	0,                    // grammar_part_OrSyntaxError : syntax_problem
	Nonterm,              // nonterm : annotations identifier nonterm_params nonterm_type reportClause ':' rules ';'
	Nonterm,              // nonterm : annotations identifier nonterm_params nonterm_type ':' rules ';'
	Nonterm,              // nonterm : annotations identifier nonterm_params reportClause ':' rules ';'
	Nonterm,              // nonterm : annotations identifier nonterm_params ':' rules ';'
	Nonterm,              // nonterm : annotations identifier nonterm_type reportClause ':' rules ';'
	Nonterm,              // nonterm : annotations identifier nonterm_type ':' rules ';'
	Nonterm,              // nonterm : annotations identifier reportClause ':' rules ';'
	Nonterm,              // nonterm : annotations identifier ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_params nonterm_type reportClause ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_params nonterm_type ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_params reportClause ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_params ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_type reportClause ':' rules ';'
	Nonterm,              // nonterm : identifier nonterm_type ':' rules ';'
	Nonterm,              // nonterm : identifier reportClause ':' rules ';'
	Nonterm,              // nonterm : identifier ':' rules ';'
	SubType,              // nonterm_type : 'returns' symref
	InterfaceType,        // nonterm_type : 'interface'
	ClassType,            // nonterm_type : 'class' implements_clause
	ClassType,            // nonterm_type : 'class'
	VoidType,             // nonterm_type : 'void'
	0,                    // nonterm_type : rawType
	Implements,           // implements_clause : 'implements' references_cs
	Assoc,                // assoc : 'left'
	Assoc,                // assoc : 'right'
	Assoc,                // assoc : 'nonassoc'
	ParamModifier,        // param_modifier : 'explicit'
	ParamModifier,        // param_modifier : 'global'
	ParamModifier,        // param_modifier : 'lookahead'
	TemplateParam,        // template_param : '%' param_modifier param_type identifier '=' param_value ';'
	TemplateParam,        // template_param : '%' param_modifier param_type identifier ';'
	TemplateParam,        // template_param : '%' param_type identifier '=' param_value ';'
	TemplateParam,        // template_param : '%' param_type identifier ';'
	DirectivePrio,        // directive : '%' assoc references ';'
	DirectiveInput,       // directive : '%' 'input' inputref_list_Comma_separated ';'
	DirectiveInterface,   // directive : '%' 'interface' identifier_list_Comma_separated ';'
	DirectiveAssert,      // directive : '%' 'assert' 'empty' rhsSet ';'
	DirectiveAssert,      // directive : '%' 'assert' 'nonempty' rhsSet ';'
	DirectiveSet,         // directive : '%' 'generate' identifier '=' rhsSet ';'
	0,                    // identifier_list_Comma_separated : identifier_list_Comma_separated ',' identifier
	0,                    // identifier_list_Comma_separated : identifier
	0,                    // inputref_list_Comma_separated : inputref_list_Comma_separated ',' inputref
	0,                    // inputref_list_Comma_separated : inputref
	Inputref,             // inputref : symref 'no-eoi'
	Inputref,             // inputref : symref
	References,           // references : symref
	References,           // references : references symref
	0,                    // references_cs : symref
	0,                    // references_cs : references_cs ',' symref
	0,                    // rules : rule0
	0,                    // rules : rules '|' rule0
	Rule,                 // rule0 : predicate rhsParts rhsSuffixopt reportClause
	Rule,                 // rule0 : predicate rhsParts rhsSuffixopt
	Rule,                 // rule0 : predicate rhsSuffixopt reportClause
	Rule,                 // rule0 : predicate rhsSuffixopt
	Rule,                 // rule0 : rhsParts rhsSuffixopt reportClause
	Rule,                 // rule0 : rhsParts rhsSuffixopt
	Rule,                 // rule0 : rhsSuffixopt reportClause
	Rule,                 // rule0 : rhsSuffixopt
	0,                    // rule0 : syntax_problem
	Predicate,            // predicate : '[' predicate_expression ']'
	RhsSuffix,            // rhsSuffix : '%' 'prec' symref
	RhsSuffix,            // rhsSuffix : '%' 'shift' symref
	ReportClause,         // reportClause : '->' identifier '/' identifier reportAs
	ReportClause,         // reportClause : '->' identifier '/' identifier
	ReportClause,         // reportClause : '->' identifier reportAs
	ReportClause,         // reportClause : '->' identifier
	ReportAs,             // reportAs : 'as' identifier
	0,                    // rhsParts : rhsPart
	0,                    // rhsParts : rhsParts rhsPart_OrSyntaxError
	0,                    // rhsPart : rhsAnnotated
	0,                    // rhsPart : command
	0,                    // rhsPart : rhsStateMarker
	0,                    // rhsPart : rhsLookahead
	0,                    // rhsPart_OrSyntaxError : rhsAnnotated
	0,                    // rhsPart_OrSyntaxError : command
	0,                    // rhsPart_OrSyntaxError : rhsStateMarker
	0,                    // rhsPart_OrSyntaxError : rhsLookahead
	0,                    // rhsPart_OrSyntaxError : syntax_problem
	0,                    // lookahead_predicate_list_And_separated : lookahead_predicate_list_And_separated '&' lookahead_predicate
	0,                    // lookahead_predicate_list_And_separated : lookahead_predicate
	RhsLookahead,         // rhsLookahead : '(?=' lookahead_predicate_list_And_separated ')'
	LookaheadPredicate,   // lookahead_predicate : '!' symref
	LookaheadPredicate,   // lookahead_predicate : symref
	StateMarker,          // rhsStateMarker : '.' identifier
	0,                    // rhsAnnotated : rhsAssignment
	RhsAnnotated,         // rhsAnnotated : annotations rhsAssignment
	0,                    // rhsAssignment : rhsOptional
	RhsAssignment,        // rhsAssignment : identifier '=' rhsOptional
	RhsPlusAssignment,    // rhsAssignment : identifier '+=' rhsOptional
	0,                    // rhsOptional : rhsCast
	RhsOptional,          // rhsOptional : rhsCast '?'
	0,                    // rhsCast : rhsPrimary
	RhsCast,              // rhsCast : rhsPrimary 'as' symref_Args
	RhsAsLiteral,         // rhsCast : rhsPrimary 'as' literal
	ListSeparator,        // listSeparator : 'separator' references
	RhsSymbol,            // rhsPrimary : symref_Args
	RhsNested,            // rhsPrimary : '(' rules ')'
	RhsPlusList,          // rhsPrimary : '(' rhsParts listSeparator ')' '+'
	RhsStarList,          // rhsPrimary : '(' rhsParts listSeparator ')' '*'
	RhsQuantifier,        // rhsPrimary : rhsPrimary '+'
	RhsQuantifier,        // rhsPrimary : rhsPrimary '*'
	RhsIgnored,           // rhsPrimary : '$' '(' rules ')'
	0,                    // rhsPrimary : rhsSet
	RhsSet,               // rhsSet : 'set' '(' setExpression ')'
	SetSymbol,            // setPrimary : identifier symref_Args
	SetSymbol,            // setPrimary : symref_Args
	SetCompound,          // setPrimary : '(' setExpression ')'
	SetComplement,        // setPrimary : '~' setPrimary
	0,                    // setExpression : setPrimary
	SetOr,                // setExpression : setExpression '|' setExpression
	SetAnd,               // setExpression : setExpression '&' setExpression
	0,                    // annotation_list : annotation_list annotation
	0,                    // annotation_list : annotation
	Annotations,          // annotations : annotation_list
	AnnotationImpl,       // annotation : '@' identifier '=' expression
	AnnotationImpl,       // annotation : '@' identifier
	0,                    // annotation : '@' syntax_problem
	0,                    // nonterm_param_list_Comma_separated : nonterm_param_list_Comma_separated ',' nonterm_param
	0,                    // nonterm_param_list_Comma_separated : nonterm_param
	NontermParams,        // nonterm_params : '<' nonterm_param_list_Comma_separated '>'
	0,                    // nonterm_param : param_ref
	InlineParameter,      // nonterm_param : identifier identifier '=' param_value
	InlineParameter,      // nonterm_param : identifier identifier
	ParamRef,             // param_ref : identifier
	0,                    // argument_list_Comma_separated : argument_list_Comma_separated ',' argument
	0,                    // argument_list_Comma_separated : argument
	0,                    // argument_list_Comma_separated_opt : argument_list_Comma_separated
	0,                    // argument_list_Comma_separated_opt :
	SymrefArgs,           // symref_args : '<' argument_list_Comma_separated_opt '>'
	ArgumentImpl,         // argument : param_ref ':' param_value
	ArgumentImpl,         // argument : param_ref
	ArgumentTrue,         // argument : '+' param_ref
	ArgumentFalse,        // argument : '~' param_ref
	ParamType,            // param_type : 'flag'
	ParamType,            // param_type : 'param'
	0,                    // param_value : literal
	0,                    // param_value : param_ref
	0,                    // predicate_primary : param_ref
	PredicateNot,         // predicate_primary : '!' param_ref
	PredicateEq,          // predicate_primary : param_ref '==' literal
	PredicateNotEq,       // predicate_primary : param_ref '!=' literal
	0,                    // predicate_expression : predicate_primary
	PredicateAnd,         // predicate_expression : predicate_expression '&&' predicate_expression
	PredicateOr,          // predicate_expression : predicate_expression '||' predicate_expression
	0,                    // expression : literal
	0,                    // expression : symref_Args
	Array,                // expression : '[' expression_list_Comma_separated_opt ']'
	0,                    // expression : syntax_problem
	0,                    // expression_list_Comma_separated : expression_list_Comma_separated ',' expression
	0,                    // expression_list_Comma_separated : expression
	0,                    // expression_list_Comma_separated_opt : expression_list_Comma_separated
	0,                    // expression_list_Comma_separated_opt :
	0,                    // rawTypeopt : rawType
	0,                    // rawTypeopt :
	0,                    // integer_literalopt : integer_literal
	0,                    // integer_literalopt :
	0,                    // lexeme_attrsopt : lexeme_attrs
	0,                    // lexeme_attrsopt :
	0,                    // commandopt : command
	0,                    // commandopt :
	0,                    // rhsSuffixopt : rhsSuffix
	0,                    // rhsSuffixopt :
}
