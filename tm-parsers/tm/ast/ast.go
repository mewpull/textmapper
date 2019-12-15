// generated by Textmapper; DO NOT EDIT

package ast

import (
	"github.com/inspirer/textmapper/tm-parsers/tm/selector"
)

// Interfaces.

type TmNode interface {
	TmNode() *Node
}

type Token struct {
	*Node
}

type NilNode struct{}

var nilInstance = &NilNode{}

// All types implement TmNode.
func (n AnnotationImpl) TmNode() *Node       { return n.Node }
func (n Annotations) TmNode() *Node          { return n.Node }
func (n ArgumentFalse) TmNode() *Node        { return n.Node }
func (n ArgumentTrue) TmNode() *Node         { return n.Node }
func (n ArgumentVal) TmNode() *Node          { return n.Node }
func (n Array) TmNode() *Node                { return n.Node }
func (n Assoc) TmNode() *Node                { return n.Node }
func (n BooleanLiteral) TmNode() *Node       { return n.Node }
func (n ClassType) TmNode() *Node            { return n.Node }
func (n Command) TmNode() *Node              { return n.Node }
func (n DirectiveAssert) TmNode() *Node      { return n.Node }
func (n DirectiveBrackets) TmNode() *Node    { return n.Node }
func (n DirectiveInput) TmNode() *Node       { return n.Node }
func (n DirectiveInterface) TmNode() *Node   { return n.Node }
func (n DirectivePrio) TmNode() *Node        { return n.Node }
func (n DirectiveSet) TmNode() *Node         { return n.Node }
func (n ExclusiveStartConds) TmNode() *Node  { return n.Node }
func (n File) TmNode() *Node                 { return n.Node }
func (n Header) TmNode() *Node               { return n.Node }
func (n Identifier) TmNode() *Node           { return n.Node }
func (n Implements) TmNode() *Node           { return n.Node }
func (n Import) TmNode() *Node               { return n.Node }
func (n InclusiveStartConds) TmNode() *Node  { return n.Node }
func (n InlineParameter) TmNode() *Node      { return n.Node }
func (n Inputref) TmNode() *Node             { return n.Node }
func (n IntegerLiteral) TmNode() *Node       { return n.Node }
func (n InterfaceType) TmNode() *Node        { return n.Node }
func (n KeyValue) TmNode() *Node             { return n.Node }
func (n Lexeme) TmNode() *Node               { return n.Node }
func (n LexemeAttribute) TmNode() *Node      { return n.Node }
func (n LexemeAttrs) TmNode() *Node          { return n.Node }
func (n LexerSection) TmNode() *Node         { return n.Node }
func (n LexerState) TmNode() *Node           { return n.Node }
func (n ListSeparator) TmNode() *Node        { return n.Node }
func (n LookaheadPredicate) TmNode() *Node   { return n.Node }
func (n NamedPattern) TmNode() *Node         { return n.Node }
func (n NoEoi) TmNode() *Node                { return n.Node }
func (n Nonterm) TmNode() *Node              { return n.Node }
func (n NontermParams) TmNode() *Node        { return n.Node }
func (n ParamModifier) TmNode() *Node        { return n.Node }
func (n ParamRef) TmNode() *Node             { return n.Node }
func (n ParamType) TmNode() *Node            { return n.Node }
func (n ParserSection) TmNode() *Node        { return n.Node }
func (n Pattern) TmNode() *Node              { return n.Node }
func (n Predicate) TmNode() *Node            { return n.Node }
func (n PredicateAnd) TmNode() *Node         { return n.Node }
func (n PredicateEq) TmNode() *Node          { return n.Node }
func (n PredicateNot) TmNode() *Node         { return n.Node }
func (n PredicateNotEq) TmNode() *Node       { return n.Node }
func (n PredicateOr) TmNode() *Node          { return n.Node }
func (n RawType) TmNode() *Node              { return n.Node }
func (n ReportAs) TmNode() *Node             { return n.Node }
func (n ReportClause) TmNode() *Node         { return n.Node }
func (n RhsAnnotated) TmNode() *Node         { return n.Node }
func (n RhsAsLiteral) TmNode() *Node         { return n.Node }
func (n RhsAssignment) TmNode() *Node        { return n.Node }
func (n RhsCast) TmNode() *Node              { return n.Node }
func (n RhsIgnored) TmNode() *Node           { return n.Node }
func (n RhsLookahead) TmNode() *Node         { return n.Node }
func (n RhsNested) TmNode() *Node            { return n.Node }
func (n RhsOptional) TmNode() *Node          { return n.Node }
func (n RhsPlusAssignment) TmNode() *Node    { return n.Node }
func (n RhsPlusList) TmNode() *Node          { return n.Node }
func (n RhsQuantifier) TmNode() *Node        { return n.Node }
func (n RhsSet) TmNode() *Node               { return n.Node }
func (n RhsStarList) TmNode() *Node          { return n.Node }
func (n RhsSuffix) TmNode() *Node            { return n.Node }
func (n RhsSymbol) TmNode() *Node            { return n.Node }
func (n Rule) TmNode() *Node                 { return n.Node }
func (n SetAnd) TmNode() *Node               { return n.Node }
func (n SetComplement) TmNode() *Node        { return n.Node }
func (n SetCompound) TmNode() *Node          { return n.Node }
func (n SetOr) TmNode() *Node                { return n.Node }
func (n SetSymbol) TmNode() *Node            { return n.Node }
func (n StartConditions) TmNode() *Node      { return n.Node }
func (n StartConditionsScope) TmNode() *Node { return n.Node }
func (n StateMarker) TmNode() *Node          { return n.Node }
func (n Stateref) TmNode() *Node             { return n.Node }
func (n StringLiteral) TmNode() *Node        { return n.Node }
func (n SubType) TmNode() *Node              { return n.Node }
func (n Symref) TmNode() *Node               { return n.Node }
func (n SymrefArgs) TmNode() *Node           { return n.Node }
func (n SyntaxProblem) TmNode() *Node        { return n.Node }
func (n TemplateParam) TmNode() *Node        { return n.Node }
func (n VoidType) TmNode() *Node             { return n.Node }
func (n Token) TmNode() *Node                { return n.Node }
func (NilNode) TmNode() *Node                { return nil }

type Annotation interface {
	TmNode
	annotationNode()
}

// annotationNode() ensures that only the following types can be
// assigned to Annotation.
//
func (AnnotationImpl) annotationNode() {}
func (SyntaxProblem) annotationNode()  {}
func (NilNode) annotationNode()        {}

type Argument interface {
	TmNode
	argumentNode()
}

// argumentNode() ensures that only the following types can be
// assigned to Argument.
//
func (ArgumentFalse) argumentNode() {}
func (ArgumentTrue) argumentNode()  {}
func (ArgumentVal) argumentNode()   {}
func (NilNode) argumentNode()       {}

type Expression interface {
	TmNode
	expressionNode()
}

// expressionNode() ensures that only the following types can be
// assigned to Expression.
//
func (Array) expressionNode()          {}
func (BooleanLiteral) expressionNode() {}
func (IntegerLiteral) expressionNode() {}
func (StringLiteral) expressionNode()  {}
func (Symref) expressionNode()         {}
func (SyntaxProblem) expressionNode()  {}
func (NilNode) expressionNode()        {}

type GrammarPart interface {
	TmNode
	grammarPartNode()
}

// grammarPartNode() ensures that only the following types can be
// assigned to GrammarPart.
//
func (DirectiveAssert) grammarPartNode()    {}
func (DirectiveInput) grammarPartNode()     {}
func (DirectiveInterface) grammarPartNode() {}
func (DirectivePrio) grammarPartNode()      {}
func (DirectiveSet) grammarPartNode()       {}
func (Nonterm) grammarPartNode()            {}
func (SyntaxProblem) grammarPartNode()      {}
func (TemplateParam) grammarPartNode()      {}
func (NilNode) grammarPartNode()            {}

type LexerPart interface {
	TmNode
	lexerPartNode()
}

// lexerPartNode() ensures that only the following types can be
// assigned to LexerPart.
//
func (DirectiveBrackets) lexerPartNode()    {}
func (ExclusiveStartConds) lexerPartNode()  {}
func (InclusiveStartConds) lexerPartNode()  {}
func (Lexeme) lexerPartNode()               {}
func (NamedPattern) lexerPartNode()         {}
func (StartConditionsScope) lexerPartNode() {}
func (SyntaxProblem) lexerPartNode()        {}
func (NilNode) lexerPartNode()              {}

type Literal interface {
	TmNode
	literalNode()
}

// literalNode() ensures that only the following types can be
// assigned to Literal.
//
func (BooleanLiteral) literalNode() {}
func (IntegerLiteral) literalNode() {}
func (StringLiteral) literalNode()  {}
func (NilNode) literalNode()        {}

type NontermParam interface {
	TmNode
	nontermParamNode()
}

// nontermParamNode() ensures that only the following types can be
// assigned to NontermParam.
//
func (InlineParameter) nontermParamNode() {}
func (ParamRef) nontermParamNode()        {}
func (NilNode) nontermParamNode()         {}

type NontermType interface {
	TmNode
	nontermTypeNode()
}

// nontermTypeNode() ensures that only the following types can be
// assigned to NontermType.
//
func (ClassType) nontermTypeNode()     {}
func (InterfaceType) nontermTypeNode() {}
func (RawType) nontermTypeNode()       {}
func (SubType) nontermTypeNode()       {}
func (VoidType) nontermTypeNode()      {}
func (NilNode) nontermTypeNode()       {}

type Option interface {
	TmNode
	optionNode()
}

// optionNode() ensures that only the following types can be
// assigned to Option.
//
func (KeyValue) optionNode()      {}
func (SyntaxProblem) optionNode() {}
func (NilNode) optionNode()       {}

type ParamValue interface {
	TmNode
	paramValueNode()
}

// paramValueNode() ensures that only the following types can be
// assigned to ParamValue.
//
func (BooleanLiteral) paramValueNode() {}
func (IntegerLiteral) paramValueNode() {}
func (ParamRef) paramValueNode()       {}
func (StringLiteral) paramValueNode()  {}
func (NilNode) paramValueNode()        {}

type PredicateExpression interface {
	TmNode
	predicateExpressionNode()
}

// predicateExpressionNode() ensures that only the following types can be
// assigned to PredicateExpression.
//
func (ParamRef) predicateExpressionNode()       {}
func (PredicateAnd) predicateExpressionNode()   {}
func (PredicateEq) predicateExpressionNode()    {}
func (PredicateNot) predicateExpressionNode()   {}
func (PredicateNotEq) predicateExpressionNode() {}
func (PredicateOr) predicateExpressionNode()    {}
func (NilNode) predicateExpressionNode()        {}

type RhsPart interface {
	TmNode
	rhsPartNode()
}

// rhsPartNode() ensures that only the following types can be
// assigned to RhsPart.
//
func (Command) rhsPartNode()           {}
func (RhsAnnotated) rhsPartNode()      {}
func (RhsAsLiteral) rhsPartNode()      {}
func (RhsAssignment) rhsPartNode()     {}
func (RhsCast) rhsPartNode()           {}
func (RhsIgnored) rhsPartNode()        {}
func (RhsLookahead) rhsPartNode()      {}
func (RhsNested) rhsPartNode()         {}
func (RhsOptional) rhsPartNode()       {}
func (RhsPlusAssignment) rhsPartNode() {}
func (RhsPlusList) rhsPartNode()       {}
func (RhsQuantifier) rhsPartNode()     {}
func (RhsSet) rhsPartNode()            {}
func (RhsStarList) rhsPartNode()       {}
func (RhsSymbol) rhsPartNode()         {}
func (StateMarker) rhsPartNode()       {}
func (SyntaxProblem) rhsPartNode()     {}
func (NilNode) rhsPartNode()           {}

type Rule0 interface {
	TmNode
	rule0Node()
}

// rule0Node() ensures that only the following types can be
// assigned to Rule0.
//
func (Rule) rule0Node()          {}
func (SyntaxProblem) rule0Node() {}
func (NilNode) rule0Node()       {}

type SetExpression interface {
	TmNode
	setExpressionNode()
}

// setExpressionNode() ensures that only the following types can be
// assigned to SetExpression.
//
func (SetAnd) setExpressionNode()        {}
func (SetComplement) setExpressionNode() {}
func (SetCompound) setExpressionNode()   {}
func (SetOr) setExpressionNode()         {}
func (SetSymbol) setExpressionNode()     {}
func (NilNode) setExpressionNode()       {}

// Types.

type AnnotationImpl struct {
	*Node
}

func (n AnnotationImpl) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n AnnotationImpl) Expression() (Expression, bool) {
	field := ToTmNode(n.Child(selector.Expression)).(Expression)
	return field, field.TmNode() != nil
}

type Annotations struct {
	*Node
}

func (n Annotations) Annotation() []Annotation {
	nodes := n.Children(selector.Annotation)
	var ret = make([]Annotation, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Annotation))
	}
	return ret
}

type ArgumentFalse struct {
	*Node
}

func (n ArgumentFalse) Name() ParamRef {
	return ParamRef{n.Child(selector.ParamRef)}
}

type ArgumentTrue struct {
	*Node
}

func (n ArgumentTrue) Name() ParamRef {
	return ParamRef{n.Child(selector.ParamRef)}
}

type ArgumentVal struct {
	*Node
}

func (n ArgumentVal) Name() ParamRef {
	return ParamRef{n.Child(selector.ParamRef)}
}

func (n ArgumentVal) Val() (ParamValue, bool) {
	field := ToTmNode(n.Child(selector.ParamRef).Next(selector.ParamValue)).(ParamValue)
	return field, field.TmNode() != nil
}

type Array struct {
	*Node
}

func (n Array) Expression() []Expression {
	nodes := n.Children(selector.Expression)
	var ret = make([]Expression, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Expression))
	}
	return ret
}

type Assoc struct {
	*Node
}

type BooleanLiteral struct {
	*Node
}

type ClassType struct {
	*Node
}

func (n ClassType) Implements() (Implements, bool) {
	field := Implements{n.Child(selector.Implements)}
	return field, field.IsValid()
}

type Command struct {
	*Node
}

type DirectiveAssert struct {
	*Node
}

func (n DirectiveAssert) RhsSet() RhsSet {
	return RhsSet{n.Child(selector.RhsSet)}
}

type DirectiveBrackets struct {
	*Node
}

func (n DirectiveBrackets) Opening() Symref {
	return Symref{n.Child(selector.Symref)}
}

func (n DirectiveBrackets) Closing() Symref {
	return Symref{n.Child(selector.Symref).Next(selector.Symref)}
}

type DirectiveInput struct {
	*Node
}

func (n DirectiveInput) InputRefs() []Inputref {
	nodes := n.Children(selector.Inputref)
	var ret = make([]Inputref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Inputref{node})
	}
	return ret
}

type DirectiveInterface struct {
	*Node
}

func (n DirectiveInterface) Ids() []Identifier {
	nodes := n.Children(selector.Identifier)
	var ret = make([]Identifier, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Identifier{node})
	}
	return ret
}

type DirectivePrio struct {
	*Node
}

func (n DirectivePrio) Assoc() Assoc {
	return Assoc{n.Child(selector.Assoc)}
}

func (n DirectivePrio) Symbols() []Symref {
	nodes := n.Children(selector.Symref)
	var ret = make([]Symref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Symref{node})
	}
	return ret
}

type DirectiveSet struct {
	*Node
}

func (n DirectiveSet) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n DirectiveSet) RhsSet() RhsSet {
	return RhsSet{n.Child(selector.RhsSet)}
}

type ExclusiveStartConds struct {
	*Node
}

func (n ExclusiveStartConds) States() []LexerState {
	nodes := n.Children(selector.LexerState)
	var ret = make([]LexerState, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, LexerState{node})
	}
	return ret
}

type File struct {
	*Node
}

func (n File) Header() Header {
	return Header{n.Child(selector.Header)}
}

func (n File) Imports() []Import {
	nodes := n.Children(selector.Import)
	var ret = make([]Import, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Import{node})
	}
	return ret
}

func (n File) Options() []Option {
	nodes := n.Children(selector.Option)
	var ret = make([]Option, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Option))
	}
	return ret
}

func (n File) Lexer() (LexerSection, bool) {
	field := LexerSection{n.Child(selector.LexerSection)}
	return field, field.IsValid()
}

func (n File) Parser() (ParserSection, bool) {
	field := ParserSection{n.Child(selector.ParserSection)}
	return field, field.IsValid()
}

type Header struct {
	*Node
}

func (n Header) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n Header) Target() (Identifier, bool) {
	field := Identifier{n.Child(selector.Identifier).Next(selector.Identifier)}
	return field, field.IsValid()
}

type Identifier struct {
	*Node
}

type Implements struct {
	*Node
}

func (n Implements) Symref() []Symref {
	nodes := n.Children(selector.Symref)
	var ret = make([]Symref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Symref{node})
	}
	return ret
}

type Import struct {
	*Node
}

func (n Import) Alias() (Identifier, bool) {
	field := Identifier{n.Child(selector.Identifier)}
	return field, field.IsValid()
}

func (n Import) Path() StringLiteral {
	return StringLiteral{n.Child(selector.StringLiteral)}
}

type InclusiveStartConds struct {
	*Node
}

func (n InclusiveStartConds) States() []LexerState {
	nodes := n.Children(selector.LexerState)
	var ret = make([]LexerState, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, LexerState{node})
	}
	return ret
}

type InlineParameter struct {
	*Node
}

func (n InlineParameter) ParamType() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n InlineParameter) Name() Identifier {
	return Identifier{n.Child(selector.Identifier).Next(selector.Identifier)}
}

func (n InlineParameter) ParamValue() (ParamValue, bool) {
	field := ToTmNode(n.Child(selector.ParamValue)).(ParamValue)
	return field, field.TmNode() != nil
}

type Inputref struct {
	*Node
}

func (n Inputref) Reference() Symref {
	return Symref{n.Child(selector.Symref)}
}

func (n Inputref) NoEoi() (NoEoi, bool) {
	field := NoEoi{n.Child(selector.NoEoi)}
	return field, field.IsValid()
}

type IntegerLiteral struct {
	*Node
}

type InterfaceType struct {
	*Node
}

type KeyValue struct {
	*Node
}

func (n KeyValue) Key() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n KeyValue) Value() Expression {
	return ToTmNode(n.Child(selector.Expression)).(Expression)
}

type Lexeme struct {
	*Node
}

func (n Lexeme) StartConditions() (StartConditions, bool) {
	field := StartConditions{n.Child(selector.StartConditions)}
	return field, field.IsValid()
}

func (n Lexeme) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n Lexeme) RawType() (RawType, bool) {
	field := RawType{n.Child(selector.RawType)}
	return field, field.IsValid()
}

func (n Lexeme) Pattern() (Pattern, bool) {
	field := Pattern{n.Child(selector.Pattern)}
	return field, field.IsValid()
}

func (n Lexeme) Priority() (IntegerLiteral, bool) {
	field := IntegerLiteral{n.Child(selector.IntegerLiteral)}
	return field, field.IsValid()
}

func (n Lexeme) Attrs() (LexemeAttrs, bool) {
	field := LexemeAttrs{n.Child(selector.LexemeAttrs)}
	return field, field.IsValid()
}

func (n Lexeme) Command() (Command, bool) {
	field := Command{n.Child(selector.Command)}
	return field, field.IsValid()
}

type LexemeAttribute struct {
	*Node
}

type LexemeAttrs struct {
	*Node
}

func (n LexemeAttrs) LexemeAttribute() LexemeAttribute {
	return LexemeAttribute{n.Child(selector.LexemeAttribute)}
}

type LexerSection struct {
	*Node
}

func (n LexerSection) LexerPart() []LexerPart {
	nodes := n.Children(selector.LexerPart)
	var ret = make([]LexerPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(LexerPart))
	}
	return ret
}

type LexerState struct {
	*Node
}

func (n LexerState) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

type ListSeparator struct {
	*Node
}

func (n ListSeparator) Separator() []Symref {
	nodes := n.Children(selector.Symref)
	var ret = make([]Symref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Symref{node})
	}
	return ret
}

type LookaheadPredicate struct {
	*Node
}

func (n LookaheadPredicate) Symref() Symref {
	return Symref{n.Child(selector.Symref)}
}

type NamedPattern struct {
	*Node
}

func (n NamedPattern) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n NamedPattern) Pattern() Pattern {
	return Pattern{n.Child(selector.Pattern)}
}

type NoEoi struct {
	*Node
}

type Nonterm struct {
	*Node
}

func (n Nonterm) Annotations() (Annotations, bool) {
	field := Annotations{n.Child(selector.Annotations)}
	return field, field.IsValid()
}

func (n Nonterm) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n Nonterm) Params() (NontermParams, bool) {
	field := NontermParams{n.Child(selector.NontermParams)}
	return field, field.IsValid()
}

func (n Nonterm) NontermType() (NontermType, bool) {
	field := ToTmNode(n.Child(selector.NontermType)).(NontermType)
	return field, field.TmNode() != nil
}

func (n Nonterm) ReportClause() (ReportClause, bool) {
	field := ReportClause{n.Child(selector.ReportClause)}
	return field, field.IsValid()
}

func (n Nonterm) Rule0() []Rule0 {
	nodes := n.Children(selector.Rule0)
	var ret = make([]Rule0, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Rule0))
	}
	return ret
}

type NontermParams struct {
	*Node
}

func (n NontermParams) List() []NontermParam {
	nodes := n.Children(selector.NontermParam)
	var ret = make([]NontermParam, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(NontermParam))
	}
	return ret
}

type ParamModifier struct {
	*Node
}

type ParamRef struct {
	*Node
}

func (n ParamRef) Identifier() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

type ParamType struct {
	*Node
}

type ParserSection struct {
	*Node
}

func (n ParserSection) GrammarPart() []GrammarPart {
	nodes := n.Children(selector.GrammarPart)
	var ret = make([]GrammarPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(GrammarPart))
	}
	return ret
}

type Pattern struct {
	*Node
}

type Predicate struct {
	*Node
}

func (n Predicate) PredicateExpression() PredicateExpression {
	return ToTmNode(n.Child(selector.PredicateExpression)).(PredicateExpression)
}

type PredicateAnd struct {
	*Node
}

func (n PredicateAnd) Left() PredicateExpression {
	return ToTmNode(n.Child(selector.PredicateExpression)).(PredicateExpression)
}

func (n PredicateAnd) Right() PredicateExpression {
	return ToTmNode(n.Child(selector.PredicateExpression).Next(selector.PredicateExpression)).(PredicateExpression)
}

type PredicateEq struct {
	*Node
}

func (n PredicateEq) ParamRef() ParamRef {
	return ParamRef{n.Child(selector.ParamRef)}
}

func (n PredicateEq) Literal() Literal {
	return ToTmNode(n.Child(selector.Literal)).(Literal)
}

type PredicateNot struct {
	*Node
}

func (n PredicateNot) ParamRef() ParamRef {
	return ParamRef{n.Child(selector.ParamRef)}
}

type PredicateNotEq struct {
	*Node
}

func (n PredicateNotEq) ParamRef() ParamRef {
	return ParamRef{n.Child(selector.ParamRef)}
}

func (n PredicateNotEq) Literal() Literal {
	return ToTmNode(n.Child(selector.Literal)).(Literal)
}

type PredicateOr struct {
	*Node
}

func (n PredicateOr) Left() PredicateExpression {
	return ToTmNode(n.Child(selector.PredicateExpression)).(PredicateExpression)
}

func (n PredicateOr) Right() PredicateExpression {
	return ToTmNode(n.Child(selector.PredicateExpression).Next(selector.PredicateExpression)).(PredicateExpression)
}

type RawType struct {
	*Node
}

type ReportAs struct {
	*Node
}

func (n ReportAs) Identifier() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

type ReportClause struct {
	*Node
}

func (n ReportClause) Action() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n ReportClause) Kind() (Identifier, bool) {
	field := Identifier{n.Child(selector.Identifier).Next(selector.Identifier)}
	return field, field.IsValid()
}

func (n ReportClause) ReportAs() (ReportAs, bool) {
	field := ReportAs{n.Child(selector.ReportAs)}
	return field, field.IsValid()
}

type RhsAnnotated struct {
	*Node
}

func (n RhsAnnotated) Annotations() Annotations {
	return Annotations{n.Child(selector.Annotations)}
}

func (n RhsAnnotated) Inner() RhsPart {
	return ToTmNode(n.Child(selector.RhsPart)).(RhsPart)
}

type RhsAsLiteral struct {
	*Node
}

func (n RhsAsLiteral) Inner() RhsPart {
	return ToTmNode(n.Child(selector.RhsPart)).(RhsPart)
}

func (n RhsAsLiteral) Literal() Literal {
	return ToTmNode(n.Child(selector.Literal)).(Literal)
}

type RhsAssignment struct {
	*Node
}

func (n RhsAssignment) Id() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n RhsAssignment) Inner() RhsPart {
	return ToTmNode(n.Child(selector.RhsPart)).(RhsPart)
}

type RhsCast struct {
	*Node
}

func (n RhsCast) Inner() RhsPart {
	return ToTmNode(n.Child(selector.RhsPart)).(RhsPart)
}

func (n RhsCast) Target() Symref {
	return Symref{n.Child(selector.Symref)}
}

type RhsIgnored struct {
	*Node
}

func (n RhsIgnored) Rule0() []Rule0 {
	nodes := n.Children(selector.Rule0)
	var ret = make([]Rule0, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Rule0))
	}
	return ret
}

type RhsLookahead struct {
	*Node
}

func (n RhsLookahead) Predicates() []LookaheadPredicate {
	nodes := n.Children(selector.LookaheadPredicate)
	var ret = make([]LookaheadPredicate, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, LookaheadPredicate{node})
	}
	return ret
}

type RhsNested struct {
	*Node
}

func (n RhsNested) Rule0() []Rule0 {
	nodes := n.Children(selector.Rule0)
	var ret = make([]Rule0, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Rule0))
	}
	return ret
}

type RhsOptional struct {
	*Node
}

func (n RhsOptional) Inner() RhsPart {
	return ToTmNode(n.Child(selector.RhsPart)).(RhsPart)
}

type RhsPlusAssignment struct {
	*Node
}

func (n RhsPlusAssignment) Id() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n RhsPlusAssignment) Inner() RhsPart {
	return ToTmNode(n.Child(selector.RhsPart)).(RhsPart)
}

type RhsPlusList struct {
	*Node
}

func (n RhsPlusList) RuleParts() []RhsPart {
	nodes := n.Children(selector.RhsPart)
	var ret = make([]RhsPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(RhsPart))
	}
	return ret
}

func (n RhsPlusList) ListSeparator() ListSeparator {
	return ListSeparator{n.Child(selector.ListSeparator)}
}

type RhsQuantifier struct {
	*Node
}

func (n RhsQuantifier) Inner() RhsPart {
	return ToTmNode(n.Child(selector.RhsPart)).(RhsPart)
}

type RhsSet struct {
	*Node
}

func (n RhsSet) Expr() SetExpression {
	return ToTmNode(n.Child(selector.SetExpression)).(SetExpression)
}

type RhsStarList struct {
	*Node
}

func (n RhsStarList) RuleParts() []RhsPart {
	nodes := n.Children(selector.RhsPart)
	var ret = make([]RhsPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(RhsPart))
	}
	return ret
}

func (n RhsStarList) ListSeparator() ListSeparator {
	return ListSeparator{n.Child(selector.ListSeparator)}
}

type RhsSuffix struct {
	*Node
}

func (n RhsSuffix) Symref() Symref {
	return Symref{n.Child(selector.Symref)}
}

type RhsSymbol struct {
	*Node
}

func (n RhsSymbol) Reference() Symref {
	return Symref{n.Child(selector.Symref)}
}

type Rule struct {
	*Node
}

func (n Rule) Predicate() (Predicate, bool) {
	field := Predicate{n.Child(selector.Predicate)}
	return field, field.IsValid()
}

func (n Rule) RhsPart() []RhsPart {
	nodes := n.Children(selector.RhsPart)
	var ret = make([]RhsPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(RhsPart))
	}
	return ret
}

func (n Rule) RhsSuffix() (RhsSuffix, bool) {
	field := RhsSuffix{n.Child(selector.RhsSuffix)}
	return field, field.IsValid()
}

func (n Rule) ReportClause() (ReportClause, bool) {
	field := ReportClause{n.Child(selector.ReportClause)}
	return field, field.IsValid()
}

type SetAnd struct {
	*Node
}

func (n SetAnd) Left() SetExpression {
	return ToTmNode(n.Child(selector.SetExpression)).(SetExpression)
}

func (n SetAnd) Right() SetExpression {
	return ToTmNode(n.Child(selector.SetExpression).Next(selector.SetExpression)).(SetExpression)
}

type SetComplement struct {
	*Node
}

func (n SetComplement) Inner() SetExpression {
	return ToTmNode(n.Child(selector.SetExpression)).(SetExpression)
}

type SetCompound struct {
	*Node
}

func (n SetCompound) Inner() SetExpression {
	return ToTmNode(n.Child(selector.SetExpression)).(SetExpression)
}

type SetOr struct {
	*Node
}

func (n SetOr) Left() SetExpression {
	return ToTmNode(n.Child(selector.SetExpression)).(SetExpression)
}

func (n SetOr) Right() SetExpression {
	return ToTmNode(n.Child(selector.SetExpression).Next(selector.SetExpression)).(SetExpression)
}

type SetSymbol struct {
	*Node
}

func (n SetSymbol) Operator() (Identifier, bool) {
	field := Identifier{n.Child(selector.Identifier)}
	return field, field.IsValid()
}

func (n SetSymbol) Symbol() Symref {
	return Symref{n.Child(selector.Symref)}
}

type StartConditions struct {
	*Node
}

func (n StartConditions) Stateref() []Stateref {
	nodes := n.Children(selector.Stateref)
	var ret = make([]Stateref, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, Stateref{node})
	}
	return ret
}

type StartConditionsScope struct {
	*Node
}

func (n StartConditionsScope) StartConditions() StartConditions {
	return StartConditions{n.Child(selector.StartConditions)}
}

func (n StartConditionsScope) LexerPart() []LexerPart {
	nodes := n.Children(selector.LexerPart)
	var ret = make([]LexerPart, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(LexerPart))
	}
	return ret
}

type StateMarker struct {
	*Node
}

func (n StateMarker) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

type Stateref struct {
	*Node
}

func (n Stateref) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

type StringLiteral struct {
	*Node
}

type SubType struct {
	*Node
}

func (n SubType) Reference() Symref {
	return Symref{n.Child(selector.Symref)}
}

type Symref struct {
	*Node
}

func (n Symref) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n Symref) Args() (SymrefArgs, bool) {
	field := SymrefArgs{n.Child(selector.SymrefArgs)}
	return field, field.IsValid()
}

type SymrefArgs struct {
	*Node
}

func (n SymrefArgs) ArgList() []Argument {
	nodes := n.Children(selector.Argument)
	var ret = make([]Argument, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, ToTmNode(node).(Argument))
	}
	return ret
}

type SyntaxProblem struct {
	*Node
}

type TemplateParam struct {
	*Node
}

func (n TemplateParam) Modifier() (ParamModifier, bool) {
	field := ParamModifier{n.Child(selector.ParamModifier)}
	return field, field.IsValid()
}

func (n TemplateParam) ParamType() ParamType {
	return ParamType{n.Child(selector.ParamType)}
}

func (n TemplateParam) Name() Identifier {
	return Identifier{n.Child(selector.Identifier)}
}

func (n TemplateParam) ParamValue() (ParamValue, bool) {
	field := ToTmNode(n.Child(selector.ParamValue)).(ParamValue)
	return field, field.TmNode() != nil
}

type VoidType struct {
	*Node
}
