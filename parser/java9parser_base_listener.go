// Code generated from Java9Parser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Java9Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJava9ParserListener is a complete listener for a parse tree produced by Java9Parser.
type BaseJava9ParserListener struct{}

var _ Java9ParserListener = &BaseJava9ParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJava9ParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJava9ParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJava9ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJava9ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJava9ParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJava9ParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseJava9ParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseJava9ParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterNumericType is called when production numericType is entered.
func (s *BaseJava9ParserListener) EnterNumericType(ctx *NumericTypeContext) {}

// ExitNumericType is called when production numericType is exited.
func (s *BaseJava9ParserListener) ExitNumericType(ctx *NumericTypeContext) {}

// EnterIntegralType is called when production integralType is entered.
func (s *BaseJava9ParserListener) EnterIntegralType(ctx *IntegralTypeContext) {}

// ExitIntegralType is called when production integralType is exited.
func (s *BaseJava9ParserListener) ExitIntegralType(ctx *IntegralTypeContext) {}

// EnterFloatingPointType is called when production floatingPointType is entered.
func (s *BaseJava9ParserListener) EnterFloatingPointType(ctx *FloatingPointTypeContext) {}

// ExitFloatingPointType is called when production floatingPointType is exited.
func (s *BaseJava9ParserListener) ExitFloatingPointType(ctx *FloatingPointTypeContext) {}

// EnterReferenceType is called when production referenceType is entered.
func (s *BaseJava9ParserListener) EnterReferenceType(ctx *ReferenceTypeContext) {}

// ExitReferenceType is called when production referenceType is exited.
func (s *BaseJava9ParserListener) ExitReferenceType(ctx *ReferenceTypeContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterClassType is called when production classType is entered.
func (s *BaseJava9ParserListener) EnterClassType(ctx *ClassTypeContext) {}

// ExitClassType is called when production classType is exited.
func (s *BaseJava9ParserListener) ExitClassType(ctx *ClassTypeContext) {}

// EnterClassType_lf_classOrInterfaceType is called when production classType_lf_classOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterClassType_lf_classOrInterfaceType(ctx *ClassType_lf_classOrInterfaceTypeContext) {
}

// ExitClassType_lf_classOrInterfaceType is called when production classType_lf_classOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitClassType_lf_classOrInterfaceType(ctx *ClassType_lf_classOrInterfaceTypeContext) {
}

// EnterClassType_lfno_classOrInterfaceType is called when production classType_lfno_classOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterClassType_lfno_classOrInterfaceType(ctx *ClassType_lfno_classOrInterfaceTypeContext) {
}

// ExitClassType_lfno_classOrInterfaceType is called when production classType_lfno_classOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitClassType_lfno_classOrInterfaceType(ctx *ClassType_lfno_classOrInterfaceTypeContext) {
}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BaseJava9ParserListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BaseJava9ParserListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterInterfaceType_lf_classOrInterfaceType is called when production interfaceType_lf_classOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterInterfaceType_lf_classOrInterfaceType(ctx *InterfaceType_lf_classOrInterfaceTypeContext) {
}

// ExitInterfaceType_lf_classOrInterfaceType is called when production interfaceType_lf_classOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitInterfaceType_lf_classOrInterfaceType(ctx *InterfaceType_lf_classOrInterfaceTypeContext) {
}

// EnterInterfaceType_lfno_classOrInterfaceType is called when production interfaceType_lfno_classOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterInterfaceType_lfno_classOrInterfaceType(ctx *InterfaceType_lfno_classOrInterfaceTypeContext) {
}

// ExitInterfaceType_lfno_classOrInterfaceType is called when production interfaceType_lfno_classOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitInterfaceType_lfno_classOrInterfaceType(ctx *InterfaceType_lfno_classOrInterfaceTypeContext) {
}

// EnterTypeVariable is called when production typeVariable is entered.
func (s *BaseJava9ParserListener) EnterTypeVariable(ctx *TypeVariableContext) {}

// ExitTypeVariable is called when production typeVariable is exited.
func (s *BaseJava9ParserListener) ExitTypeVariable(ctx *TypeVariableContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseJava9ParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseJava9ParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterDims is called when production dims is entered.
func (s *BaseJava9ParserListener) EnterDims(ctx *DimsContext) {}

// ExitDims is called when production dims is exited.
func (s *BaseJava9ParserListener) ExitDims(ctx *DimsContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseJava9ParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseJava9ParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterTypeParameterModifier is called when production typeParameterModifier is entered.
func (s *BaseJava9ParserListener) EnterTypeParameterModifier(ctx *TypeParameterModifierContext) {}

// ExitTypeParameterModifier is called when production typeParameterModifier is exited.
func (s *BaseJava9ParserListener) ExitTypeParameterModifier(ctx *TypeParameterModifierContext) {}

// EnterTypeBound is called when production typeBound is entered.
func (s *BaseJava9ParserListener) EnterTypeBound(ctx *TypeBoundContext) {}

// ExitTypeBound is called when production typeBound is exited.
func (s *BaseJava9ParserListener) ExitTypeBound(ctx *TypeBoundContext) {}

// EnterAdditionalBound is called when production additionalBound is entered.
func (s *BaseJava9ParserListener) EnterAdditionalBound(ctx *AdditionalBoundContext) {}

// ExitAdditionalBound is called when production additionalBound is exited.
func (s *BaseJava9ParserListener) ExitAdditionalBound(ctx *AdditionalBoundContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseJava9ParserListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseJava9ParserListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeArgumentList is called when production typeArgumentList is entered.
func (s *BaseJava9ParserListener) EnterTypeArgumentList(ctx *TypeArgumentListContext) {}

// ExitTypeArgumentList is called when production typeArgumentList is exited.
func (s *BaseJava9ParserListener) ExitTypeArgumentList(ctx *TypeArgumentListContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BaseJava9ParserListener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BaseJava9ParserListener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterWildcard is called when production wildcard is entered.
func (s *BaseJava9ParserListener) EnterWildcard(ctx *WildcardContext) {}

// ExitWildcard is called when production wildcard is exited.
func (s *BaseJava9ParserListener) ExitWildcard(ctx *WildcardContext) {}

// EnterWildcardBounds is called when production wildcardBounds is entered.
func (s *BaseJava9ParserListener) EnterWildcardBounds(ctx *WildcardBoundsContext) {}

// ExitWildcardBounds is called when production wildcardBounds is exited.
func (s *BaseJava9ParserListener) ExitWildcardBounds(ctx *WildcardBoundsContext) {}

// EnterModuleName is called when production moduleName is entered.
func (s *BaseJava9ParserListener) EnterModuleName(ctx *ModuleNameContext) {}

// ExitModuleName is called when production moduleName is exited.
func (s *BaseJava9ParserListener) ExitModuleName(ctx *ModuleNameContext) {}

// EnterPackageName is called when production packageName is entered.
func (s *BaseJava9ParserListener) EnterPackageName(ctx *PackageNameContext) {}

// ExitPackageName is called when production packageName is exited.
func (s *BaseJava9ParserListener) ExitPackageName(ctx *PackageNameContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseJava9ParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseJava9ParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterPackageOrTypeName is called when production packageOrTypeName is entered.
func (s *BaseJava9ParserListener) EnterPackageOrTypeName(ctx *PackageOrTypeNameContext) {}

// ExitPackageOrTypeName is called when production packageOrTypeName is exited.
func (s *BaseJava9ParserListener) ExitPackageOrTypeName(ctx *PackageOrTypeNameContext) {}

// EnterExpressionName is called when production expressionName is entered.
func (s *BaseJava9ParserListener) EnterExpressionName(ctx *ExpressionNameContext) {}

// ExitExpressionName is called when production expressionName is exited.
func (s *BaseJava9ParserListener) ExitExpressionName(ctx *ExpressionNameContext) {}

// EnterMethodName is called when production methodName is entered.
func (s *BaseJava9ParserListener) EnterMethodName(ctx *MethodNameContext) {}

// ExitMethodName is called when production methodName is exited.
func (s *BaseJava9ParserListener) ExitMethodName(ctx *MethodNameContext) {}

// EnterAmbiguousName is called when production ambiguousName is entered.
func (s *BaseJava9ParserListener) EnterAmbiguousName(ctx *AmbiguousNameContext) {}

// ExitAmbiguousName is called when production ambiguousName is exited.
func (s *BaseJava9ParserListener) ExitAmbiguousName(ctx *AmbiguousNameContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseJava9ParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseJava9ParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterOrdinaryCompilation is called when production ordinaryCompilation is entered.
func (s *BaseJava9ParserListener) EnterOrdinaryCompilation(ctx *OrdinaryCompilationContext) {}

// ExitOrdinaryCompilation is called when production ordinaryCompilation is exited.
func (s *BaseJava9ParserListener) ExitOrdinaryCompilation(ctx *OrdinaryCompilationContext) {}

// EnterModularCompilation is called when production modularCompilation is entered.
func (s *BaseJava9ParserListener) EnterModularCompilation(ctx *ModularCompilationContext) {}

// ExitModularCompilation is called when production modularCompilation is exited.
func (s *BaseJava9ParserListener) ExitModularCompilation(ctx *ModularCompilationContext) {}

// EnterPackageDeclaration is called when production packageDeclaration is entered.
func (s *BaseJava9ParserListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {}

// ExitPackageDeclaration is called when production packageDeclaration is exited.
func (s *BaseJava9ParserListener) ExitPackageDeclaration(ctx *PackageDeclarationContext) {}

// EnterPackageModifier is called when production packageModifier is entered.
func (s *BaseJava9ParserListener) EnterPackageModifier(ctx *PackageModifierContext) {}

// ExitPackageModifier is called when production packageModifier is exited.
func (s *BaseJava9ParserListener) ExitPackageModifier(ctx *PackageModifierContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseJava9ParserListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseJava9ParserListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterSingleTypeImportDeclaration is called when production singleTypeImportDeclaration is entered.
func (s *BaseJava9ParserListener) EnterSingleTypeImportDeclaration(ctx *SingleTypeImportDeclarationContext) {
}

// ExitSingleTypeImportDeclaration is called when production singleTypeImportDeclaration is exited.
func (s *BaseJava9ParserListener) ExitSingleTypeImportDeclaration(ctx *SingleTypeImportDeclarationContext) {
}

// EnterTypeImportOnDemandDeclaration is called when production typeImportOnDemandDeclaration is entered.
func (s *BaseJava9ParserListener) EnterTypeImportOnDemandDeclaration(ctx *TypeImportOnDemandDeclarationContext) {
}

// ExitTypeImportOnDemandDeclaration is called when production typeImportOnDemandDeclaration is exited.
func (s *BaseJava9ParserListener) ExitTypeImportOnDemandDeclaration(ctx *TypeImportOnDemandDeclarationContext) {
}

// EnterSingleStaticImportDeclaration is called when production singleStaticImportDeclaration is entered.
func (s *BaseJava9ParserListener) EnterSingleStaticImportDeclaration(ctx *SingleStaticImportDeclarationContext) {
}

// ExitSingleStaticImportDeclaration is called when production singleStaticImportDeclaration is exited.
func (s *BaseJava9ParserListener) ExitSingleStaticImportDeclaration(ctx *SingleStaticImportDeclarationContext) {
}

// EnterStaticImportOnDemandDeclaration is called when production staticImportOnDemandDeclaration is entered.
func (s *BaseJava9ParserListener) EnterStaticImportOnDemandDeclaration(ctx *StaticImportOnDemandDeclarationContext) {
}

// ExitStaticImportOnDemandDeclaration is called when production staticImportOnDemandDeclaration is exited.
func (s *BaseJava9ParserListener) ExitStaticImportOnDemandDeclaration(ctx *StaticImportOnDemandDeclarationContext) {
}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseJava9ParserListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseJava9ParserListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterModuleDeclaration is called when production moduleDeclaration is entered.
func (s *BaseJava9ParserListener) EnterModuleDeclaration(ctx *ModuleDeclarationContext) {}

// ExitModuleDeclaration is called when production moduleDeclaration is exited.
func (s *BaseJava9ParserListener) ExitModuleDeclaration(ctx *ModuleDeclarationContext) {}

// EnterModuleDirective is called when production moduleDirective is entered.
func (s *BaseJava9ParserListener) EnterModuleDirective(ctx *ModuleDirectiveContext) {}

// ExitModuleDirective is called when production moduleDirective is exited.
func (s *BaseJava9ParserListener) ExitModuleDirective(ctx *ModuleDirectiveContext) {}

// EnterRequiresModifier is called when production requiresModifier is entered.
func (s *BaseJava9ParserListener) EnterRequiresModifier(ctx *RequiresModifierContext) {}

// ExitRequiresModifier is called when production requiresModifier is exited.
func (s *BaseJava9ParserListener) ExitRequiresModifier(ctx *RequiresModifierContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseJava9ParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseJava9ParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterNormalClassDeclaration is called when production normalClassDeclaration is entered.
func (s *BaseJava9ParserListener) EnterNormalClassDeclaration(ctx *NormalClassDeclarationContext) {}

// ExitNormalClassDeclaration is called when production normalClassDeclaration is exited.
func (s *BaseJava9ParserListener) ExitNormalClassDeclaration(ctx *NormalClassDeclarationContext) {}

// EnterClassModifier is called when production classModifier is entered.
func (s *BaseJava9ParserListener) EnterClassModifier(ctx *ClassModifierContext) {}

// ExitClassModifier is called when production classModifier is exited.
func (s *BaseJava9ParserListener) ExitClassModifier(ctx *ClassModifierContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseJava9ParserListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseJava9ParserListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameterList is called when production typeParameterList is entered.
func (s *BaseJava9ParserListener) EnterTypeParameterList(ctx *TypeParameterListContext) {}

// ExitTypeParameterList is called when production typeParameterList is exited.
func (s *BaseJava9ParserListener) ExitTypeParameterList(ctx *TypeParameterListContext) {}

// EnterSuperclass is called when production superclass is entered.
func (s *BaseJava9ParserListener) EnterSuperclass(ctx *SuperclassContext) {}

// ExitSuperclass is called when production superclass is exited.
func (s *BaseJava9ParserListener) ExitSuperclass(ctx *SuperclassContext) {}

// EnterSuperinterfaces is called when production superinterfaces is entered.
func (s *BaseJava9ParserListener) EnterSuperinterfaces(ctx *SuperinterfacesContext) {}

// ExitSuperinterfaces is called when production superinterfaces is exited.
func (s *BaseJava9ParserListener) ExitSuperinterfaces(ctx *SuperinterfacesContext) {}

// EnterInterfaceTypeList is called when production interfaceTypeList is entered.
func (s *BaseJava9ParserListener) EnterInterfaceTypeList(ctx *InterfaceTypeListContext) {}

// ExitInterfaceTypeList is called when production interfaceTypeList is exited.
func (s *BaseJava9ParserListener) ExitInterfaceTypeList(ctx *InterfaceTypeListContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseJava9ParserListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseJava9ParserListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseJava9ParserListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseJava9ParserListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterClassMemberDeclaration is called when production classMemberDeclaration is entered.
func (s *BaseJava9ParserListener) EnterClassMemberDeclaration(ctx *ClassMemberDeclarationContext) {}

// ExitClassMemberDeclaration is called when production classMemberDeclaration is exited.
func (s *BaseJava9ParserListener) ExitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseJava9ParserListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseJava9ParserListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterFieldModifier is called when production fieldModifier is entered.
func (s *BaseJava9ParserListener) EnterFieldModifier(ctx *FieldModifierContext) {}

// ExitFieldModifier is called when production fieldModifier is exited.
func (s *BaseJava9ParserListener) ExitFieldModifier(ctx *FieldModifierContext) {}

// EnterVariableDeclaratorList is called when production variableDeclaratorList is entered.
func (s *BaseJava9ParserListener) EnterVariableDeclaratorList(ctx *VariableDeclaratorListContext) {}

// ExitVariableDeclaratorList is called when production variableDeclaratorList is exited.
func (s *BaseJava9ParserListener) ExitVariableDeclaratorList(ctx *VariableDeclaratorListContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseJava9ParserListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseJava9ParserListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseJava9ParserListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseJava9ParserListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseJava9ParserListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseJava9ParserListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterUnannType is called when production unannType is entered.
func (s *BaseJava9ParserListener) EnterUnannType(ctx *UnannTypeContext) {}

// ExitUnannType is called when production unannType is exited.
func (s *BaseJava9ParserListener) ExitUnannType(ctx *UnannTypeContext) {}

// EnterUnannPrimitiveType is called when production unannPrimitiveType is entered.
func (s *BaseJava9ParserListener) EnterUnannPrimitiveType(ctx *UnannPrimitiveTypeContext) {}

// ExitUnannPrimitiveType is called when production unannPrimitiveType is exited.
func (s *BaseJava9ParserListener) ExitUnannPrimitiveType(ctx *UnannPrimitiveTypeContext) {}

// EnterUnannReferenceType is called when production unannReferenceType is entered.
func (s *BaseJava9ParserListener) EnterUnannReferenceType(ctx *UnannReferenceTypeContext) {}

// ExitUnannReferenceType is called when production unannReferenceType is exited.
func (s *BaseJava9ParserListener) ExitUnannReferenceType(ctx *UnannReferenceTypeContext) {}

// EnterUnannClassOrInterfaceType is called when production unannClassOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterUnannClassOrInterfaceType(ctx *UnannClassOrInterfaceTypeContext) {
}

// ExitUnannClassOrInterfaceType is called when production unannClassOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitUnannClassOrInterfaceType(ctx *UnannClassOrInterfaceTypeContext) {
}

// EnterUnannClassType is called when production unannClassType is entered.
func (s *BaseJava9ParserListener) EnterUnannClassType(ctx *UnannClassTypeContext) {}

// ExitUnannClassType is called when production unannClassType is exited.
func (s *BaseJava9ParserListener) ExitUnannClassType(ctx *UnannClassTypeContext) {}

// EnterUnannClassType_lf_unannClassOrInterfaceType is called when production unannClassType_lf_unannClassOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterUnannClassType_lf_unannClassOrInterfaceType(ctx *UnannClassType_lf_unannClassOrInterfaceTypeContext) {
}

// ExitUnannClassType_lf_unannClassOrInterfaceType is called when production unannClassType_lf_unannClassOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitUnannClassType_lf_unannClassOrInterfaceType(ctx *UnannClassType_lf_unannClassOrInterfaceTypeContext) {
}

// EnterUnannClassType_lfno_unannClassOrInterfaceType is called when production unannClassType_lfno_unannClassOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterUnannClassType_lfno_unannClassOrInterfaceType(ctx *UnannClassType_lfno_unannClassOrInterfaceTypeContext) {
}

// ExitUnannClassType_lfno_unannClassOrInterfaceType is called when production unannClassType_lfno_unannClassOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitUnannClassType_lfno_unannClassOrInterfaceType(ctx *UnannClassType_lfno_unannClassOrInterfaceTypeContext) {
}

// EnterUnannInterfaceType is called when production unannInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterUnannInterfaceType(ctx *UnannInterfaceTypeContext) {}

// ExitUnannInterfaceType is called when production unannInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitUnannInterfaceType(ctx *UnannInterfaceTypeContext) {}

// EnterUnannInterfaceType_lf_unannClassOrInterfaceType is called when production unannInterfaceType_lf_unannClassOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterUnannInterfaceType_lf_unannClassOrInterfaceType(ctx *UnannInterfaceType_lf_unannClassOrInterfaceTypeContext) {
}

// ExitUnannInterfaceType_lf_unannClassOrInterfaceType is called when production unannInterfaceType_lf_unannClassOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitUnannInterfaceType_lf_unannClassOrInterfaceType(ctx *UnannInterfaceType_lf_unannClassOrInterfaceTypeContext) {
}

// EnterUnannInterfaceType_lfno_unannClassOrInterfaceType is called when production unannInterfaceType_lfno_unannClassOrInterfaceType is entered.
func (s *BaseJava9ParserListener) EnterUnannInterfaceType_lfno_unannClassOrInterfaceType(ctx *UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext) {
}

// ExitUnannInterfaceType_lfno_unannClassOrInterfaceType is called when production unannInterfaceType_lfno_unannClassOrInterfaceType is exited.
func (s *BaseJava9ParserListener) ExitUnannInterfaceType_lfno_unannClassOrInterfaceType(ctx *UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext) {
}

// EnterUnannTypeVariable is called when production unannTypeVariable is entered.
func (s *BaseJava9ParserListener) EnterUnannTypeVariable(ctx *UnannTypeVariableContext) {}

// ExitUnannTypeVariable is called when production unannTypeVariable is exited.
func (s *BaseJava9ParserListener) ExitUnannTypeVariable(ctx *UnannTypeVariableContext) {}

// EnterUnannArrayType is called when production unannArrayType is entered.
func (s *BaseJava9ParserListener) EnterUnannArrayType(ctx *UnannArrayTypeContext) {}

// ExitUnannArrayType is called when production unannArrayType is exited.
func (s *BaseJava9ParserListener) ExitUnannArrayType(ctx *UnannArrayTypeContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseJava9ParserListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseJava9ParserListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterMethodModifier is called when production methodModifier is entered.
func (s *BaseJava9ParserListener) EnterMethodModifier(ctx *MethodModifierContext) {}

// ExitMethodModifier is called when production methodModifier is exited.
func (s *BaseJava9ParserListener) ExitMethodModifier(ctx *MethodModifierContext) {}

// EnterMethodHeader is called when production methodHeader is entered.
func (s *BaseJava9ParserListener) EnterMethodHeader(ctx *MethodHeaderContext) {}

// ExitMethodHeader is called when production methodHeader is exited.
func (s *BaseJava9ParserListener) ExitMethodHeader(ctx *MethodHeaderContext) {}

// EnterResult is called when production result is entered.
func (s *BaseJava9ParserListener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *BaseJava9ParserListener) ExitResult(ctx *ResultContext) {}

// EnterMethodDeclarator is called when production methodDeclarator is entered.
func (s *BaseJava9ParserListener) EnterMethodDeclarator(ctx *MethodDeclaratorContext) {}

// ExitMethodDeclarator is called when production methodDeclarator is exited.
func (s *BaseJava9ParserListener) ExitMethodDeclarator(ctx *MethodDeclaratorContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseJava9ParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseJava9ParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseJava9ParserListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseJava9ParserListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseJava9ParserListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseJava9ParserListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BaseJava9ParserListener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BaseJava9ParserListener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterLastFormalParameter is called when production lastFormalParameter is entered.
func (s *BaseJava9ParserListener) EnterLastFormalParameter(ctx *LastFormalParameterContext) {}

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (s *BaseJava9ParserListener) ExitLastFormalParameter(ctx *LastFormalParameterContext) {}

// EnterReceiverParameter is called when production receiverParameter is entered.
func (s *BaseJava9ParserListener) EnterReceiverParameter(ctx *ReceiverParameterContext) {}

// ExitReceiverParameter is called when production receiverParameter is exited.
func (s *BaseJava9ParserListener) ExitReceiverParameter(ctx *ReceiverParameterContext) {}

// EnterThrows_ is called when production throws_ is entered.
func (s *BaseJava9ParserListener) EnterThrows_(ctx *Throws_Context) {}

// ExitThrows_ is called when production throws_ is exited.
func (s *BaseJava9ParserListener) ExitThrows_(ctx *Throws_Context) {}

// EnterExceptionTypeList is called when production exceptionTypeList is entered.
func (s *BaseJava9ParserListener) EnterExceptionTypeList(ctx *ExceptionTypeListContext) {}

// ExitExceptionTypeList is called when production exceptionTypeList is exited.
func (s *BaseJava9ParserListener) ExitExceptionTypeList(ctx *ExceptionTypeListContext) {}

// EnterExceptionType is called when production exceptionType is entered.
func (s *BaseJava9ParserListener) EnterExceptionType(ctx *ExceptionTypeContext) {}

// ExitExceptionType is called when production exceptionType is exited.
func (s *BaseJava9ParserListener) ExitExceptionType(ctx *ExceptionTypeContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseJava9ParserListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseJava9ParserListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterInstanceInitializer is called when production instanceInitializer is entered.
func (s *BaseJava9ParserListener) EnterInstanceInitializer(ctx *InstanceInitializerContext) {}

// ExitInstanceInitializer is called when production instanceInitializer is exited.
func (s *BaseJava9ParserListener) ExitInstanceInitializer(ctx *InstanceInitializerContext) {}

// EnterStaticInitializer is called when production staticInitializer is entered.
func (s *BaseJava9ParserListener) EnterStaticInitializer(ctx *StaticInitializerContext) {}

// ExitStaticInitializer is called when production staticInitializer is exited.
func (s *BaseJava9ParserListener) ExitStaticInitializer(ctx *StaticInitializerContext) {}

// EnterConstructorDeclaration is called when production constructorDeclaration is entered.
func (s *BaseJava9ParserListener) EnterConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// ExitConstructorDeclaration is called when production constructorDeclaration is exited.
func (s *BaseJava9ParserListener) ExitConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// EnterConstructorModifier is called when production constructorModifier is entered.
func (s *BaseJava9ParserListener) EnterConstructorModifier(ctx *ConstructorModifierContext) {}

// ExitConstructorModifier is called when production constructorModifier is exited.
func (s *BaseJava9ParserListener) ExitConstructorModifier(ctx *ConstructorModifierContext) {}

// EnterConstructorDeclarator is called when production constructorDeclarator is entered.
func (s *BaseJava9ParserListener) EnterConstructorDeclarator(ctx *ConstructorDeclaratorContext) {}

// ExitConstructorDeclarator is called when production constructorDeclarator is exited.
func (s *BaseJava9ParserListener) ExitConstructorDeclarator(ctx *ConstructorDeclaratorContext) {}

// EnterSimpleTypeName is called when production simpleTypeName is entered.
func (s *BaseJava9ParserListener) EnterSimpleTypeName(ctx *SimpleTypeNameContext) {}

// ExitSimpleTypeName is called when production simpleTypeName is exited.
func (s *BaseJava9ParserListener) ExitSimpleTypeName(ctx *SimpleTypeNameContext) {}

// EnterConstructorBody is called when production constructorBody is entered.
func (s *BaseJava9ParserListener) EnterConstructorBody(ctx *ConstructorBodyContext) {}

// ExitConstructorBody is called when production constructorBody is exited.
func (s *BaseJava9ParserListener) ExitConstructorBody(ctx *ConstructorBodyContext) {}

// EnterExplicitConstructorInvocation is called when production explicitConstructorInvocation is entered.
func (s *BaseJava9ParserListener) EnterExplicitConstructorInvocation(ctx *ExplicitConstructorInvocationContext) {
}

// ExitExplicitConstructorInvocation is called when production explicitConstructorInvocation is exited.
func (s *BaseJava9ParserListener) ExitExplicitConstructorInvocation(ctx *ExplicitConstructorInvocationContext) {
}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseJava9ParserListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseJava9ParserListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseJava9ParserListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseJava9ParserListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumConstantList is called when production enumConstantList is entered.
func (s *BaseJava9ParserListener) EnterEnumConstantList(ctx *EnumConstantListContext) {}

// ExitEnumConstantList is called when production enumConstantList is exited.
func (s *BaseJava9ParserListener) ExitEnumConstantList(ctx *EnumConstantListContext) {}

// EnterEnumConstant is called when production enumConstant is entered.
func (s *BaseJava9ParserListener) EnterEnumConstant(ctx *EnumConstantContext) {}

// ExitEnumConstant is called when production enumConstant is exited.
func (s *BaseJava9ParserListener) ExitEnumConstant(ctx *EnumConstantContext) {}

// EnterEnumConstantModifier is called when production enumConstantModifier is entered.
func (s *BaseJava9ParserListener) EnterEnumConstantModifier(ctx *EnumConstantModifierContext) {}

// ExitEnumConstantModifier is called when production enumConstantModifier is exited.
func (s *BaseJava9ParserListener) ExitEnumConstantModifier(ctx *EnumConstantModifierContext) {}

// EnterEnumBodyDeclarations is called when production enumBodyDeclarations is entered.
func (s *BaseJava9ParserListener) EnterEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// ExitEnumBodyDeclarations is called when production enumBodyDeclarations is exited.
func (s *BaseJava9ParserListener) ExitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseJava9ParserListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseJava9ParserListener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterNormalInterfaceDeclaration is called when production normalInterfaceDeclaration is entered.
func (s *BaseJava9ParserListener) EnterNormalInterfaceDeclaration(ctx *NormalInterfaceDeclarationContext) {
}

// ExitNormalInterfaceDeclaration is called when production normalInterfaceDeclaration is exited.
func (s *BaseJava9ParserListener) ExitNormalInterfaceDeclaration(ctx *NormalInterfaceDeclarationContext) {
}

// EnterInterfaceModifier is called when production interfaceModifier is entered.
func (s *BaseJava9ParserListener) EnterInterfaceModifier(ctx *InterfaceModifierContext) {}

// ExitInterfaceModifier is called when production interfaceModifier is exited.
func (s *BaseJava9ParserListener) ExitInterfaceModifier(ctx *InterfaceModifierContext) {}

// EnterExtendsInterfaces is called when production extendsInterfaces is entered.
func (s *BaseJava9ParserListener) EnterExtendsInterfaces(ctx *ExtendsInterfacesContext) {}

// ExitExtendsInterfaces is called when production extendsInterfaces is exited.
func (s *BaseJava9ParserListener) ExitExtendsInterfaces(ctx *ExtendsInterfacesContext) {}

// EnterInterfaceBody is called when production interfaceBody is entered.
func (s *BaseJava9ParserListener) EnterInterfaceBody(ctx *InterfaceBodyContext) {}

// ExitInterfaceBody is called when production interfaceBody is exited.
func (s *BaseJava9ParserListener) ExitInterfaceBody(ctx *InterfaceBodyContext) {}

// EnterInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is entered.
func (s *BaseJava9ParserListener) EnterInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {
}

// ExitInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is exited.
func (s *BaseJava9ParserListener) ExitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {
}

// EnterConstantDeclaration is called when production constantDeclaration is entered.
func (s *BaseJava9ParserListener) EnterConstantDeclaration(ctx *ConstantDeclarationContext) {}

// ExitConstantDeclaration is called when production constantDeclaration is exited.
func (s *BaseJava9ParserListener) ExitConstantDeclaration(ctx *ConstantDeclarationContext) {}

// EnterConstantModifier is called when production constantModifier is entered.
func (s *BaseJava9ParserListener) EnterConstantModifier(ctx *ConstantModifierContext) {}

// ExitConstantModifier is called when production constantModifier is exited.
func (s *BaseJava9ParserListener) ExitConstantModifier(ctx *ConstantModifierContext) {}

// EnterInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is entered.
func (s *BaseJava9ParserListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
}

// ExitInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is exited.
func (s *BaseJava9ParserListener) ExitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
}

// EnterInterfaceMethodModifier is called when production interfaceMethodModifier is entered.
func (s *BaseJava9ParserListener) EnterInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) {}

// ExitInterfaceMethodModifier is called when production interfaceMethodModifier is exited.
func (s *BaseJava9ParserListener) ExitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) {}

// EnterAnnotationTypeDeclaration is called when production annotationTypeDeclaration is entered.
func (s *BaseJava9ParserListener) EnterAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {
}

// ExitAnnotationTypeDeclaration is called when production annotationTypeDeclaration is exited.
func (s *BaseJava9ParserListener) ExitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {
}

// EnterAnnotationTypeBody is called when production annotationTypeBody is entered.
func (s *BaseJava9ParserListener) EnterAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// ExitAnnotationTypeBody is called when production annotationTypeBody is exited.
func (s *BaseJava9ParserListener) ExitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// EnterAnnotationTypeMemberDeclaration is called when production annotationTypeMemberDeclaration is entered.
func (s *BaseJava9ParserListener) EnterAnnotationTypeMemberDeclaration(ctx *AnnotationTypeMemberDeclarationContext) {
}

// ExitAnnotationTypeMemberDeclaration is called when production annotationTypeMemberDeclaration is exited.
func (s *BaseJava9ParserListener) ExitAnnotationTypeMemberDeclaration(ctx *AnnotationTypeMemberDeclarationContext) {
}

// EnterAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is entered.
func (s *BaseJava9ParserListener) EnterAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// ExitAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is exited.
func (s *BaseJava9ParserListener) ExitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// EnterAnnotationTypeElementModifier is called when production annotationTypeElementModifier is entered.
func (s *BaseJava9ParserListener) EnterAnnotationTypeElementModifier(ctx *AnnotationTypeElementModifierContext) {
}

// ExitAnnotationTypeElementModifier is called when production annotationTypeElementModifier is exited.
func (s *BaseJava9ParserListener) ExitAnnotationTypeElementModifier(ctx *AnnotationTypeElementModifierContext) {
}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseJava9ParserListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseJava9ParserListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseJava9ParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseJava9ParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterNormalAnnotation is called when production normalAnnotation is entered.
func (s *BaseJava9ParserListener) EnterNormalAnnotation(ctx *NormalAnnotationContext) {}

// ExitNormalAnnotation is called when production normalAnnotation is exited.
func (s *BaseJava9ParserListener) ExitNormalAnnotation(ctx *NormalAnnotationContext) {}

// EnterElementValuePairList is called when production elementValuePairList is entered.
func (s *BaseJava9ParserListener) EnterElementValuePairList(ctx *ElementValuePairListContext) {}

// ExitElementValuePairList is called when production elementValuePairList is exited.
func (s *BaseJava9ParserListener) ExitElementValuePairList(ctx *ElementValuePairListContext) {}

// EnterElementValuePair is called when production elementValuePair is entered.
func (s *BaseJava9ParserListener) EnterElementValuePair(ctx *ElementValuePairContext) {}

// ExitElementValuePair is called when production elementValuePair is exited.
func (s *BaseJava9ParserListener) ExitElementValuePair(ctx *ElementValuePairContext) {}

// EnterElementValue is called when production elementValue is entered.
func (s *BaseJava9ParserListener) EnterElementValue(ctx *ElementValueContext) {}

// ExitElementValue is called when production elementValue is exited.
func (s *BaseJava9ParserListener) ExitElementValue(ctx *ElementValueContext) {}

// EnterElementValueArrayInitializer is called when production elementValueArrayInitializer is entered.
func (s *BaseJava9ParserListener) EnterElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// ExitElementValueArrayInitializer is called when production elementValueArrayInitializer is exited.
func (s *BaseJava9ParserListener) ExitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// EnterElementValueList is called when production elementValueList is entered.
func (s *BaseJava9ParserListener) EnterElementValueList(ctx *ElementValueListContext) {}

// ExitElementValueList is called when production elementValueList is exited.
func (s *BaseJava9ParserListener) ExitElementValueList(ctx *ElementValueListContext) {}

// EnterMarkerAnnotation is called when production markerAnnotation is entered.
func (s *BaseJava9ParserListener) EnterMarkerAnnotation(ctx *MarkerAnnotationContext) {}

// ExitMarkerAnnotation is called when production markerAnnotation is exited.
func (s *BaseJava9ParserListener) ExitMarkerAnnotation(ctx *MarkerAnnotationContext) {}

// EnterSingleElementAnnotation is called when production singleElementAnnotation is entered.
func (s *BaseJava9ParserListener) EnterSingleElementAnnotation(ctx *SingleElementAnnotationContext) {}

// ExitSingleElementAnnotation is called when production singleElementAnnotation is exited.
func (s *BaseJava9ParserListener) ExitSingleElementAnnotation(ctx *SingleElementAnnotationContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseJava9ParserListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseJava9ParserListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterVariableInitializerList is called when production variableInitializerList is entered.
func (s *BaseJava9ParserListener) EnterVariableInitializerList(ctx *VariableInitializerListContext) {}

// ExitVariableInitializerList is called when production variableInitializerList is exited.
func (s *BaseJava9ParserListener) ExitVariableInitializerList(ctx *VariableInitializerListContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJava9ParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJava9ParserListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatements is called when production blockStatements is entered.
func (s *BaseJava9ParserListener) EnterBlockStatements(ctx *BlockStatementsContext) {}

// ExitBlockStatements is called when production blockStatements is exited.
func (s *BaseJava9ParserListener) ExitBlockStatements(ctx *BlockStatementsContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseJava9ParserListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseJava9ParserListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is entered.
func (s *BaseJava9ParserListener) EnterLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// ExitLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is exited.
func (s *BaseJava9ParserListener) ExitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseJava9ParserListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseJava9ParserListener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
}

// EnterStatement is called when production statement is entered.
func (s *BaseJava9ParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJava9ParserListener) ExitStatement(ctx *StatementContext) {}

// EnterStatementNoShortIf is called when production statementNoShortIf is entered.
func (s *BaseJava9ParserListener) EnterStatementNoShortIf(ctx *StatementNoShortIfContext) {}

// ExitStatementNoShortIf is called when production statementNoShortIf is exited.
func (s *BaseJava9ParserListener) ExitStatementNoShortIf(ctx *StatementNoShortIfContext) {}

// EnterStatementWithoutTrailingSubstatement is called when production statementWithoutTrailingSubstatement is entered.
func (s *BaseJava9ParserListener) EnterStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) {
}

// ExitStatementWithoutTrailingSubstatement is called when production statementWithoutTrailingSubstatement is exited.
func (s *BaseJava9ParserListener) ExitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) {
}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseJava9ParserListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseJava9ParserListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *BaseJava9ParserListener) EnterLabeledStatement(ctx *LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *BaseJava9ParserListener) ExitLabeledStatement(ctx *LabeledStatementContext) {}

// EnterLabeledStatementNoShortIf is called when production labeledStatementNoShortIf is entered.
func (s *BaseJava9ParserListener) EnterLabeledStatementNoShortIf(ctx *LabeledStatementNoShortIfContext) {
}

// ExitLabeledStatementNoShortIf is called when production labeledStatementNoShortIf is exited.
func (s *BaseJava9ParserListener) ExitLabeledStatementNoShortIf(ctx *LabeledStatementNoShortIfContext) {
}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseJava9ParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseJava9ParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterStatementExpression is called when production statementExpression is entered.
func (s *BaseJava9ParserListener) EnterStatementExpression(ctx *StatementExpressionContext) {}

// ExitStatementExpression is called when production statementExpression is exited.
func (s *BaseJava9ParserListener) ExitStatementExpression(ctx *StatementExpressionContext) {}

// EnterIfThenStatement is called when production ifThenStatement is entered.
func (s *BaseJava9ParserListener) EnterIfThenStatement(ctx *IfThenStatementContext) {}

// ExitIfThenStatement is called when production ifThenStatement is exited.
func (s *BaseJava9ParserListener) ExitIfThenStatement(ctx *IfThenStatementContext) {}

// EnterIfThenElseStatement is called when production ifThenElseStatement is entered.
func (s *BaseJava9ParserListener) EnterIfThenElseStatement(ctx *IfThenElseStatementContext) {}

// ExitIfThenElseStatement is called when production ifThenElseStatement is exited.
func (s *BaseJava9ParserListener) ExitIfThenElseStatement(ctx *IfThenElseStatementContext) {}

// EnterIfThenElseStatementNoShortIf is called when production ifThenElseStatementNoShortIf is entered.
func (s *BaseJava9ParserListener) EnterIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) {
}

// ExitIfThenElseStatementNoShortIf is called when production ifThenElseStatementNoShortIf is exited.
func (s *BaseJava9ParserListener) ExitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) {
}

// EnterAssertStatement is called when production assertStatement is entered.
func (s *BaseJava9ParserListener) EnterAssertStatement(ctx *AssertStatementContext) {}

// ExitAssertStatement is called when production assertStatement is exited.
func (s *BaseJava9ParserListener) ExitAssertStatement(ctx *AssertStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseJava9ParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseJava9ParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchBlock is called when production switchBlock is entered.
func (s *BaseJava9ParserListener) EnterSwitchBlock(ctx *SwitchBlockContext) {}

// ExitSwitchBlock is called when production switchBlock is exited.
func (s *BaseJava9ParserListener) ExitSwitchBlock(ctx *SwitchBlockContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BaseJava9ParserListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BaseJava9ParserListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// EnterSwitchLabels is called when production switchLabels is entered.
func (s *BaseJava9ParserListener) EnterSwitchLabels(ctx *SwitchLabelsContext) {}

// ExitSwitchLabels is called when production switchLabels is exited.
func (s *BaseJava9ParserListener) ExitSwitchLabels(ctx *SwitchLabelsContext) {}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseJava9ParserListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseJava9ParserListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterEnumConstantName is called when production enumConstantName is entered.
func (s *BaseJava9ParserListener) EnterEnumConstantName(ctx *EnumConstantNameContext) {}

// ExitEnumConstantName is called when production enumConstantName is exited.
func (s *BaseJava9ParserListener) ExitEnumConstantName(ctx *EnumConstantNameContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseJava9ParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseJava9ParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterWhileStatementNoShortIf is called when production whileStatementNoShortIf is entered.
func (s *BaseJava9ParserListener) EnterWhileStatementNoShortIf(ctx *WhileStatementNoShortIfContext) {}

// ExitWhileStatementNoShortIf is called when production whileStatementNoShortIf is exited.
func (s *BaseJava9ParserListener) ExitWhileStatementNoShortIf(ctx *WhileStatementNoShortIfContext) {}

// EnterDoStatement is called when production doStatement is entered.
func (s *BaseJava9ParserListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production doStatement is exited.
func (s *BaseJava9ParserListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseJava9ParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseJava9ParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForStatementNoShortIf is called when production forStatementNoShortIf is entered.
func (s *BaseJava9ParserListener) EnterForStatementNoShortIf(ctx *ForStatementNoShortIfContext) {}

// ExitForStatementNoShortIf is called when production forStatementNoShortIf is exited.
func (s *BaseJava9ParserListener) ExitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) {}

// EnterBasicForStatement is called when production basicForStatement is entered.
func (s *BaseJava9ParserListener) EnterBasicForStatement(ctx *BasicForStatementContext) {}

// ExitBasicForStatement is called when production basicForStatement is exited.
func (s *BaseJava9ParserListener) ExitBasicForStatement(ctx *BasicForStatementContext) {}

// EnterBasicForStatementNoShortIf is called when production basicForStatementNoShortIf is entered.
func (s *BaseJava9ParserListener) EnterBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) {
}

// ExitBasicForStatementNoShortIf is called when production basicForStatementNoShortIf is exited.
func (s *BaseJava9ParserListener) ExitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) {
}

// EnterForInit is called when production forInit is entered.
func (s *BaseJava9ParserListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseJava9ParserListener) ExitForInit(ctx *ForInitContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseJava9ParserListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseJava9ParserListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterStatementExpressionList is called when production statementExpressionList is entered.
func (s *BaseJava9ParserListener) EnterStatementExpressionList(ctx *StatementExpressionListContext) {}

// ExitStatementExpressionList is called when production statementExpressionList is exited.
func (s *BaseJava9ParserListener) ExitStatementExpressionList(ctx *StatementExpressionListContext) {}

// EnterEnhancedForStatement is called when production enhancedForStatement is entered.
func (s *BaseJava9ParserListener) EnterEnhancedForStatement(ctx *EnhancedForStatementContext) {}

// ExitEnhancedForStatement is called when production enhancedForStatement is exited.
func (s *BaseJava9ParserListener) ExitEnhancedForStatement(ctx *EnhancedForStatementContext) {}

// EnterEnhancedForStatementNoShortIf is called when production enhancedForStatementNoShortIf is entered.
func (s *BaseJava9ParserListener) EnterEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) {
}

// ExitEnhancedForStatementNoShortIf is called when production enhancedForStatementNoShortIf is exited.
func (s *BaseJava9ParserListener) ExitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) {
}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseJava9ParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseJava9ParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseJava9ParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseJava9ParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseJava9ParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseJava9ParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseJava9ParserListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseJava9ParserListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterSynchronizedStatement is called when production synchronizedStatement is entered.
func (s *BaseJava9ParserListener) EnterSynchronizedStatement(ctx *SynchronizedStatementContext) {}

// ExitSynchronizedStatement is called when production synchronizedStatement is exited.
func (s *BaseJava9ParserListener) ExitSynchronizedStatement(ctx *SynchronizedStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseJava9ParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseJava9ParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatches is called when production catches is entered.
func (s *BaseJava9ParserListener) EnterCatches(ctx *CatchesContext) {}

// ExitCatches is called when production catches is exited.
func (s *BaseJava9ParserListener) ExitCatches(ctx *CatchesContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BaseJava9ParserListener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BaseJava9ParserListener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterCatchFormalParameter is called when production catchFormalParameter is entered.
func (s *BaseJava9ParserListener) EnterCatchFormalParameter(ctx *CatchFormalParameterContext) {}

// ExitCatchFormalParameter is called when production catchFormalParameter is exited.
func (s *BaseJava9ParserListener) ExitCatchFormalParameter(ctx *CatchFormalParameterContext) {}

// EnterCatchType is called when production catchType is entered.
func (s *BaseJava9ParserListener) EnterCatchType(ctx *CatchTypeContext) {}

// ExitCatchType is called when production catchType is exited.
func (s *BaseJava9ParserListener) ExitCatchType(ctx *CatchTypeContext) {}

// EnterFinally_ is called when production finally_ is entered.
func (s *BaseJava9ParserListener) EnterFinally_(ctx *Finally_Context) {}

// ExitFinally_ is called when production finally_ is exited.
func (s *BaseJava9ParserListener) ExitFinally_(ctx *Finally_Context) {}

// EnterTryWithResourcesStatement is called when production tryWithResourcesStatement is entered.
func (s *BaseJava9ParserListener) EnterTryWithResourcesStatement(ctx *TryWithResourcesStatementContext) {
}

// ExitTryWithResourcesStatement is called when production tryWithResourcesStatement is exited.
func (s *BaseJava9ParserListener) ExitTryWithResourcesStatement(ctx *TryWithResourcesStatementContext) {
}

// EnterResourceSpecification is called when production resourceSpecification is entered.
func (s *BaseJava9ParserListener) EnterResourceSpecification(ctx *ResourceSpecificationContext) {}

// ExitResourceSpecification is called when production resourceSpecification is exited.
func (s *BaseJava9ParserListener) ExitResourceSpecification(ctx *ResourceSpecificationContext) {}

// EnterResourceList is called when production resourceList is entered.
func (s *BaseJava9ParserListener) EnterResourceList(ctx *ResourceListContext) {}

// ExitResourceList is called when production resourceList is exited.
func (s *BaseJava9ParserListener) ExitResourceList(ctx *ResourceListContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseJava9ParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseJava9ParserListener) ExitResource(ctx *ResourceContext) {}

// EnterVariableAccess is called when production variableAccess is entered.
func (s *BaseJava9ParserListener) EnterVariableAccess(ctx *VariableAccessContext) {}

// ExitVariableAccess is called when production variableAccess is exited.
func (s *BaseJava9ParserListener) ExitVariableAccess(ctx *VariableAccessContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseJava9ParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseJava9ParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterPrimaryNoNewArray is called when production primaryNoNewArray is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray(ctx *PrimaryNoNewArrayContext) {}

// ExitPrimaryNoNewArray is called when production primaryNoNewArray is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray(ctx *PrimaryNoNewArrayContext) {}

// EnterPrimaryNoNewArray_lf_arrayAccess is called when production primaryNoNewArray_lf_arrayAccess is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray_lf_arrayAccess(ctx *PrimaryNoNewArray_lf_arrayAccessContext) {
}

// ExitPrimaryNoNewArray_lf_arrayAccess is called when production primaryNoNewArray_lf_arrayAccess is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray_lf_arrayAccess(ctx *PrimaryNoNewArray_lf_arrayAccessContext) {
}

// EnterPrimaryNoNewArray_lfno_arrayAccess is called when production primaryNoNewArray_lfno_arrayAccess is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray_lfno_arrayAccess(ctx *PrimaryNoNewArray_lfno_arrayAccessContext) {
}

// ExitPrimaryNoNewArray_lfno_arrayAccess is called when production primaryNoNewArray_lfno_arrayAccess is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray_lfno_arrayAccess(ctx *PrimaryNoNewArray_lfno_arrayAccessContext) {
}

// EnterPrimaryNoNewArray_lf_primary is called when production primaryNoNewArray_lf_primary is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray_lf_primary(ctx *PrimaryNoNewArray_lf_primaryContext) {
}

// ExitPrimaryNoNewArray_lf_primary is called when production primaryNoNewArray_lf_primary is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray_lf_primary(ctx *PrimaryNoNewArray_lf_primaryContext) {
}

// EnterPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is called when production primaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext) {
}

// ExitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is called when production primaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext) {
}

// EnterPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is called when production primaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext) {
}

// ExitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is called when production primaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext) {
}

// EnterPrimaryNoNewArray_lfno_primary is called when production primaryNoNewArray_lfno_primary is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray_lfno_primary(ctx *PrimaryNoNewArray_lfno_primaryContext) {
}

// ExitPrimaryNoNewArray_lfno_primary is called when production primaryNoNewArray_lfno_primary is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray_lfno_primary(ctx *PrimaryNoNewArray_lfno_primaryContext) {
}

// EnterPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is called when production primaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext) {
}

// ExitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is called when production primaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext) {
}

// EnterPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is called when production primaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is entered.
func (s *BaseJava9ParserListener) EnterPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext) {
}

// ExitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is called when production primaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is exited.
func (s *BaseJava9ParserListener) ExitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext) {
}

// EnterClassLiteral is called when production classLiteral is entered.
func (s *BaseJava9ParserListener) EnterClassLiteral(ctx *ClassLiteralContext) {}

// ExitClassLiteral is called when production classLiteral is exited.
func (s *BaseJava9ParserListener) ExitClassLiteral(ctx *ClassLiteralContext) {}

// EnterClassInstanceCreationExpression is called when production classInstanceCreationExpression is entered.
func (s *BaseJava9ParserListener) EnterClassInstanceCreationExpression(ctx *ClassInstanceCreationExpressionContext) {
}

// ExitClassInstanceCreationExpression is called when production classInstanceCreationExpression is exited.
func (s *BaseJava9ParserListener) ExitClassInstanceCreationExpression(ctx *ClassInstanceCreationExpressionContext) {
}

// EnterClassInstanceCreationExpression_lf_primary is called when production classInstanceCreationExpression_lf_primary is entered.
func (s *BaseJava9ParserListener) EnterClassInstanceCreationExpression_lf_primary(ctx *ClassInstanceCreationExpression_lf_primaryContext) {
}

// ExitClassInstanceCreationExpression_lf_primary is called when production classInstanceCreationExpression_lf_primary is exited.
func (s *BaseJava9ParserListener) ExitClassInstanceCreationExpression_lf_primary(ctx *ClassInstanceCreationExpression_lf_primaryContext) {
}

// EnterClassInstanceCreationExpression_lfno_primary is called when production classInstanceCreationExpression_lfno_primary is entered.
func (s *BaseJava9ParserListener) EnterClassInstanceCreationExpression_lfno_primary(ctx *ClassInstanceCreationExpression_lfno_primaryContext) {
}

// ExitClassInstanceCreationExpression_lfno_primary is called when production classInstanceCreationExpression_lfno_primary is exited.
func (s *BaseJava9ParserListener) ExitClassInstanceCreationExpression_lfno_primary(ctx *ClassInstanceCreationExpression_lfno_primaryContext) {
}

// EnterTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is entered.
func (s *BaseJava9ParserListener) EnterTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// ExitTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is exited.
func (s *BaseJava9ParserListener) ExitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// EnterFieldAccess is called when production fieldAccess is entered.
func (s *BaseJava9ParserListener) EnterFieldAccess(ctx *FieldAccessContext) {}

// ExitFieldAccess is called when production fieldAccess is exited.
func (s *BaseJava9ParserListener) ExitFieldAccess(ctx *FieldAccessContext) {}

// EnterFieldAccess_lf_primary is called when production fieldAccess_lf_primary is entered.
func (s *BaseJava9ParserListener) EnterFieldAccess_lf_primary(ctx *FieldAccess_lf_primaryContext) {}

// ExitFieldAccess_lf_primary is called when production fieldAccess_lf_primary is exited.
func (s *BaseJava9ParserListener) ExitFieldAccess_lf_primary(ctx *FieldAccess_lf_primaryContext) {}

// EnterFieldAccess_lfno_primary is called when production fieldAccess_lfno_primary is entered.
func (s *BaseJava9ParserListener) EnterFieldAccess_lfno_primary(ctx *FieldAccess_lfno_primaryContext) {
}

// ExitFieldAccess_lfno_primary is called when production fieldAccess_lfno_primary is exited.
func (s *BaseJava9ParserListener) ExitFieldAccess_lfno_primary(ctx *FieldAccess_lfno_primaryContext) {
}

// EnterArrayAccess is called when production arrayAccess is entered.
func (s *BaseJava9ParserListener) EnterArrayAccess(ctx *ArrayAccessContext) {}

// ExitArrayAccess is called when production arrayAccess is exited.
func (s *BaseJava9ParserListener) ExitArrayAccess(ctx *ArrayAccessContext) {}

// EnterArrayAccess_lf_primary is called when production arrayAccess_lf_primary is entered.
func (s *BaseJava9ParserListener) EnterArrayAccess_lf_primary(ctx *ArrayAccess_lf_primaryContext) {}

// ExitArrayAccess_lf_primary is called when production arrayAccess_lf_primary is exited.
func (s *BaseJava9ParserListener) ExitArrayAccess_lf_primary(ctx *ArrayAccess_lf_primaryContext) {}

// EnterArrayAccess_lfno_primary is called when production arrayAccess_lfno_primary is entered.
func (s *BaseJava9ParserListener) EnterArrayAccess_lfno_primary(ctx *ArrayAccess_lfno_primaryContext) {
}

// ExitArrayAccess_lfno_primary is called when production arrayAccess_lfno_primary is exited.
func (s *BaseJava9ParserListener) ExitArrayAccess_lfno_primary(ctx *ArrayAccess_lfno_primaryContext) {
}

// EnterMethodInvocation is called when production methodInvocation is entered.
func (s *BaseJava9ParserListener) EnterMethodInvocation(ctx *MethodInvocationContext) {}

// ExitMethodInvocation is called when production methodInvocation is exited.
func (s *BaseJava9ParserListener) ExitMethodInvocation(ctx *MethodInvocationContext) {}

// EnterMethodInvocation_lf_primary is called when production methodInvocation_lf_primary is entered.
func (s *BaseJava9ParserListener) EnterMethodInvocation_lf_primary(ctx *MethodInvocation_lf_primaryContext) {
}

// ExitMethodInvocation_lf_primary is called when production methodInvocation_lf_primary is exited.
func (s *BaseJava9ParserListener) ExitMethodInvocation_lf_primary(ctx *MethodInvocation_lf_primaryContext) {
}

// EnterMethodInvocation_lfno_primary is called when production methodInvocation_lfno_primary is entered.
func (s *BaseJava9ParserListener) EnterMethodInvocation_lfno_primary(ctx *MethodInvocation_lfno_primaryContext) {
}

// ExitMethodInvocation_lfno_primary is called when production methodInvocation_lfno_primary is exited.
func (s *BaseJava9ParserListener) ExitMethodInvocation_lfno_primary(ctx *MethodInvocation_lfno_primaryContext) {
}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseJava9ParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseJava9ParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterMethodReference is called when production methodReference is entered.
func (s *BaseJava9ParserListener) EnterMethodReference(ctx *MethodReferenceContext) {}

// ExitMethodReference is called when production methodReference is exited.
func (s *BaseJava9ParserListener) ExitMethodReference(ctx *MethodReferenceContext) {}

// EnterMethodReference_lf_primary is called when production methodReference_lf_primary is entered.
func (s *BaseJava9ParserListener) EnterMethodReference_lf_primary(ctx *MethodReference_lf_primaryContext) {
}

// ExitMethodReference_lf_primary is called when production methodReference_lf_primary is exited.
func (s *BaseJava9ParserListener) ExitMethodReference_lf_primary(ctx *MethodReference_lf_primaryContext) {
}

// EnterMethodReference_lfno_primary is called when production methodReference_lfno_primary is entered.
func (s *BaseJava9ParserListener) EnterMethodReference_lfno_primary(ctx *MethodReference_lfno_primaryContext) {
}

// ExitMethodReference_lfno_primary is called when production methodReference_lfno_primary is exited.
func (s *BaseJava9ParserListener) ExitMethodReference_lfno_primary(ctx *MethodReference_lfno_primaryContext) {
}

// EnterArrayCreationExpression is called when production arrayCreationExpression is entered.
func (s *BaseJava9ParserListener) EnterArrayCreationExpression(ctx *ArrayCreationExpressionContext) {}

// ExitArrayCreationExpression is called when production arrayCreationExpression is exited.
func (s *BaseJava9ParserListener) ExitArrayCreationExpression(ctx *ArrayCreationExpressionContext) {}

// EnterDimExprs is called when production dimExprs is entered.
func (s *BaseJava9ParserListener) EnterDimExprs(ctx *DimExprsContext) {}

// ExitDimExprs is called when production dimExprs is exited.
func (s *BaseJava9ParserListener) ExitDimExprs(ctx *DimExprsContext) {}

// EnterDimExpr is called when production dimExpr is entered.
func (s *BaseJava9ParserListener) EnterDimExpr(ctx *DimExprContext) {}

// ExitDimExpr is called when production dimExpr is exited.
func (s *BaseJava9ParserListener) ExitDimExpr(ctx *DimExprContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseJava9ParserListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseJava9ParserListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJava9ParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJava9ParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseJava9ParserListener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseJava9ParserListener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *BaseJava9ParserListener) EnterLambdaParameters(ctx *LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *BaseJava9ParserListener) ExitLambdaParameters(ctx *LambdaParametersContext) {}

// EnterInferredFormalParameterList is called when production inferredFormalParameterList is entered.
func (s *BaseJava9ParserListener) EnterInferredFormalParameterList(ctx *InferredFormalParameterListContext) {
}

// ExitInferredFormalParameterList is called when production inferredFormalParameterList is exited.
func (s *BaseJava9ParserListener) ExitInferredFormalParameterList(ctx *InferredFormalParameterListContext) {
}

// EnterLambdaBody is called when production lambdaBody is entered.
func (s *BaseJava9ParserListener) EnterLambdaBody(ctx *LambdaBodyContext) {}

// ExitLambdaBody is called when production lambdaBody is exited.
func (s *BaseJava9ParserListener) ExitLambdaBody(ctx *LambdaBodyContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseJava9ParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseJava9ParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseJava9ParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseJava9ParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterLeftHandSide is called when production leftHandSide is entered.
func (s *BaseJava9ParserListener) EnterLeftHandSide(ctx *LeftHandSideContext) {}

// ExitLeftHandSide is called when production leftHandSide is exited.
func (s *BaseJava9ParserListener) ExitLeftHandSide(ctx *LeftHandSideContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseJava9ParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseJava9ParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseJava9ParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseJava9ParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterConditionalOrExpression is called when production conditionalOrExpression is entered.
func (s *BaseJava9ParserListener) EnterConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// ExitConditionalOrExpression is called when production conditionalOrExpression is exited.
func (s *BaseJava9ParserListener) ExitConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// EnterConditionalAndExpression is called when production conditionalAndExpression is entered.
func (s *BaseJava9ParserListener) EnterConditionalAndExpression(ctx *ConditionalAndExpressionContext) {
}

// ExitConditionalAndExpression is called when production conditionalAndExpression is exited.
func (s *BaseJava9ParserListener) ExitConditionalAndExpression(ctx *ConditionalAndExpressionContext) {
}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BaseJava9ParserListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BaseJava9ParserListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BaseJava9ParserListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BaseJava9ParserListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseJava9ParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseJava9ParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseJava9ParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseJava9ParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseJava9ParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseJava9ParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseJava9ParserListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseJava9ParserListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseJava9ParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseJava9ParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseJava9ParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseJava9ParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseJava9ParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseJava9ParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPreIncrementExpression is called when production preIncrementExpression is entered.
func (s *BaseJava9ParserListener) EnterPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// ExitPreIncrementExpression is called when production preIncrementExpression is exited.
func (s *BaseJava9ParserListener) ExitPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// EnterPreDecrementExpression is called when production preDecrementExpression is entered.
func (s *BaseJava9ParserListener) EnterPreDecrementExpression(ctx *PreDecrementExpressionContext) {}

// ExitPreDecrementExpression is called when production preDecrementExpression is exited.
func (s *BaseJava9ParserListener) ExitPreDecrementExpression(ctx *PreDecrementExpressionContext) {}

// EnterUnaryExpressionNotPlusMinus is called when production unaryExpressionNotPlusMinus is entered.
func (s *BaseJava9ParserListener) EnterUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) {
}

// ExitUnaryExpressionNotPlusMinus is called when production unaryExpressionNotPlusMinus is exited.
func (s *BaseJava9ParserListener) ExitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) {
}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseJava9ParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseJava9ParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPostIncrementExpression is called when production postIncrementExpression is entered.
func (s *BaseJava9ParserListener) EnterPostIncrementExpression(ctx *PostIncrementExpressionContext) {}

// ExitPostIncrementExpression is called when production postIncrementExpression is exited.
func (s *BaseJava9ParserListener) ExitPostIncrementExpression(ctx *PostIncrementExpressionContext) {}

// EnterPostIncrementExpression_lf_postfixExpression is called when production postIncrementExpression_lf_postfixExpression is entered.
func (s *BaseJava9ParserListener) EnterPostIncrementExpression_lf_postfixExpression(ctx *PostIncrementExpression_lf_postfixExpressionContext) {
}

// ExitPostIncrementExpression_lf_postfixExpression is called when production postIncrementExpression_lf_postfixExpression is exited.
func (s *BaseJava9ParserListener) ExitPostIncrementExpression_lf_postfixExpression(ctx *PostIncrementExpression_lf_postfixExpressionContext) {
}

// EnterPostDecrementExpression is called when production postDecrementExpression is entered.
func (s *BaseJava9ParserListener) EnterPostDecrementExpression(ctx *PostDecrementExpressionContext) {}

// ExitPostDecrementExpression is called when production postDecrementExpression is exited.
func (s *BaseJava9ParserListener) ExitPostDecrementExpression(ctx *PostDecrementExpressionContext) {}

// EnterPostDecrementExpression_lf_postfixExpression is called when production postDecrementExpression_lf_postfixExpression is entered.
func (s *BaseJava9ParserListener) EnterPostDecrementExpression_lf_postfixExpression(ctx *PostDecrementExpression_lf_postfixExpressionContext) {
}

// ExitPostDecrementExpression_lf_postfixExpression is called when production postDecrementExpression_lf_postfixExpression is exited.
func (s *BaseJava9ParserListener) ExitPostDecrementExpression_lf_postfixExpression(ctx *PostDecrementExpression_lf_postfixExpressionContext) {
}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseJava9ParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseJava9ParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseJava9ParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseJava9ParserListener) ExitIdentifier(ctx *IdentifierContext) {}
