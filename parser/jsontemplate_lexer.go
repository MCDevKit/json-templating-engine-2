// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JsonTemplateLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var jsontemplatelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsontemplatelexerLexerInit() {
	staticData := &jsontemplatelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'<'", "'<='", "'=='", "'>'", "'>='", "'!='", "'&&'", "'||'", "'!'",
		"'+'", "'-'", "'*'", "'/'", "'('", "')'", "'['", "']'", "'#'", "'?'",
		"'='", "'??'", "'..'", "'...'", "'as'", "','", "'=>'", "':'", "';'",
		"'.'", "'{'", "'}'", "'null'", "'false'", "'true'", "'return'", "'for'",
		"'in'", "'if'", "'else'", "'while'", "'do'", "'break'", "'continue'",
	}
	staticData.symbolicNames = []string{
		"", "Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual", "NotEqual",
		"And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide", "LeftParen",
		"RightParen", "LeftBracket", "RightBracket", "Iteration", "Question",
		"Literal", "NullCoalescing", "Range", "Spread", "As", "Comma", "Arrow",
		"Colon", "Semicolon", "Dot", "LeftBrace", "RightBrace", "Null", "False",
		"True", "Return", "For", "In", "If", "Else", "While", "Do", "Break",
		"Continue", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"Less", "LessOrEqual", "Equal", "Greater", "GreaterOrEqual", "NotEqual",
		"And", "Or", "Not", "Add", "Subtract", "Multiply", "Divide", "LeftParen",
		"RightParen", "LeftBracket", "RightBracket", "Iteration", "Question",
		"Literal", "NullCoalescing", "Range", "Spread", "As", "Comma", "Arrow",
		"Colon", "Semicolon", "Dot", "LeftBrace", "RightBrace", "Null", "False",
		"True", "Return", "For", "In", "If", "Else", "While", "Do", "Break",
		"Continue", "ESCAPED_STRING", "STRING", "NUMBER", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 47, 277, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1,
		43, 5, 43, 236, 8, 43, 10, 43, 12, 43, 239, 9, 43, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 5, 43, 246, 8, 43, 10, 43, 12, 43, 249, 9, 43, 1, 43, 3,
		43, 252, 8, 43, 1, 44, 1, 44, 5, 44, 256, 8, 44, 10, 44, 12, 44, 259, 9,
		44, 1, 45, 4, 45, 262, 8, 45, 11, 45, 12, 45, 263, 1, 45, 1, 45, 4, 45,
		268, 8, 45, 11, 45, 12, 45, 269, 3, 45, 272, 8, 45, 1, 46, 1, 46, 1, 46,
		1, 46, 0, 0, 47, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71,
		36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89,
		45, 91, 46, 93, 47, 1, 0, 6, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92,
		4, 0, 36, 36, 65, 90, 95, 95, 97, 122, 5, 0, 36, 36, 48, 57, 65, 90, 95,
		95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 285, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0,
		0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0,
		0, 1, 95, 1, 0, 0, 0, 3, 97, 1, 0, 0, 0, 5, 100, 1, 0, 0, 0, 7, 103, 1,
		0, 0, 0, 9, 105, 1, 0, 0, 0, 11, 108, 1, 0, 0, 0, 13, 111, 1, 0, 0, 0,
		15, 114, 1, 0, 0, 0, 17, 117, 1, 0, 0, 0, 19, 119, 1, 0, 0, 0, 21, 121,
		1, 0, 0, 0, 23, 123, 1, 0, 0, 0, 25, 125, 1, 0, 0, 0, 27, 127, 1, 0, 0,
		0, 29, 129, 1, 0, 0, 0, 31, 131, 1, 0, 0, 0, 33, 133, 1, 0, 0, 0, 35, 135,
		1, 0, 0, 0, 37, 137, 1, 0, 0, 0, 39, 139, 1, 0, 0, 0, 41, 141, 1, 0, 0,
		0, 43, 144, 1, 0, 0, 0, 45, 147, 1, 0, 0, 0, 47, 151, 1, 0, 0, 0, 49, 154,
		1, 0, 0, 0, 51, 156, 1, 0, 0, 0, 53, 159, 1, 0, 0, 0, 55, 161, 1, 0, 0,
		0, 57, 163, 1, 0, 0, 0, 59, 165, 1, 0, 0, 0, 61, 167, 1, 0, 0, 0, 63, 169,
		1, 0, 0, 0, 65, 174, 1, 0, 0, 0, 67, 180, 1, 0, 0, 0, 69, 185, 1, 0, 0,
		0, 71, 192, 1, 0, 0, 0, 73, 196, 1, 0, 0, 0, 75, 199, 1, 0, 0, 0, 77, 202,
		1, 0, 0, 0, 79, 207, 1, 0, 0, 0, 81, 213, 1, 0, 0, 0, 83, 216, 1, 0, 0,
		0, 85, 222, 1, 0, 0, 0, 87, 251, 1, 0, 0, 0, 89, 253, 1, 0, 0, 0, 91, 261,
		1, 0, 0, 0, 93, 273, 1, 0, 0, 0, 95, 96, 5, 60, 0, 0, 96, 2, 1, 0, 0, 0,
		97, 98, 5, 60, 0, 0, 98, 99, 5, 61, 0, 0, 99, 4, 1, 0, 0, 0, 100, 101,
		5, 61, 0, 0, 101, 102, 5, 61, 0, 0, 102, 6, 1, 0, 0, 0, 103, 104, 5, 62,
		0, 0, 104, 8, 1, 0, 0, 0, 105, 106, 5, 62, 0, 0, 106, 107, 5, 61, 0, 0,
		107, 10, 1, 0, 0, 0, 108, 109, 5, 33, 0, 0, 109, 110, 5, 61, 0, 0, 110,
		12, 1, 0, 0, 0, 111, 112, 5, 38, 0, 0, 112, 113, 5, 38, 0, 0, 113, 14,
		1, 0, 0, 0, 114, 115, 5, 124, 0, 0, 115, 116, 5, 124, 0, 0, 116, 16, 1,
		0, 0, 0, 117, 118, 5, 33, 0, 0, 118, 18, 1, 0, 0, 0, 119, 120, 5, 43, 0,
		0, 120, 20, 1, 0, 0, 0, 121, 122, 5, 45, 0, 0, 122, 22, 1, 0, 0, 0, 123,
		124, 5, 42, 0, 0, 124, 24, 1, 0, 0, 0, 125, 126, 5, 47, 0, 0, 126, 26,
		1, 0, 0, 0, 127, 128, 5, 40, 0, 0, 128, 28, 1, 0, 0, 0, 129, 130, 5, 41,
		0, 0, 130, 30, 1, 0, 0, 0, 131, 132, 5, 91, 0, 0, 132, 32, 1, 0, 0, 0,
		133, 134, 5, 93, 0, 0, 134, 34, 1, 0, 0, 0, 135, 136, 5, 35, 0, 0, 136,
		36, 1, 0, 0, 0, 137, 138, 5, 63, 0, 0, 138, 38, 1, 0, 0, 0, 139, 140, 5,
		61, 0, 0, 140, 40, 1, 0, 0, 0, 141, 142, 5, 63, 0, 0, 142, 143, 5, 63,
		0, 0, 143, 42, 1, 0, 0, 0, 144, 145, 5, 46, 0, 0, 145, 146, 5, 46, 0, 0,
		146, 44, 1, 0, 0, 0, 147, 148, 5, 46, 0, 0, 148, 149, 5, 46, 0, 0, 149,
		150, 5, 46, 0, 0, 150, 46, 1, 0, 0, 0, 151, 152, 5, 97, 0, 0, 152, 153,
		5, 115, 0, 0, 153, 48, 1, 0, 0, 0, 154, 155, 5, 44, 0, 0, 155, 50, 1, 0,
		0, 0, 156, 157, 5, 61, 0, 0, 157, 158, 5, 62, 0, 0, 158, 52, 1, 0, 0, 0,
		159, 160, 5, 58, 0, 0, 160, 54, 1, 0, 0, 0, 161, 162, 5, 59, 0, 0, 162,
		56, 1, 0, 0, 0, 163, 164, 5, 46, 0, 0, 164, 58, 1, 0, 0, 0, 165, 166, 5,
		123, 0, 0, 166, 60, 1, 0, 0, 0, 167, 168, 5, 125, 0, 0, 168, 62, 1, 0,
		0, 0, 169, 170, 5, 110, 0, 0, 170, 171, 5, 117, 0, 0, 171, 172, 5, 108,
		0, 0, 172, 173, 5, 108, 0, 0, 173, 64, 1, 0, 0, 0, 174, 175, 5, 102, 0,
		0, 175, 176, 5, 97, 0, 0, 176, 177, 5, 108, 0, 0, 177, 178, 5, 115, 0,
		0, 178, 179, 5, 101, 0, 0, 179, 66, 1, 0, 0, 0, 180, 181, 5, 116, 0, 0,
		181, 182, 5, 114, 0, 0, 182, 183, 5, 117, 0, 0, 183, 184, 5, 101, 0, 0,
		184, 68, 1, 0, 0, 0, 185, 186, 5, 114, 0, 0, 186, 187, 5, 101, 0, 0, 187,
		188, 5, 116, 0, 0, 188, 189, 5, 117, 0, 0, 189, 190, 5, 114, 0, 0, 190,
		191, 5, 110, 0, 0, 191, 70, 1, 0, 0, 0, 192, 193, 5, 102, 0, 0, 193, 194,
		5, 111, 0, 0, 194, 195, 5, 114, 0, 0, 195, 72, 1, 0, 0, 0, 196, 197, 5,
		105, 0, 0, 197, 198, 5, 110, 0, 0, 198, 74, 1, 0, 0, 0, 199, 200, 5, 105,
		0, 0, 200, 201, 5, 102, 0, 0, 201, 76, 1, 0, 0, 0, 202, 203, 5, 101, 0,
		0, 203, 204, 5, 108, 0, 0, 204, 205, 5, 115, 0, 0, 205, 206, 5, 101, 0,
		0, 206, 78, 1, 0, 0, 0, 207, 208, 5, 119, 0, 0, 208, 209, 5, 104, 0, 0,
		209, 210, 5, 105, 0, 0, 210, 211, 5, 108, 0, 0, 211, 212, 5, 101, 0, 0,
		212, 80, 1, 0, 0, 0, 213, 214, 5, 100, 0, 0, 214, 215, 5, 111, 0, 0, 215,
		82, 1, 0, 0, 0, 216, 217, 5, 98, 0, 0, 217, 218, 5, 114, 0, 0, 218, 219,
		5, 101, 0, 0, 219, 220, 5, 97, 0, 0, 220, 221, 5, 107, 0, 0, 221, 84, 1,
		0, 0, 0, 222, 223, 5, 99, 0, 0, 223, 224, 5, 111, 0, 0, 224, 225, 5, 110,
		0, 0, 225, 226, 5, 116, 0, 0, 226, 227, 5, 105, 0, 0, 227, 228, 5, 110,
		0, 0, 228, 229, 5, 117, 0, 0, 229, 230, 5, 101, 0, 0, 230, 86, 1, 0, 0,
		0, 231, 237, 5, 34, 0, 0, 232, 233, 5, 92, 0, 0, 233, 236, 9, 0, 0, 0,
		234, 236, 8, 0, 0, 0, 235, 232, 1, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236,
		239, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 240,
		1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 240, 252, 5, 34, 0, 0, 241, 247, 5, 39,
		0, 0, 242, 243, 5, 92, 0, 0, 243, 246, 9, 0, 0, 0, 244, 246, 8, 1, 0, 0,
		245, 242, 1, 0, 0, 0, 245, 244, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247,
		245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249, 247,
		1, 0, 0, 0, 250, 252, 5, 39, 0, 0, 251, 231, 1, 0, 0, 0, 251, 241, 1, 0,
		0, 0, 252, 88, 1, 0, 0, 0, 253, 257, 7, 2, 0, 0, 254, 256, 7, 3, 0, 0,
		255, 254, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257,
		258, 1, 0, 0, 0, 258, 90, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 260, 262, 7,
		4, 0, 0, 261, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 261, 1, 0, 0,
		0, 263, 264, 1, 0, 0, 0, 264, 271, 1, 0, 0, 0, 265, 267, 5, 46, 0, 0, 266,
		268, 7, 4, 0, 0, 267, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 267,
		1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 272, 1, 0, 0, 0, 271, 265, 1, 0,
		0, 0, 271, 272, 1, 0, 0, 0, 272, 92, 1, 0, 0, 0, 273, 274, 7, 5, 0, 0,
		274, 275, 1, 0, 0, 0, 275, 276, 6, 46, 0, 0, 276, 94, 1, 0, 0, 0, 10, 0,
		235, 237, 245, 247, 251, 257, 263, 269, 271, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JsonTemplateLexerInit initializes any static state used to implement JsonTemplateLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonTemplateLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonTemplateLexerInit() {
	staticData := &jsontemplatelexerLexerStaticData
	staticData.once.Do(jsontemplatelexerLexerInit)
}

// NewJsonTemplateLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonTemplateLexer(input antlr.CharStream) *JsonTemplateLexer {
	JsonTemplateLexerInit()
	l := new(JsonTemplateLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &jsontemplatelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "JsonTemplate.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonTemplateLexer tokens.
const (
	JsonTemplateLexerLess           = 1
	JsonTemplateLexerLessOrEqual    = 2
	JsonTemplateLexerEqual          = 3
	JsonTemplateLexerGreater        = 4
	JsonTemplateLexerGreaterOrEqual = 5
	JsonTemplateLexerNotEqual       = 6
	JsonTemplateLexerAnd            = 7
	JsonTemplateLexerOr             = 8
	JsonTemplateLexerNot            = 9
	JsonTemplateLexerAdd            = 10
	JsonTemplateLexerSubtract       = 11
	JsonTemplateLexerMultiply       = 12
	JsonTemplateLexerDivide         = 13
	JsonTemplateLexerLeftParen      = 14
	JsonTemplateLexerRightParen     = 15
	JsonTemplateLexerLeftBracket    = 16
	JsonTemplateLexerRightBracket   = 17
	JsonTemplateLexerIteration      = 18
	JsonTemplateLexerQuestion       = 19
	JsonTemplateLexerLiteral        = 20
	JsonTemplateLexerNullCoalescing = 21
	JsonTemplateLexerRange          = 22
	JsonTemplateLexerSpread         = 23
	JsonTemplateLexerAs             = 24
	JsonTemplateLexerComma          = 25
	JsonTemplateLexerArrow          = 26
	JsonTemplateLexerColon          = 27
	JsonTemplateLexerSemicolon      = 28
	JsonTemplateLexerDot            = 29
	JsonTemplateLexerLeftBrace      = 30
	JsonTemplateLexerRightBrace     = 31
	JsonTemplateLexerNull           = 32
	JsonTemplateLexerFalse          = 33
	JsonTemplateLexerTrue           = 34
	JsonTemplateLexerReturn         = 35
	JsonTemplateLexerFor            = 36
	JsonTemplateLexerIn             = 37
	JsonTemplateLexerIf             = 38
	JsonTemplateLexerElse           = 39
	JsonTemplateLexerWhile          = 40
	JsonTemplateLexerDo             = 41
	JsonTemplateLexerBreak          = 42
	JsonTemplateLexerContinue       = 43
	JsonTemplateLexerESCAPED_STRING = 44
	JsonTemplateLexerSTRING         = 45
	JsonTemplateLexerNUMBER         = 46
	JsonTemplateLexerWS             = 47
)
