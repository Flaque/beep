// Code generated from parser/BeepBoop.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 126,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 7, 3, 50, 10, 3, 12, 3, 14, 3, 53, 11, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 6, 4, 59, 10, 4, 13, 4, 14, 4, 60, 5, 4, 63, 10, 4, 3, 5, 6,
	5, 66, 10, 5, 13, 5, 14, 5, 67, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 6, 21, 118,
	10, 21, 13, 21, 14, 21, 119, 3, 22, 6, 22, 123, 10, 22, 13, 22, 14, 22,
	124, 2, 2, 23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 3, 2, 6, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11,
	34, 34, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 2, 131, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 3, 45, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 62, 3, 2, 2, 2, 9,
	65, 3, 2, 2, 2, 11, 71, 3, 2, 2, 2, 13, 74, 3, 2, 2, 2, 15, 77, 3, 2, 2,
	2, 17, 81, 3, 2, 2, 2, 19, 86, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2, 23, 97,
	3, 2, 2, 2, 25, 100, 3, 2, 2, 2, 27, 102, 3, 2, 2, 2, 29, 104, 3, 2, 2,
	2, 31, 106, 3, 2, 2, 2, 33, 108, 3, 2, 2, 2, 35, 110, 3, 2, 2, 2, 37, 112,
	3, 2, 2, 2, 39, 114, 3, 2, 2, 2, 41, 117, 3, 2, 2, 2, 43, 122, 3, 2, 2,
	2, 45, 46, 7, 38, 2, 2, 46, 4, 3, 2, 2, 2, 47, 51, 7, 37, 2, 2, 48, 50,
	10, 2, 2, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2,
	51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 8,
	3, 2, 2, 55, 6, 3, 2, 2, 2, 56, 63, 9, 2, 2, 2, 57, 59, 9, 2, 2, 2, 58,
	57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2,
	2, 61, 63, 3, 2, 2, 2, 62, 56, 3, 2, 2, 2, 62, 58, 3, 2, 2, 2, 63, 8, 3,
	2, 2, 2, 64, 66, 9, 3, 2, 2, 65, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67,
	65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 8, 5, 3,
	2, 70, 10, 3, 2, 2, 2, 71, 72, 7, 107, 2, 2, 72, 73, 7, 104, 2, 2, 73,
	12, 3, 2, 2, 2, 74, 75, 7, 102, 2, 2, 75, 76, 7, 113, 2, 2, 76, 14, 3,
	2, 2, 2, 77, 78, 7, 103, 2, 2, 78, 79, 7, 112, 2, 2, 79, 80, 7, 102, 2,
	2, 80, 16, 3, 2, 2, 2, 81, 82, 7, 104, 2, 2, 82, 83, 7, 119, 2, 2, 83,
	84, 7, 112, 2, 2, 84, 85, 7, 101, 2, 2, 85, 18, 3, 2, 2, 2, 86, 87, 7,
	116, 2, 2, 87, 88, 7, 103, 2, 2, 88, 89, 7, 118, 2, 2, 89, 90, 7, 119,
	2, 2, 90, 91, 7, 116, 2, 2, 91, 92, 7, 112, 2, 2, 92, 20, 3, 2, 2, 2, 93,
	94, 7, 104, 2, 2, 94, 95, 7, 113, 2, 2, 95, 96, 7, 116, 2, 2, 96, 22, 3,
	2, 2, 2, 97, 98, 7, 107, 2, 2, 98, 99, 7, 117, 2, 2, 99, 24, 3, 2, 2, 2,
	100, 101, 7, 45, 2, 2, 101, 26, 3, 2, 2, 2, 102, 103, 7, 47, 2, 2, 103,
	28, 3, 2, 2, 2, 104, 105, 7, 44, 2, 2, 105, 30, 3, 2, 2, 2, 106, 107, 7,
	49, 2, 2, 107, 32, 3, 2, 2, 2, 108, 109, 7, 63, 2, 2, 109, 34, 3, 2, 2,
	2, 110, 111, 7, 126, 2, 2, 111, 36, 3, 2, 2, 2, 112, 113, 7, 42, 2, 2,
	113, 38, 3, 2, 2, 2, 114, 115, 7, 43, 2, 2, 115, 40, 3, 2, 2, 2, 116, 118,
	9, 4, 2, 2, 117, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 117, 3, 2,
	2, 2, 119, 120, 3, 2, 2, 2, 120, 42, 3, 2, 2, 2, 121, 123, 9, 5, 2, 2,
	122, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124,
	125, 3, 2, 2, 2, 125, 44, 3, 2, 2, 2, 9, 2, 51, 60, 62, 67, 119, 124, 4,
	8, 2, 2, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'$'", "", "", "", "'if'", "'do'", "'end'", "'func'", "'return'", "'for'",
	"'is'", "'+'", "'-'", "'*'", "'/'", "'='", "'|'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "COMMENT", "NEWLINE", "WHITESPACE", "IF", "DO", "END", "FUNC",
	"RETURN", "FOR", "IS", "PLUS", "MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE",
	"LPAREN", "RPAREN", "INT", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "COMMENT", "NEWLINE", "WHITESPACE", "IF", "DO", "END", "FUNC",
	"RETURN", "FOR", "IS", "PLUS", "MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE",
	"LPAREN", "RPAREN", "INT", "STRING",
}

type BeepBoopLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewBeepBoopLexer(input antlr.CharStream) *BeepBoopLexer {

	l := new(BeepBoopLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "BeepBoop.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BeepBoopLexer tokens.
const (
	BeepBoopLexerT__0       = 1
	BeepBoopLexerCOMMENT    = 2
	BeepBoopLexerNEWLINE    = 3
	BeepBoopLexerWHITESPACE = 4
	BeepBoopLexerIF         = 5
	BeepBoopLexerDO         = 6
	BeepBoopLexerEND        = 7
	BeepBoopLexerFUNC       = 8
	BeepBoopLexerRETURN     = 9
	BeepBoopLexerFOR        = 10
	BeepBoopLexerIS         = 11
	BeepBoopLexerPLUS       = 12
	BeepBoopLexerMINUS      = 13
	BeepBoopLexerMULT       = 14
	BeepBoopLexerDIVIDE     = 15
	BeepBoopLexerASSIGN     = 16
	BeepBoopLexerPIPE       = 17
	BeepBoopLexerLPAREN     = 18
	BeepBoopLexerRPAREN     = 19
	BeepBoopLexerINT        = 20
	BeepBoopLexerSTRING     = 21
)
