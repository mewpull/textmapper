%{
#include <stdio.h>
%}

%start Program

%right else
%token space
%token LineTerminatorSequence
%token MultiLineComment
%token SingleLineComment
%token Identifier
%token break
%token do
%token instanceof
%token typeof
%token case
%token new
%token var
%token catch
%token finally
%token return
%token void
%token continue
%token for
%token switch
%token while
%token debugger
%token function
%token this
%token with
%token default
%token if
%token throw
%token delete
%token in
%token try
%token class
%token enum
%token extends
%token super
%token const
%token export
%token import
%token Lcurly
%token Rcurly
%token Lparen
%token Rparen
%token Lsquare
%token Rsquare
%token Dot
%token Semicolon
%token Comma
%token Less
%token Greater
%token LessEqual
%token GreaterEqual
%token EqualEqual
%token ExclamationEqual
%token EqualEqualEqual
%token ExclamationEqualEqual
%token Plus
%token Minus
%token Mult
%token Percent
%token PlusPlus
%token MinusMinus
%token LessLess
%token GreaterGreater
%token GreaterGreaterGreater
%token Ampersand
%token Or
%token Xor
%token Exclamation
%token Tilde
%token AmpersandAmpersand
%token OrOr
%token Questionmark
%token Colon
%token Equal
%token PlusEqual
%token MinusEqual
%token MultEqual
%token PercentEqual
%token LessLessEqual
%token GreaterGreaterEqual
%token GreaterGreaterGreaterEqual
%token AmpersandEqual
%token OrEqual
%token XorEqual
%token null
%token true
%token false
%token NumericLiteral
%token StringLiteral
%token RegularExpressionLiteral
%token Slash
%token SlashEqual

%locations
%%

IdentifierName :
  Identifier
| break
| do
| instanceof
| typeof
| case
| else
| new
| var
| catch
| finally
| return
| void
| continue
| for
| switch
| while
| debugger
| function
| this
| with
| default
| if
| throw
| delete
| in
| try
| class
| enum
| extends
| super
| const
| export
| import
;

Literal :
  null
| true
| false
| NumericLiteral
| StringLiteral
| RegularExpressionLiteral
;

PrimaryExpression :
  this
| Identifier
| Literal
| ArrayLiteral
| ObjectLiteral
| Lparen Expression Rparen
;

ArrayLiteral :
  Lsquare Elisionopt Rsquare
| Lsquare ElementList Rsquare
| Lsquare ElementList Comma Elisionopt Rsquare
;

ElementList :
  Elisionopt AssignmentExpression
| ElementList Comma Elisionopt AssignmentExpression
;

Elision :
  Comma
| Elision Comma
;

ObjectLiteral :
  Lcurly Rcurly
| Lcurly PropertyNameAndValueList Rcurly
| Lcurly PropertyNameAndValueList Comma Rcurly
;

PropertyNameAndValueList :
  PropertyAssignment
| PropertyNameAndValueList Comma PropertyAssignment
;

PropertyAssignment :
  PropertyName Colon AssignmentExpression
| Identifier PropertyName Lparen Rparen Lcurly FunctionBody Rcurly
| Identifier PropertyName Lparen PropertySetParameterList Rparen Lcurly FunctionBody Rcurly
;

PropertyName :
  IdentifierName
| StringLiteral
| NumericLiteral
;

PropertySetParameterList :
  Identifier
;

MemberExpression :
  PrimaryExpression
| FunctionExpression
| MemberExpression Lsquare Expression Rsquare
| MemberExpression Dot IdentifierName
| new MemberExpression Arguments
;

NewExpression :
  MemberExpression
| new NewExpression
;

CallExpression :
  MemberExpression Arguments
| CallExpression Arguments
| CallExpression Lsquare Expression Rsquare
| CallExpression Dot IdentifierName
;

Arguments :
  Lparen Rparen
| Lparen ArgumentList Rparen
;

ArgumentList :
  AssignmentExpression
| ArgumentList Comma AssignmentExpression
;

LeftHandSideExpression :
  NewExpression
| CallExpression
;

PostfixExpression :
  LeftHandSideExpression
| LeftHandSideExpression PlusPlus
| LeftHandSideExpression MinusMinus
;

UnaryExpression :
  PostfixExpression
| delete UnaryExpression
| void UnaryExpression
| typeof UnaryExpression
| PlusPlus UnaryExpression
| MinusMinus UnaryExpression
| Plus UnaryExpression
| Minus UnaryExpression
| Tilde UnaryExpression
| Exclamation UnaryExpression
;

MultiplicativeExpression :
  UnaryExpression
| MultiplicativeExpression Mult UnaryExpression
| MultiplicativeExpression Slash UnaryExpression
| MultiplicativeExpression Percent UnaryExpression
;

AdditiveExpression :
  MultiplicativeExpression
| AdditiveExpression Plus MultiplicativeExpression
| AdditiveExpression Minus MultiplicativeExpression
;

ShiftExpression :
  AdditiveExpression
| ShiftExpression LessLess AdditiveExpression
| ShiftExpression GreaterGreater AdditiveExpression
| ShiftExpression GreaterGreaterGreater AdditiveExpression
;

RelationalExpression :
  ShiftExpression
| RelationalExpression Less ShiftExpression
| RelationalExpression Greater ShiftExpression
| RelationalExpression LessEqual ShiftExpression
| RelationalExpression GreaterEqual ShiftExpression
| RelationalExpression instanceof ShiftExpression
| RelationalExpression in ShiftExpression
;

RelationalExpression_NoIn :
  ShiftExpression
| RelationalExpression_NoIn Less ShiftExpression
| RelationalExpression_NoIn Greater ShiftExpression
| RelationalExpression_NoIn LessEqual ShiftExpression
| RelationalExpression_NoIn GreaterEqual ShiftExpression
| RelationalExpression_NoIn instanceof ShiftExpression
;

EqualityExpression :
  RelationalExpression
| EqualityExpression EqualEqual RelationalExpression
| EqualityExpression ExclamationEqual RelationalExpression
| EqualityExpression EqualEqualEqual RelationalExpression
| EqualityExpression ExclamationEqualEqual RelationalExpression
;

EqualityExpression_NoIn :
  RelationalExpression_NoIn
| EqualityExpression_NoIn EqualEqual RelationalExpression_NoIn
| EqualityExpression_NoIn ExclamationEqual RelationalExpression_NoIn
| EqualityExpression_NoIn EqualEqualEqual RelationalExpression_NoIn
| EqualityExpression_NoIn ExclamationEqualEqual RelationalExpression_NoIn
;

BitwiseANDExpression :
  EqualityExpression
| BitwiseANDExpression Ampersand EqualityExpression
;

BitwiseANDExpression_NoIn :
  EqualityExpression_NoIn
| BitwiseANDExpression_NoIn Ampersand EqualityExpression_NoIn
;

BitwiseXORExpression :
  BitwiseANDExpression
| BitwiseXORExpression Xor BitwiseANDExpression
;

BitwiseXORExpression_NoIn :
  BitwiseANDExpression_NoIn
| BitwiseXORExpression_NoIn Xor BitwiseANDExpression_NoIn
;

BitwiseORExpression :
  BitwiseXORExpression
| BitwiseORExpression Or BitwiseXORExpression
;

BitwiseORExpression_NoIn :
  BitwiseXORExpression_NoIn
| BitwiseORExpression_NoIn Or BitwiseXORExpression_NoIn
;

LogicalANDExpression :
  BitwiseORExpression
| LogicalANDExpression AmpersandAmpersand BitwiseORExpression
;

LogicalANDExpression_NoIn :
  BitwiseORExpression_NoIn
| LogicalANDExpression_NoIn AmpersandAmpersand BitwiseORExpression_NoIn
;

LogicalORExpression :
  LogicalANDExpression
| LogicalORExpression OrOr LogicalANDExpression
;

LogicalORExpression_NoIn :
  LogicalANDExpression_NoIn
| LogicalORExpression_NoIn OrOr LogicalANDExpression_NoIn
;

ConditionalExpression :
  LogicalORExpression
| LogicalORExpression Questionmark AssignmentExpression Colon AssignmentExpression
;

ConditionalExpression_NoIn :
  LogicalORExpression_NoIn
| LogicalORExpression_NoIn Questionmark AssignmentExpression_NoIn Colon AssignmentExpression_NoIn
;

AssignmentExpression :
  ConditionalExpression
| LeftHandSideExpression AssignmentOperator AssignmentExpression
;

AssignmentExpression_NoIn :
  ConditionalExpression_NoIn
| LeftHandSideExpression AssignmentOperator AssignmentExpression_NoIn
;

AssignmentOperator :
  Equal
| MultEqual
| SlashEqual
| PercentEqual
| PlusEqual
| MinusEqual
| LessLessEqual
| GreaterGreaterEqual
| GreaterGreaterGreaterEqual
| AmpersandEqual
| XorEqual
| OrEqual
;

Expression :
  AssignmentExpression
| Expression Comma AssignmentExpression
;

Expression_NoIn :
  AssignmentExpression_NoIn
| Expression_NoIn Comma AssignmentExpression_NoIn
;

Statement :
  Block
| VariableStatement
| EmptyStatement
| ExpressionStatement
| IfStatement
| IterationStatement
| ContinueStatement
| BreakStatement
| ReturnStatement
| WithStatement
| LabelledStatement
| SwitchStatement
| ThrowStatement
| TryStatement
| DebuggerStatement
;

Block :
  Lcurly StatementList Rcurly
| Lcurly Rcurly
;

StatementList :
  Statement
| StatementList Statement
;

VariableStatement :
  var VariableDeclarationList Semicolon
;

VariableDeclarationList :
  VariableDeclaration
| VariableDeclarationList Comma VariableDeclaration
;

VariableDeclarationList_NoIn :
  VariableDeclaration_NoIn
| VariableDeclarationList_NoIn Comma VariableDeclaration_NoIn
;

VariableDeclaration :
  Identifier Initialiseropt
;

VariableDeclaration_NoIn :
  Identifier Initialiseropt_NoIn
;

Initialiser :
  Equal AssignmentExpression
;

Initialiser_NoIn :
  Equal AssignmentExpression_NoIn
;

EmptyStatement :
  Semicolon
;

ExpressionStatement :
  PlusPlus Expression Semicolon
;

IfStatement :
  if Lparen Expression Rparen Statement else Statement
| if Lparen Expression Rparen Statement
;

IterationStatement :
  do Statement while Lparen Expression Rparen Semicolon
| while Lparen Expression Rparen Statement
| for Lparen Expressionopt_NoIn Semicolon Expressionopt Semicolon Expressionopt Rparen Statement
| for Lparen var VariableDeclarationList_NoIn Semicolon Expressionopt Semicolon Expressionopt Rparen Statement
| for Lparen LeftHandSideExpression in Expression Rparen Statement
| for Lparen var VariableDeclaration_NoIn in Expression Rparen Statement
;

ContinueStatement :
  continue Identifier Semicolon
| continue Semicolon
;

BreakStatement :
  break Identifier Semicolon
| break Semicolon
;

ReturnStatement :
  return Expressionopt Semicolon
;

WithStatement :
  with Lparen Expression Rparen Statement
;

SwitchStatement :
  switch Lparen Expression Rparen CaseBlock
;

CaseBlock :
  Lcurly CaseClausesopt Rcurly
| Lcurly CaseClausesopt DefaultClause CaseClausesopt Rcurly
;

CaseClauses :
  CaseClause
| CaseClauses CaseClause
;

CaseClause :
  case Expression Colon StatementListopt
;

DefaultClause :
  default Colon StatementListopt
;

LabelledStatement :
  Identifier Colon Statement
;

ThrowStatement :
  throw Expression Semicolon
;

TryStatement :
  try Block Catch
| try Block Finally
| try Block Catch Finally
;

Catch :
  catch Lparen Identifier Rparen Block
;

Finally :
  finally Block
;

DebuggerStatement :
  debugger Semicolon
;

FunctionDeclaration :
  function Identifier Lparen FormalParameterListopt Rparen Lcurly FunctionBody Rcurly
;

FunctionExpression :
  function Identifier Lparen FormalParameterListopt Rparen Lcurly FunctionBody Rcurly
| function Lparen FormalParameterListopt Rparen Lcurly FunctionBody Rcurly
;

FormalParameterList :
  Identifier
| FormalParameterList Comma Identifier
;

FunctionBody :
  SourceElementsopt
;

Program :
  SourceElementsopt
;

SourceElements :
  SourceElement
| SourceElements SourceElement
;

SourceElement :
  Statement
| FunctionDeclaration
;

Elisionopt :
  %empty
| Elision
;

Initialiseropt :
  %empty
| Initialiser
;

Initialiseropt_NoIn :
  %empty
| Initialiser_NoIn
;

Expressionopt :
  %empty
| Expression
;

Expressionopt_NoIn :
  %empty
| Expression_NoIn
;

CaseClausesopt :
  %empty
| CaseClauses
;

StatementListopt :
  %empty
| StatementList
;

FormalParameterListopt :
  %empty
| FormalParameterList
;

SourceElementsopt :
  %empty
| SourceElements
;

%%

