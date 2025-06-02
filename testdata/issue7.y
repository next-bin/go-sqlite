%{
// java 20
package parser


%}

%union{
  node  *Node
  nodes []*Node
  anys  []any

  token *Token
}
%token<token> ID LineComment BlockComment

/* primitive */
%token<token> VOID "void"
%token<token> BOOLEAN "boolean" BYTE "byte" SHORT "short"  CHAR "char"  INT "int"
%token<token> LONG "long" FLOAT "float" DOUBLE "double"
/* literal */
%token<token> LSTRING LTEXT LINT LLONG LFLOAT LDOUBLE LCHAR
%token<token> LTRUE "true" LFALSE "false" NULL "null"
/* reference */
%token<token> THIS "this" SUPER "super"
/* module declaration */
%token<token> USES "uses" EXPORTS "exports" PROVIDES "provides" OPENS "OPENS"
%token<token> MODULE "module" REQUIRES "requires" TO "to" WITH "with" OPEN "open"
/* declaration */
%token<token> PACKAGE "package"  IMPORT "import" PERMITS "permits"
%token<token> VAR "var"
%token<token> RECORD "record" INTERFACE "interface" CLASS "class"  ENUM "enum"
%token<token> THROWS "throws"  IMPLEMENTS "implements"  EXTENDS "extends"
%token<token> SEALED "sealed"  NONSEALED "non-sealed"
/* modifier */
%token<token> PUBLIC "public" PRIVATE "private" PROTECTED "protected"
%token<token> CONST "const" ABSTRACT "abstract" FINAL "final" STATIC "static"
%token<token> STRICTFP "strictfp" VOLATILE "volatile" NATIVE "native"
%token<token> TRANSIENT "transient" SYNCHRONIZED "synchronized"
/* operation */
%token<token> INSTANCEOF "instanceof" RETURN "return" THROW "throw"
%token<token> TRY "try" CATCH "catch" FINALLY "finally"
%token<token> WHILE "while" DO "do" FOR "for" BREAK "break" CONTINUE "continue"
%token<token> IF "if" ELSE "else"
%token<token> NEW "new" GOTO "goto"
%token<token> SWITCH "switch" CASE "case" DEFAULT "default"
%token<token> ASSERT "assert"
%token<token> YIELD "yield"  WHEN "when" /* 17+ */
/* declare mark */
%token<token> HASH "#"
/* separator */
%token<token> AT "@"
%token<token> LPAREN "(" RPAREN ")"
%token<token> LBRACE "{" RBRACE "}"
%token<token> LBRACKET "[" RBRACKET "]"
%token<token> SEMI ";" COMMA ","
%token<token> DOT "." ELLIPSIS "..."
%token<token> COLON ":" COLONCOLON "::"
%token<token> ARROW "->"
%token<token> QUERY "?"
/* operator */
%token<token> BANG "!"
%token<token> ASSIGN "=" PLUS "+" MINUS "-"
%token<token> STAR "*" SLASH "/" BAR "|"
%token<token> TILDE "~" AMP "&"
%token<token> PERC "%" CARET "^" LT "<" GT ">"

/* names */
%type<node> identifier contextualKeyword contextualKeywordMinusForTypeIdentifier
%type<node> contextualKeywordMinusForUnqualifiedMethodIdentifier
%type<node> literal integer float boolean character string text null

%type<node> primitiveType booleanType numericType floatingPointType integerType
%type<node> referenceType classOrInterfaceType typeVariable arrayType
%type<anys> coit interfaceType classType

%type<node> qualifiedName qualifiedStar packagePrefixName
%type<nodes> identifierList

%type<node> compilationUnit packageDecl importDecl importDecls classDecls

%type<node> assignmentOperator leftHandSide switchExpression lambdaBody lambdaParameter lambdaParameters
%type<node> arrayAccess assignmentExpression assignment conditionalOrExpression conditionalAndExpression inclusiveOrExpression
%type<node> exclusiveOrExpression andExpression equalityExpression  pattern relationalExpression shiftExpression
%type<node> additiveExpression multiplicativeExpression castExpression unaryExpressionNotPlusMinus
%type<node> preDecrementExpression preIncrementExpression unaryExpression
%type<node> postDecrementExpression postfixExpression postIncrementExpression pfE
%type<node> methodReference methodInvocation fieldAccess dimExpr dimExprs arrayCreationExpression annoIdentifier
%type<node> typeArgumentsOrDiamond classOrInterfaceTypeToInstantiate unqualifiedClassInstanceCreationExpression
%type<node> classInstanceCreationExpression classLiteral primaryNoNewArray methodName primary resourceSpecification
%type<node> arrayCreationExpressionWithInitializer lambdaExpression yieldStatement expression resource
%type<node> catchType catchFormalParameter tryWithResourcesStatement catchClause finallyBlock tryStatement
%type<anys> pNNA
%type<nodes> lambdaParameterList identifierTypeList annoIdentifierList resourceList


%%
/* special perpose */
semi
    : ";" {$$=N(JSemi,$1)}
    ;
/* elements */

identifier
    : ID {$$=N(JID,$1)}
    ;
contextualKeyword
    : "exports" {$$=N(JID,$1)}
    | "module" {$$=N(JID,$1)}
    | "non-sealed" {$$=N(JID,$1)}
    | "open" {$$=N(JID,$1)}
    | "opens" {$$=N(JID,$1)}
    | "permits" {$$=N(JID,$1)}
    | "provides" {$$=N(JID,$1)}
    | "record" {$$=N(JID,$1)}
    | "requires" {$$=N(JID,$1)}
    | "sealed" {$$=N(JID,$1)}
    | "to" {$$=N(JID,$1)}
    | "transitive" {$$=N(JID,$1)}
    | "uses" {$$=N(JID,$1)}
    | "var" {$$=N(JID,$1)}
    | "with" {$$=N(JID,$1)}
    | "yield" {$$=N(JID,$1)}
    ;
contextualKeywordMinusForTypeIdentifier
    : "exports" {$$=N(JID,$1)}
    | "module" {$$=N(JID,$1)}
    | "non-sealed" {$$=N(JID,$1)}
    | "open" {$$=N(JID,$1)}
    | "opens" {$$=N(JID,$1)}
    | "provides" {$$=N(JID,$1)}
    | "requires" {$$=N(JID,$1)}
    | "to" {$$=N(JID,$1)}
    | "transitive" {$$=N(JID,$1)}
    | "uses" {$$=N(JID,$1)}
    | "with" {$$=N(JID,$1)}
    ;
contextualKeywordMinusForUnqualifiedMethodIdentifier
    : "exports" {$$=N(JID,$1)}
    | "module" {$$=N(JID,$1)}
    | "non-sealed" {$$=N(JID,$1)}
    | "open" {$$=N(JID,$1)}
    | "opens" {$$=N(JID,$1)}
    | "permits" {$$=N(JID,$1)}
    | "provides" {$$=N(JID,$1)}
    | "record" {$$=N(JID,$1)}
    | "requires" {$$=N(JID,$1)}
    | "sealed" {$$=N(JID,$1)}
    | "to" {$$=N(JID,$1)}
    | "transitive" {$$=N(JID,$1)}
    | "uses" {$$=N(JID,$1)}
    | "var" {$$=N(JID,$1)}
    | "with" {$$=N(JID,$1)}
    ;
/* 3.10 */
literal
    : integer {$$=$1}
    | float {$$=$1}
    | boolean {$$=$1}
    | character {$$=$1}
    | string {$$=$1}
    | text {$$=$1}
    | null {$$=$1}
    ;
integer
    : LINT {$$=N(JLiteral,$1)}
    | LONG {$$=N(JLiteral,$1)}
    ;
float
    : LFLOAT {$$=N(JLiteral,$1)}
    | LDOUBLE {$$=N(JLiteral,$1)}
    ;
boolean
    : LTRUE {$$=N(JLiteral,$1)}
    | LFALSE {$$=N(JLiteral,$1)}
    ;
character
    : LCHAR {$$=N(JLiteral,$1)}
    ;
string
    : LSTRING {$$=N(JLiteral,$1)}
    ;
text
    : LTEXT {$$=N(JLiteral,$1)}
    ;
null
    : "null" {$$=N(JLiteral,$1).set(FNull)}
    ;
/*  4.2 */
primitiveType
    : numericType {$$=$1}
    | booleanType {$$=$1}
    ;
booleanType
    : "boolean" {$$=N(JPrimitiveType,$1)}
    ;
numericType
    : integerType {$$=$1}
    | floatingPointType {$$=$1}
    ;
floatingPointType
    : "float"  {$$=N(JPrimitiveType,$1)}
    | "double" {$$=N(JPrimitiveType,$1)}
    ;
integerType
    : "byte"  {$$=N(JPrimitiveType,$1)}
    | "short"   {$$=N(JPrimitiveType,$1)}
    | "char"   {$$=N(JPrimitiveType,$1)}
    | "int"    {$$=N(JPrimitiveType,$1)}
    | "long"  {$$=N(JPrimitiveType,$1)}
    ;
/* 4.3 */
referenceType
    : classOrInterfaceType {$$=$1}
    | typeVariable {$$=$1}
    | arrayType {$$=$1}
    ;
/* '.' annotation* typeIdentifier typeArguments? coit? */
coit
    : "." annotations typeIdentifier typeArguments  {$$=Anys($1,$2,$3,$4)}
    | "." typeIdentifier typeArguments  {$$=Anys($1,$2,$3)}
    | "." typeIdentifier  {$$=Anys($1,$2)}
    | "." typeIdentifier {$$=Anys($1)}
    | "." annotations typeIdentifier  {$$=Anys($1,$2,$3)}
    | "." annotations typeIdentifier {$$=Anys($1,$2)}
    | "." typeIdentifier typeArguments  {$$=Anys($1,$2,$3)}
    | "." typeIdentifier  {$$=Anys($1,$2)}
    | "." annotations typeIdentifier typeArguments {$$=Anys($1,$2,$3)}
    | coit coit {$$=append($1,$2...}
    ;

classOrInterfaceType
    : typeIdentifier {$$=N(JClassOrInteraceType,$1)}
    | annotations typeIdentifier {$$=N(JClassOrInteraceType,$1,$2)}
    | packagePrefixName annotations typeIdentifier {$$=N(JClassOrInteraceType,$1,$2,$3)}
    | typeIdentifier typeArguments {$$=N(JClassOrInteraceType,$1,$2)}
    | annotations typeIdentifier typeArguments {$$=N(JClassOrInteraceType,$1,$2,$3)}
    | packagePrefixName typeIdentifier typeArguments {$$=N(JClassOrInteraceType,$1,$2,$3)}
    | packagePrefixName annotations typeIdentifier typeArguments {$$=N(JClassOrInteraceType,$1,$2,$3,$4)}
    | classOrInterfaceType coit {$$=$1.add($2...)}
    ;

classType
    : annotations typeIdentifier typeArguments {$$=N(JClassType,$1,$2,$3)}
    | typeIdentifier typeArguments {$$=N(JClassType,$1,$2)}
    | annotations typeIdentifier {$$=N(JClassType,$1,$2)}
    | typeIdentifier {$$=N(JClassType,$1)}
    | packagePrefixName annotations typeIdentifier typeArguments {$$=N(JClassType,$1,$2,$3,$4)}
    | classOrInterfaceType "." annotations typeIdentifier typeArguments {$$=N(JClassType,$1,$2,$3,$4,$5)}
    | classOrInterfaceType "." annotations typeIdentifier {$$=N(JClassType,$1,$2,$3,$4)}
    | classOrInterfaceType "." typeIdentifier typeArguments {$$=N(JClassType,$1,$2,$3,$4)}
    | classOrInterfaceType "." typeIdentifier {$$=N(JClassType,$1,$2,$3)}
    ;
interfaceType
    : classType {$$=$1}
    ;
typeVariable
    : annotations typeIdentifier {$$=N(JTypeVariable,$1,$2)}
    ;
arrayType
    : primitiveType dims {$$=N(JArrayType,$1,$2)}
    | classType dims {$$=N(JArrayType,$1,$2)}
    | typeVariable dims {$$=N(JArrayType,$1,$2)}
    ;
dims
    : dim {$$=N(JDims,$1)}
    | dims dim {$$=$1.add($2)}
    ;
dim
    : annotations "[" "]" {$$=N(JDim,$1,$2,$3)}
    | "[" "]" {$$=N(JDim,$1,$2)}
    ;
/*4.4*/
typeParameter
    : annotations typeIdentifier typeBound {$$=N(JTypeParameter,$1,$2,$3)}
    | annotations typeIdentifier {$$=N(JTypeParameter,$1,$2)}
    | typeIdentifier typeBound {$$=N(JTypeParameter,$1,$2)}
    | typeIdentifier {$$=N(JTypeParameter,$1)}
    ;
typeBound
    : "extends" typeVariable  {$$=N(JTypeBound,$1,$2)}
    | "extends" classOrInterfaceType {$$=N(JTypeBound,$1,$2)}
    | typeBound intersectionBound {$$=$1.add($2)}
    ;
intersectionBounds
    : intersectionBound {$$=$1}
    | intersectionBounds intersectionBound{$$=$1.add($2)}
    ;
intersectionBound
    : "&" interfaceType {$$=N(JIntersectionBound,$1)}
    ;
/*4.5.1*/
typeArguments
    : "<" typeArgumentList ">" {$$=N(JTypeArguments,$1,$2,$3)}
    ;
typeArgumentList
    : typeArgument {$$=NS($1)}
    | typeArgumentList "," typeArgument {$$=append($1,$3)}
    ;
typeArgument
    :referenceType {$$=$1}
    |wildcard {$$=$1}
    ;
wildcard
    : annotations "?" wildcardBounds {$$=N(JWildcard,$1,$2,$3)}
    | "?" wildcardBound {$$=N(JWildcard,$1,$2)}
    | annotations "?" {$$=N(JWildcard,$1,$2)}
    | "?" {$$=N(JWildcard,$1)}
    ;
wildcardBounds
    : "super" referenceType {$$=N(JWildcardBounds,$1,$2).set(FSuper)}
    | "extends" referenceType {$$=N(JWildcardBounds,$1,$2).set(FExtends)}
    ;

typeIdentifier
    : Identifier {$$=$1}
    | contextualKeywordMinusForTypeIdentifier {$$=$1}
    ;
/*6.5*/
/*moduleName*/
/*packageName*/
packagePrefixName
    : identifier "." {$$=N(JQualified,$1,$2).set(FPrefix)}
    | packagePrefixName identifier "." {$$=$1.add($2,$3)}
    ;

typeName
    : packagePrefixName "." typeIdentifier {$$=N(JTypeName,$1,$2,$3)}
    ;
/*packageOrTypeName*/
/*expressionName*/
methodName
    : unqualifiedMethodIdentifier
    ;
/*ambiguousName*/
/*7.3              =================================*/
compilationUnit
    : ordinaryCompilationUnit {$$=$1}
    ;
ordinaryCompilationUnit
    : packageDeclaration importDeclarationList topLevelClassOrInterfaceDeclarationList {$$=N(JCompilationUnit,$1,$2,$3)}
    |  importDeclarationList topLevelClassOrInterfaceDeclarationList {$$=N(JCompilationUnit,$1,$2)}
    |  topLevelClassOrInterfaceDeclarationList {$$=N(JCompilationUnit,$1)}
    |  packageDeclaration {$$=N(JCompilationUnit,$1)}
    ;

/*7.4*/
packageDeclaration
    : annotations "package" qualifiedName ";" {$$=N(JPacakageDeclaration,$1,$2,$3)}
    ;
/*7.5*/
importDeclaration
    : "import" qualifiedStar ";" {$$=N(JImportDeclaration,$1,$2,$3)}
    | "import" "static" qualifiedStar ";" {$$=N(JImportDeclaration,$1,$2,$3,$4).set(FStatic)}
    ;
/*7.6*/
topLevelClassOrInterfaceDeclaration
    : classDeclaration {$$=$1}
    | interfaceDeclaration {$$=$1}
    | semi {$$=$1}
    ;

/* 8.1 */
classDeclaration
    : normalClassDeclaration {$$=$1}
    | enumDeclaration {$$=$1}
    | recordDeclaration {$$=$1}
    ;
/* ! Todo */
normalClassDeclaration
    : "class" typeIdentifier classBody { $$=N(JClassDeclaraation,$1,$2,$3)}
    | classModifiers "class" typeIdentifier classBody { $$=N(JClassDeclaraation,$1,$2,$3,$4)}
    | "class" typeIdentifier typeParameters classBody { $$=N(JClassDeclaraation,$1,$2,$3,$4)}
    | classModifiers "class" typeIdentifier typeParameters classBody { $$=N(JClassDeclaraation,$1,$2,$3,$4,$5)}
    | classModifiers "class" typeIdentifier classExtends classBody { $$=N(JClassDeclaraation,$1,$2,$3,$4,$5)}
    | "class" typeIdentifier typeParameters classExtends classBody { $$=N(JClassDeclaraation,$1,$2,$3,$4,$5)}
    | "class" typeIdentifier typeParameters classExtends classImplements classPermits classBody
    { $$=N(JClassDeclaraation,$1,$2,$3,$4,$5,$6,$7)}
    | "class" typeIdentifier typeParameters classExtends classImplements classPermits classBody
    { $$=N(JClassDeclaraation,$1,$2,$3,$4,$5,$6,$7)}
    | classModifiers "class" typeIdentifier typeParameters classExtends classImplements classPermits classBody
     {$$=N(JClassDeclaraation,$1,$2,$3,$4,$5,$6,$7,$8)}
    ;
classModifiers
    : classModifier {$$=N(JModifiers,$1)}
    | classModifiers classModifier{$$=$1.add($2)}
    ;
classModifier
    : annotation {$$=$1}
    | "public" {$$=N(JModifier,$1)}
    | "protected" {$$=N(JModifier,$1)}
    | "private" {$$=N(JModifier,$1)}
    | "abstract" {$$=N(JModifier,$1)}
    | "static" {$$=N(JModifier,$1)}
    | "final" {$$=N(JModifier,$1)}
    | "sealed" {$$=N(JModifier,$1)}
    | "non-sealed" {$$=N(JModifier,$1)}
    | "strictfp" {$$=N(JModifier,$1)}
    ;
typeParameters
    : "<" typeParameterList ">" {$$=N(JTypeParameters,$1,$2,$3)}
    ;

typeParameterList
    : typeParameter {$$=NS($1)}
    | typeParameterList "," typeParameter{$$=append($1,$3)}
    ;
classExtends
    : "extends" classType {$$=N(JClassExtends,$1,$2).set(FExtends)}
    ;
classImplements
    : "implements" interfaceTypeList {$$=N(JClassExtends,$1,$2).set(FImplements)}
    ;
interfaceTypeList
    : interfaceType {$$=NS($1)}
    | interfaceTypeList "," interfaceType{$$=append($1,$3)}
    ;
classPermits
    : "permits" typeNameList {$$=N(JClassExtends,$1,$2).set(FPermits)}
    ;
typeNameList
    : typeName {$$=NS($1)}
    | typeNameList typeName{$$=append($1,$2)}
    ;
classBody
    : "{" classBodyDeclarationList "}" {$$=N(JTypeBody,$1,$2,$3)}
    ;
classBodyDeclarationList
    : classBodyDeclaration {$$=NS($1)}
    | classBodyDeclarationList classBodyDeclaration{$$=append($1,$2)}
    ;
classBodyDeclaration
    : classMemberDeclaration {$$=$1}
    | instanceInitializer {$$=$1}
    | staticInitializer {$$=$1}
    | constructorDeclaration {$$=$1}
    ;
classMemberDeclaration
    : fieldDeclaration {$$=$1}
    | methodDeclaration {$$=$1}
    | classDeclaration {$$=$1}
    | interfaceDeclaration {$$=$1}
    | semi {$$=$1}
    ;
/*8.3 */
fieldDeclaration
    : fieldModifiers unannType variableDeclaratorList semi {$$=N(JFieldDeclaration,$1,$2,$3,$4)}
    | unannType variableDeclaratorList semi {$$=N(JFieldDeclaration,$1,$2,$3)}
    ;
fieldModifiers
    : fieldModifier {$$=N(JModifiers,$1)}
    | fieldModifiers fieldModifier{$$=$1.add($2)}
    ;
fieldModifier
    : annotation {$$=$1}
    | "public" {$$=N(JModifier,$1)}
    | "protected" {$$=N(JModifier,$1)}
    | "private" {$$=N(JModifier,$1)}
    | "static" {$$=N(JModifier,$1)}
    | "final" {$$=N(JModifier,$1)}
    | "transient" {$$=N(JModifier,$1)}
    | "volatile" {$$=N(JModifier,$1)}
    ;
variableDeclaratorList
    : variableDeclarator {$$=NS($1)}
    | variableDeclaratorList variableDeclarator{$$=append($1,$2)}
    ;
variableDeclarator
    : variableDeclaratorId {$$=N(JVariableDeclarator,$1)}
    | variableDeclaratorId "=" variableInitializer {$$=N(JVariableDeclarator,$1,$3)}
    ;
variableDeclaratorId
    : identifier {$$=$1}
    | identifier dims {$$=N(JVariableDeclaratorId,$1,$2)}
    ;
variableInitializer
    : expression {$$=$1}
    | arrayInitializer {$$=$1}
    ;
unannType
    : unannPrimitiveType {$$=$1}
    | unannReferenceType {$$=$1}
    ;
unannPrimitiveType
    : numericType {$$=$1}
    | booleanType {$$=$1}
    ;
unannReferenceType
    : unannClassOrInterfaceType {$$=$1}
    | unannTypeVariable {$$=$1}
    | unannArrayType {$$=$1}
    ;
unannClassOrInterfaceType
    : packagePrefixAnnotated genericTypeName  {$$=N(JClassOrInterfaceType,$1,$2)}
    | unannClassOrInterfaceType uCOIT {$$=$1.add($2...)}
    ;
uCOIT
    : "." annotations genericTypeName {$$=Anys($1,$2,$3,$4)}
    | "." genericTypeName {$$=Anys($1,$2,$3)}
    ;
packagePrefixAnnotated
    : packagePrefixName {$$=$1}
    | packagePrefixAnnotated "." annotations {$$=$1.add($2)}
    ;
unannClassType
    : genericTypeName {$$=$1}
    | packagePrefixAnnotated genericTypeName {$$=N(JClassType,$1,$2)}
    | unannClassOrInterfaceType "." annotations genericTypeName {$$=N(JClassType,$1,$2,$3 )}
    ;
genericTypeName
    : typeIdentifier {$$=N(JGenericTypeName,$1)}
    | typeIdentifier typeArguments {$$=N(JGenericTypeName,$1,$2)}
    ;
unannInterfaceType
    : unannClassType {$$=$1}
    ;
unannTypeVariable
    : typeIdentifier {$$=$1}
    ;
unannArrayType
    : unannPrimitiveType dims {$$=N(JArrayType,$1,$2)}
    | unannClassOrInterfaceType  dims {$$=N(JArrayType,$1,$2)}
    | unannTypeVariable dims {$$=N(JArrayType,$1,$2)}
    ;
/*8.4*/
methodDeclaration
    : methodModifiers methodHeader methodBody {$$=N(JMethodDeclaration,$1,$2,$3)}
    | methodHeader methodBody {$$=N(JMethodDeclaration,$1,$2,$3)}
    ;
methodModifiers
    : methodModifier {$$=N(JModifiers,$1)}
    | methodModifiers methodModifier{$$=$1.add($2)}
    ;
methodModifier
    : annotation {$$=$1}
    | "public" {$$=N(JModifier,$1)}
    | "protected" {$$=N(JModifier,$1)}
    | "private" {$$=N(JModifier,$1)}
    | "abstract" {$$=N(JModifier,$1)}
    | "static" {$$=N(JModifier,$1)}
    | "final" {$$=N(JModifier,$1)}
    | "synchronized" {$$=N(JModifier,$1)}
    | "native" {$$=N(JModifier,$1)}
    | "strictfp" {$$=N(JModifier,$1)}
    ;
methodHeader
    : typeParameters annotations resultType methodDeclarator  {$$=N(JMethodHeader,$1,$2,$3,$4)}
    | annotations resultType methodDeclarator  {$$=N(JMethodHeader,$1,$2,$3)}
    | typeParameters resultType methodDeclarator  {$$=N(JMethodHeader,$1,$2,$3)}
    | resultType methodDeclarator  {$$=N(JMethodHeader,$1,$2)}
    | methodHeader throwsT {$$=$1.add($2)}
    ;
methodDeclarator
    : identifier parameters
    | methodDeclarator dims {$$=$1.add($2)}
    ;
parameters
    :"("")" {$$=N(JParameters,$1,$2)}
    |"(" receiverParameter "," formalParameterList ")" {$$=N(JParameters,$1,$2,$3,$4,$5)}
    |"(" receiverParameter "," variableArityParameter ")" {$$=N(JParameters,$1,$2,$3,$4,$5)}
    |"(" formalParameterList ")" {$$=N(JParameters,$1,$2,$3)}
    |"(" variableArityParameter ")" {$$=N(JParameters,$1,$2,$3)}
    |"(" formalParameterList "," variableArityParameter ")" {$$=N(JParameters,$1,$2,$3,$4,$5)}
   ;
receiverParameter
    : annotations unannType identifier "." "this" {$$=N(JReceiverParameter,$1,$2,$3,$4,$5)}
    | unannType identifier "." "this" {$$=N(JReceiverParameter,$1,$2,$3,$4)}
    | annotations unannType  "this" {$$=N(JReceiverParameter,$1,$2,$3)}
    | unannType "this" {$$=N(JReceiverParameter,$1,$2)}
    ;
formalParameterList
    : formalParameter {$$=NS($1)}
    | formalParameterList "," formalParameter{$$=append($1,$3)}
    ;
formalParameter
    : variableModifiers unannType variableDeclaratorId {$$=N(JParameter,$1,$2,$3)}
    | unannType variableDeclaratorId {$$=N(JParameter,$1,$2)}
    ;
variableArityParameter
    : variableModifiers unannType annotations "..." identifier {$$=N(JParameter,$1,$2,$3,$4,$5).set(FVararg)}
    | variableModifiers unannType "..." identifier {$$=N(JParameter,$1,$2,$3,$4,$5).set(FVararg)}
    | unannType annotations "..." identifier {$$=N(JParameter,$1,$2,$3,$4).set(FVararg)}
    | unannType "..." identifier {$$=N(JParameter,$1,$2,$3).set(FVararg)}
    ;
resultType
    : voidType
    | unannType
    ;
variableModifiers
    : variableModifier {$$=N(JModifiers,$1)}
    | variableModifiers variableModifier{$$=$1.add($2)}
    ;
variableModifier
    : annotation {$$=$1}
    | "final" {$$=N(JModifier,$1)}
    ;
throwsT
    : "throws" exceptionTypeList {$$=N(JThrowList,$1,$2)}
    ;
exceptionTypeList
    : exceptionType {$$=NS($1)}
    | exceptionTypeList "," exceptionType{$$=append($1,$3)}
    ;
exceptionType
    : classType {$$=$1}
    | typeVariable {$$=$1}
    ;
methodBody
    : block {$$=$1}
    | semi {$$=$1}
    ;
/*8.6*/
instanceInitializer
    : block {$$=$1}
    ;
staticInitializer
    : "static" block {$$=$1.set(FStatic)}
    ;
constructorDeclaration
    : constructorModifiers constructorDeclarator throwsT constructorBody {$$=N(JContructorDeclaration,$1,$2,$3,$4)}
    | constructorDeclarator throwsT constructorBody {$$=N(JContructorDeclaration,$1,$2,$3)}
    | constructorModifiers constructorDeclarator constructorBody {$$=N(JContructorDeclaration,$1,$2,$3)}
    | constructorDeclarator constructorBody {$$=N(JContructorDeclaration,$1,$2)}
    ;
constructorModifiers
    : constructorModifier {$$=N(,$1)}
    | constructorModifiers constructorModifier{$$=$1.add($2)}
    ;
constructorModifier
    : annotation {$$=$1}
    | "public" {$$=N(JModifier,$1)}
    | "protected" {$$=N(JModifier,$1)}
    | "private" {$$=N(JModifier,$1)}
    ;
constructorDeclarator
    : typeParameters simpleTypeName parameters {$$=N(JConstructorDeclarator,$1,$2,$3)}
    | simpleTypeName parameters {$$=N(JConstructorDeclarator,$1,$2)}
    ;
simpleTypeName
    : typeIdentifier {$$=$1}
    ;
constructorBody
    : "{" explicitConstructorInvocation blockStatements "}" {$$=N(JConstructorBody,$1,$2,$3,$4)}
    | "{" explicitConstructorInvocation "}" {$$=N(JConstructorBody,$1,$2,$3)}
    | "{" blockStatements "}" {$$=N(JConstructorBody,$1,$2,$3)}
    | "{" "}" {$$=N(JConstructorBody,$1,$2)}
    ;
explicitConstructorInvocation
    : typeArguments "this" arguments ";" {$$=N(JExplicitConstructorInvokcation,$1,$2,$3,$4).set(FThis)}
    | typeArguments "super" arguments ";" {$$=N(JExplicitConstructorInvokcation,$1,$2,$3,$4).set(FSuper)}
    | "this" arguments ";" {$$=N(JExplicitConstructorInvokcation,$1,$2,$3).set(FThis)}
    | "super" arguments ";" {$$=N(JExplicitConstructorInvokcation,$1,$2,$3).set(FSuper)}
    | primary "." typeArguments "super" arguments ";" {$$=N(JExplicitConstructorInvokcation,$1,$2,$3,$4,$5,$6)}
    | expressionName "." typeArguments "super" arguments ";" {$$=N(JExplicitConstructorInvokcation,$1,$2,$3,$4,$5,$6)}
    | primary "." "super" arguments ";" {$$=N(JExplicitConstructorInvokcation,$1,$2,$3,$4,$5)}
    | expressionName "."  "super" arguments ";" {$$=N(JExplicitConstructorInvokcation,$1,$2,$3,$4,$5)}
    ;


arguments
    :"(" argumentList ")" {$$=N(JArguments,$1,$2,$3)}
    |"("  ")" {$$=N(JArguments,$1,$2)}
    ;
argumentList
    : argument {$$=NS($1)}
    | argumentList "," argument{$$=append($1,$3)}
    ;
argument
    :
    ;
/*8.9*/
enumDeclaration
    : classModifiers "enum" typeIdentifier classImplements enumBody {$$=N(JEnumDeclaration,$1,$2,$3,$4,$5)}
    | "enum" typeIdentifier classImplements enumBody {$$=N(JEnumDeclaration,$1,$2,$3,$4)}
    | classModifiers "enum" typeIdentifier enumBody {$$=N(JEnumDeclaration,$1,$2,$3,$4)}
    | "enum" typeIdentifier enumBody {$$=N(JEnumDeclaration,$1,$2,$3)}
    ;
enumBody
    : "{"  "}" {$$=N(JTypeBody,$1,$2).set(FEnum)}
    | "{" "," "}" {$$=N(JTypeBody,$1,$2,$3).set(FEnum)}
    | "{" enumConstantList "}" {$$=N(JTypeBody,$1,$2,$3).set(FEnum)}
    | "{" enumConstantList "," "}" {$$=N(JTypeBody,$1,$2,$4).set(FEnum)}
    | "{" enumConstantList "," enumBodyDeclarations "}" {$$=N(JTypeBody,$1,$2,$3,$4,$5).set(FEnum)}
    | "{" enumBodyDeclarations "}" {$$=N(JTypeBody,$1,$2,$3).set(FEnum)}
    ;
enumConstantList
    : enumConstant {$$=NS($1)}
    | enumConstantList "," enumConstant{$$=append($1,$3)}
    ;
enumConstant
    : identifier   {$$=N(JEnumConstant,$1)}
    | identifier arguments  {$$=N(JEnumConstant,$1,$2)}
    | identifier classBody  {$$=N(JEnumConstant,$1,$2)}
    | annotations identifier arguments  {$$=N(JEnumConstant,$1,$2,$3)}
    | annotations identifier classBody  {$$=N(JEnumConstant,$1,$2,$3)}
    | annotations identifier arguments classBody {$$=N(JEnumConstant,$1,$2,$3,$4)}
    ;
enumBodyDeclarations
    : semi {$$=$1}
    | ";" classBodyDeclarations {$$=N(JEnumBodyDeclarations,$1,$2)}
    ;
/*8.10*/
recordDeclaration
    :  "record" typeIdentifier recordHeader recordBody {$$=N(JRecordDeclaration,$1,$2,$3,$4)}
    | classModifiers "record" typeIdentifier recordHeader recordBody {$$=N(JRecordDeclaration,$1,$2,$3,$4,$5)}
    | classModifiers "record" typeIdentifier typeParameters recordHeader recordBody {$$=N(JRecordDeclaration,$1,$2,$3,$4,$5,$6)}
    | classModifiers "record" typeIdentifier recordHeader classImplements recordBody {$$=N(JRecordDeclaration,$1,$2,$3,$4,$5,$6)}
    | "record" typeIdentifier recordHeader typeParameters classImplements recordBody {$$=N(JRecordDeclaration,$1,$2,$3,$4,$5,$6)}
    | "record" typeIdentifier recordHeader classImplements recordBody {$$=N(JRecordDeclaration,$1,$2,$3,$4,$5)}
    | "record" typeIdentifier typeParameters recordHeader  recordBody {$$=N(JRecordDeclaration,$1,$2,$3,$4,$5)}
    | classModifiers "record" typeIdentifier typeParameters recordHeader classImplements recordBody {$$=N(JRecordDeclaration,$1,$2,$3,$4,$5,$6,$7)}
    ;

recordHeader
    : "("")" {$$=N(JRecordHeader,$1,$2)}
    | "("recordComponentList")" {$$=N(JRecordHeader,$1,$2,$3)}
    | "("variableArityRecordComponent")" {$$=N(JRecordHeader,$1,$2,$3)}
    | "("recordComponentList "," variableArityRecordComponent")" {$$=N(JRecordHeader,$1,$2,$3,$4,$5)}
    ;
recordComponentList
    : recordComponent {$$=NS($1)}
    | recordComponentList "," recordComponent{$$=append($1,$3)}
    ;
recordComponent
    : annotations unannType identifier {$$=N(JRecordComponent,$1,$2,$3)}
    |  unannType identifier {$$=N(JRecordComponent,$1,$2)}
    ;
variableArityRecordComponent
    : annotations unannType annotations "..." identifier {$$=N(JRecordComponent,$1,$2,$3,$4,$5).set(FVararg)}
    | unannType annotations "..." identifier {$$=N(JRecordComponent,$1,$2,$3,$4).set(FVararg)}
    | annotations unannType "..." identifier {$$=N(JRecordComponent,$1,$2,$3,$4).set(FVararg)}
    | unannType "..." identifier {$$=N(JRecordComponent,$1,$2,$3).set(FVararg)}
    ;
recordBody
    : "{"  "}" {$$=N(JTypeBody,$1,$2).set(FRecord)}
    | "{" recordBodyDeclarationList "}" {$$=N(JTypeBody,$1,$2).set(FRecord)}
    ;
recordBodyDeclarationList
    : recordBodyDeclaration {$$=NS($1)}
    | recordBodyDeclarationList  recordBodyDeclaration{$$=append($1,$2)}
    ;
recordBodyDeclaration
    : classBodyDeclaration {$$=$1}
    | compactConstructorDeclaration {$$=$1}
    ;
compactConstructorDeclaration
    : constructorModifiers typeIdentifier constructorBody {$$=N(JCompactConstructorDeclaration,$1,$2,$3)}
    | typeIdentifier constructorBody {$$=N(JCompactConstructorDeclaration,$1,$2,$3)}
    ;
/* 9.1*/
interfaceDeclaration
    : normalInterfaceDeclaration {$$=$1}
//    | annotationInterfaceDeclaration {$$=$1}
    ;
normalInterfaceDeclaration
    : "interface" typeIdentifier interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3)}
    |  "interface" typeIdentifier typeParameters interfacePermits interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5)}
    |  "interface" typeIdentifier typeParameters interfaceExtends interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5)}
    |  "interface" typeIdentifier interfaceExtends interfacePermits interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5)}
    |  "interface" typeIdentifier typeParameters interfaceExtends interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5)}
    | interfaceModifiers "interface" typeIdentifier interfaceExtends interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5)}
    | interfaceModifiers "interface" typeIdentifier typeParameters interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5)}
    | interfaceModifiers "interface" typeIdentifier interfacePermits interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5)}
    | interfaceModifiers "interface" typeIdentifier interfaceExtends interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5)}
    | interfaceModifiers "interface" typeIdentifier typeParameters interfacePermits interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5,$6)}
    | interfaceModifiers "interface" typeIdentifier typeParameters interfaceExtends interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5,$6)}
    | interfaceModifiers "interface" typeIdentifier interfaceExtends interfacePermits interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5,$6)}
    | interfaceModifiers "interface" typeIdentifier typeParameters interfaceExtends interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5,$6)}
    | "interface" typeIdentifier typeParameters interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4)}
    | "interface" typeIdentifier interfacePermits interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4)}
    | "interface" typeIdentifier interfaceExtends interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4)}
    | interfaceModifiers "interface" typeIdentifier interfaceBody  {$$=N(JInterfaceDeclaration,$1,$2,$3,$4)}
    | interfaceModifiers "interface" typeIdentifier typeParameters interfaceExtends interfacePermits interfaceBody
     {$$=N(JInterfaceDeclaration,$1,$2,$3,$4,$5,$6,$7)}
    ;
interfaceModifiers
    : interfaceModifier {$$=N(JModifiers,$1)}
    | interfaceModifiers interfaceModifier{$$=$1.add($2)}
    ;
interfaceModifier
    : annotation {$$=$1}
    | "public" {$$=N(JModifier,$1)}
    | "protected" {$$=N(JModifier,$1)}
    | "private" {$$=N(JModifier,$1)}
    | "abstract" {$$=N(JModifier,$1)}
    | "static" {$$=N(JModifier,$1)}
    | "sealed" {$$=N(JModifier,$1)}
    | "non-sealed" {$$=N(JModifier,$1)}
    | "strictfp" {$$=N(JModifier,$1)}
    ;
interfaceExtends
    : "extends" interfaceTypeList {$$=N(JClassExtends,$1,$2).set(FInterface).set(FImplements)}
    ;

interfacePermits
    : classPermits {$$=$1.set(FInterface)}
    ;

interfaceBody
    : "{" "}" {$$=N(JTypeBody,$1,$2).set(FInterface)}
    | "{" interfaceMemberDeclarationList "}" {$$=N(JTypeBody,$1,$2,$3).set(FInterface)}
    ;
interfaceMemberDeclarationList
    : interfaceMemberDeclaration {$$=NS($1)}
    | interfaceMemberDeclarationList interfaceMemberDeclaration{$$=append($1,$2)}
    ;
interfaceMemberDeclaration
    : constantDeclaration {$$=$1}
    | interfaceMethodDeclaration {$$=$1}
    | classDeclaration {$$=$1}
    | interfaceDeclaration {$$=$1}
    | semi {$$=$1}
    ;
/* 9.3*/
constantDeclaration
    : unannType variableDeclaratorList ";" {$$=N(JConstantDeclaration,$1,$2)}
    |constantModifiers unannType variableDeclaratorList ";" {$$=N(JConstantDeclaration,$1,$2,$3)}
    ;
constantModifiers
    : constantModifier {$$=N(JModifiers,$1)}
    | constantModifiers constantModifier{$$=$1.add($2)}
    ;
constantModifier
    : annotation {$$=$1}
    | "public" {$$=N(JModifier,$1)}
    | "static" {$$=N(JModifier,$1)}
    | "final" {$$=N(JModifier,$1)}
    ;
/*9.4 */
interfaceMethodDeclaration
    :  methodHeader methodBody {$$=N(JMethodDeclaration,$1,$2).set(FInterface)}
    | interfaceMethodModifiers methodHeader methodBody {$$=N(JMethodDeclaration,$1,$2,$3).set(FInterface)}
    ;
interfaceMethodModifiers
    : interfaceMethodModifier {$$=N(JModifiers,$1)}
    | interfaceMethodModifiers interfaceMethodModifier{$$=$1.add($2)}
    ;
interfaceMethodModifier
    : annotation {$$=$1}
    | "public" {$$=N(JModifier,$1)}
    | "private" {$$=N(JModifier,$1)}
    | "abstract" {$$=N(JModifier,$1)}
    | "default" {$$=N(JModifier,$1)}
    | "static" {$$=N(JModifier,$1)}
    | "strictfp" {$$=N(JModifier,$1)}
    ;
/* 9.6*/
annotationInterfaceDeclaration
    : "@" "interface" typeIdentifier annotationInterfaceBody {$$=N(JAnnotationDeclaration,$1,$2,$3,$4)}
    | interfaceModifiers "@" "interface" typeIdentifier annotationInterfaceBody {$$=N(JAnnotationDeclaration,$1,$2,$3,$4,$5)}
    ;

annotationInterfaceBody
    : "{""}" {$$=N(JTypeBody,$1,$2).set(FAnnotation)}
    | "{" annotationInterfaceMemberDeclarationList"}" {$$=N(JTypeBody,$1,$2,$3).set(FAnnotation)}
    ;
 annotationInterfaceMemberDeclarationList
    :  annotationInterfaceMemberDeclaration {$$=NS($1)}
    |  annotationInterfaceMemberDeclarationList  annotationInterfaceMemberDeclaration{$$=append($1,$2)}
    ;
annotationInterfaceMemberDeclaration
    : annotationInterfaceElementDeclaration {$$=$1}
    | constantDeclaration {$$=$1}
    | classDeclaration {$$=$1}
    | interfaceDeclaration {$$=$1}
    | semi {$$=$1}
    ;

annotationInterfaceElementDeclaration
    : unannType identifier "("")"  ";"  {$$=N(JAnnotationElementDeclaration,$1,$2,$3,$4,$5)}
    | unannType identifier "("")" dims defaultValue ";" {$$=N(JAnnotationElementDeclaration,$1,$2,$3,$4,$5,$6,$7)}
    |annotationInterfaceElementModifiers unannType identifier "("")" dims ";"
    {$$=N(JAnnotationElementDeclaration,$1,$2,$3,$4,$5,$6,$7)}
    |annotationInterfaceElementModifiers unannType identifier "("")" defaultValue ";"
    {$$=N(JAnnotationElementDeclaration,$1,$2,$3,$4,$5,$6,$7)}
    |annotationInterfaceElementModifiers unannType identifier "("")" dims defaultValue ";"
    {$$=N(JAnnotationElementDeclaration,$1,$2,$3,$4,$5,$6,$7,$8)}
    ;
annotationInterfaceElementModifiers
    : annotationInterfaceElementModifier {$$=N(JModifiers,$1)}
    | annotationInterfaceElementModifiers annotationInterfaceElementModifier{$$=$1.add($2)}
    ;
annotationInterfaceElementModifier
    : annotation {$$=$1}
    | "public" {$$=N(JModifier,$1)}
    | "abstract" {$$=N(JModifier,$1)}
    ;

defaultValue
    : "default" elementValue {$$=N(JDefaultValue,$1,$2)}
    ;
/* 9.7 */
annotations
    : annotation {$$=N(JAnnotations,$1)}
    | annotations annotation{$$=$1.add($2)}
    ;
annotation
    : normalAnnotation {$$=$1}
    | markerAnnotation {$$=$1}
    | singleElementAnnotation {$$=$1}
    ;

normalAnnotation
    : "@" typeName "(" ")" {$$=N(JAnnotation,$1,$2,$3,$4)}
    |"@" typeName "(" elementValuePairList ")" {$$=N(JAnnotation,$1,$2,$3,$4,$5)}
    ;
elementValuePairList
    : elementValuePair {$$=NS($1)}
    | elementValuePairList "," elementValuePair {$$=append($1,$3)}
    ;

elementValuePair
    : identifier "=" elementValue {$$=N(JElementValuePair,$1,$2,$3)}
    ;

elementValue
    : conditionalExpression {$$=$1}
    | elementValueArrayInitializer {$$=$1}
    | annotation {$$=$1}
    ;

elementValueArrayInitializer
    : "{" "}" {$$=N(JElementValueArrayInitializer,$2)}
    | "{" "," "}" {$$=N(JElementValueArrayInitializer,$1,$2,$3)}
    | "{" elementValueList "}" {$$=N(JElementValueArrayInitializer,$1,$2,$3)}
    | "{" elementValueList "," "}" {$$=N(JElementValueArrayInitializer,$1,$2,$3,$4)}
    ;
elementValueList
    : elementValue {$$=NS($1)}
    | elementValueList "," elementValue{$$=append($1,$3)}
    ;

markerAnnotation
    : "@" typeName {$$=N(JAnnotation,$1,$2)}
    ;

singleElementAnnotation
    : "@" typeName "(" elementValue ")" {$$=N(JAnnotation,$1,$2,$3,$4,$5)}
    ;
/* 10.6*/
arrayInitializer
    : "{" "}" {$$=N(JArrayInitializer,$1,$2)}
    | "{" "," "}" {$$=N(JArrayInitializer,$1,$2,$3)}
    | "{" variableInitializerList "}" {$$=N(JArrayInitializer,$1,$2,$3)}
    | "{" variableInitializerList "," "}" {$$=N(JArrayInitializer,$1,$2,$3,$4)}
    ;
variableInitializerList
    : variableInitializer {$$=NS($1)}
    | variableInitializerList "," variableInitializer{$$=append($1,$3)}
    ;
/* 14.2*/
block
    : "{" "}" {$$=N(JBlock,$1,$2)}
    | "{" blockStatementList "}" {$$=N(JBlock,$1,$2,$3)}
    ;
blockStatementList
    : blockStatement {$$=NS($1)}
    | blockStatementList blockStatement{$$=append($1,$2)}
    ;
blockStatement
    : localClassOrInterfaceDeclaration {$$=$1}
    | localVariableDeclarationStatement {$$=$1}
    | statement {$$=$1}
    ;
/* 14.3*/
localClassOrInterfaceDeclaration
    : classDeclaration {$$=$1}
    | normalInterfaceDeclaration {$$=$1}
    ;
/* 14.4*/
localVariableDeclaration
    : variableModifiers localVariableType {$$=N(JLocalVariableDeclaration,$1,$2)}
    | localVariableType {$$=N(JLocalVariableDeclaration,$1)}
    | variableModifiers localVariableType variableDeclaratorList {$$=N(JLocalVariableDeclaration,$1,$2,$3)}
    ;

localVariableType
    : unannType {$$=$1}
    | "var" {$$=N(JVar,$1)}
    ;

localVariableDeclarationStatement
    : localVariableDeclaration ";"
    ;
/* 14.5*/
statement
    : statementWithoutTrailingSubstatement {$$=$1}
    | labeledStatement {$$=$1}
    | ifThenStatement {$$=$1}
    | ifThenElseStatement {$$=$1}
    | whileStatement {$$=$1}
    | forStatement {$$=$1}
    ;

statementNoShortIf
    : statementWithoutTrailingSubstatement {$$=$1}
    | labeledStatementNoShortIf {$$=$1}
    | ifThenElseStatementNoShortIf {$$=$1}
    | whileStatementNoShortIf {$$=$1}
    | forStatementNoShortIf {$$=$1}
    ;

statementWithoutTrailingSubstatement
    : block {$$=$1}
    | emptyStatement_ {$$=$1}
    | expressionStatement {$$=$1}
    | assertStatement {$$=$1}
    | switchStatement {$$=$1}
    | doStatement {$$=$1}
    | breakStatement {$$=$1}
    | continueStatement {$$=$1}
    | returnStatement {$$=$1}
    | synchronizedStatement {$$=$1}
    | throwStatement {$$=$1}
    | tryStatement {$$=$1}
    | yieldStatement {$$=$1}
    ;

/*  14.6 */
emptyStatement_
    : semi {$$=$1}
    ;
/* 14.7 */
labeledStatement
    : identifier ":" statement {$$=N(JLabeledStmt,$1,$2)}
    ;

labeledStatementNoShortIf
    : identifier ":" statementNoShortIf {$$=N(JLabeledStmt,$1,$2)}
    ;
/* 14.8 */
expressionStatement
    : statementExpression ";" {$$=N(JExprStmt,$1,$2)}
    ;

statementExpression
    : assignment {$$=$1}
    | preIncrementExpression {$$=$1}
    | preDecrementExpression {$$=$1}
    | postIncrementExpression {$$=$1}
    | postDecrementExpression {$$=$1}
    | methodInvocation {$$=$1}
    | classInstanceCreationExpression {$$=$1}
    ;
/* 14.9 */
ifThenStatement
    : "if" "(" expression ")" statement {$$=N(JIfThenStmt,$1,$2,$3,$4)}
    ;

ifThenElseStatement
    : "if" "(" expression ")" statementNoShortIf "else" statement {$$=N(JIfThenStmt,$1,$2,$3,$4,$5,$6,$7)}
    ;

ifThenElseStatementNoShortIf
    : "if" "(" expression ")" statementNoShortIf "else" statementNoShortIf {$$=N(JIfThenStmt,$1,$2,$3,$4,$5,$6,$7)}
    ;
/*14.10*/
assertStatement
    : "assert" expression  ";" {$$=N(JAssertStmt,$1,$2,$3)}
    | "assert" expression ":" expression ";" {$$=N(JAssertStmt,$1,$2,$3,$4,$5)}
    ;
/*14.11*/
switchStatement
    : "switch" "(" expression ")" switchBlock {$$=N(JSwitchStmt,$1,$2,$3,$4,$5)}
    ;

switchBlock
    : "{" switchRuleList "}" {$$=N(JSwitchBlock,$1,$2,$3)}
    | "{" switchBlockStatementGroupList switchLabelList "}" {$$=N(JSwitchBlock,$1,$2,$3,$4)}
    | "{" switchBlockStatementGroupList "}" {$$=N(JSwitchBlock,$1,$2,$3,$4)}
    ;
switchRuleList
    : switchRule {$$=NS($1)}
    | switchRuleList switchRule{$$=append($1,$2)}
    ;
switchRule
    : "case" caseConstantList "->" expression ";" {$$=N(JSwitchCaseRule,$1,$2,$3,$4,$5)}
    | "case" caseConstantList "->" block {$$=N(JSwitchCaseRule,$1,$2,$3,$4)}
    | "case" caseConstantList "->" throwStatement {$$=N(JSwitchCaseRule,$1,$2,$3,$4)}
    | "default"  "->" expression ";" {$$=N(JSwitchDefaultRule,$1,$2,$3,$4)}
    | "default"  "->" block {$$=N(JSwitchDefaultRule,$1,$2,$3)}
    | "default"  "->" throwStatement {$$=N(JSwitchDefaultRule,$1,$2,$3)}
    ;
switchBlockStatementGroupList
    : switchBlockStatementGroup {$$=NS($1)}
    | switchBlockStatementGroupList switchBlockStatementGroup{$$=append($1,$2)}
    ;
switchBlockStatementGroup
    : switchLabelList blockStatementList {$$=N(JSwitchBlockGroup,$1,$2)}
    ;
switchLabelList
    : switchLabel {$$=NS($1)}
    | switchLabelList switchLabel{$$=append($1,$2)}
    ;
switchLabel
    : "case" caseConstantList ":" {$$=N(JSwitchCaseLabel,$1,$2,$3)}
    | "default" ":" {$$=N(JSwitchDefaultLabel,$1,$2)}
    ;
caseConstantList
    : caseConstant {$$=NS($1)}
    | caseConstantList "," caseConstant{$$=append($1,$3)}
    ;
caseConstant
    : conditionalExpression {$$=$1}
    ;
/*  14.12 */
whileStatement
    : "while" "(" expression ")" statement {$$=N(JWhileStmt,$1,$2,$3,$4,$5)}
    ;
whileStatementNoShortIf
    : "while" "(" expression ")" statementNoShortIf {$$=N(JWhileStmt,$1,$2,$3,$4,$5)}
    ;
/* 14.13*/
doStatement
    : "do" statement "while" "(" expression ")" ";" {$$=N(JDoStmt,$1,$2,$3,$4,$5,$6,$7)}
    ;
/* 14.14*/
forStatement
    : basicForStatement {$$=$1}
    | enhancedForStatement {$$=$1}
    ;

forStatementNoShortIf
    : basicForStatementNoShortIf {$$=$1}
    | enhancedForStatementNoShortIf {$$=$1}
    ;
forHeader
    :"for" "("  ";"  ";"  ")" {$$=Anys($1,$2,$3,$4,$5)}
    |"for" "(" forInit ";"  ";"  ")" {$$=Anys($1,$2,$3,$4,$5,$6)}
    |"for" "("";" expression ";"  ")" {$$=Anys($1,$2,$3,$4,$5,$6)}
    |"for" "(" ";"  ";" forUpdate ")" {$$=Anys($1,$2,$3,$4,$5,$6)}
    |"for" "(" forInit ";" expression ";"  ")" {$$=Anys($1,$2,$3,$4,$5,$6,$7)}
    |"for" "("  ";" expression ";" forUpdate ")" {$$=Anys($1,$2,$3,$4,$5,$6,$7)}
    |"for" "(" forInit ";"  ";" forUpdate ")" {$$=Anys($1,$2,$3,$4,$5,$6,$7)}
    |"for" "(" forInit ";" expression ";" forUpdate ")" {$$=Anys($1,$2,$3,$4,$5,$6,$7,$8)}
    ;
basicForStatement
    : forHeader statement {$$=N(JForStmt,$1[0]).add($1...).add($2)}
    ;

basicForStatementNoShortIf
    : forHeader statementNoShortIf {$$=N(JForStmt,$1[0]).add($1...).add($2)}
    ;

forInit
    : statementExpressionList {$$=$1}
    | localVariableDeclaration {$$=$1}
    ;

forUpdate
    : statementExpressionList {$$=$1}
    ;
statementExpressionList
    : statementExpression {$$=NS($1)}
    | statementExpressionList "," statementExpression{$$=append($1,$3)}
    ;

enhancedForStatement
    : "for" "(" localVariableDeclaration ":" expression ")" statement {$$=N(JForStmt,$1,$2,$3,$4,$5,$6,$7)}
    ;

enhancedForStatementNoShortIf
    : "for" "(" localVariableDeclaration ":" expression ")" statementNoShortIf {$$=N(JForStmt,$1,$2,$3,$4,$5,$6,$7)}
    ;
/* 14.15*/
breakStatement
    : "break"  ";" {$$=N(JBreakStmt,$1,$2)}
    | "break" identifier ";" {$$=N(JBreakStmt,$1,$2,$3)}
    ;
/* 14.16*/
continueStatement
    : "continue"  ";" {$$=N(JContinueStmt,$1,$2)}
    | "continue" identifier ";" {$$=N(JContinueStmt,$1,$2,$3)}
    ;
returnStatement
    : "return"  ";" {$$=N(JReturnStmt,$1,$2)}
    | "return" expression ";" {$$=N(JReturnStmt,$1,$2,$3)}
    ;
throwStatement
    : "throw" expression ";" {$$=N(JThrowStmt,$1,$2,$3)}
    ;
synchronizedStatement
    : "synchronized" "(" expression ")" block {$$=N(JSyncStmt,$1,$2,$3,$4,$5)}
    ;
/* 14.20*/
tryStatement
    : "try" block catches {$$=N(JTryStmt,$1,$2,$3)}
    | "try" block catches finallyBlock {$$=N(JTryStmt,$1,$2,$3,$4)}
    | "try" block finallyBlock {$$=N(JTryStmt,$1,$2,$3)}
    | tryWithResourcesStatement {$$=$1}
    ;
catches
    : catchClause {$$=N(JCatchClauses,$1)}
    | catches catchClause{$$=$1.add($2)}
    ;
catchClause
    : "catch" "(" catchFormalParameter ")" block {$$=N(JCatchClause,$1,$2,$3,$4,$5)}
    ;
catchFormalParameter
    : variableModifiers catchType variableDeclaratorId {$$=N(JCatchFormalParameter,$1,$2,$3)}
    | catchType variableDeclaratorId {$$=N(JCatchFormalParameter,$1,$2)}
    ;
catchType
    : unannClassType {$$=N(JCatchType,$1)}
    |catchType "|" classType {$$=$1.add($3)}
    ;
finallyBlock
    : "finally" block {$$=N(JFinallyBlock,$1,$2)}
    ;
tryWithResourcesStatement
    :"try" resourceSpecification block catches {$$=N(JTryResourceStmt,$1,$2,$3,$4)}
    |"try" resourceSpecification block finallyBlock {$$=N(JTryResourceStmt,$1,$2,$3,$4)}
    |"try" resourceSpecification block catches finallyBlock {$$=N(JTryResourceStmt,$1,$2,$3,$4,$5)}
    ;
resourceSpecification
    : "(" resourceList ")" {$$=N(JResourceSpec,$1,$2,$3)}
    | "(" resourceList ";" ")" {$$=N(JResourceSpec,$1,$2,$3)}
    ;
resourceList
    : resource {$$=NS($1)}
    | resourceList ";" resource{$$=append($1,$3)}
    ;
resource
    : localVariableDeclaration {$$=$1}
    | expressionName {$$=$1}
    | fieldAccess  {$$=$1}
    ;
yieldStatement
    : "yield" expression ";" {$$=N(JYieldStmt,$1,$2,$3)}
    ;
pattern
    : localVariableDeclaration {$$=$1}
    ;

expression
    : lambdaExpression {$$=$1}
    | assignmentExpression {$$=$1}
    ;
primary
    : primaryNoNewArray {$$=$1}
    | arrayCreationExpression {$$=$1}
    ;
primaryNoNewArray
    : literal {$$=N(JLiteralExpr,$1)}
    | classLiteral {$$=N(JLiteralExpr,$1)}
    | "this" {$$=N(JThisExpr,$1)}
    | typeName "." "this" {$$=N(JThisExpr,$1,$2,$3)}
    | "(" expression ")" {$$=N(JParenExpr,$1,$2,$3)}
    | unqualifiedClassInstanceCreationExpression {$$=N(JObjectCreationExpr,$1)}
    | expressionName "." unqualifiedClassInstanceCreationExpression {$$=N(JObjectCreationExpr,$1,$2,$3)}
    | arrayCreationExpression "." unqualifiedClassInstanceCreationExpression {$$=N(JObjectCreationExpr,$1,$2,$3)}
    | arrayCreationExpression "." identifier  {$$=N(JMemberAccessExpr,$1,$2)}
    | "super" "." identifier  {$$=N(JMemberAccessExpr,$1,$2,$3).set(FSuper)}
    | typeName "." "super" "." identifier {$$=N(JMemberAccessExpr,$1,$2,$3,$4,$5).set(FSuper)}
    | expressionName "[" expression "]" {$$=N(JArrayAccessExpr,$1,$2,$3,$4)}
    | arrayCreationExpressionWithInitializer "[" expression "]" {$$=N(JArrayAccessExpr,$1,$2,$3,$4)}
    | methodName arguments {$$=N(JMethodInvocationExpr,$1,$2)}
    | typeName "."  identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4)}
    | typeName "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5)}
    | expressionName "."  identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4)}
    | expressionName "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5)}
    | arrayCreationExpression "."  identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4)}
    | arrayCreationExpression "." typeArguments identifier arguments  {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5)}
    | "super" "." identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4).set(FSuper)}
    | "super" "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5).set(FSuper)}
    | typeName "." "super" "."  identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5,$6).set(FSuper)}
    | typeName "." "super" "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5,$6,$7).set(FSuper)}
    | expressionName "::"  identifier {$$=N(JMethodRefernceExpr,$1,$2,$3)}
    | expressionName "::" typeArguments identifier {$$=N(JMethodRefernceExpr,$1,$2,$3,$4)}
    | arrayCreationExpression "::" identifier {$$=N(JMethodRefernceExpr,$1,$2,$3)}
    | arrayCreationExpression "::" typeArguments identifier {$$=N(JMethodRefernceExpr,$1,$2,$3,$4)}
    | referenceType "::"  identifier {$$=N(JMethodRefernceExpr,$1,$2,$3)}
    | referenceType "::" typeArguments identifier{$$=N(JMethodRefernceExpr,$1,$2,$3,$4)}
    | "super" "::"  identifier  {$$=N(JMethodRefernceExpr,$1,$2,$3).set(FSuper)}
    | "super" "::" typeArguments identifier  {$$=N(JMethodRefernceExpr,$1,$2,$3,$4).set(FSuper)}
    | typeName "." "super" "::" typeArguments identifier {$$=N(JMethodRefernceExpr,$1,$2,$3,$4,$5,$6).set(FSuper)}
    | typeName "." "super" "::" typeArguments identifier {$$=N(JMethodRefernceExpr,$1,$2,$3,$4,$5).set(FSuper)}
    | classType "::"  "new" {$$=N(JCreatorReferenceExpr,$1,$2,$3,$4)}
    | classType "::" typeArguments "new" {$$=N(JCreatorReferenceExpr,$1,$2,$3)}
    | arrayType "::" "new" {$$=N(JCreatorReferenceExpr,$1,$2,$3)}
    | primaryNoNewArray pNNA {$$=$1.add($2...)}
    ;
pNNA
    : "." unqualifiedClassInstanceCreationExpression {$$=Anys($1,$2)}
    | "." identifier {$$=Anys($1,$2)}
    | "."  identifier arguments {$$=Anys($1,$2,$3)}
    | "." typeArguments identifier arguments {$$=Anys($1,$2,$3,$4)}
    | "[" expression "]" {$$=Anys($1,$2,$3)}
    | "::" identifier{$$=Anys($1,$2)}
    | "::" typeArguments identifier{$$=Anys($1,$2,$3)}
    | pNNA pNNA {$$=append($1,$2...)}
    ;
classLiteral
    : typeName "." "class" {$$=N(JLiteral,$1,$2,$3)}
    | typeName dims "." "class" {$$=N(JLiteral,$1,$2,$3,$4)}
    | numericType "." "class" {$$=N(JLiteral,$1,$2,$3)}
    | numericType dims "." "class" {$$=N(JLiteral,$1,$2,$3,$4)}
    | "boolean" "." "class" {$$=N(JLiteral,$1,$2,$3)}
    | "boolean" dims "." "class" {$$=N(JLiteral,$1,$2,$3,$4)}
    | "void" "." "class" {$$=N(JLiteral,$1,$2,$3)}
    ;
classInstanceCreationExpression
    : unqualifiedClassInstanceCreationExpression {$$=$1}
    | expressionName "." unqualifiedClassInstanceCreationExpression {$$=N(JObjectCreationExpr,$1,$2,$3)}
    | primary "." unqualifiedClassInstanceCreationExpression {$$=N(JObjectCreationExpr,$1,$2,$3)}
    ;

unqualifiedClassInstanceCreationExpression
    : "new"  classOrInterfaceTypeToInstantiate arguments {$$=N(JObjectCreationExpr,$1,$2,$3)}
    | "new" typeArguments classOrInterfaceTypeToInstantiate arguments {$$=N(JObjectCreationExpr,$1,$2,$3,$4)}
    | "new" classOrInterfaceTypeToInstantiate arguments classBody {$$=N(JObjectCreationExpr,$1,$2,$3,$4)}
    | "new" typeArguments classOrInterfaceTypeToInstantiate arguments classBody {$$=N(JObjectCreationExpr,$1,$2,$3,$4,$5)}
    ;

classOrInterfaceTypeToInstantiate
    : annoIdentifierList {$$=N(JClassInstatiateType,$1)}
    | annoIdentifierList typeArgumentsOrDiamond {$$=N(JClassInstatiateType,$1,$2)}
    ;
annoIdentifierList
    : annoIdentifier {$$=NS($1)} 
    | annoIdentifierList "." annoIdentifier{$$=append($1,$3)}
    ;
annoIdentifier
    : annotations identifier {$$=N(JAnnoIdentifier,$1,$2)}
    ;
typeArgumentsOrDiamond
    : typeArguments
    | "<" ">" {$$=N(JTypeArguments,$1,$2)}
    ;
arrayCreationExpression
      : "new" primitiveType dimExprs  {$$=N(JArrayCreationExpr,$1,$2,$3)}
      | "new" primitiveType dimExprs dims {$$=N(JArrayCreationExpr,$1,$2,$3,$4)}
      | "new" classType dimExprs  {$$=N(JArrayCreationExpr,$1,$2,$3)}
      | "new" classType dimExprs dims {$$=N(JArrayCreationExpr,$1,$2,$3,$4)}
      | "new" primitiveType dims arrayInitializer {$$=N(JArrayCreationExpr,$1,$2,$3,$4)}
      | "new" classOrInterfaceType dims arrayInitializer {$$=N(JArrayCreationExpr,$1,$2,$3,$4)}
    ;

dimExprs
    : dimExpr {$$=N(,$1)}
    | dimExprs dimExpr{$$=$1.add($2)}
    ;

dimExpr
    :  "[" expression "]" {$$=N(JDimExpr,$1,$2,$3,$4)}
    | annotations "[" expression "]" {$$=N(JDimExpr,$1,$2,$3)}
    ;

arrayAccess
    : expressionName "[" expression "]"
    | primaryNoNewArray "[" expression "]"
    | arrayCreationExpressionWithInitializer "[" expression "]"
    ;

fieldAccess
    : primary "." identifier {$$=N(JMemberAccessExpr,$1,$2,$3)}
    | "super" "." identifier {$$=N(JMemberAccessExpr,$1,$2,$3).set(FSuper)}
    | typeName "." "super" "." identifier {$$=N(JMemberAccessExpr,$1,$2,$3,$4,$5).set(FSuper)}
    ;
methodInvocation
    : methodName arguments {$$=N(JMethodInvocationExpr,$1,$2)}
    | typeName "."  identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4)}
    | typeName "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5)}
    | expressionName "."  identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4)}
    | expressionName "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5)}
    | primary "." identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4)}
    | primary "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5)}
    | "super" "." identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4).set(FSuper)}
    | "super" "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5).set(FSuper)}
    | typeName "." "super" "."  identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5,$5).set(FSuper)}
    | typeName "." "super" "." typeArguments identifier arguments {$$=N(JMethodInvocationExpr,$1,$2,$3,$4,$5,$6,$7).set(FSuper)}
    ;
methodReference
    : expressionName "::"  identifier {$$=N(JMethodReferenceExpr,$1,$2,$3)}
    | expressionName "::" typeArguments identifier {$$=N(JMethodReferenceExpr,$1,$2,$3,$4)}
    | primary "::"  identifier {$$=N(JMethodReferenceExpr,$1,$2,$3)}
    | primary "::" typeArguments identifier {$$=N(JMethodReferenceExpr,$1,$2,$3,$4)}
    | referenceType "::"  identifier {$$=N(JMethodReferenceExpr,$1,$2,$3)}
    | referenceType "::" typeArguments identifier {$$=N(JMethodReferenceExpr,$1,$2,$3,$4)}
    | "super" "::"  identifier {$$=N(JMethodReferenceExpr,$1,$2,$3).set(FSuper)}
    | "super" "::" typeArguments identifier {$$=N(JMethodReferenceExpr,$1,$2,$3,$4).set(FSuper)}
    | typeName "." "super" "::"  identifier  {$$=N(JMethodReferenceExpr,$1,$2,$3,$4).set(FSuper)}
    | typeName "." "super" "::" typeArguments identifier  {$$=N(JMethodReferenceExpr,$1,$2,$3,$4).set(FSuper)}
    | classType "::"  "new" {$$=N(JCreatorReferenceExpr,$1,$2,$3)}
    | classType "::" typeArguments "new" {$$=N(JCreatorReferenceExpr,$1,$2,$3,$4)}
    | arrayType "::" "new" {$$=N(JCreatorReferenceExpr,$1,$2,$3)}
    ;
postfixExpression
    : primary  {$$=$1}
    | expressionName  {$$=$1}
    | primary pfE {$$=N(JPostfixExpr,$1,$2)}
    | expressionName pfE {$$=N(JPostfixExpr,$1,$2)}
    ;
pfE
    : "+""+" {$$=N(JPostfixOperator,$1,$2).set(FPostPlus)}
    |  "-""-" {$$=N(JPostfixOperator,$1,$2).set(FPostMinus)}
    | pfE pfE {$$=$1.add($2)}
    ;

postIncrementExpression
    : postfixExpression  "+""+" {$$=N(JPostfixExpr,$1,N(JPostOperater,$2,$3).set(FIncr))}
    ;

postDecrementExpression
    : postfixExpression  "-""-" {$$=N(JPostfixExpr,$1,N(JPostOperater,$2,$3).set(FDecr))}
    ;
unaryExpression
    : preIncrementExpression {$$=$1}
    | preDecrementExpression {$$=$1}
    | "+" unaryExpression {$$=N(JUnaryExpr,$1,$2).set(FUnaryPlus)}
    | "-" unaryExpression {$$=N(JUnaryExpr,$1,$2).set(FUnaryMinus)}
    | unaryExpressionNotPlusMinus {$$=$1}
    ;
preIncrementExpression
    :  "+""+" unaryExpression {$$=N(JPrefixExpr,$1,$2,$3).set(FIncr)}
    ;

preDecrementExpression
    : "-""-" unaryExpression {$$=N(JPrefixExpr,$1,$2,$3).set(FDecr)}
    ;

unaryExpressionNotPlusMinus
    : postfixExpression {$$=$1}
    | "~" unaryExpression {$$=N(JPrefixExpr,$1,$2).set(FBitNot)}
    | "!" unaryExpression  {$$=N(JPrefixExpr,$1,$2).set(FNot)}
    | castExpression {$$=$1}
    | switchExpression {$$=$1}
    ;
castExpression
    : "(" primitiveType ")" unaryExpression {$$=N(JCastExpr,$1,$2,$3,$4)}
    | "(" referenceType  ")" unaryExpressionNotPlusMinus {$$=N(JCastExpr,$1,$2,$3,$4)}
    | "(" referenceType intersectionBounds ")" unaryExpressionNotPlusMinus {$$=N(JCastExpr,$1,$2,$3,$4,$5)}
    | "(" referenceType intersectionBounds ")" lambdaExpression {$$=N(JCastExpr,$1,$2,$3,$4,$5)}
    | "(" referenceType  ")" lambdaExpression {$$=N(JCastExpr,$1,$2,$3,$4)}
    ;
multiplicativeExpression
    : unaryExpression {$$=$1}
    | multiplicativeExpression "*" unaryExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FMul)}
    | multiplicativeExpression "/" unaryExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FDiv)}
    | multiplicativeExpression "%" unaryExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FRem)}
    ;
additiveExpression
    : multiplicativeExpression {$$=$1}
    | additiveExpression "+" multiplicativeExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FPlus)}
    | additiveExpression "-" multiplicativeExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FMinus)}
    ;
shiftExpression
    : additiveExpression {$$=$1}
    | shiftExpression "<" "<" additiveExpression {$$=N(JBinaryExpr,$1,$2,$3,$4).set(FSHL)}
    | shiftExpression ">" ">" additiveExpression {$$=N(JBinaryExpr,$1,$2,$3,$4).set(FSHR)}
    | shiftExpression ">" ">" ">" additiveExpression {$$=N(JBinaryExpr,$1,$2,$3,$4,$5).set(FUSHR)}
    ;
relationalExpression
    : shiftExpression {$$=$1}
    | relationalExpression "<" shiftExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FLT)}
    | relationalExpression ">" shiftExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FGT)}
    | relationalExpression "<""=" shiftExpression {$$=N(JBinaryExpr,$1,$2,$3,$4).set(FLE)}
    | relationalExpression ">""=" shiftExpression {$$=N(JBinaryExpr,$1,$2,$3,$4).set(FGE)}
    //      | instanceofExpression
    | relationalExpression "instanceof" referenceType {$$=N(JBinaryExpr,$1,$2,$3).set(FInstanceOf)}
    | relationalExpression "instanceof"  pattern {$$=N(JBinaryExpr,$1,$2,$3).set(FInstanceOf)}
    // Solves left recursion with instanceofExpression.
    ;
equalityExpression
    : relationalExpression {$$=$1}
    | equalityExpression "=""=" relationalExpression {$$=N(JBinaryExpr,$1,$2,$3,$4).set(FEQ)}
    | equalityExpression "!""=" relationalExpression {$$=N(JBinaryExpr,$1,$2,$3,$4).set(FNEQ)}
    ;
andExpression
    : equalityExpression {$$=$1}
    | andExpression "&" equalityExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FBitAnd)}
    ;
exclusiveOrExpression
    : andExpression {$$=$1}
    | exclusiveOrExpression "^" andExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FBitXor)}
    ;
inclusiveOrExpression
    : exclusiveOrExpression {$$=$1}
    | inclusiveOrExpression "|" exclusiveOrExpression {$$=N(JBinaryExpr,$1,$2,$3).set(FBitOr)}
    ;
conditionalAndExpression
    : inclusiveOrExpression {$$=$1}
    | conditionalAndExpression "&""&" inclusiveOrExpression {$$=N(JBinaryExpr,$1,$2,$3,$4).set(FAND)}
    ;
conditionalOrExpression
    : conditionalAndExpression {$$=$1}
    | conditionalOrExpression "|""|" conditionalAndExpression {$$=N(JBinaryExpr,$1,$2,$3,$4).set(FOR)}
    ;
conditionalExpression
    : conditionalOrExpression {$$=$1}
    | conditionalOrExpression "?" expression ":" conditionalExpression {$$=N(JTernaryExpr,$1,$2,$3,$4,$5)}
    | conditionalOrExpression "?" expression ":" lambdaExpression {$$=N(JTernaryExpr,$1,$2,$3,$4,$5)}
    ;
assignmentExpression
    : conditionalExpression {$$=$1}
    | assignment {$$=$1}
    ;
assignment
    : leftHandSide assignmentOperator expression {$$=N(JAssignExpr,$1,$2,$3)}
    ;
leftHandSide
    : expressionName {$$=$1}
    | fieldAccess {$$=$1}
    | arrayAccess {$$=$1}
    ;
assignmentOperator
    : "=" {$$=N(JAssignOp,$1)}
    | "*""=" {$$=N(JAssignMul,$1,$2)}
    | "/""=" {$$=N(JAssignDiv,$1,$2)}
    | "%""=" {$$=N(JAssignRem,$1,$2)}
    | "+""=" {$$=N(JAssignPlus,$1,$2)}
    | "-""=" {$$=N(JAssignMinus,$1,$2)}
    | "<""<""=" {$$=N(JAssignSHL,$1,$2,$3)}
    | ">"">""=" {$$=N(JAssignSHR,$1,$2,$3)}
    | ">"">"">""=" {$$=N(JAssignUSHR,$1,$2,$3,$4)}
    | "&""="{$$=N(JAssignAND,$1,$2)}
    | "^""="{$$=N(JAssignXOR,$1,$2)}
    | "|""="{$$=N(JAssignOR,$1,$2)}
    ;
lambdaExpression
    : lambdaParameters "->" lambdaBody {$$=N(JLambdaExpr,$1,$2,$3)}
    ;
lambdaParameters
    :  "(" ")" {$$=N(JLambdaParameters,$1,$2)}
    | "(" identifierTypeList ")" {$$=N(JLambdaParameters,$1,$2,$3)}
    | "(" lambdaParameterList ")" {$$=N(JLambdaParameters,$1,$2,$3)}
    | "(" variableArityParameter ")" {$$=N(JLambdaParameters,$1,$2,$3)}
    | "(" lambdaParameterList "," variableArityParameter ")" {$$=N(JLambdaParameters,$1,$2,$3,$4,$5)}
    | identifier {$$=N(JLambdaParameters,$1)}
    ;
identifierTypeList
    : identifier {$$=NS($1)}
    | identifierTypeList "," identifier {$$=append($1,$3)}
    ;
lambdaParameterList
    : lambdaParameter {$$=NS($1)}
    | lambdaParameterList "," lambdaParameter{$$=append($1,$3)}
    ;
lambdaParameter
    :  localVariableType variableDeclaratorId {$$=N(JLambdaParameter,$1,$2)}
    | variableModifiers localVariableType variableDeclaratorId {$$=N(JLambdaParameter,$1,$2,$3)}
    ;
lambdaBody
    : expression {$$=$1}
    | block {$$=$1}
    ;
switchExpression
    : "switch" "(" expression ")" switchBlock {$$=N(JSwitchExpr,$1,$2,$3,$4,$5)}
    ;

/*=============extra ==============*/
identifierList
    : identifier {$$=NS($1)}
    | identifierList identifier {$$=append($1,$2)}
    ;

qualifiedName
    : identifier {$$=N(JQualified,$1)}
    | qualifiedName "." identifier {$$=$1.add($3)}
    ;

qualifiedStar:
    identifier {$$=N(JQualified,$1)}
    |qualifiedStar "." identifier {$$=$1.add($3)}
    |qualifiedStar "." STAR {$$=$1.add($3)};

%%
const (
 	JSemi Type=iota
    JCompilationUnit

    JClassDeclaraation
    JClassExtends
    JTypeBody

    JFieldDeclaration
    JPacakageDeclaration

    JMethodDeclaration
    JMethodHeader

    JContructorDeclaration
    JCompactConstructorDeclaration
    JConstructorBody
    JExplicitConstructorInvokcation

    JEnumDeclaration
    JEnumConstant
    JEnumBodyDeclarations

    JInterfaceDeclaration

    JRecordDeclaration
    JRecordHeader
    JRecordComponent

    JConstantDeclaration

    JAnnotationDeclaration
    JAnnotationElementDeclaration
    JDefaultValue

    JAnnotations
    JAnnotation
    JElementValuePair
    JElementValueArrayInitializer
    JArrayInitializer
    JLocalVariableDeclaration
    JVar

    JLabeledStmt
    JExprStmt
    JIfThenStmt
    JAssertStmt
    JSwitchStmt
    JSwitchBlock
    JSwitchBlockGroup
    JSwitchCaseLabel
    JSwitchDefaultLabel
    JSwitchCaseRule
    JSwitchDefaultRule
    JWhileStmt
    JDoStmt
    JForStmt
    JBreakStmt
    JContinueStmt
    JReturnStmt
    JThrowStmt
    JSyncStmt
    JTryStmt
    JCatchClause
    JCatchClauses
    JCatchType
    JFinallyBlock
    JTryResourceStmt
    JResourceSpec
    JYieldStmt

    JParameters
    JReceiverParameter
    JParameter
    JThrowList
    JArguments

    JID
    JLiteral
    JPrimitiveType
    JQualified
    JClassOrInteraceType
    JClassType
    JTypeVariable
    JArrayType
    JTypeParameter
    JTypeBound
    JIntersectionBound
    JTypeArguments
    JWildcard
    JWildcardBounds
    JTypeName

    JGenericTypeName
    JVariableDeclarator
    JVariableDeclaratorId
    JDim
    JDims

    JModifier
    JModifiers
    
    JCreatorReferenceExpr
    JMethodReferenceExpr
    JMethodInvocationExpr
    JArrayAccessExpr
    JMemberAccessExpr
    JObjectCreationExpr
    JParenExpr
    JThisExpr
    JAnnoIdentifier
    JClassInstatiateType
    JArrayCreationExpr
    JDimExpr
    JDimExprs
    JPostfixOperator
    JPostfixExpr
    JPrefixExpr
    JUnaryExpr
    JCastExpr
    JBinaryExpr
    JTernaryExpr
    JAssignExpr
    JAssignOp
    JAssignMul
    JAssignDiv
    JAssignRem
    JAssignPlus
    JAssignMinus
    JAssignSHL
    JAssignSHR
    JAssignUSHR
    JAssignOR
    JAssignXOR
    JAssignAND
    JLambdaExpr
    JSwitchExpr
    JLambdaParameters
    JLambdaParameter
)
const (
 	FNone Modifier =iota
 	FNull
 	FStatic
 	FPrefix
 	FSuper
 	FThis
 	FExtends
 	FPermits
 	FImplements
 	FVararg
 	FInterface
 	FAnnotation
 	FRecord
 	FEnum
 	FIncr
 	FDecr
 	FNot
 	FBitNot
 	FUnaryPlus
    FUnaryMinus
    FMul
    FDiv
    FRem
    FPlus
    FMinus
    FSHR
    FSHL
    FUSHR
    FGT
    FLT
    FGE
    FLE
    FEQ
    FNEQ
    FOR
    FAND
    FBitAnd
    FBitXor
    FBitOr
    FInstanceOf
)
func NS(a ...*Node)[]*Node{
	return a
}
func Anys(a ...any)[]any{
	return a
}
